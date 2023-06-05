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
	"syscall"

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
)

var logger = flogging.MustGetLogger("history")

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

// DataKey Local index entry struct
type DataKey struct {
	Key        string `json:"key"`
	PrevBlock  uint64 `json:"prev"`
	TranNumber uint64 `json:"tx"`
}

// DataKeySai structure
type DataKeySai struct {
	Key         string `json:"k"`
	BlockNumber uint64 `json:"bn"`
	TranNumber  uint64 `json:"tn"`
}

func lockFile(file *os.File) {
	syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
}

func unlockFile(file *os.File) {
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(err)
	}
}

// Commit implements method in HistoryDB interface
func (d *DB) Commit(block *common.Block) error {

	blockNo := block.Header.Number
	// SET THE STARTING tranNo to 0
	var tranNo uint64

	dbBatch := d.levelDB.NewUpdateBatch()

	logger.Debugf("Channel [%s]: Updating history database for blockNo [%v] with [%d] transactions",
		d.name, blockNo, len(block.Data.Data))

	// OPEN THE SAI FILE
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

	// OPEN THE GLOBAL INDEX FILE
	globalIndexFile, err := os.OpenFile("/var/GI-Storage/globalIndex.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}
	defer closeFile(globalIndexFile)

	// READ CURRENT GLOBAL INDEX INTO MAP
	lockFile(globalIndexFile)
	indexBytes, err := io.ReadAll(globalIndexFile)
	if err != nil {
		log.Println(err)
	}
	unlockFile(globalIndexFile)
	lastGI := make(map[string]uint64)
	json.Unmarshal(indexBytes, &lastGI)

	// OPEN localIndex FILE
	localFileName := "/var/LI-Storage/localIndex-" + strconv.FormatUint(blockNo, 10) + ".json"
	localIndexFile, err := os.OpenFile(localFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
	}
	defer closeFile(localIndexFile)

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
					prev, found := lastGI[kvWrite.Key]
					lastGI[kvWrite.Key] = blockNo
					// If the key doesn't exist, prev will be the blockNo
					if !found {
						prev = blockNo
					}

					dataKey := constructDataKey(ns, kvWrite.Key, blockNo, tranNo)

					// CREATE A NEW DataKey FOR SAI
					dkSAI := DataKeySai{
						kvWrite.Key,
						blockNo,
						tranNo,
					}

					// CONVERT THE DATAKEYSAI INSTANCE TO JSON
					jsonBytesSAI, err := json.Marshal(dkSAI)
					if err != nil {
						log.Println(err)
					}

					// CONVERT THE JSONBYTESSAI TO A STRING
					jsonStringSAI := string(jsonBytesSAI) + "\n"

					// WRITE SAI
					if _, err := file.WriteString(jsonStringSAI); err != nil {
						log.Println(err)
					}

					// Create a new DataKey instance and set its fields
					dk := DataKey{
						Key:        kvWrite.Key,
						PrevBlock:  prev,
						TranNumber: tranNo,
					}

					// Convert the DataKey instance to json
					jsonBytes, err := json.Marshal(dk)
					if err != nil {
						log.Println(err)
					}

					jsonString := string(jsonBytes) + "\n"

					lockFile(localIndexFile)
					_, err = localIndexFile.WriteString(jsonString)
					if err != nil {
						log.Println(err)
					}
					unlockFile(localIndexFile)

					// No value is required, write an empty byte array (emptyValue) since Put() of nil is not allowed
					dbBatch.Put(dataKey, emptyValue)
				}
			}

		} else {
			logger.Debugf("Skipping transaction [%d] since it is not an endorsement transaction\n", tranNo)
		}
		tranNo++
	}

	// LOCK GLOBAL INDEX FILE
	lockFile(globalIndexFile)
	indexBytes, err = io.ReadAll(globalIndexFile)
	if err != nil {
		log.Println(err)
	}

	currentGI := make(map[string]uint64)
	json.Unmarshal(indexBytes, &currentGI)

	for key, record := range lastGI {
		currentGI[key] = record
	}

	globalIndexFile.Seek(0, 0)
	globalIndexFile.Truncate(0)
	indexBytes, err = json.Marshal(currentGI)
	if err != nil {
		log.Println(err)
	}

	_, err = globalIndexFile.Write(indexBytes)
	if err != nil {
		log.Println(err)
	}
	unlockFile(globalIndexFile)

	// add savepoint for recovery purpose
	height := version.NewHeight(blockNo, tranNo)
	dbBatch.Put(savePointKey, height.ToBytes())

	// write the block's history records and savepoint to LevelDB
	// Setting snyc to true as a precaution, false may be an ok optimization after further testing.
	if err := d.levelDB.WriteBatch(dbBatch, true); err != nil {
		return err
	}

	logger.Debugf("Channel [%s]: Updates committed to history database for blockNo [%v]", d.name, blockNo)
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
