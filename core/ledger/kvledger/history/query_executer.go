/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package history

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/ledger/queryresult"
	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	protoutil "github.com/hyperledger/fabric/protoutil"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb/iterator"
)

// QueryExecutor is a query executor against the LevelDB history DB
type QueryExecutor struct {
	levelDB    *leveldbhelper.DBHandle
	blockStore *blkstorage.BlockStore
}

// GetHistoryForKey implements method in interface `ledger.HistoryQueryExecutor`
func (q *QueryExecutor) GetHistoryForKey(namespace string, key string) (commonledger.ResultsIterator, error) {
	rangeScan := constructRangeScan(namespace, key)
	dbItr, err := q.levelDB.GetIterator(rangeScan.startKey, rangeScan.endKey)
	if err != nil {
		return nil, err
	}

	// By default, dbItr is in the orderer of oldest to newest and its cursor is at the beginning of the entries.
	// Need to call Last() and Next() to move the cursor to the end of the entries so that we can iterate
	// the entries in the order of newest to oldest.
	if dbItr.Last() {
		dbItr.Next()
	}
	return &historyScanner{rangeScan, namespace, key, dbItr, q.blockStore, newSaiIter()}, nil
}

// historyScanner implements ResultsIterator for iterating through history results
type historyScanner struct {
	rangeScan        *rangeScan
	namespace        string
	key              string
	dbItr            iterator.Iterator
	blockStore       *blkstorage.BlockStore
	getNextBlockTran func(key string) (uint64, uint64)
}

func findGlobalBlockNum(globalBytes *[]byte, key string) (uint64, uint64) {
	globalIndex := make(map[string][]uint64)
	json.Unmarshal(*globalBytes, &globalIndex)
	return globalIndex[key][0], globalIndex[key][1]
}

func decodeLocalBlockTran(localFile *os.File, key string, tranNum uint64) (uint64, uint64) {
	liKey := key + "," + strconv.FormatUint(tranNum, 10)
	decoder := json.NewDecoder(localFile)
	for decoder.More() {
		var curr LIEntry
		err := decoder.Decode(&curr)
		if err != nil {
			log.Println(err)
		}
		if liKey == curr.KT {
			return curr.Prev[0], curr.Prev[1]
		}
	}
	return 0, 0
}

func newSaiIter() func(key string) (uint64, uint64) {
	first := true
	var prevBlockNum uint64
	var currBlockNum uint64
	var tranNum uint64
	return func(key string) (uint64, uint64) {
		// If first flag hasn't been set to false, use global index to find most recent block
		if first {
			first = false
			globalIndexFile, err := os.Open("/var/GI-Storage/globalIndex.json")
			if err != nil {
				log.Println(err)
			}
			defer CloseFile(globalIndexFile)
			LockFile(globalIndexFile)
			globalBytes, err := io.ReadAll(globalIndexFile)
			if err != nil {
				log.Println(err)
			}
			UnlockFile(globalIndexFile)
			prevBlockNum, tranNum = findGlobalBlockNum(&globalBytes, key)
		}
		currBlockNum = prevBlockNum
		localFileName := "/var/LI-Storage/localIndex-" + strconv.FormatUint(prevBlockNum, 10) + ".json"
		localIndexFile, err := os.Open(localFileName)
		if err != nil {
			log.Println(err)
		}
		defer CloseFile(localIndexFile)
		LockFile(localIndexFile)
		prevBlockNum, tranNum = decodeLocalBlockTran(localIndexFile, key, tranNum)
		UnlockFile(localIndexFile)
		return currBlockNum, tranNum
	}
}

// Next iterates to the next key, in the order of newest to oldest, from history scanner.
// It decodes blockNumTranNumBytes to get blockNum and tranNum,
// loads the block:tran from block storage, finds the key and returns the result.
func (scanner *historyScanner) Next() (commonledger.QueryResult, error) {
	// call Prev because history query result is returned from newest to oldest
	// if !scanner.dbItr.Prev() {
	// 	return nil, nil
	// }

	// historyKey := scanner.dbItr.Key()
	// blockNum, tranNum, err := scanner.rangeScan.decodeBlockNumTranNum(historyKey)
	// if err != nil {
	// 	return nil, err
	// }

	blockNum, tranNum := scanner.getNextBlockTran(scanner.key)

	logger.Debugf("Found history record for namespace:%s key:%s at blockNumTranNum %v:%v\n",
		scanner.namespace, scanner.key, blockNum, tranNum)

	// Get the transaction from block storage that is associated with this history record
	tranEnvelope, err := scanner.blockStore.RetrieveTxByBlockNumTranNum(blockNum, tranNum)
	if err != nil {
		return nil, err
	}

	// Get the txid, key write value, timestamp, and delete indicator associated with this transaction
	queryResult, err := getKeyModificationFromTran(tranEnvelope, scanner.namespace, scanner.key)
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		// should not happen, but make sure there is inconsistency between historydb and statedb
		logger.Errorf("No namespace or key is found for namespace %s and key %s with decoded blockNum %d and tranNum %d", scanner.namespace, scanner.key, blockNum, tranNum)
		return nil, errors.Errorf("no namespace or key is found for namespace %s and key %s with decoded blockNum %d and tranNum %d", scanner.namespace, scanner.key, blockNum, tranNum)
	}
	logger.Debugf("Found historic key value for namespace:%s key:%s from transaction %s",
		scanner.namespace, scanner.key, queryResult.(*queryresult.KeyModification).TxId)
	return queryResult, nil
}

func (scanner *historyScanner) Close() {
	scanner.dbItr.Release()
}

// getTxIDandKeyWriteValueFromTran inspects a transaction for writes to a given key
func getKeyModificationFromTran(tranEnvelope *common.Envelope, namespace string, key string) (commonledger.QueryResult, error) {
	logger.Debugf("Entering getKeyModificationFromTran()\n", namespace, key)

	// extract action from the envelope
	payload, err := protoutil.UnmarshalPayload(tranEnvelope.Payload)
	if err != nil {
		return nil, err
	}

	tx, err := protoutil.UnmarshalTransaction(payload.Data)
	if err != nil {
		return nil, err
	}

	_, respPayload, err := protoutil.GetPayloads(tx.Actions[0])
	if err != nil {
		return nil, err
	}

	chdr, err := protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return nil, err
	}

	txID := chdr.TxId
	timestamp := chdr.Timestamp

	txRWSet := &rwsetutil.TxRwSet{}

	// Get the Result from the Action and then Unmarshal
	// it into a TxReadWriteSet using custom unmarshalling
	if err = txRWSet.FromProtoBytes(respPayload.Results); err != nil {
		return nil, err
	}

	// look for the namespace and key by looping through the transaction's ReadWriteSets
	for _, nsRWSet := range txRWSet.NsRwSets {
		if nsRWSet.NameSpace == namespace {
			// got the correct namespace, now find the key write
			for _, kvWrite := range nsRWSet.KvRwSet.Writes {
				if kvWrite.Key == key {
					return &queryresult.KeyModification{TxId: txID, Value: kvWrite.Value,
						Timestamp: timestamp, IsDelete: rwsetutil.IsKVWriteDelete(kvWrite)}, nil
				}
			} // end keys loop
			logger.Debugf("key [%s] not found in namespace [%s]'s writeset", key, namespace)
			return nil, nil
		} // end if
	} //end namespaces loop
	logger.Debugf("namespace [%s] not found in transaction's ReadWriteSets", namespace)
	return nil, nil
}
