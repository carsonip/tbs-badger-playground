package main

import (
	"flag"
	"log"

	"github.com/dgraph-io/badger/v2"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Please specify a subcommand.")
	}
	cmd, args := args[0], args[1:]

	switch cmd {
	case "setup":
		// Create a ~2GB database
		flag := flag.NewFlagSet("run", flag.ExitOnError)
		var (
			dir = flag.String("dir", ".", "badger storage dir")
		)
		flag.Parse(args)
		args = flag.Args()
		setup(*dir)
	case "run":
		flag := flag.NewFlagSet("run", flag.ExitOnError)
		var (
			prefetch = flag.String("prefetch", "false", "prefetch behavior; possible values: false, true, hybrid")
			hit      = flag.Bool("hit", false, "search for key that exists")
			dir      = flag.String("dir", ".", "badger storage dir")
			openOnly = flag.Bool("open-only", false, "open db without reading; ignores -prefetch and -hit")
		)
		flag.Parse(args)
		args = flag.Args()
		var db *badger.DB
		var err error
		memdiff(func() {
			db, err = openBadger(*dir, 0)
			if err != nil {
				log.Fatal(err)
			}
			if *openOnly {
				return
			}
			var prefetchVal prefetchValue
			switch *prefetch {
			case "false":
				prefetchVal = falsePrefetch
			case "true":
				prefetchVal = truePrefetch
			case "hybrid":
				prefetchVal = hybridPrefetch
			}
			run(db, prefetchVal, *hit)
		})
	default:
		log.Fatalf("Unrecognized command %q", cmd)
	}
}
