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
=== ./prefetch run -dir baz -prefetch=false -hit=false
Running scope as unit: run-rf5e886f5558f44cd84a49c6563a261dc.scope
badger 2024/06/25 15:18:37 INFO: All 33 tables opened in 20ms
badger 2024/06/25 15:18:37 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:18:37 INFO: Set nextTxnTs to 9900
badger 2024/06/25 15:18:37 INFO: Deleting empty file: baz/000001.vlog
2024/06/25 15:18:39 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12722644 kB	->	MemFree:        11544208 kB	=	-1178436 kB
MemAvailable:   21442476 kB	->	MemAvailable:   21269064 kB	=	-173412 kB
Buffers:            3276 kB	->	Buffers:            3876 kB	=	600 kB
Cached:          6725020 kB	->	Cached:          7729780 kB	=	1004760 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24315836 kB	->	Active:         25489064 kB	=	1173228 kB
Inactive:       17335712 kB	->	Inactive:       17354044 kB	=	18332 kB
Active(anon):   22889440 kB	->	Active(anon):   23075980 kB	=	186540 kB
Inactive(anon): 16013020 kB	->	Inactive(anon): 16013016 kB	=	-4 kB
Active(file):    1426396 kB	->	Active(file):    2413084 kB	=	986688 kB
Inactive(file):  1322692 kB	->	Inactive(file):  1341028 kB	=	18336 kB
Unevictable:     2794420 kB	->	Unevictable:     2794328 kB	=	-92 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              4200 kB	->	Dirty:              4312 kB	=	112 kB
Writeback:           164 kB	->	Writeback:             4 kB	=	-160 kB
AnonPages:      37707164 kB	->	AnonPages:      37894544 kB	=	187380 kB
Mapped:          3103232 kB	->	Mapped:          3730496 kB	=	627264 kB
Shmem:           4122784 kB	->	Shmem:           4122684 kB	=	-100 kB
KReclaimable:    6702084 kB	->	KReclaimable:    6702084 kB	=	0 kB
Slab:            7499564 kB	->	Slab:            7499564 kB	=	0 kB
SReclaimable:    6702084 kB	->	SReclaimable:    6702084 kB	=	0 kB
SUnreclaim:       797480 kB	->	SUnreclaim:       797480 kB	=	0 kB
KernelStack:       90576 kB	->	KernelStack:       90816 kB	=	240 kB
PageTables:       285920 kB	->	PageTables:       288464 kB	=	2544 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   70912300 kB	->	Committed_AS:   71227628 kB	=	315328 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237660 kB	->	VmallocUsed:      237868 kB	=	208 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1304576 kB	->	ShmemHugePages:  1304576 kB	=	0 kB
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
0.39user 0.58system 0:02.30elapsed 42%CPU (0avgtext+0avgdata 765568maxresident)k
1908352inputs+0outputs (4670major+41716minor)pagefaults 0swaps
259:0 rbytes=977088512 wbytes=0 rios=4882 wios=0 dbytes=0 dios=0
252:0 rbytes=977088512 wbytes=0 rios=4882 wios=0 dbytes=0 dios=0
252:1 rbytes=977088512 wbytes=0 rios=4882 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir baz -prefetch=true -hit=false
Running scope as unit: run-r52bda7296820493ca4b94a22d3811da8.scope
badger 2024/06/25 15:18:42 INFO: All 33 tables opened in 18ms
badger 2024/06/25 15:18:42 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:18:42 INFO: Set nextTxnTs to 9900
2024/06/25 15:18:47 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12744156 kB	->	MemFree:        11101232 kB	=	-1642924 kB
MemAvailable:   21459496 kB	->	MemAvailable:   20845032 kB	=	-614464 kB
Buffers:            2656 kB	->	Buffers:            3420 kB	=	764 kB
Cached:          6728988 kB	->	Cached:          7779092 kB	=	1050104 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24322548 kB	->	Active:         25929044 kB	=	1606496 kB
Inactive:       17315184 kB	->	Inactive:       17337956 kB	=	22772 kB
Active(anon):   22879544 kB	->	Active(anon):   23480336 kB	=	600792 kB
Inactive(anon): 16013016 kB	->	Inactive(anon): 16013016 kB	=	0 kB
Active(file):    1443004 kB	->	Active(file):    2448708 kB	=	1005704 kB
Inactive(file):  1302168 kB	->	Inactive(file):  1324940 kB	=	22772 kB
Unevictable:     2794264 kB	->	Unevictable:     2815908 kB	=	21644 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              5152 kB	->	Dirty:              5240 kB	=	88 kB
Writeback:             0 kB	->	Writeback:           120 kB	=	120 kB
AnonPages:      37690580 kB	->	AnonPages:      38291704 kB	=	601124 kB
Mapped:          3105856 kB	->	Mapped:          4068056 kB	=	962200 kB
Shmem:           4129860 kB	->	Shmem:           4151448 kB	=	21588 kB
KReclaimable:    6701508 kB	->	KReclaimable:    6701492 kB	=	-16 kB
Slab:            7498964 kB	->	Slab:            7498916 kB	=	-48 kB
SReclaimable:    6701508 kB	->	SReclaimable:    6701492 kB	=	-16 kB
SUnreclaim:       797456 kB	->	SUnreclaim:       797424 kB	=	-32 kB
KernelStack:       90528 kB	->	KernelStack:       90944 kB	=	416 kB
PageTables:       285904 kB	->	PageTables:       289400 kB	=	3496 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   70928156 kB	->	Committed_AS:   72010568 kB	=	1082412 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237596 kB	->	VmallocUsed:      238124 kB	=	528 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1304576 kB	->	ShmemHugePages:  1320960 kB	=	16384 kB
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
9.14user 1.59system 0:04.67elapsed 229%CPU (0avgtext+0avgdata 1567500maxresident)k
1974376inputs+0outputs (5627major+168499minor)pagefaults 0swaps
259:0 rbytes=1010892800 wbytes=0 rios=4670 wios=0 dbytes=0 dios=0
252:0 rbytes=1010892800 wbytes=0 rios=4670 wios=0 dbytes=0 dios=0
252:1 rbytes=1010892800 wbytes=0 rios=4670 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir baz -prefetch=hybrid -hit=false
Running scope as unit: run-rd9d7c0095a394977b6910287fbd7f86b.scope
badger 2024/06/25 15:18:50 INFO: All 33 tables opened in 18ms
badger 2024/06/25 15:18:50 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:18:50 INFO: Set nextTxnTs to 9900
2024/06/25 15:18:52 0 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12753132 kB	->	MemFree:        11589264 kB	=	-1163868 kB
MemAvailable:   21460960 kB	->	MemAvailable:   21328392 kB	=	-132568 kB
Buffers:            2632 kB	->	Buffers:            6492 kB	=	3860 kB
Cached:          6732676 kB	->	Cached:          7758852 kB	=	1026176 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24293572 kB	->	Active:         25427224 kB	=	1133652 kB
Inactive:       17311572 kB	->	Inactive:       17356288 kB	=	44716 kB
Active(anon):   22854408 kB	->	Active(anon):   23001476 kB	=	147068 kB
Inactive(anon): 16013016 kB	->	Inactive(anon): 16013016 kB	=	0 kB
Active(file):    1439164 kB	->	Active(file):    2425748 kB	=	986584 kB
Inactive(file):  1298556 kB	->	Inactive(file):  1343272 kB	=	44716 kB
Unevictable:     2808856 kB	->	Unevictable:     2804656 kB	=	-4200 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              5620 kB	->	Dirty:              5236 kB	=	-384 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      37669012 kB	->	AnonPages:      37813916 kB	=	144904 kB
Mapped:          3103040 kB	->	Mapped:          3727596 kB	=	624556 kB
Shmem:           4141248 kB	->	Shmem:           4139524 kB	=	-1724 kB
KReclaimable:    6701448 kB	->	KReclaimable:    6701448 kB	=	0 kB
Slab:            7498816 kB	->	Slab:            7498824 kB	=	8 kB
SReclaimable:    6701448 kB	->	SReclaimable:    6701448 kB	=	0 kB
SUnreclaim:       797368 kB	->	SUnreclaim:       797376 kB	=	8 kB
KernelStack:       90560 kB	->	KernelStack:       90896 kB	=	336 kB
PageTables:       285832 kB	->	PageTables:       288368 kB	=	2536 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   70931708 kB	->	Committed_AS:   71173028 kB	=	241320 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237692 kB	->	VmallocUsed:      238028 kB	=	336 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1314816 kB	->	ShmemHugePages:  1312768 kB	=	-2048 kB
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
0.41user 0.56system 0:02.31elapsed 42%CPU (0avgtext+0avgdata 765440maxresident)k
1927704inputs+0outputs (4751major+41186minor)pagefaults 0swaps
259:0 rbytes=986996736 wbytes=0 rios=4873 wios=0 dbytes=0 dios=0
252:0 rbytes=986996736 wbytes=0 rios=4873 wios=0 dbytes=0 dios=0
252:1 rbytes=986996736 wbytes=0 rios=4873 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir baz -prefetch=false -hit=true
Running scope as unit: run-r98785ecb1e054be3992a2a0433ef83e8.scope
badger 2024/06/25 15:18:55 INFO: All 33 tables opened in 17ms
badger 2024/06/25 15:18:55 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:18:55 INFO: Set nextTxnTs to 9900
2024/06/25 15:18:58 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12723040 kB	->	MemFree:        11088784 kB	=	-1634256 kB
MemAvailable:   21450260 kB	->	MemAvailable:   20863008 kB	=	-587252 kB
Buffers:            2800 kB	->	Buffers:           14944 kB	=	12144 kB
Cached:          6748140 kB	->	Cached:          7772104 kB	=	1023964 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24311492 kB	->	Active:         25903904 kB	=	1592412 kB
Inactive:       17329968 kB	->	Inactive:       17378976 kB	=	49008 kB
Active(anon):   22871300 kB	->	Active(anon):   23465700 kB	=	594400 kB
Inactive(anon): 16013016 kB	->	Inactive(anon): 16013016 kB	=	0 kB
Active(file):    1440192 kB	->	Active(file):    2438204 kB	=	998012 kB
Inactive(file):  1316952 kB	->	Inactive(file):  1365960 kB	=	49008 kB
Unevictable:     2804568 kB	->	Unevictable:     2793624 kB	=	-10944 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              5780 kB	->	Dirty:              2872 kB	=	-2908 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      37685476 kB	->	AnonPages:      38279960 kB	=	594484 kB
Mapped:          3102548 kB	->	Mapped:          4081220 kB	=	978672 kB
Shmem:           4136960 kB	->	Shmem:           4126020 kB	=	-10940 kB
KReclaimable:    6701416 kB	->	KReclaimable:    6701400 kB	=	-16 kB
Slab:            7498800 kB	->	Slab:            7498760 kB	=	-40 kB
SReclaimable:    6701416 kB	->	SReclaimable:    6701400 kB	=	-16 kB
SUnreclaim:       797384 kB	->	SUnreclaim:       797360 kB	=	-24 kB
KernelStack:       90656 kB	->	KernelStack:       90880 kB	=	224 kB
PageTables:       285800 kB	->	PageTables:       289224 kB	=	3424 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   70922512 kB	->	Committed_AS:   71618216 kB	=	695704 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237708 kB	->	VmallocUsed:      237932 kB	=	224 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1312768 kB	->	ShmemHugePages:  1304576 kB	=	-8192 kB
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
3.70user 0.98system 0:03.05elapsed 153%CPU (0avgtext+0avgdata 1582008maxresident)k
1974568inputs+0outputs (5446major+170436minor)pagefaults 0swaps
259:0 rbytes=1010991104 wbytes=0 rios=4813 wios=0 dbytes=0 dios=0
252:0 rbytes=1010991104 wbytes=0 rios=4813 wios=0 dbytes=0 dios=0
252:1 rbytes=1010991104 wbytes=0 rios=4813 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir baz -prefetch=true -hit=true
Running scope as unit: run-r67d00d8689154e9fbf4dd55f04ffe089.scope
badger 2024/06/25 15:19:01 INFO: All 33 tables opened in 19ms
badger 2024/06/25 15:19:01 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:19:01 INFO: Set nextTxnTs to 9900
2024/06/25 15:19:08 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12711888 kB	->	MemFree:        11028792 kB	=	-1683096 kB
MemAvailable:   21451104 kB	->	MemAvailable:   20794808 kB	=	-656296 kB
Buffers:            5920 kB	->	Buffers:            8488 kB	=	2568 kB
Cached:          6746440 kB	->	Cached:          7828068 kB	=	1081628 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24324088 kB	->	Active:         25926888 kB	=	1602800 kB
Inactive:       17337452 kB	->	Inactive:       17362848 kB	=	25396 kB
Active(anon):   22879244 kB	->	Active(anon):   23480644 kB	=	601400 kB
Inactive(anon): 16013016 kB	->	Inactive(anon): 16013016 kB	=	0 kB
Active(file):    1444844 kB	->	Active(file):    2446244 kB	=	1001400 kB
Inactive(file):  1324436 kB	->	Inactive(file):  1349832 kB	=	25396 kB
Unevictable:     2793496 kB	->	Unevictable:     2851032 kB	=	57536 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              2900 kB	->	Dirty:               692 kB	=	-2208 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      37693148 kB	->	AnonPages:      38294520 kB	=	601372 kB
Mapped:          3111168 kB	->	Mapped:          4079248 kB	=	968080 kB
Shmem:           4125888 kB	->	Shmem:           4183428 kB	=	57540 kB
KReclaimable:    6701276 kB	->	KReclaimable:    6701280 kB	=	4 kB
Slab:            7498700 kB	->	Slab:            7498712 kB	=	12 kB
SReclaimable:    6701276 kB	->	SReclaimable:    6701280 kB	=	4 kB
SUnreclaim:       797424 kB	->	SUnreclaim:       797432 kB	=	8 kB
KernelStack:       90904 kB	->	KernelStack:       91120 kB	=	216 kB
PageTables:       286272 kB	->	PageTables:       289420 kB	=	3148 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   71000660 kB	->	Committed_AS:   71680236 kB	=	679576 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237964 kB	->	VmallocUsed:      238012 kB	=	48 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1304576 kB	->	ShmemHugePages:  1355776 kB	=	51200 kB
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
14.68user 2.02system 0:06.88elapsed 242%CPU (0avgtext+0avgdata 1589024maxresident)k
1974728inputs+0outputs (5517major+170283minor)pagefaults 0swaps
259:0 rbytes=1011073024 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
252:0 rbytes=1011073024 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
252:1 rbytes=1011073024 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir baz -prefetch=hybrid -hit=true
Running scope as unit: run-rcdbb71219c1f4a58b15658305535402c.scope
badger 2024/06/25 15:19:11 INFO: All 33 tables opened in 18ms
badger 2024/06/25 15:19:11 INFO: Discard stats nextEmptySlot: 0
badger 2024/06/25 15:19:11 INFO: Set nextTxnTs to 9900
2024/06/25 15:19:18 891000 values read
MemTotal:       65496312 kB	->	MemTotal:       65496312 kB	=	0 kB
MemFree:        12689464 kB	->	MemFree:        11067792 kB	=	-1621672 kB
MemAvailable:   21409484 kB	->	MemAvailable:   20820708 kB	=	-588776 kB
Buffers:            5540 kB	->	Buffers:            7808 kB	=	2268 kB
Cached:          6765780 kB	->	Cached:          7796352 kB	=	1030572 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         24335460 kB	->	Active:         25933580 kB	=	1598120 kB
Inactive:       17322828 kB	->	Inactive:       17349108 kB	=	26280 kB
Active(anon):   22895060 kB	->	Active(anon):   23486568 kB	=	591508 kB
Inactive(anon): 16013016 kB	->	Inactive(anon): 16013016 kB	=	0 kB
Active(file):    1440400 kB	->	Active(file):    2447012 kB	=	1006612 kB
Inactive(file):  1309812 kB	->	Inactive(file):  1336092 kB	=	26280 kB
Unevictable:     2831832 kB	->	Unevictable:     2831832 kB	=	0 kB
Mlocked:             376 kB	->	Mlocked:             376 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              1776 kB	->	Dirty:              1172 kB	=	-604 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      37708976 kB	->	AnonPages:      38301328 kB	=	592352 kB
Mapped:          3105632 kB	->	Mapped:          4092964 kB	=	987332 kB
Shmem:           4164252 kB	->	Shmem:           4164228 kB	=	-24 kB
KReclaimable:    6701148 kB	->	KReclaimable:    6701152 kB	=	4 kB
Slab:            7498556 kB	->	Slab:            7498504 kB	=	-52 kB
SReclaimable:    6701148 kB	->	SReclaimable:    6701152 kB	=	4 kB
SUnreclaim:       797408 kB	->	SUnreclaim:       797352 kB	=	-56 kB
KernelStack:       90672 kB	->	KernelStack:       91008 kB	=	336 kB
PageTables:       285972 kB	->	PageTables:       289376 kB	=	3404 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32748156 kB	->	CommitLimit:    32748156 kB	=	0 kB
Committed_AS:   71023780 kB	->	Committed_AS:   71676244 kB	=	652464 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      237788 kB	->	VmallocUsed:      238236 kB	=	448 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            32960 kB	->	Percpu:            32960 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:     79872 kB	->	AnonHugePages:     79872 kB	=	0 kB
ShmemHugePages:  1341440 kB	->	ShmemHugePages:  1341440 kB	=	0 kB
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
14.76user 2.05system 0:06.92elapsed 242%CPU (0avgtext+0avgdata 1590280maxresident)k
1974576inputs+0outputs (5509major+170799minor)pagefaults 0swaps
259:0 rbytes=1010995200 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
252:0 rbytes=1010995200 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
252:1 rbytes=1010995200 wbytes=0 rios=4647 wios=0 dbytes=0 dios=0
===
```
