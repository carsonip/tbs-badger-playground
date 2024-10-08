package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/dgraph-io/badger/v4"
)

const (
	defaultValueLogFileSize = 64 * 1024 * 1024
)

func openBadger(storageDir string, valueLogFileSize int64) (*badger.DB, error) {
	// Tunable memory options:
	//  - NumMemtables - default 5 in-mem tables (MaxTableSize default)
	//  - NumLevelZeroTables - default 5 - number of L0 tables before compaction starts.
	//  - NumLevelZeroTablesStall - number of L0 tables before writing stalls (waiting for compaction).
	//  - IndexCacheSize - default all in mem, Each table has its own bloom filter and each bloom filter is approximately of 5 MB.
	//  - BaseTableSize - Default 2MB
	if valueLogFileSize <= 0 {
		valueLogFileSize = defaultValueLogFileSize
	}
	const tableLimit = 4
	badgerOpts := badger.DefaultOptions(storageDir).
		WithNumMemtables(tableLimit).                // in-memory tables.
		WithNumLevelZeroTables(tableLimit).          // L0 tables.
		WithNumLevelZeroTablesStall(tableLimit * 3). // Maintain the default 1-to-3 ratio before stalling.
		WithBaseTableSize(int64(16 << 20)).
		WithValueLogFileSize(valueLogFileSize) // vlog file size.

	return badger.Open(badgerOpts)
}

func populate(db *badger.DB) error {
	// 1000:1000, 1000:1100, ..., 1000:9900, 1010:1000, ..., 99990:9900
	txn := db.NewTransaction(true)
	buf := make([]byte, valueSize)
	for i := 1000; i < 100000; i += 10 {
		for j := 1000; j < 10000; j += 100 {
			rand.Read(buf)
			entry := badger.NewEntry([]byte(fmt.Sprintf("%d:%d", i, j)), buf)
			if err := txn.SetEntry(entry); err != nil {
				return err
			}
		}
		if err := txn.Commit(); err != nil {
			return err
		}
		txn = db.NewTransaction(true)
	}
	return nil
}

func setup(dir string) {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	db, err := openBadger(dir, 0)
	if err != nil {
		log.Fatal(err)
	}
	if err := populate(db); err != nil {
		log.Fatal(err)
	}

	// close and reopen to report size
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	db, err = openBadger(dir, 0)
	if err != nil {
		log.Fatal(err)
	}
	lsm, vlog := db.Size()
	fmt.Printf("lsm size: %d vlog size: %d\n", lsm, vlog)
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}
