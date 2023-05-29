/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package history

import (
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/ledger/blkstorage"
	"github.com/hyperledger/fabric/common/ledger/dataformat"
	"github.com/hyperledger/fabric/common/ledger/util/leveldbhelper"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/internal/version"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/internal/pkg/txflags"
	protoutil "github.com/hyperledger/fabric/protoutil"
	"io/ioutil"
	"log"
	"os"
)

var logger = flogging.MustGetLogger("history")

type DataKey struct {
	Key         string `json:"k"`
	BlockNumber uint64 `json:"bn"`
	TranNumber  uint64 `json:"tn"`
}

type BlockNumber struct {
	Bn uint64 `json:"bn"`
}

// DBProvider provides handle to HistoryDB for a given channel
type DBProvider struct {
	leveldbProvider *leveldbhelper.Provider
}

// NewDBProvider instantiates DBProvider
func NewDBProvider(path string) (*DBProvider, error) {
	logger.Debugf("constructing HistoryDBProvider dbPath=%s", path)
	levelDBProvider, err := leveldbhelper.NewProvider(
		&leveldbhelper.Conf{
			DBPath:         path,
			ExpectedFormat: dataformat.CurrentFormat,
		},
	)
	if err != nil {
		return nil, err
	}
	return &DBProvider{
		leveldbProvider: levelDBProvider,
	}, nil
}

// GetDBHandle gets the handle to a named database
func (p *DBProvider) GetDBHandle(name string) (*DB, error) {
	return &DB{
			levelDB: p.leveldbProvider.GetDBHandle(name),
			name:    name,
		},
		nil
}

// Close closes the underlying db
func (p *DBProvider) Close() {
	p.leveldbProvider.Close()
}

// DB maintains and provides access to history data for a particular channel
type DB struct {
	levelDB *leveldbhelper.DBHandle
	name    string
}

