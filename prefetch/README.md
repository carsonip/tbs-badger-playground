# Prefetch

This experiment shows the effect of iterator prefetch on resource usage.

Setup: creates a database with keys which consists of 2 parts. The first part, which mimics a trace ID in TBS, ranges from 10000 to 99990, at increments of 10, i.e. 10000, 10010, ..., 99980, 99990. The second part, which mimics a transaction ID / span ID, ranges from 1000 to 9900, at increments of 100, i.e. 1000, 1100, ..., 9800, 9900. The total number of combinations are 9000 * 90 = 810000 keys.

Run: searches for the KV in the database using a key prefix, i.e. the first part of the key, and it will either always hit when `-hit=true` or always miss when `-hit=false`. Observe the memory / disk IO footprint on different `-prefetch` settings.

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

Flags available are `-prefetch`, `-hit`, `-open-only`.

e.g. 
```shell
./prefetch run -dir foo -prefetch=false -hit=false
```

or to run all with `time` and `systemd` cgroup disk IO monitoring:
```shell
sudo ./run.sh
```

## Results

Command (running and stripping out ANSI escape codes):
```shell
sudo ./run.sh 2>&1 | sed -r "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2};?)?)?[mGK]//g" > run.out
```

Output:
```
=== ./prefetch run -dir foo -prefetch=false -hit=false
Running scope as unit: run-rb397772b79704fe3a0ca598af08d03a7.scope
badger 2024/06/20 22:47:26 INFO: All 2 tables opened in 26ms
badger 2024/06/20 22:47:26 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:47:26 INFO: Replay took: 8.873µs
2024/06/20 22:47:26 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        19966060 kB	->	MemFree:        19871840 kB	=	-94220 kB
MemAvailable:   31131728 kB	->	MemAvailable:   31103544 kB	=	-28184 kB
Buffers:          249464 kB	->	Buffers:          273448 kB	=	23984 kB
Cached:          6811036 kB	->	Cached:          6852620 kB	=	41584 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20119892 kB	->	Active:         20173692 kB	=	53800 kB
Inactive:       13121348 kB	->	Inactive:       13149448 kB	=	28100 kB
Active(anon):   17930676 kB	->	Active(anon):   17946540 kB	=	15864 kB
Inactive(anon): 11604516 kB	->	Inactive(anon): 11604516 kB	=	0 kB
Active(file):    2189216 kB	->	Active(file):    2227152 kB	=	37936 kB
Inactive(file):  1516832 kB	->	Inactive(file):  1544932 kB	=	28100 kB
Unevictable:     2575644 kB	->	Unevictable:     2575644 kB	=	0 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             35684 kB	->	Dirty:             36116 kB	=	432 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      28756192 kB	->	AnonPages:      28771864 kB	=	15672 kB
Mapped:          3696776 kB	->	Mapped:          3730296 kB	=	33520 kB
Shmem:           3499848 kB	->	Shmem:           3499844 kB	=	-4 kB
KReclaimable:    8190960 kB	->	KReclaimable:    8190960 kB	=	0 kB
Slab:            8982472 kB	->	Slab:            8982472 kB	=	0 kB
SReclaimable:    8190960 kB	->	SReclaimable:    8190960 kB	=	0 kB
SUnreclaim:       791512 kB	->	SUnreclaim:       791512 kB	=	0 kB
KernelStack:       80296 kB	->	KernelStack:       80584 kB	=	288 kB
PageTables:       218908 kB	->	PageTables:       219716 kB	=	808 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   55951708 kB	->	Committed_AS:   56037652 kB	=	85944 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226608 kB	->	VmallocUsed:      226736 kB	=	128 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    200704 kB	->	AnonHugePages:    200704 kB	=	0 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
0.10user 0.04system 0:00.22elapsed 66%CPU (0avgtext+0avgdata 61568maxresident)k
71472inputs+16outputs (1218major+6732minor)pagefaults 0swaps
259:0 rbytes=36605952 wbytes=0 rios=311 wios=0 dbytes=0 dios=0
252:0 rbytes=36605952 wbytes=4096 rios=311 wios=1 dbytes=0 dios=0
252:1 rbytes=36605952 wbytes=0 rios=311 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=false
Running scope as unit: run-r395ba7464fc9452c86ac176a17032b67.scope
badger 2024/06/20 22:47:30 INFO: All 2 tables opened in 23ms
badger 2024/06/20 22:47:30 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:47:30 INFO: Replay took: 8.476µs
2024/06/20 22:47:36 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        20014716 kB	->	MemFree:        17615672 kB	=	-2399044 kB
MemAvailable:   31145040 kB	->	MemAvailable:   31112992 kB	=	-32048 kB
Buffers:          236648 kB	->	Buffers:          873460 kB	=	636812 kB
Cached:          6787992 kB	->	Cached:          8516692 kB	=	1728700 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20126752 kB	->	Active:         20211172 kB	=	84420 kB
Inactive:       13071456 kB	->	Inactive:       15379480 kB	=	2308024 kB
Active(anon):   17922960 kB	->	Active(anon):   17948392 kB	=	25432 kB
Inactive(anon): 11604516 kB	->	Inactive(anon): 11604516 kB	=	0 kB
Active(file):    2203792 kB	->	Active(file):    2262780 kB	=	58988 kB
Inactive(file):  1466940 kB	->	Inactive(file):  3774964 kB	=	2308024 kB
Unevictable:     2575536 kB	->	Unevictable:     2574896 kB	=	-640 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             36096 kB	->	Dirty:             36472 kB	=	376 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      28748512 kB	->	AnonPages:      28773932 kB	=	25420 kB
Mapped:          3694248 kB	->	Mapped:          5381384 kB	=	1687136 kB
Shmem:           3499740 kB	->	Shmem:           3499112 kB	=	-628 kB
KReclaimable:    8190932 kB	->	KReclaimable:    8190916 kB	=	-16 kB
Slab:            8982348 kB	->	Slab:            8982292 kB	=	-56 kB
SReclaimable:    8190932 kB	->	SReclaimable:    8190916 kB	=	-16 kB
SUnreclaim:       791416 kB	->	SUnreclaim:       791376 kB	=	-40 kB
KernelStack:       80080 kB	->	KernelStack:       80384 kB	=	304 kB
PageTables:       218700 kB	->	PageTables:       223008 kB	=	4308 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   55946248 kB	->	Committed_AS:   56121624 kB	=	175376 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226496 kB	->	VmallocUsed:      226608 kB	=	112 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    200704 kB	->	AnonHugePages:    200704 kB	=	0 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
6.65user 9.85system 0:06.63elapsed 248%CPU (0avgtext+0avgdata 1730688maxresident)k
3374232inputs+16outputs (606659major+35655minor)pagefaults 0swaps
259:0 rbytes=1727619072 wbytes=0 rios=413154 wios=0 dbytes=0 dios=0
252:0 rbytes=1727619072 wbytes=4096 rios=413154 wios=1 dbytes=0 dios=0
252:1 rbytes=1727619072 wbytes=0 rios=413154 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=hybrid -hit=false
Running scope as unit: run-r86327bfe25bb40719ab303384a3ab8bf.scope
badger 2024/06/20 22:47:40 INFO: All 2 tables opened in 24ms
badger 2024/06/20 22:47:40 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:47:40 INFO: Replay took: 10.921µs
2024/06/20 22:47:40 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        19991188 kB	->	MemFree:        19919144 kB	=	-72044 kB
MemAvailable:   31137768 kB	->	MemAvailable:   31125188 kB	=	-12580 kB
Buffers:          266732 kB	->	Buffers:          285256 kB	=	18524 kB
Cached:          6774040 kB	->	Cached:          6814552 kB	=	40512 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20110156 kB	->	Active:         20168752 kB	=	58596 kB
Inactive:       13093564 kB	->	Inactive:       13115964 kB	=	22400 kB
Active(anon):   17912292 kB	->	Active(anon):   17933824 kB	=	21532 kB
Inactive(anon): 11604380 kB	->	Inactive(anon): 11604380 kB	=	0 kB
Active(file):    2197864 kB	->	Active(file):    2234928 kB	=	37064 kB
Inactive(file):  1489184 kB	->	Inactive(file):  1511584 kB	=	22400 kB
Unevictable:     2575004 kB	->	Unevictable:     2575004 kB	=	0 kB
Mlocked:             384 kB	->	Mlocked:             384 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             37132 kB	->	Dirty:             37144 kB	=	12 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      28738108 kB	->	AnonPages:      28759224 kB	=	21116 kB
Mapped:          3696208 kB	->	Mapped:          3729496 kB	=	33288 kB
Shmem:           3499208 kB	->	Shmem:           3499208 kB	=	0 kB
KReclaimable:    8190872 kB	->	KReclaimable:    8190872 kB	=	0 kB
Slab:            8982136 kB	->	Slab:            8982136 kB	=	0 kB
SReclaimable:    8190872 kB	->	SReclaimable:    8190872 kB	=	0 kB
SUnreclaim:       791264 kB	->	SUnreclaim:       791264 kB	=	0 kB
KernelStack:       79960 kB	->	KernelStack:       79960 kB	=	0 kB
PageTables:       218704 kB	->	PageTables:       218676 kB	=	-28 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   55948184 kB	->	Committed_AS:   56037588 kB	=	89404 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226512 kB	->	VmallocUsed:      226368 kB	=	-144 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    200704 kB	->	AnonHugePages:    200704 kB	=	0 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
0.08user 0.04system 0:00.19elapsed 68%CPU (0avgtext+0avgdata 61952maxresident)k
71456inputs+16outputs (1213major+6752minor)pagefaults 0swaps
259:0 rbytes=36597760 wbytes=0 rios=309 wios=0 dbytes=0 dios=0
252:0 rbytes=36597760 wbytes=4096 rios=309 wios=1 dbytes=0 dios=0
252:1 rbytes=36597760 wbytes=0 rios=309 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=false -hit=true
Running scope as unit: run-rdab3b246b5224fa89461f3e3f0cd0cb3.scope
badger 2024/06/20 22:47:43 INFO: All 2 tables opened in 27ms
badger 2024/06/20 22:47:43 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:47:43 INFO: Replay took: 12.4µs
2024/06/20 22:48:17 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        19987668 kB	->	MemFree:        17370256 kB	=	-2617412 kB
MemAvailable:   31128264 kB	->	MemAvailable:   31118324 kB	=	-9940 kB
Buffers:          242404 kB	->	Buffers:          961708 kB	=	719304 kB
Cached:          6793016 kB	->	Cached:          8678380 kB	=	1885364 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20146664 kB	->	Active:         20176340 kB	=	29676 kB
Inactive:       13073120 kB	->	Inactive:       15644016 kB	=	2570896 kB
Active(anon):   17934340 kB	->	Active(anon):   17927428 kB	=	-6912 kB
Inactive(anon): 11604380 kB	->	Inactive(anon): 11604380 kB	=	0 kB
Active(file):    2212324 kB	->	Active(file):    2248912 kB	=	36588 kB
Inactive(file):  1468740 kB	->	Inactive(file):  4039636 kB	=	2570896 kB
Unevictable:     2574876 kB	->	Unevictable:     2574512 kB	=	-364 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             36988 kB	->	Dirty:               352 kB	=	-36636 kB
Writeback:           276 kB	->	Writeback:             0 kB	=	-276 kB
AnonPages:      28759412 kB	->	AnonPages:      28752232 kB	=	-7180 kB
Mapped:          3702624 kB	->	Mapped:          5496552 kB	=	1793928 kB
Shmem:           3499080 kB	->	Shmem:           3498720 kB	=	-360 kB
KReclaimable:    8190872 kB	->	KReclaimable:    8190860 kB	=	-12 kB
Slab:            8982096 kB	->	Slab:            8982180 kB	=	84 kB
SReclaimable:    8190872 kB	->	SReclaimable:    8190860 kB	=	-12 kB
SUnreclaim:       791224 kB	->	SUnreclaim:       791320 kB	=	96 kB
KernelStack:       80328 kB	->	KernelStack:       80168 kB	=	-160 kB
PageTables:       218648 kB	->	PageTables:       222628 kB	=	3980 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   56314512 kB	->	Committed_AS:   56032008 kB	=	-282504 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226848 kB	->	VmallocUsed:      226528 kB	=	-320 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    206848 kB	->	AnonHugePages:    186368 kB	=	-20480 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
1.24user 4.15system 0:34.10elapsed 15%CPU (0avgtext+0avgdata 1829632maxresident)k
3600160inputs+16outputs (440841major+8513minor)pagefaults 0swaps
259:0 rbytes=1843294208 wbytes=0 rios=441398 wios=0 dbytes=0 dios=0
252:0 rbytes=1843294208 wbytes=4096 rios=441398 wios=1 dbytes=0 dios=0
252:1 rbytes=1843294208 wbytes=0 rios=441398 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=true
Running scope as unit: run-rd382748b9cc74c1bb51c4152796ae674.scope
badger 2024/06/20 22:48:21 INFO: All 2 tables opened in 19ms
badger 2024/06/20 22:48:21 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:48:21 INFO: Replay took: 1.959µs
2024/06/20 22:48:27 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        20256400 kB	->	MemFree:        18399288 kB	=	-1857112 kB
MemAvailable:   31156844 kB	->	MemAvailable:   31130932 kB	=	-25912 kB
Buffers:           13000 kB	->	Buffers:           19540 kB	=	6540 kB
Cached:          6779876 kB	->	Cached:          8604440 kB	=	1824564 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20099640 kB	->	Active:         20183704 kB	=	84064 kB
Inactive:       12846996 kB	->	Inactive:       14629212 kB	=	1782216 kB
Active(anon):   17901332 kB	->	Active(anon):   17936412 kB	=	35080 kB
Inactive(anon): 11604380 kB	->	Inactive(anon): 11604380 kB	=	0 kB
Active(file):    2198308 kB	->	Active(file):    2247292 kB	=	48984 kB
Inactive(file):  1242616 kB	->	Inactive(file):  3024832 kB	=	1782216 kB
Unevictable:     2574044 kB	->	Unevictable:     2574556 kB	=	512 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              1196 kB	->	Dirty:              2356 kB	=	1160 kB
Writeback:             0 kB	->	Writeback:             4 kB	=	4 kB
AnonPages:      28726076 kB	->	AnonPages:      28761048 kB	=	34972 kB
Mapped:          3698344 kB	->	Mapped:          5496756 kB	=	1798412 kB
Shmem:           3498244 kB	->	Shmem:           3498764 kB	=	520 kB
KReclaimable:    8190860 kB	->	KReclaimable:    8190860 kB	=	0 kB
Slab:            8982244 kB	->	Slab:            8982164 kB	=	-80 kB
SReclaimable:    8190860 kB	->	SReclaimable:    8190860 kB	=	0 kB
SUnreclaim:       791384 kB	->	SUnreclaim:       791304 kB	=	-80 kB
KernelStack:       80008 kB	->	KernelStack:       80144 kB	=	136 kB
PageTables:       219220 kB	->	PageTables:       222992 kB	=	3772 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   55964092 kB	->	Committed_AS:   56121876 kB	=	157784 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226288 kB	->	VmallocUsed:      226368 kB	=	80 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    186368 kB	->	AnonHugePages:    186368 kB	=	0 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
8.24user 10.55system 0:06.60elapsed 284%CPU (0avgtext+0avgdata 1844992maxresident)k
3615608inputs+16outputs (659248major+38629minor)pagefaults 0swaps
259:0 rbytes=1851203584 wbytes=0 rios=441473 wios=0 dbytes=0 dios=0
252:0 rbytes=1851203584 wbytes=4096 rios=441473 wios=1 dbytes=0 dios=0
252:1 rbytes=1851203584 wbytes=0 rios=441473 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=hybrid -hit=true
Running scope as unit: run-rfbd15cc32b034d8e8d7a2f4f1ace0a01.scope
badger 2024/06/20 22:48:31 INFO: All 2 tables opened in 19ms
badger 2024/06/20 22:48:31 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 22:48:31 INFO: Replay took: 2.377µs
2024/06/20 22:48:37 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        20301492 kB	->	MemFree:        18364940 kB	=	-1936552 kB
MemAvailable:   31176420 kB	->	MemAvailable:   31104996 kB	=	-71424 kB
Buffers:            1868 kB	->	Buffers:            7836 kB	=	5968 kB
Cached:          6764884 kB	->	Cached:          8624576 kB	=	1859692 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         20081488 kB	->	Active:         20209868 kB	=	128380 kB
Inactive:       12831428 kB	->	Inactive:       14636768 kB	=	1805340 kB
Active(anon):   17893096 kB	->	Active(anon):   17961692 kB	=	68596 kB
Inactive(anon): 11604380 kB	->	Inactive(anon): 11604380 kB	=	0 kB
Active(file):    2188392 kB	->	Active(file):    2248176 kB	=	59784 kB
Inactive(file):  1227048 kB	->	Inactive(file):  3032388 kB	=	1805340 kB
Unevictable:     2574492 kB	->	Unevictable:     2574556 kB	=	64 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              2688 kB	->	Dirty:              1104 kB	=	-1584 kB
Writeback:             0 kB	->	Writeback:            24 kB	=	24 kB
AnonPages:      28717848 kB	->	AnonPages:      28786812 kB	=	68964 kB
Mapped:          3697436 kB	->	Mapped:          5496888 kB	=	1799452 kB
Shmem:           3498692 kB	->	Shmem:           3498768 kB	=	76 kB
KReclaimable:    8190828 kB	->	KReclaimable:    8190832 kB	=	4 kB
Slab:            8982084 kB	->	Slab:            8982120 kB	=	36 kB
SReclaimable:    8190828 kB	->	SReclaimable:    8190832 kB	=	4 kB
SUnreclaim:       791256 kB	->	SUnreclaim:       791288 kB	=	32 kB
KernelStack:       79576 kB	->	KernelStack:       79520 kB	=	-56 kB
PageTables:       218592 kB	->	PageTables:       222684 kB	=	4092 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   55943300 kB	->	Committed_AS:   56119764 kB	=	176464 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      225952 kB	->	VmallocUsed:      226112 kB	=	160 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            29888 kB	->	Percpu:            29888 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    186368 kB	->	AnonHugePages:    190464 kB	=	4096 kB
ShmemHugePages:  2324480 kB	->	ShmemHugePages:  2324480 kB	=	0 kB
ShmemPmdMapped:        0 kB	->	ShmemPmdMapped:        0 kB	=	0 kB
FileHugePages:         0 kB	->	FileHugePages:         0 kB	=	0 kB
FilePmdMapped:         0 kB	->	FilePmdMapped:         0 kB	=	0 kB
Unaccepted:            0 kB	->	Unaccepted:            0 kB	=	0 kB
HugePages_Total:       0	->	HugePages_Total:       0	=	0 
HugePages_Free:        0	->	HugePages_Free:        0	=	0 
HugePages_Rsvd:        0	->	HugePages_Rsvd:        0	=	0 
HugePages_Surp:        0	->	HugePages_Surp:        0	=	0 
Hugepagesize:       2048 kB	->	Hugepagesize:       2048 kB	=	0 kB
Hugetlb:               0 kB	->	Hugetlb:               0 kB	=	0 kB
DirectMap4k:      684064 kB	->	DirectMap4k:      684064 kB	=	0 kB
DirectMap2M:    30429184 kB	->	DirectMap2M:    30429184 kB	=	0 kB
DirectMap1G:    35651584 kB	->	DirectMap1G:    35651584 kB	=	0 kB
8.44user 10.34system 0:06.72elapsed 279%CPU (0avgtext+0avgdata 1844736maxresident)k
3615608inputs+16outputs (658568major+38578minor)pagefaults 0swaps
259:0 rbytes=1851203584 wbytes=0 rios=441474 wios=0 dbytes=0 dios=0
252:0 rbytes=1851203584 wbytes=4096 rios=441474 wios=1 dbytes=0 dios=0
252:1 rbytes=1851203584 wbytes=0 rios=441474 wios=0 dbytes=0 dios=0
===
```
