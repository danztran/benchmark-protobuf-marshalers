# Benchmark Summary

This repository benchmarks various serialization libraries for marshaling and unmarshaling operations.

## TLDR

For overall performance across marshaling, unmarshaling, and copying operations, Sonic consistently ranks highest in terms of speed and efficiency. However, if HTML escaping is enabled, Sonic's performance can become significantly slower. In such cases, it is recommended to use Goccy as an alternative for better performance.

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
BenchmarkMarshal/Large/JsonBuiltin-12             269888              4501 ns/op        2050.08 MB/s        9991 B/op         16 allocs/op
BenchmarkMarshal/Large/ProtoJson-12                37171             32594 ns/op         286.49 MB/s       61640 B/op        949 allocs/op
BenchmarkMarshal/Large/Proto-12                   451491              2428 ns/op        2189.11 MB/s        5824 B/op         29 allocs/op
BenchmarkMarshal/Large/Msgpack-12                 153603              8439 ns/op        1027.51 MB/s       32768 B/op         11 allocs/op
BenchmarkMarshal/Large/Jsoniter-12                272140              4350 ns/op        2121.16 MB/s        9729 B/op          6 allocs/op
BenchmarkMarshal/Large/GoccyDefault-12            382527              3253 ns/op        2836.57 MB/s        9502 B/op          1 allocs/op
BenchmarkMarshal/Large/GoccyUnordered-12          408310              3133 ns/op        2945.42 MB/s        9488 B/op          1 allocs/op
BenchmarkMarshal/Large/GoccyFastest-12            376828              3083 ns/op        2992.94 MB/s        9487 B/op          1 allocs/op
BenchmarkMarshal/Large/SonicStd-12                251235              4611 ns/op        2001.12 MB/s       26057 B/op          3 allocs/op
BenchmarkMarshal/Large/SonicDefault-12            515186              2360 ns/op        3910.53 MB/s       10025 B/op          2 allocs/op
BenchmarkMarshal/Large/SonicCompact-12            525940              2381 ns/op        3875.32 MB/s       10060 B/op          2 allocs/op
BenchmarkMarshal/Large/SonicFast-12               471136              2363 ns/op        3905.21 MB/s       10050 B/op          2 allocs/op
BenchmarkMarshal/Medium/JsonBuiltin-12           3981360               304.0 ns/op      2108.66 MB/s         704 B/op          1 allocs/op
BenchmarkMarshal/Medium/ProtoJson-12              608092              1930 ns/op         338.38 MB/s        3789 B/op         69 allocs/op
BenchmarkMarshal/Medium/Proto-12                 7680972               158.5 ns/op      2271.34 MB/s         384 B/op          1 allocs/op
BenchmarkMarshal/Medium/Msgpack-12               2231942               521.9 ns/op      1168.90 MB/s        2032 B/op          6 allocs/op
BenchmarkMarshal/Medium/Jsoniter-12              3591643               335.3 ns/op      1911.49 MB/s         712 B/op          2 allocs/op
BenchmarkMarshal/Medium/GoccyDefault-12          7142762               166.0 ns/op      3862.00 MB/s         704 B/op          1 allocs/op
BenchmarkMarshal/Medium/GoccyUnordered-12        6999511               167.9 ns/op      3816.82 MB/s         704 B/op          1 allocs/op
BenchmarkMarshal/Medium/GoccyFastest-12          7187436               160.7 ns/op      3987.70 MB/s         704 B/op          1 allocs/op
BenchmarkMarshal/Medium/SonicStd-12              3587187               345.9 ns/op      1853.26 MB/s        2059 B/op          3 allocs/op
BenchmarkMarshal/Medium/SonicDefault-12          7262254               169.7 ns/op      3776.44 MB/s         785 B/op          2 allocs/op
BenchmarkMarshal/Medium/SonicCompact-12          6978620               168.7 ns/op      3799.79 MB/s         786 B/op          2 allocs/op
BenchmarkMarshal/Medium/SonicFast-12             7323219               167.5 ns/op      3826.20 MB/s         784 B/op          2 allocs/op
BenchmarkMarshal/Small/JsonBuiltin-12           12129409                94.86 ns/op     2076.64 MB/s         208 B/op          1 allocs/op
BenchmarkMarshal/Small/ProtoJson-12              2899084               417.4 ns/op       469.61 MB/s         920 B/op         19 allocs/op
BenchmarkMarshal/Small/Proto-12                 19097130                62.32 ns/op     2198.43 MB/s         144 B/op          1 allocs/op
BenchmarkMarshal/Small/Msgpack-12                8981721               132.4 ns/op      1329.63 MB/s         496 B/op          4 allocs/op
BenchmarkMarshal/Small/Jsoniter-12              11362836               103.6 ns/op      1901.01 MB/s         216 B/op          2 allocs/op
BenchmarkMarshal/Small/GoccyDefault-12          21288254                54.24 ns/op     3631.76 MB/s         208 B/op          1 allocs/op
BenchmarkMarshal/Small/GoccyUnordered-12        22536423                53.02 ns/op     3715.38 MB/s         208 B/op          1 allocs/op
BenchmarkMarshal/Small/GoccyFastest-12          21980794                54.55 ns/op     3611.19 MB/s         208 B/op          1 allocs/op
BenchmarkMarshal/Small/SonicStd-12               9446162               117.5 ns/op      1675.97 MB/s         648 B/op          3 allocs/op
BenchmarkMarshal/Small/SonicDefault-12          21005424                58.43 ns/op     3371.79 MB/s         242 B/op          2 allocs/op
BenchmarkMarshal/Small/SonicCompact-12          20294385                57.33 ns/op     3436.32 MB/s         242 B/op          2 allocs/op
BenchmarkMarshal/Small/SonicFast-12             21316912                58.15 ns/op     3387.68 MB/s         242 B/op          2 allocs/op
BenchmarkUnmarshal/Large/JsonBuiltin-12            62550             19061 ns/op         400.14 MB/s       14704 B/op        301 allocs/op
BenchmarkUnmarshal/Large/ProtoJson-12              38061             31739 ns/op         242.04 MB/s       26313 B/op       1355 allocs/op
BenchmarkUnmarshal/Large/Proto-12                 250274              5051 ns/op         839.21 MB/s       14360 B/op        297 allocs/op
BenchmarkUnmarshal/Large/Msgpack-12               104936             10859 ns/op         649.34 MB/s       15188 B/op        318 allocs/op
BenchmarkUnmarshal/Large/Jsoniter-12              175658              6889 ns/op        1107.09 MB/s       14484 B/op        311 allocs/op
BenchmarkUnmarshal/Large/GoccyDefault-12          194214              6063 ns/op        1257.85 MB/s       17059 B/op        116 allocs/op
BenchmarkUnmarshal/Large/GoccyUnordered-12        150088              7129 ns/op        1069.84 MB/s       17055 B/op        116 allocs/op
BenchmarkUnmarshal/Large/GoccyFastest-12          174104              6855 ns/op        1112.66 MB/s       17054 B/op        116 allocs/op
BenchmarkUnmarshal/Large/SonicStd-12              175226              6211 ns/op        1227.94 MB/s       23147 B/op        277 allocs/op
BenchmarkUnmarshal/Large/SonicDefault-12          189003              5440 ns/op        1402.03 MB/s       18766 B/op        140 allocs/op
BenchmarkUnmarshal/Large/SonicCompact-12          219082              5122 ns/op        1488.96 MB/s       18705 B/op        140 allocs/op
BenchmarkUnmarshal/Large/SonicFast-12             247057              5336 ns/op        1429.48 MB/s       18690 B/op        140 allocs/op
BenchmarkUnmarshal/Medium/JsonBuiltin-12         1254496               848.2 ns/op       391.43 MB/s         816 B/op         18 allocs/op
BenchmarkUnmarshal/Medium/ProtoJson-12            967340              1376 ns/op         245.60 MB/s        1192 B/op         60 allocs/op
BenchmarkUnmarshal/Medium/Proto-12               6328732               177.7 ns/op      1012.87 MB/s         520 B/op         11 allocs/op
BenchmarkUnmarshal/Medium/Msgpack-12             2838138               424.7 ns/op       744.11 MB/s         648 B/op         15 allocs/op
BenchmarkUnmarshal/Medium/Jsoniter-12            5062372               239.1 ns/op      1388.46 MB/s         520 B/op         11 allocs/op
BenchmarkUnmarshal/Medium/GoccyDefault-12        5511987               222.4 ns/op      1492.72 MB/s         672 B/op          5 allocs/op
BenchmarkUnmarshal/Medium/GoccyUnordered-12      5347778               216.7 ns/op      1532.12 MB/s         672 B/op          5 allocs/op
BenchmarkUnmarshal/Medium/GoccyFastest-12        5306853               221.2 ns/op      1500.69 MB/s         672 B/op          5 allocs/op
BenchmarkUnmarshal/Medium/SonicStd-12            4842526               249.6 ns/op      1330.30 MB/s         940 B/op         12 allocs/op
BenchmarkUnmarshal/Medium/SonicDefault-12        5451633               202.5 ns/op      1639.30 MB/s         750 B/op          6 allocs/op
BenchmarkUnmarshal/Medium/SonicCompact-12        5587334               206.5 ns/op      1607.42 MB/s         749 B/op          6 allocs/op
BenchmarkUnmarshal/Medium/SonicFast-12           5693005               198.7 ns/op      1670.76 MB/s         751 B/op          6 allocs/op
BenchmarkUnmarshal/Small/JsonBuiltin-12          2934370               418.7 ns/op       470.48 MB/s         504 B/op         10 allocs/op
BenchmarkUnmarshal/Small/ProtoJson-12            2181120               538.0 ns/op       364.33 MB/s         664 B/op         26 allocs/op
BenchmarkUnmarshal/Small/Proto-12               13172190                90.17 ns/op     1519.34 MB/s         288 B/op          6 allocs/op
BenchmarkUnmarshal/Small/Msgpack-12              6867572               172.7 ns/op      1018.94 MB/s         336 B/op          7 allocs/op
BenchmarkUnmarshal/Small/Jsoniter-12            10142586               121.4 ns/op      1623.24 MB/s         288 B/op          6 allocs/op
BenchmarkUnmarshal/Small/GoccyDefault-12        11486623               102.8 ns/op      1917.22 MB/s         336 B/op          2 allocs/op
BenchmarkUnmarshal/Small/GoccyUnordered-12      11248878               104.2 ns/op      1891.42 MB/s         336 B/op          2 allocs/op
BenchmarkUnmarshal/Small/GoccyFastest-12        11496760               102.7 ns/op      1919.04 MB/s         336 B/op          2 allocs/op
BenchmarkUnmarshal/Small/SonicStd-12             8137474               138.6 ns/op      1420.90 MB/s         559 B/op          8 allocs/op
BenchmarkUnmarshal/Small/SonicDefault-12        11791794               107.8 ns/op      1827.66 MB/s         402 B/op          3 allocs/op
BenchmarkUnmarshal/Small/SonicCompact-12        10760235               104.9 ns/op      1877.60 MB/s         403 B/op          3 allocs/op
BenchmarkUnmarshal/Small/SonicFast-12           11418880               102.9 ns/op      1913.79 MB/s         403 B/op          3 allocs/op
```