// Commit implements method in HistoryDB interface
func (d *DB) Commit(block *common.Block) error {

	blockNo := block.Header.Number
	//Set the starting tranNo to 0
	var tranNo uint64

	dbBatch := d.levelDB.NewUpdateBatch()

	logger.Debugf("Channel [%s]: Updating history database for blockNo [%v] with [%d] transactions",
		d.name, blockNo, len(block.Data.Data))

	// Open the file in append mode or create
	file, err := os.OpenFile("/var/SAIStorage/saiStorage.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	// Create a map to keep track of the latest value of each key
	blockRecord := make(map[string]BlockNumber)

	// Get the invalidation byte array for the block
	txsFilter := txflags.ValidationFlags(block.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER])

	// write each tran's write set to history db
	for _, envBytes := range block.Data.Data {

		// If the tran is marked as invalid, skip it
		if txsFilter.IsInvalid(int(tranNo)) {
			logger.Debugf("Channel [%s]: Skipping history write for invalid transaction number %d",
				d.name, tranNo)
			tranNo++
			continue
		}

		env, err := protoutil.GetEnvelopeFromBlock(envBytes)
		if err != nil {
			return err
		}

		payload, err := protoutil.UnmarshalPayload(env.Payload)
		if err != nil {
			return err
		}

		chdr, err := protoutil.UnmarshalChannelHeader(payload.Header.ChannelHeader)
		if err != nil {
			return err
		}

		if common.HeaderType(chdr.Type) == common.HeaderType_ENDORSER_TRANSACTION {
			// extract RWSet from transaction
			respPayload, err := protoutil.GetActionFromEnvelope(envBytes)
			if err != nil {
				return err
			}
			txRWSet := &rwsetutil.TxRwSet{}
			if err = txRWSet.FromProtoBytes(respPayload.Results); err != nil {
				return err
			}
			// add a history record for each write
			for _, nsRWSet := range txRWSet.NsRwSets {
				ns := nsRWSet.NameSpace

				for _, kvWrite := range nsRWSet.KvRwSet.Writes {
					dataKey := constructDataKey(ns, kvWrite.Key, blockNo, tranNo)

					// Create a new DataKey instance and set its fields
					dk := DataKey{
						kvWrite.Key,
						blockNo,
						tranNo,
					}

					blockRecord[kvWrite.Key] = BlockNumber{Bn: blockNo}

					// Convert the DataKey instance to json
					jsonBytes, err := json.Marshal(dk)
					if err != nil {
						log.Println(err)
					}

					// Convert the JSON bytes to a string and append a newline
					jsonString := string(jsonBytes) + "\n"

					// Write a newline character to the file
					if _, err := file.WriteString(jsonString); err != nil {
						log.Println(err)
					}

					// No value is required, write an empty byte array (emptyValue) since Put() of nil is not allowed
					dbBatch.Put(dataKey, emptyValue)
				}
			}

		} else {
			logger.Debugf("Skipping transaction [%d] since it is not an endorsement transaction\n", tranNo)
		}
		tranNo++
	}

	// add savepoint for recovery purpose
	height := version.NewHeight(blockNo, tranNo)
	dbBatch.Put(savePointKey, height.ToBytes())

	// write the block's history records and savepoint to LevelDB
	// Setting snyc to true as a precaution, false may be an ok optimization after further testing.
	if err := d.levelDB.WriteBatch(dbBatch, true); err != nil {
		return err
	}

	logger.Debugf("Channel [%s]: Updates committed to history database for blockNo [%v]", d.name, blockNo)

	// After committing the block, serialize the lastRecord map and write it to globalIndex.json
	indexFile, err := os.OpenFile("/var/SAIStorage/globalIndex.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(indexFile)

	// NEW
	indexBytes, err := ioutil.ReadAll(indexFile)
	if err != nil {
		log.Println(err)
	}

	// NEW
	existingRecords := make(map[string]BlockNumber)
	json.Unmarshal(indexBytes, &existingRecords)

	// NEW
	for key, record := range blockRecord {
		existingRecords[key] = record
	}

	// NEW
	indexFile.Seek(0, 0)
	indexFile.Truncate(0)
	indexBytes, err = json.Marshal(existingRecords)
	if err != nil {
		log.Println(err)
	}

	if _, err := indexFile.Write(indexBytes); err != nil {
		log.Println(err)
	}

	return nil
}

// NewQueryExecutor implements method in HistoryDB interface
func (d *DB) NewQueryExecutor(blockStore *blkstorage.BlockStore) (ledger.HistoryQueryExecutor, error) {
	return &QueryExecutor{d.levelDB, blockStore}, nil
}

// GetLastSavepoint implements returns the height till which the history is present in the db
func (d *DB) GetLastSavepoint() (*version.Height, error) {
	versionBytes, err := d.levelDB.Get(savePointKey)
	if err != nil || versionBytes == nil {
		return nil, err
	}
	height, _, err := version.NewHeightFromBytes(versionBytes)
	if err != nil {
		return nil, err
	}
	return height, nil
}

// ShouldRecover implements method in interface kvledger.Recoverer
func (d *DB) ShouldRecover(lastAvailableBlock uint64) (bool, uint64, error) {
	savepoint, err := d.GetLastSavepoint()
	if err != nil {
		return false, 0, err
	}
	if savepoint == nil {
		return true, 0, nil
	}
	return savepoint.BlockNum != lastAvailableBlock, savepoint.BlockNum + 1, nil
}

// Name returns the name of the database that manages historical states.
func (d *DB) Name() string {
	return "history"
}

// CommitLostBlock implements method in interface kvledger.Recoverer
func (d *DB) CommitLostBlock(blockAndPvtdata *ledger.BlockAndPvtData) error {
	block := blockAndPvtdata.Block

	// log every 1000th block at Info level so that history rebuild progress can be tracked in production envs.
	if block.Header.Number%1000 == 0 {
		logger.Infof("Recommitting block [%d] to history database", block.Header.Number)
	} else {
		logger.Debugf("Recommitting block [%d] to history database", block.Header.Number)
	}

	if err := d.Commit(block); err != nil {
		return err
	}
	return nil
}
