# Benchmark Summary

This repository benchmarks various serialization libraries for marshaling and unmarshaling operations. The benchmarks were run on an Apple M1 machine.

## TLDR

For overall performance across marshaling, unmarshaling, and copying operations, GoccyJson consistently ranks highest in terms of speed and efficiency. It has the best performance in marshaling and unmarshaling, and second-best in copying operations.

## System Information

- **OS**: Darwin
- **Architecture**: ARM64
- **CPU**: Apple M1

## Benchmark Results

**Notice**: Every time the test is run, the data is generated randomly for each run but remains consistent across all iterations within that run.

### Marshaling

| Library   | ops (Rank)  | ns/op (Rank) | B/op (Rank) | allocs/op (Rank) |
| --------- | ----------- | ------------ | ----------- | ---------------- |
| Json      | 1304946 (4) | 1091 (4)     | 464 (4)     | 4 (3)            |
| Proto     | 2152027 (2) | 539.1 (2)    | 304 (1)     | 5 (4)            |
| ProtoJson | 358879 (6)  | 3341 (6)     | 2539 (6)    | 51 (6)           |
| Jsoniter  | 1405380 (3) | 845.7 (3)    | 632 (5)     | 6 (5)            |
| GoccyJson | 3021697 (1) | 396.2 (1)    | 384 (2)     | 1 (1)            |
| Msgpack   | 1358685 (5) | 892.8 (5)    | 1016 (3)    | 6 (5)            |

### Unmarshaling

| Library   | ops (Rank) | ns/op (Rank) | B/op (Rank) | allocs/op (Rank) |
| --------- | ---------- | ------------ | ----------- | ---------------- |
| Json      | 154759 (5) | 7764 (5)     | 2256 (5)    | 52 (4)           |
| Proto     | 492438 (2) | 2468 (2)     | 2232 (4)    | 62 (5)           |
| ProtoJson | 119209 (6) | 10170 (6)    | 3408 (6)    | 146 (6)          |
| Jsoniter  | 470722 (3) | 2501 (3)     | 2048 (2)    | 51 (3)           |
| GoccyJson | 657730 (1) | 1779 (1)     | 2146 (3)    | 25 (1)           |
| Msgpack   | 399151 (4) | 3019 (4)     | 1976 (1)    | 40 (2)           |

### Copying

| Library    | ops (Rank) | ns/op (Rank) | B/op (Rank) | allocs/op (Rank) |
| ---------- | ---------- | ------------ | ----------- | ---------------- |
| Json       | 89764 (6)  | 11617 (6)    | 4810 (6)    | 80 (5)           |
| Proto      | 211557 (4) | 5601 (4)     | 4569 (5)    | 116 (6)          |
| ProtoJson  | 56898 (7)  | 19863 (7)    | 11029 (7)   | 295 (7)          |
| Jsoniter   | 235582 (3) | 4970 (3)     | 4226 (3)    | 68 (4)           |
| GoccyJson  | 317118 (2) | 3703 (2)     | 4163 (2)    | 32 (1)           |
| Msgpack    | 242764 (5) | 5758 (5)     | 4373 (4)    | 51 (3)           |
| ProtoClone | 651907 (1) | 1980 (1)     | 1913 (1)    | 33 (2)           |
