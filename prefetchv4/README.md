# Prefetch v4

Same as prefetch but for badger v4 instead of v2.

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
Running as unit: run-rc381262fe5994a91bdab0f71adf56c4d.scope; invocation ID: 70250a2bc16147b7a41da2634cc4869d
badger 2024/10/08 16:03:07 INFO: All 33 tables opened in 21ms
badger 2024/10/08 16:03:07 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:07 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:07 INFO: Deleting empty file: foo/000001.vlog
2024/10/08 16:03:09 0 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20607188 kB	->	MemFree:        17336840 kB	=	-3270348 kB
MemAvailable:   29567412 kB	->	MemAvailable:   27323972 kB	=	-2243440 kB
Buffers:            3440 kB	->	Buffers:            5808 kB	=	2368 kB
Cached:          7229728 kB	->	Cached:          8254076 kB	=	1024348 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25041748 kB	->	Active:         28055036 kB	=	3013288 kB
Inactive:        9635596 kB	->	Inactive:        9684588 kB	=	48992 kB
Active(anon):   22499332 kB	->	Active(anon):   24535000 kB	=	2035668 kB
Inactive(anon):  8898336 kB	->	Inactive(anon):  8898336 kB	=	0 kB
Active(file):    2542416 kB	->	Active(file):    3520036 kB	=	977620 kB
Inactive(file):   737260 kB	->	Inactive(file):   786252 kB	=	48992 kB
Unevictable:     2173624 kB	->	Unevictable:     2173624 kB	=	0 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             10568 kB	->	Dirty:             10384 kB	=	-184 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      29618020 kB	->	AnonPages:      31653852 kB	=	2035832 kB
Mapped:          3496488 kB	->	Mapped:          4220812 kB	=	724324 kB
Shmem:           4089020 kB	->	Shmem:           4089024 kB	=	4 kB
KReclaimable:    6412756 kB	->	KReclaimable:    6412712 kB	=	-44 kB
Slab:            7294700 kB	->	Slab:            7294844 kB	=	144 kB
SReclaimable:    6412756 kB	->	SReclaimable:    6412712 kB	=	-44 kB
SUnreclaim:       881944 kB	->	SUnreclaim:       882132 kB	=	188 kB
KernelStack:       71120 kB	->	KernelStack:       71616 kB	=	496 kB
PageTables:       225088 kB	->	PageTables:       231384 kB	=	6296 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60544980 kB	->	Committed_AS:   63709784 kB	=	3164804 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      220980 kB	->	VmallocUsed:      221204 kB	=	224 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
3.27user 3.62system 0:02.26elapsed 305%CPU (0avgtext+0avgdata 2821744maxresident)k
1919368inputs+928512outputs (3406major+702981minor)pagefaults 0swaps
259:0 rbytes=982728704 wbytes=0 rios=9142 wios=1 dbytes=0 dios=0
252:0 rbytes=982728704 wbytes=475381760 rios=4873 wios=496 dbytes=0 dios=0
252:1 rbytes=982728704 wbytes=475348992 rios=4873 wios=459 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=false
Running as unit: run-raa0700d2f802454aa780b136b65e87a6.scope; invocation ID: d389cfe1dac547a5838b3ed063b48c2c
badger 2024/10/08 16:03:12 INFO: All 23 tables opened in 16ms
badger 2024/10/08 16:03:12 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:12 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:12 INFO: Deleting empty file: foo/000002.vlog
2024/10/08 16:03:14 0 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20327036 kB	->	MemFree:        19466300 kB	=	-860736 kB
MemAvailable:   29297348 kB	->	MemAvailable:   29432016 kB	=	134668 kB
Buffers:            2512 kB	->	Buffers:            3148 kB	=	636 kB
Cached:          7241264 kB	->	Cached:          8236204 kB	=	994940 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25057228 kB	->	Active:         26118428 kB	=	1061200 kB
Inactive:        9636044 kB	->	Inactive:        9657340 kB	=	21296 kB
Active(anon):   22505064 kB	->	Active(anon):   22592156 kB	=	87092 kB
Inactive(anon):  8898336 kB	->	Inactive(anon):  8898336 kB	=	0 kB
Active(file):    2552164 kB	->	Active(file):    3526272 kB	=	974108 kB
Inactive(file):   737708 kB	->	Inactive(file):   759004 kB	=	21296 kB
Unevictable:     2174072 kB	->	Unevictable:     2174136 kB	=	64 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             11060 kB	->	Dirty:              8396 kB	=	-2664 kB
Writeback:             0 kB	->	Writeback:            84 kB	=	84 kB
AnonPages:      29623668 kB	->	AnonPages:      29710896 kB	=	87228 kB
Mapped:          3496712 kB	->	Mapped:          4086592 kB	=	589880 kB
Shmem:           4089472 kB	->	Shmem:           4089536 kB	=	64 kB
KReclaimable:    6412648 kB	->	KReclaimable:    6412648 kB	=	0 kB
Slab:            7294848 kB	->	Slab:            7294872 kB	=	24 kB
SReclaimable:    6412648 kB	->	SReclaimable:    6412648 kB	=	0 kB
SUnreclaim:       882200 kB	->	SUnreclaim:       882224 kB	=	24 kB
KernelStack:       71344 kB	->	KernelStack:       71488 kB	=	144 kB
PageTables:       225132 kB	->	PageTables:       227560 kB	=	2428 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60553268 kB	->	Committed_AS:   60780308 kB	=	227040 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      220964 kB	->	VmallocUsed:      221060 kB	=	96 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
0.23user 0.67system 0:02.28elapsed 39%CPU (0avgtext+0avgdata 684800maxresident)k
1910000inputs+48outputs (4577major+28910minor)pagefaults 0swaps
259:0 rbytes=977932288 wbytes=0 rios=9370 wios=0 dbytes=0 dios=0
252:0 rbytes=977932288 wbytes=12288 rios=4869 wios=3 dbytes=0 dios=0
252:1 rbytes=977932288 wbytes=0 rios=4869 wios=0 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=hybrid -hit=false
Running as unit: run-r8b38f0d796a54510a978b4a6bcb05261.scope; invocation ID: f9e03aa1d65c4711bd711bd9985f1b1a
badger 2024/10/08 16:03:17 INFO: All 23 tables opened in 16ms
badger 2024/10/08 16:03:17 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:17 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:17 INFO: Deleting empty file: foo/000003.vlog
2024/10/08 16:03:20 0 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20450492 kB	->	MemFree:        19509844 kB	=	-940648 kB
MemAvailable:   29411416 kB	->	MemAvailable:   29447380 kB	=	35964 kB
Buffers:            2120 kB	->	Buffers:            5480 kB	=	3360 kB
Cached:          7232904 kB	->	Cached:          8206044 kB	=	973140 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25062076 kB	->	Active:         26112872 kB	=	1050796 kB
Inactive:        9628380 kB	->	Inactive:        9645980 kB	=	17600 kB
Active(anon):   22510724 kB	->	Active(anon):   22602548 kB	=	91824 kB
Inactive(anon):  8898336 kB	->	Inactive(anon):  8898328 kB	=	-8 kB
Active(file):    2551352 kB	->	Active(file):    3510324 kB	=	958972 kB
Inactive(file):   730044 kB	->	Inactive(file):   747652 kB	=	17608 kB
Unevictable:     2173560 kB	->	Unevictable:     2173624 kB	=	64 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              9384 kB	->	Dirty:              9460 kB	=	76 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      29629452 kB	->	AnonPages:      29721256 kB	=	91804 kB
Mapped:          3496648 kB	->	Mapped:          4085280 kB	=	588632 kB
Shmem:           4088956 kB	->	Shmem:           4089024 kB	=	68 kB
KReclaimable:    6411736 kB	->	KReclaimable:    6411768 kB	=	32 kB
Slab:            7293660 kB	->	Slab:            7293864 kB	=	204 kB
SReclaimable:    6411736 kB	->	SReclaimable:    6411768 kB	=	32 kB
SUnreclaim:       881924 kB	->	SUnreclaim:       882096 kB	=	172 kB
KernelStack:       71376 kB	->	KernelStack:       71504 kB	=	128 kB
PageTables:       225184 kB	->	PageTables:       227376 kB	=	2192 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60556624 kB	->	Committed_AS:   60780612 kB	=	223988 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      220964 kB	->	VmallocUsed:      221156 kB	=	192 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
0.24user 0.65system 0:02.23elapsed 40%CPU (0avgtext+0avgdata 685568maxresident)k
1910256inputs+48outputs (4576major+28993minor)pagefaults 0swaps
259:0 rbytes=978063360 wbytes=0 rios=9373 wios=2 dbytes=0 dios=0
252:0 rbytes=978063360 wbytes=4096 rios=4872 wios=3 dbytes=0 dios=0
252:1 rbytes=978063360 wbytes=0 rios=4872 wios=2 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=false -hit=true
Running as unit: run-r1ce7e4207a0646b39771814d06d117f0.scope; invocation ID: 593df81350254301ae91d4f27e700984
badger 2024/10/08 16:03:23 INFO: All 23 tables opened in 17ms
badger 2024/10/08 16:03:23 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:23 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:23 INFO: Deleting empty file: foo/000004.vlog
2024/10/08 16:03:26 891000 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20385544 kB	->	MemFree:        18899808 kB	=	-1485736 kB
MemAvailable:   29371420 kB	->	MemAvailable:   28872384 kB	=	-499036 kB
Buffers:            1692 kB	->	Buffers:            2628 kB	=	936 kB
Cached:          7258548 kB	->	Cached:          8244324 kB	=	985776 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25073852 kB	->	Active:         26724932 kB	=	1651080 kB
Inactive:        9641376 kB	->	Inactive:        9648540 kB	=	7164 kB
Active(anon):   22510648 kB	->	Active(anon):   23182192 kB	=	671544 kB
Inactive(anon):  8898328 kB	->	Inactive(anon):  8898328 kB	=	0 kB
Active(file):    2563204 kB	->	Active(file):    3542740 kB	=	979536 kB
Inactive(file):   743048 kB	->	Inactive(file):   750212 kB	=	7164 kB
Unevictable:     2174136 kB	->	Unevictable:     2174136 kB	=	0 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             10316 kB	->	Dirty:             10196 kB	=	-120 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      29629316 kB	->	AnonPages:      30300944 kB	=	671628 kB
Mapped:          3496908 kB	->	Mapped:          4463228 kB	=	966320 kB
Shmem:           4089536 kB	->	Shmem:           4089536 kB	=	0 kB
KReclaimable:    6411832 kB	->	KReclaimable:    6411832 kB	=	0 kB
Slab:            7293996 kB	->	Slab:            7294296 kB	=	300 kB
SReclaimable:    6411832 kB	->	SReclaimable:    6411832 kB	=	0 kB
SUnreclaim:       882164 kB	->	SUnreclaim:       882464 kB	=	300 kB
KernelStack:       71280 kB	->	KernelStack:       71696 kB	=	416 kB
PageTables:       225132 kB	->	PageTables:       228892 kB	=	3760 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60556696 kB	->	Committed_AS:   61432544 kB	=	875848 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      221012 kB	->	VmallocUsed:      221316 kB	=	304 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
3.47user 1.20system 0:03.23elapsed 144%CPU (0avgtext+0avgdata 1651416maxresident)k
1954264inputs+48outputs (1076major+192704minor)pagefaults 0swaps
259:0 rbytes=1000595456 wbytes=0 rios=8551 wios=2 dbytes=0 dios=0
252:0 rbytes=1000595456 wbytes=4096 rios=4842 wios=3 dbytes=0 dios=0
252:1 rbytes=1000595456 wbytes=0 rios=4842 wios=2 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=true -hit=true
Running as unit: run-r08eba68bbcd84497b0c1ef680bef2657.scope; invocation ID: b8dd923d6aa440e7ac0c62bccebeae54
badger 2024/10/08 16:03:29 INFO: All 23 tables opened in 18ms
badger 2024/10/08 16:03:29 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:29 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:29 INFO: Deleting empty file: foo/000005.vlog
2024/10/08 16:03:34 891000 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20364824 kB	->	MemFree:        18849172 kB	=	-1515652 kB
MemAvailable:   29344988 kB	->	MemAvailable:   28889604 kB	=	-455384 kB
Buffers:            4884 kB	->	Buffers:            6364 kB	=	1480 kB
Cached:          7249172 kB	->	Cached:          8309048 kB	=	1059876 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25076068 kB	->	Active:         26738604 kB	=	1662536 kB
Inactive:        9634468 kB	->	Inactive:        9709140 kB	=	74672 kB
Active(anon):   22511888 kB	->	Active(anon):   23188836 kB	=	676948 kB
Inactive(anon):  8898328 kB	->	Inactive(anon):  8898328 kB	=	0 kB
Active(file):    2564180 kB	->	Active(file):    3549768 kB	=	985588 kB
Inactive(file):   736140 kB	->	Inactive(file):   810812 kB	=	74672 kB
Unevictable:     2174008 kB	->	Unevictable:     2174680 kB	=	672 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:             11672 kB	->	Dirty:              4596 kB	=	-7076 kB
Writeback:            56 kB	->	Writeback:             0 kB	=	-56 kB
AnonPages:      29631096 kB	->	AnonPages:      30307516 kB	=	676420 kB
Mapped:          3508012 kB	->	Mapped:          4463456 kB	=	955444 kB
Shmem:           4089372 kB	->	Shmem:           4090044 kB	=	672 kB
KReclaimable:    6412052 kB	->	KReclaimable:    6412060 kB	=	8 kB
Slab:            7294080 kB	->	Slab:            7294332 kB	=	252 kB
SReclaimable:    6412052 kB	->	SReclaimable:    6412060 kB	=	8 kB
SUnreclaim:       882028 kB	->	SUnreclaim:       882272 kB	=	244 kB
KernelStack:       71648 kB	->	KernelStack:       71568 kB	=	-80 kB
PageTables:       225256 kB	->	PageTables:       228612 kB	=	3356 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60563312 kB	->	Committed_AS:   61439064 kB	=	875752 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      221244 kB	->	VmallocUsed:      221172 kB	=	-72 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
9.11user 1.87system 0:04.78elapsed 229%CPU (0avgtext+0avgdata 1661132maxresident)k
1954440inputs+48outputs (1096major+194752minor)pagefaults 0swaps
259:0 rbytes=1000685568 wbytes=0 rios=8552 wios=2 dbytes=0 dios=0
252:0 rbytes=1000685568 wbytes=4096 rios=4843 wios=3 dbytes=0 dios=0
252:1 rbytes=1000685568 wbytes=8192 rios=4843 wios=4 dbytes=0 dios=0
===
=== ./prefetch run -dir foo -prefetch=hybrid -hit=true
Running as unit: run-r0364b19dc8244a95b8afa51a0314a381.scope; invocation ID: 8eb21e495547444697b48c4c282ebecf
badger 2024/10/08 16:03:37 INFO: All 23 tables opened in 18ms
badger 2024/10/08 16:03:37 INFO: Discard stats nextEmptySlot: 0
badger 2024/10/08 16:03:37 INFO: Set nextTxnTs to 9900
badger 2024/10/08 16:03:37 INFO: Deleting empty file: foo/000006.vlog
2024/10/08 16:03:42 891000 values read
MemTotal:       65504044 kB	->	MemTotal:       65504044 kB	=	0 kB
MemFree:        20436900 kB	->	MemFree:        18886788 kB	=	-1550112 kB
MemAvailable:   29397620 kB	->	MemAvailable:   28888188 kB	=	-509432 kB
Buffers:            4656 kB	->	Buffers:            7416 kB	=	2760 kB
Cached:          7229584 kB	->	Cached:          8267388 kB	=	1037804 kB
SwapCached:            0 kB	->	SwapCached:            0 kB	=	0 kB
Active:         25054920 kB	->	Active:         26738068 kB	=	1683148 kB
Inactive:        9633080 kB	->	Inactive:        9672940 kB	=	39860 kB
Active(anon):   22508760 kB	->	Active(anon):   23191112 kB	=	682352 kB
Inactive(anon):  8898328 kB	->	Inactive(anon):  8898328 kB	=	0 kB
Active(file):    2546160 kB	->	Active(file):    3546956 kB	=	1000796 kB
Inactive(file):   734752 kB	->	Inactive(file):   774612 kB	=	39860 kB
Unevictable:     2173656 kB	->	Unevictable:     2173592 kB	=	-64 kB
Mlocked:             480 kB	->	Mlocked:             480 kB	=	0 kB
SwapTotal:             0 kB	->	SwapTotal:             0 kB	=	0 kB
SwapFree:              0 kB	->	SwapFree:              0 kB	=	0 kB
Zswap:                 0 kB	->	Zswap:                 0 kB	=	0 kB
Zswapped:              0 kB	->	Zswapped:              0 kB	=	0 kB
Dirty:              6120 kB	->	Dirty:               660 kB	=	-5460 kB
Writeback:             0 kB	->	Writeback:             0 kB	=	0 kB
AnonPages:      29627624 kB	->	AnonPages:      30309856 kB	=	682232 kB
Mapped:          3496352 kB	->	Mapped:          4469572 kB	=	973220 kB
Shmem:           4089016 kB	->	Shmem:           4088956 kB	=	-60 kB
KReclaimable:    6412016 kB	->	KReclaimable:    6412040 kB	=	24 kB
Slab:            7293820 kB	->	Slab:            7294412 kB	=	592 kB
SReclaimable:    6412016 kB	->	SReclaimable:    6412040 kB	=	24 kB
SUnreclaim:       881804 kB	->	SUnreclaim:       882372 kB	=	568 kB
KernelStack:       71296 kB	->	KernelStack:       71584 kB	=	288 kB
PageTables:       225048 kB	->	PageTables:       228624 kB	=	3576 kB
SecPageTables:         0 kB	->	SecPageTables:         0 kB	=	0 kB
NFS_Unstable:          0 kB	->	NFS_Unstable:          0 kB	=	0 kB
Bounce:                0 kB	->	Bounce:                0 kB	=	0 kB
WritebackTmp:          0 kB	->	WritebackTmp:          0 kB	=	0 kB
CommitLimit:    32752020 kB	->	CommitLimit:    32752020 kB	=	0 kB
Committed_AS:   60567996 kB	->	Committed_AS:   61522548 kB	=	954552 kB
VmallocTotal:   34359738367 kB	->	VmallocTotal:   34359738367 kB	=	0 kB
VmallocUsed:      220996 kB	->	VmallocUsed:      221236 kB	=	240 kB
VmallocChunk:          0 kB	->	VmallocChunk:          0 kB	=	0 kB
Percpu:            30400 kB	->	Percpu:            30400 kB	=	0 kB
HardwareCorrupted:     0 kB	->	HardwareCorrupted:     0 kB	=	0 kB
AnonHugePages:         0 kB	->	AnonHugePages:         0 kB	=	0 kB
ShmemHugePages:  2027520 kB	->	ShmemHugePages:  2027520 kB	=	0 kB
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
DirectMap4k:      606196 kB	->	DirectMap4k:      606196 kB	=	0 kB
DirectMap2M:    18972672 kB	->	DirectMap2M:    18972672 kB	=	0 kB
DirectMap1G:    47185920 kB	->	DirectMap1G:    47185920 kB	=	0 kB
9.50user 1.96system 0:04.95elapsed 231%CPU (0avgtext+0avgdata 1658384maxresident)k
1954280inputs+48outputs (1093major+194520minor)pagefaults 0swaps
259:0 rbytes=1000603648 wbytes=0 rios=8553 wios=2 dbytes=0 dios=0
252:0 rbytes=1000603648 wbytes=4096 rios=4844 wios=3 dbytes=0 dios=0
252:1 rbytes=1000603648 wbytes=8192 rios=4844 wios=4 dbytes=0 dios=0
===
```
