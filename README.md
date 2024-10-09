# Benchmark Summary

This repository benchmarks various serialization libraries for marshaling and unmarshaling operations.

## TLDR

For overall performance across marshaling, unmarshaling, and copying operations, Sonic consistently ranks highest in terms of speed and efficiency.

## System Information

- **OS**: Linux
- **Architecture**: AMD64
- **CPU**: AMD Ryzen 5 5600X 6-Core Processor

## Benchmark Results

**Notice**: Every time the test is run, the data is generated randomly for each run but remains consistent across all iterations within that run.

```
goos: linux
goarch: amd64
pkg: github.com/danztran/benchmark-protobuf-marshalers
cpu: AMD Ryzen 5 5600X 6-Core Processor
BenchmarkCopy/Json-12                      46923             25735 ns/op          67.34 MB/s        6709 B/op        106 allocs/op
BenchmarkCopy/Proto-12                    137398              9050 ns/op         139.34 MB/s        6104 B/op        127 allocs/op
BenchmarkCopy/ProtoJson-12                 21988             54339 ns/op          33.05 MB/s       20823 B/op        588 allocs/op
BenchmarkCopy/Msgpack-12                   80422             15930 ns/op          98.43 MB/s        8190 B/op         86 allocs/op
BenchmarkCopy/Jsoniter-12                 118756              9693 ns/op         178.80 MB/s        6321 B/op         97 allocs/op
BenchmarkCopy/GoccyJson-12                182484              6509 ns/op         266.25 MB/s        6238 B/op         45 allocs/op
BenchmarkCopy/Sonic-12                    212044              5576 ns/op         310.79 MB/s        6949 B/op         39 allocs/op
BenchmarkCopy/ProtoClone-12               423745              2745 ns/op         459.42 MB/s        2904 B/op         48 allocs/op
BenchmarkMarshal/Json-12                   64881             18294 ns/op         263.92 MB/s       10184 B/op        142 allocs/op
BenchmarkMarshal/ProtoJson-12              24888             48155 ns/op         102.44 MB/s       32163 B/op        724 allocs/op
BenchmarkMarshal/Msgpack-12               149388              7850 ns/op         567.63 MB/s       16385 B/op         10 allocs/op
BenchmarkMarshal/Proto-12                  65554             18504 ns/op         243.41 MB/s        9344 B/op        281 allocs/op
BenchmarkMarshal/Jsoniter-12              141916              8357 ns/op         577.73 MB/s        5119 B/op          6 allocs/op
BenchmarkMarshal/GoccyJson-12             120000             10035 ns/op         481.14 MB/s        4883 B/op          1 allocs/op
BenchmarkMarshal/Sonic-12                 306511              3789 ns/op        1274.12 MB/s        4995 B/op          2 allocs/op
BenchmarkUnmarshal/Json-12                 50607             23591 ns/op          95.84 MB/s        8751 B/op        130 allocs/op
BenchmarkUnmarshal/Proto-12                98931             11887 ns/op         170.11 MB/s        9830 B/op        209 allocs/op
BenchmarkUnmarshal/ProtoJson-12            41936             28864 ns/op          80.13 MB/s       11886 B/op        357 allocs/op
BenchmarkUnmarshal/Jsoniter-12            122516              9592 ns/op         235.71 MB/s        8913 B/op        152 allocs/op
BenchmarkUnmarshal/GoccyJson-12           160792              7320 ns/op         308.87 MB/s        8779 B/op         77 allocs/op
BenchmarkUnmarshal/Msgpack-12             157371              7410 ns/op         281.23 MB/s        5708 B/op         90 allocs/op
BenchmarkUnmarshal/Sonic-12               216990              5459 ns/op         414.15 MB/s        8147 B/op         23 allocs/op
```
