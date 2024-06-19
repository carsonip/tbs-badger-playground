# Prefetch

This experiment shows the effect of iterator prefetch on memory usage.

## Run

### Compile

```shell
go build .
```

### Setup

Creates an example database of size ~2GB.

```shell
./prefetch setup -dir foo
```

### Run

Observe memory usage with various flags `-prefetch`, `-hit`, `-open-only`.

e.g. 
```shell
./prefetch run -dir foo -prefetch=false -hit=false
```

## Results

```
$ ./prefetch run -dir foo -prefetch=true -hit=false | grep Mapped:
badger 2024/06/19 16:11:20 INFO: All 2 tables opened in 15ms
badger 2024/06/19 16:11:20 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:11:20 INFO: Replay took: 1.881µs
2024/06/19 16:11:22 0 values read
Mapped:          3385780 kB     ->      Mapped:          5182740 kB     =       1796960 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB
```
```
$ ./prefetch run -dir foo -prefetch=false -hit=false | grep Mapped:
badger 2024/06/19 16:11:34 INFO: All 2 tables opened in 15ms
badger 2024/06/19 16:11:34 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:11:34 INFO: Replay took: 2.842µs
2024/06/19 16:11:34 0 values read
Mapped:          3386096 kB     ->      Mapped:          3421428 kB     =       35332 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB
```
```
$ ./prefetch run -dir foo -prefetch=false -hit=true | grep Mapped:
badger 2024/06/19 16:12:18 INFO: All 2 tables opened in 15ms
badger 2024/06/19 16:12:18 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:12:18 INFO: Replay took: 1.845µs
2024/06/19 16:12:19 891000 values read
Mapped:          3389064 kB     ->      Mapped:          5190016 kB     =       1800952 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB
```
```
$ ./prefetch run -dir foo -prefetch=true -hit=true | grep Mapped:
badger 2024/06/19 16:12:33 INFO: All 2 tables opened in 23ms
badger 2024/06/19 16:12:33 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:12:33 INFO: Replay took: 2.967µs
2024/06/19 16:12:36 891000 values read
Mapped:          3389196 kB     ->      Mapped:          5187340 kB     =       1798144 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB
```
