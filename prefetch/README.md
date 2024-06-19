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

or to run all, use
```shell
./run.sh
```

## Results

```
$ ./run.sh
+ ./prefetch run -dir foo -prefetch=false -hit=false
+ grep Mapped:
badger 2024/06/19 16:20:30 INFO: All 2 tables opened in 36ms
badger 2024/06/19 16:20:30 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:20:30 INFO: Replay took: 3.389µs
2024/06/19 16:20:30 0 values read
Mapped:          3388184 kB     ->      Mapped:          3424140 kB     =       35956 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB

real    0m0.110s
user    0m0.108s
sys     0m0.020s
+ grep Mapped:
+ ./prefetch run -dir foo -prefetch=true -hit=false
badger 2024/06/19 16:20:30 INFO: All 2 tables opened in 24ms
badger 2024/06/19 16:20:30 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:20:30 INFO: Replay took: 3.114µs
2024/06/19 16:20:32 0 values read
Mapped:          3388184 kB     ->      Mapped:          5185212 kB     =       1797028 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB

real    0m2.021s
user    0m4.420s
sys     0m0.890s
+ ./prefetch run -dir foo -prefetch=false -hit=true
+ grep Mapped:
badger 2024/06/19 16:20:32 INFO: All 2 tables opened in 18ms
badger 2024/06/19 16:20:32 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:20:32 INFO: Replay took: 3.054µs
2024/06/19 16:20:33 891000 values read
Mapped:          3390324 kB     ->      Mapped:          5188940 kB     =       1798616 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB

real    0m1.055s
user    0m0.742s
sys     0m0.330s
+ ./prefetch run -dir foo -prefetch=true -hit=true
+ grep Mapped:
badger 2024/06/19 16:20:33 INFO: All 2 tables opened in 21ms
badger 2024/06/19 16:20:33 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/19 16:20:33 INFO: Replay took: 3.321µs
2024/06/19 16:20:36 891000 values read
Mapped:          3388452 kB     ->      Mapped:          5190184 kB     =       1801732 kB
ShmemPmdMapped:        0 kB     ->      ShmemPmdMapped:        0 kB     =       0 kB
FilePmdMapped:         0 kB     ->      FilePmdMapped:         0 kB     =       0 kB

real    0m3.174s
user    0m6.845s
sys     0m1.159s
```
