package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"log"
)

const valueSize = 2000

type prefetchValue int

const (
	falsePrefetch prefetchValue = iota
	truePrefetch
	hybridPrefetch
)

func run(db *badger.DB, prefetch prefetchValue, hit bool) {
	var cnt int
	txn := db.NewTransaction(false)
	for k := 1000; k < 100000; k += 10 {
		first := k
		if !hit {
			first++ // so that key is 1001 instead of 1000, which doesn't exist
		}
		readcnt, err := read(txn, prefetch, first)
		if err != nil {
			log.Fatal(err)
		}
		cnt += readcnt
	}
	log.Printf("%d values read", cnt)
}

func read(txn *badger.Txn, prefetch prefetchValue, first int) (int, error) {
	opts := badger.DefaultIteratorOptions
	opts.Prefix = []byte(fmt.Sprintf("%d:", first))

	if prefetch == hybridPrefetch {
		opts.PrefetchValues = false
		iter := txn.NewIterator(opts)
		iter.Rewind()
		if !iter.Valid() {
			iter.Close()
			return 0, nil
		}
		iter.Close()
	}

	opts.PrefetchValues = prefetch == truePrefetch || prefetch == hybridPrefetch
	iter := txn.NewIterator(opts)
	defer iter.Close()
	var cnt int
	buf := make([]byte, valueSize)
	for iter.Rewind(); iter.Valid(); iter.Next() {
		item := iter.Item()
		if item.IsDeletedOrExpired() {
			continue
		}
		if err := item.Value(func(data []byte) error {
			copy(buf, data)
			cnt++
			return nil
		}); err != nil {
			return cnt, err
		}
	}
	return cnt, nil
}
