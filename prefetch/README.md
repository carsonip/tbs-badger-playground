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
Running scope as unit: run-r26cb263fef864e9dab0abbccefb60717.scope
pid: 2982516
badger 2024/06/20 14:58:45 INFO: All 2 tables opened in 22ms
badger 2024/06/20 14:58:45 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 14:58:45 INFO: Replay took: 2.867µs
2024/06/20 14:58:45 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        22563788 kB	->	MemFree:        22496504 kB	=	-67284 kB
MemAvailable:   33329556 kB	->	MemAvailable:   33305716 kB	=	-23840 kB
Buffers:            2076 kB	->	Buffers:            2232 kB	=	156 kB
Cached:          6129672 kB	->	Cached:          6173740 kB	=	44068 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         17957636 kB	->	Active:         18037848 kB	=	80212 kB
Inactive:       13226296 kB	->	Inactive:       13229884 kB	=	3588 kB
Active(anon):   15812712 kB	->	Active(anon):   15853068 kB	=	40356 kB
Inactive(anon): 11982432 kB	->	Inactive(anon): 11982432 kB	=	0 kB
Active(file):    2144924 kB	->	Active(file):    2184780 kB	=	39856 kB
Inactive(file):  1243864 kB	->	Inactive(file):  1247452 kB	=	3588 kB
Unevictable:     2112620 kB	->	Unevictable:     2112620 kB	=	0 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              1284 kB	->	Dirty:              1492 kB	=	208 kB
Writeback:             4 kB	->	Writeback:            16 kB	=	12 kB
AnonPages:      27107220 kB	->	AnonPages:      27147588 kB	=	40368 kB
Mapped:          3675860 kB	->	Mapped:          3709884 kB	=	34024 kB
Shmem:           2884896 kB	->	Shmem:           2884904 kB	=	8 kB
KReclaimable:    8108320 kB	->	KReclaimable:    8108320 kB	=	0 kB
Slab:            8891688 kB	->	Slab:            8891688 kB	=	0 kB
SReclaimable:    8108320 kB	->	SReclaimable:    8108320 kB	=	0 kB
SUnreclaim:       783368 kB	->	SUnreclaim:       783368 kB	=	0 kB
KernelStack:       79376 kB	->	KernelStack:       79080 kB	=	-296 kB
PageTables:       213064 kB	->	PageTables:       212976 kB	=	-88 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   53882576 kB	->	Committed_AS:   53997604 kB	=	115028 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226432 kB	->	VmallocUsed:      226080 kB	=	-352 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            31424 kB	->	Percpu:            31424 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    186368 kB	->	AnonHugePages:    186368 kB	=	0 kB
ShmemHugePages:  1589248 kB	->	ShmemHugePages:  1589248 kB	=	0 kB
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
0.08user 0.04system 0:00.20elapsed 61%CPU (0avgtext+0avgdata 66304maxresident)k
86664inputs+16outputs (471major+7006minor)pagefaults 0swaps
259:0 rbytes=44384256 wbytes=0 rios=359 wios=0 dbytes=0 dios=0
252:0 rbytes=44384256 wbytes=4096 rios=359 wios=1 dbytes=0 dios=0
252:1 rbytes=44384256 wbytes=0 rios=359 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=false
Running scope as unit: run-rc41699af790b4b6b86cb0dccf3a75249.scope
pid: 2982652
badger 2024/06/20 14:58:48 INFO: All 2 tables opened in 24ms
badger 2024/06/20 14:58:48 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 14:58:48 INFO: Replay took: 3.044µs
2024/06/20 14:58:53 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        22643716 kB	->	MemFree:        20856952 kB	=	-1786764 kB
MemAvailable:   33372116 kB	->	MemAvailable:   33333492 kB	=	-38624 kB
Buffers:            2616 kB	->	Buffers:            6636 kB	=	4020 kB
Cached:          6090924 kB	->	Cached:          7835424 kB	=	1744500 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         17919700 kB	->	Active:         18039012 kB	=	119312 kB
Inactive:       13198780 kB	->	Inactive:       14876140 kB	=	1677360 kB
Active(anon):   15784560 kB	->	Active(anon):   15833080 kB	=	48520 kB
Inactive(anon): 11982432 kB	->	Inactive(anon): 11982432 kB	=	0 kB
Active(file):    2135140 kB	->	Active(file):    2205932 kB	=	70792 kB
Inactive(file):  1216348 kB	->	Inactive(file):  2893708 kB	=	1677360 kB
Unevictable:     2112620 kB	->	Unevictable:     2112620 kB	=	0 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              1220 kB	->	Dirty:              1296 kB	=	76 kB
Writeback:           220 kB	->	Writeback:             0 kB	=	-220 kB
AnonPages:      27079516 kB	->	AnonPages:      27127796 kB	=	48280 kB
Mapped:          3675344 kB	->	Mapped:          5368528 kB	=	1693184 kB
Shmem:           2884896 kB	->	Shmem:           2884900 kB	=	4 kB
KReclaimable:    8108252 kB	->	KReclaimable:    8108240 kB	=	-12 kB
Slab:            8891604 kB	->	Slab:            8891576 kB	=	-28 kB
SReclaimable:    8108252 kB	->	SReclaimable:    8108240 kB	=	-12 kB
SUnreclaim:       783352 kB	->	SUnreclaim:       783336 kB	=	-16 kB
KernelStack:       78976 kB	->	KernelStack:       79200 kB	=	224 kB
PageTables:       212660 kB	->	PageTables:       216672 kB	=	4012 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   53859812 kB	->	Committed_AS:   54394460 kB	=	534648 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      225936 kB	->	VmallocUsed:      226368 kB	=	432 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            31424 kB	->	Percpu:            31424 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    186368 kB	->	AnonHugePages:    186368 kB	=	0 kB
ShmemHugePages:  1589248 kB	->	ShmemHugePages:  1589248 kB	=	0 kB
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
5.47user 9.86system 0:05.63elapsed 272%CPU (0avgtext+0avgdata 1732096maxresident)k
3389424inputs+16outputs (611477major+36196minor)pagefaults 0swaps
259:0 rbytes=1735397376 wbytes=0 rios=413202 wios=0 dbytes=0 dios=0
252:0 rbytes=1735397376 wbytes=4096 rios=413202 wios=1 dbytes=0 dios=0
252:1 rbytes=1735397376 wbytes=0 rios=413202 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=false -hit=true
Running scope as unit: run-rb516682d99824e9db7137d24022b61e2.scope
pid: 2983334
badger 2024/06/20 14:58:57 INFO: All 2 tables opened in 22ms
badger 2024/06/20 14:58:57 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 14:58:57 INFO: Replay took: 3.424µs
2024/06/20 14:59:29 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        22608816 kB	->	MemFree:        20760004 kB	=	-1848812 kB
MemAvailable:   33369148 kB	->	MemAvailable:   33407760 kB	=	38612 kB
Buffers:            3932 kB	->	Buffers:           12756 kB	=	8824 kB
Cached:          6124032 kB	->	Cached:          8000240 kB	=	1876208 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         17943068 kB	->	Active:         17957100 kB	=	14032 kB
Inactive:       13213400 kB	->	Inactive:       15044372 kB	=	1830972 kB
Active(anon):   15790324 kB	->	Active(anon):   15747876 kB	=	-42448 kB
Inactive(anon): 11982432 kB	->	Inactive(anon): 11982432 kB	=	0 kB
Active(file):    2152744 kB	->	Active(file):    2209224 kB	=	56480 kB
Inactive(file):  1230968 kB	->	Inactive(file):  3061940 kB	=	1830972 kB
Unevictable:     2113132 kB	->	Unevictable:     2112492 kB	=	-640 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              2580 kB	->	Dirty:              1008 kB	=	-1572 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      27085320 kB	->	AnonPages:      27042868 kB	=	-42452 kB
Mapped:          3675600 kB	->	Mapped:          5476548 kB	=	1800948 kB
Shmem:           2885408 kB	->	Shmem:           2884708 kB	=	-700 kB
KReclaimable:    8107960 kB	->	KReclaimable:    8107932 kB	=	-28 kB
Slab:            8891304 kB	->	Slab:            8891236 kB	=	-68 kB
SReclaimable:    8107960 kB	->	SReclaimable:    8107932 kB	=	-28 kB
SUnreclaim:       783344 kB	->	SUnreclaim:       783304 kB	=	-40 kB
KernelStack:       78864 kB	->	KernelStack:       79104 kB	=	240 kB
PageTables:       212672 kB	->	PageTables:       216460 kB	=	3788 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   53862164 kB	->	Committed_AS:   53930144 kB	=	67980 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      225856 kB	->	VmallocUsed:      226048 kB	=	192 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            31424 kB	->	Percpu:            31424 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    186368 kB	->	AnonHugePages:    172032 kB	=	-14336 kB
ShmemHugePages:  1589248 kB	->	ShmemHugePages:  1589248 kB	=	0 kB
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
1.23user 3.97system 0:32.79elapsed 15%CPU (0avgtext+0avgdata 1831296maxresident)k
3615616inputs+16outputs (440122major+8660minor)pagefaults 0swaps
259:0 rbytes=1851207680 wbytes=0 rios=441480 wios=0 dbytes=0 dios=0
252:0 rbytes=1851207680 wbytes=4096 rios=441480 wios=1 dbytes=0 dios=0
252:1 rbytes=1851207680 wbytes=0 rios=441480 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=true
Running scope as unit: run-rd043cffbc1d042b9855b403258268791.scope
pid: 2985834
badger 2024/06/20 14:59:33 INFO: All 2 tables opened in 21ms
badger 2024/06/20 14:59:33 INFO: Replaying file id: 26 at offset: 60607284
badger 2024/06/20 14:59:33 INFO: Replay took: 2.171µs
2024/06/20 14:59:39 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        22680600 kB	->	MemFree:        20824204 kB	=	-1856396 kB
MemAvailable:   33426624 kB	->	MemAvailable:   33415144 kB	=	-11480 kB
Buffers:            2540 kB	->	Buffers:            7276 kB	=	4736 kB
Cached:          6109332 kB	->	Cached:          7949004 kB	=	1839672 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         17879276 kB	->	Active:         17962792 kB	=	83516 kB
Inactive:       13203804 kB	->	Inactive:       14990036 kB	=	1786232 kB
Active(anon):   15731032 kB	->	Active(anon):   15755896 kB	=	24864 kB
Inactive(anon): 11982432 kB	->	Inactive(anon): 11982432 kB	=	0 kB
Active(file):    2148244 kB	->	Active(file):    2206896 kB	=	58652 kB
Inactive(file):  1221372 kB	->	Inactive(file):  3007604 kB	=	1786232 kB
Unevictable:     2112492 kB	->	Unevictable:     2112492 kB	=	0 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              2180 kB	->	Dirty:              1048 kB	=	-1132 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      27025732 kB	->	AnonPages:      27050860 kB	=	25128 kB
Mapped:          3674752 kB	->	Mapped:          5473780 kB	=	1799028 kB
Shmem:           2884704 kB	->	Shmem:           2884708 kB	=	4 kB
KReclaimable:    8107748 kB	->	KReclaimable:    8107780 kB	=	32 kB
Slab:            8891148 kB	->	Slab:            8891212 kB	=	64 kB
SReclaimable:    8107748 kB	->	SReclaimable:    8107780 kB	=	32 kB
SUnreclaim:       783400 kB	->	SUnreclaim:       783432 kB	=	32 kB
KernelStack:       79328 kB	->	KernelStack:       79216 kB	=	-112 kB
PageTables:       213104 kB	->	PageTables:       216556 kB	=	3452 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   53940972 kB	->	Committed_AS:   54030920 kB	=	89948 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      226368 kB	->	VmallocUsed:      226112 kB	=	-256 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            31424 kB	->	Percpu:            31424 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:    172032 kB	->	AnonHugePages:    174080 kB	=	2048 kB
ShmemHugePages:  1589248 kB	->	ShmemHugePages:  1589248 kB	=	0 kB
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
8.30user 10.46system 0:06.58elapsed 284%CPU (0avgtext+0avgdata 1845376maxresident)k
3615520inputs+16outputs (658751major+38656minor)pagefaults 0swaps
259:0 rbytes=1851158528 wbytes=0 rios=441478 wios=0 dbytes=0 dios=0
252:0 rbytes=1851158528 wbytes=4096 rios=441478 wios=1 dbytes=0 dios=0
252:1 rbytes=1851158528 wbytes=0 rios=441478 wios=0 dbytes=0 dios=0
===
```
