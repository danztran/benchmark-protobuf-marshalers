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
BenchmarkCopy/JsonCopy-12                  22311             53796 ns/op           14730 B/op        184 allocs/op
BenchmarkCopy/ProtoCopy-12                 87434             13602 ns/op           13016 B/op        185 allocs/op
BenchmarkCopy/ProtoJsonCopy-12              9289            119195 ns/op           49611 B/op       1354 allocs/op
BenchmarkCopy/MsgpackCopy-12               37466             31896 ns/op           17515 B/op        175 allocs/op
BenchmarkCopy/JsoniterCopy-12              62808             19288 ns/op           14589 B/op        179 allocs/op
BenchmarkCopy/GoccyJsonCopy-12            102632             11791 ns/op           14315 B/op         85 allocs/op
BenchmarkCopy/SonicCopy-12                 81001             14581 ns/op           26456 B/op        172 allocs/op
BenchmarkCopy/ProtoCloneCopy-12           211635              5627 ns/op            7512 B/op         90 allocs/op
BenchmarkMarshal/Json-12                  157515              7426 ns/op            4242 B/op          6 allocs/op
BenchmarkMarshal/Proto-12                 288513              4066 ns/op            2816 B/op          9 allocs/op
BenchmarkMarshal/ProtoJson-12              26308             45548 ns/op           33749 B/op        635 allocs/op
BenchmarkMarshal/Jsoniter-12              174031              6978 ns/op            4347 B/op          6 allocs/op
BenchmarkMarshal/GoccyJson-12             366488              3172 ns/op            4102 B/op          1 allocs/op
BenchmarkMarshal/Msgpack-12               109818             10722 ns/op            8189 B/op          9 allocs/op
BenchmarkMarshal/Sonic-12                 360078              2975 ns/op            4181 B/op          2 allocs/op
BenchmarkUnmarshal/Json-12                 26130             45682 ns/op           10416 B/op        178 allocs/op
BenchmarkUnmarshal/Proto-12               130900              9152 ns/op           10200 B/op        176 allocs/op
BenchmarkUnmarshal/ProtoJson-12            16750             71788 ns/op           15816 B/op        719 allocs/op
BenchmarkUnmarshal/Jsoniter-12            105039             11418 ns/op           10150 B/op        173 allocs/op
BenchmarkUnmarshal/GoccyJson-12           152190              7755 ns/op           10096 B/op         84 allocs/op
BenchmarkUnmarshal/Msgpack-12              57506             20738 ns/op            9310 B/op        166 allocs/op
BenchmarkUnmarshal/Sonic-12               143558              8256 ns/op           11665 B/op         87 allocs/op
```
