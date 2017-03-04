# metser

This repo hosts various benchmarks for serializing and deserializing metrics. It is intended to
give a broad glimpse of the tradeoffs, for example serialization time vs the size of serialized
data, between different serialization approaches. The structs I used to perform these benchmarks
can be seen below:

```
type Type int8

type ID []byte

type Counter struct {
	ID    ID
	Value int64
}

type BatchTimer struct {
	ID     ID
	Values []float64
}

type Gauge struct {
	ID    ID
	Value float64
}

type MetricUnion struct {
	Type          Type
	ID            ID
	CounterVal    int64
	BatchTimerVal []float64
	GaugeVal      float64
}
```

Note that I chose not to use any `optional` or `union` fields in the `MetricUnion` struct since
these fields often require pointers and I want to avoid heap allocations. Some of the protocol
definitions and serialization methods could be improved (for example, for
gopkg.in/vmihailenco/msgpack.v2 one could use the type of the metric union to determine how the
metric should be serialized so that one doesn't have to encode each metric field) however, as
the point of this exercise was to only get the general picture of the landscape of
serialization methods, I didn't perform any optimizations special to a given technique.

The serialization methods I looked at are:

- [gopkg.in/vmihailenco/msgpack.v2](https://github.com/vmihailenco/msgpack)
- [github.com/tinylib/msgp](https://github.com/tinylib/msgp)
- [github.com/ugorji/go/codec (messagepack)](https://github.com/ugorji/go/tree/master/codec)
- [github.com/google/flatbuffers/go](https://github.com/google/flatbuffers)
- [github.com/golang/protobuf](https://github.com/golang/protobuf)
- [github.com/gogo/protobuf](https://github.com/gogo/protobuf)
- [github.com/DeDiS/protobuf](https://github.com/DeDiS/protobuf)
- [zombiezen.com/go/capnproto2](https://godoc.org/zombiezen.com/go/capnproto2)
- [git.apache.org/thrift.git/lib/go/thrift](https://github.com/apache/thrift/tree/master/lib/go)
- [encoding/gob](https://golang.org/pkg/encoding/gob/)
- [encoding/json](https://golang.org/pkg/encoding/json/)
- [github.com/mailru/easyjson](https://github.com/mailru/easyjson)
- [labix.org/v2/mgo/bson](https://godoc.org/labix.org/v2/mgo/bson)
- [github.com/davecgh/go-xdr/xdr2](https://github.com/davecgh/go-xdr)
- [github.com/andyleap/gencode](https://github.com/andyleap/gencode)
- [github.com/pascaldekloe/colfer](https://github.com/pascaldekloe/colfer)
- [github.com/Sereal/Sereal/Go/sereal](https://github.com/Sereal/Sereal)
- [github.com/hprose/hprose-golang](https://github.com/hprose/hprose-golang)

This work was largely inspired by the work done
[here](https://github.com/alecthomas/go_serialization_benchmarks) which also looked at serialization
methods in Go, so anyone interested in analyzing the tradeoffs between different serialization
methods might like reading over the results there as well.

### Marshal with Buffer

These first benchmarks look at the time it takes to marshal the different metrics into byte
slices. These libraries offer the ability to pass the underlying byte slice to the `Marshal`
method. I decided to seperate these out since they allow one to pool these byte slice across
calls to `Marshal` thereby saving allocations and reducing GC.

```
BenchmarkVmihailencoMsgpackMarshalWithBufferCounter-4           20000000                70.1 ns/op
BenchmarkVmihailencoMsgpackMarshalWithBufferBatchTimer-4        10000000               211 ns/op
BenchmarkVmihailencoMsgpackMarshalWithBufferGauge-4             20000000                85.4 ns/op
BenchmarkVmihailencoMsgpackMarshalWithBufferMetricUnion-4        5000000               274 ns/op
BenchmarkMsgpMarshalWithBufferCounter-4                         50000000                20.5 ns/op
BenchmarkMsgpMarshalWithBufferBatchTimer-4                      20000000                69.8 ns/op
BenchmarkMsgpMarshalWithBufferGauge-4                           50000000                27.9 ns/op
BenchmarkMsgpMarshalWithBufferMetricUnion-4                     20000000                94.4 ns/op
BenchmarkCodecMsgpackMarshalWithBufferCounter-4                  2000000               586 ns/op
BenchmarkCodecMsgpackMarshalWithBufferBatchTimer-4               2000000               985 ns/op
BenchmarkCodecMsgpackMarshalWithBufferGauge-4                    2000000               531 ns/op
BenchmarkCodecMsgpackMarshalWithBufferMetricUnion-4              1000000              1101 ns/op
BenchmarkFlatBuffersMarshalWithBufferCounter-4                   3000000               674 ns/op
BenchmarkFlatBuffersMarshalWithBufferBatchTimer-4                2000000               738 ns/op
BenchmarkFlatBuffersMarshalWithBufferGauge-4                     3000000               525 ns/op
BenchmarkFlatBuffersMarshalWithBufferMetricUnion-4               2000000               706 ns/op
BenchmarkGoGoProtobufMarshalWithBufferCounter-4                 100000000               16.0 ns/op
BenchmarkGoGoProtobufMarshalWithBufferBatchTimer-4              50000000                41.8 ns/op
BenchmarkGoGoProtobufMarshalWithBufferGauge-4                   100000000               17.2 ns/op
BenchmarkGoGoProtobufMarshalWithBufferMetricUnion-4             30000000                40.4 ns/op
BenchmarkCapnpMarshalWithBufferCounter-4                        10000000               123 ns/op
BenchmarkCapnpMarshalWithBufferBatchTimer-4                     10000000               125 ns/op
BenchmarkCapnpMarshalWithBufferGauge-4                          10000000               121 ns/op
BenchmarkCapnpfMarshalWithBufferMetricUnion-4                   10000000               122 ns/op
BenchmarkThriftMarshalWithBufferCounter-4                        5000000               222 ns/op
BenchmarkThriftMarshalWithBufferBatchTimer-4                     3000000               400 ns/op
BenchmarkThriftMarshalWithBufferGauge-4                          5000000               224 ns/op
BenchmarkThriftMarshalWithBufferMetricUnion-4                    2000000               658 ns/op
BenchmarkGencodeMarshalWithBufferCounter-4                      100000000               14.4 ns/op
BenchmarkGencodeMarshalWithBufferBatchTimer-4                   50000000                33.4 ns/op
BenchmarkGencodeMarshalWithBufferGauge-4                        100000000               17.3 ns/op
BenchmarkGencodeMarshalWithBufferMetricUnion-4                  50000000                38.3 ns/op
BenchmarkColferMarshalWithBufferCounter-4                       100000000               13.2 ns/op
BenchmarkColferMarshalWithBufferBatchTimer-4                    50000000                34.2 ns/op
BenchmarkColferMarshalWithBufferGauge-4                         100000000               16.5 ns/op
BenchmarkColferMarshalWithBufferMetricUnion-4                   30000000                38.4 ns/op
```

### Marshal without buffer

These benchmarks also look at Marshaling metrics but to level the playing field for serialization
methods which don't expose the ability to pass `Marshal` a byte slice we allocate a new one each
time we call `Marshal` for those methods which do take a byte slice as an argument.

```
BenchmarkVmihailencoMsgpackMarshalWithoutBufferCounter-4         5000000               365 ns/op
BenchmarkVmihailencoMsgpackMarshalWithoutBufferBatchTimer-4      3000000               628 ns/op
BenchmarkVmihailencoMsgpackMarshalWithoutBufferGauge-4           3000000               398 ns/op
BenchmarkVmihailencoMsgpackMarshalWithoutBufferMetricUnion-4     2000000               671 ns/op
BenchmarkMsgpMarshalWithoutBufferCounter-4                      10000000               145 ns/op
BenchmarkMsgpMarshalWithoutBufferBatchTimer-4                   10000000               271 ns/op
BenchmarkMsgpMarshalWithoutBufferGauge-4                        10000000               190 ns/op
BenchmarkMsgpMarshalWithoutBufferMetricUnion-4                  10000000               182 ns/op
BenchmarkCodecMsgpackMarshalWithoutBufferCounter-4               1000000              1965 ns/op
BenchmarkCodecMsgpackMarshalWithoutBufferBatchTimer-4            1000000              2462 ns/op
BenchmarkCodecMsgpackMarshalWithoutBufferGauge-4                  500000              2158 ns/op
BenchmarkCodecMsgpackMarshalWithoutBufferMetricUnion-4            500000              2842 ns/op
BenchmarkFlatBuffersMarshalWithoutBufferCounter-4                2000000               741 ns/op
BenchmarkFlatBuffersMarshalWithoutBufferBatchTimer-4             2000000              1027 ns/op
BenchmarkFlatBuffersMarshalWithoutBufferGauge-4                  2000000               904 ns/op
BenchmarkFlatBuffersMarshalWithoutBufferMetricUnion-4            1000000              1199 ns/op
BenchmarkProtobufMarshalWithoutBufferCounter-4                   3000000               471 ns/op
BenchmarkProtobufMarshalWithoutBufferBatchTimer-4                2000000               739 ns/op
BenchmarkProtobufMarshalWithoutBufferGauge-4                     5000000               386 ns/op
BenchmarkProtobufMarshalWithoutBufferMetricUnion-4               2000000               780 ns/op
BenchmarkGoGoProtobufMarshalWithoutBufferCounter-4              20000000                60.7 ns/op
BenchmarkGoGoProtobufMarshalWithoutBufferBatchTimer-4           20000000                93.8 ns/op
BenchmarkGoGoProtobufMarshalWithoutBufferGauge-4                20000000                59.6 ns/op
BenchmarkGoGoProtobufMarshalWithoutBufferMetricUnion-4          20000000                97.2 ns/op
BenchmarkDedisProtobufMarshalWithoutBufferCounter-4              3000000               576 ns/op
BenchmarkDedisProtobufMarshalWithoutBufferBatchTimer-4           2000000              1008 ns/op
BenchmarkDedisProtobufMarshalWithoutBufferGauge-4                2000000               592 ns/op
BenchmarkDedisProtobufMarshalWithoutBufferMetricUnion-4          1000000              1304 ns/op
BenchmarkCapnpMarshalWithoutBufferCounter-4                     10000000               136 ns/op
BenchmarkCapnpfMarshalWithoutBufferBatchTimer-4                 10000000               145 ns/op
BenchmarkCapnpfMarshalWithoutBufferGauge-4                      10000000               149 ns/op
BenchmarkCapnpfMarshalWithoutBufferMetricUnion-4                10000000               156 ns/op
BenchmarkThriftMarshalWithoutBufferCounter-4                     2000000               561 ns/op
BenchmarkThriftMarshalWithoutBufferBatchTimer-4                  2000000               726 ns/op
BenchmarkThriftMarshalWithoutBufferGauge-4                       3000000               557 ns/op
BenchmarkThriftMarshalWithoutBufferMetricUnion-4                 1000000              1038 ns/op
BenchmarkGobMarshalWithoutBufferCounter-4                        2000000               664 ns/op
BenchmarkGobMarshalWithoutBufferBatchTimer-4                     2000000               924 ns/op
BenchmarkGobMarshalWithoutBufferGauge-4                          2000000               682 ns/op
BenchmarkGobMarshalWithoutBufferMetricUnion-4                    1000000              1036 ns/op
BenchmarkJSONMarshalWithoutBufferCounter-4                       1000000              1466 ns/op
BenchmarkJSONMarshalWithoutBufferBatchTimer-4                    1000000              2011 ns/op
BenchmarkJSONMarshalWithoutBufferGauge-4                         1000000              1596 ns/op
BenchmarkJSONMarshalWithoutBufferMetricUnion-4                    500000              3052 ns/op
BenchmarkEasyjsonMarshalWithoutBufferCounter-4                   5000000               254 ns/op
BenchmarkEasyjsonMarshalWithoutBufferBatchTimer-4                2000000               601 ns/op
BenchmarkEasyjsonMarshalWithoutBufferGauge-4                     5000000               305 ns/op
BenchmarkEasyjsonMarshalWithoutBufferMetricUnion-4               1000000              1035 ns/op
BenchmarkBSONMarshalWithoutBufferCounter-4                       2000000               607 ns/op
BenchmarkBSONMarshalWithoutBufferBatchTimer-4                    1000000              1393 ns/op
BenchmarkBSONMarshalWithoutBufferGauge-4                         2000000               656 ns/op
BenchmarkBSONMarshalWithoutBufferMetricUnion-4                    500000              2010 ns/op
BenchmarkXDRMarshalWithoutBufferCounter-4                        2000000               734 ns/op
BenchmarkXDRMarshalWithoutBufferBatchTimer-4                     1000000              1230 ns/op
BenchmarkXDRMarshalWithoutBufferGauge-4                          2000000               753 ns/op
BenchmarkXDRMarshalWithoutBufferMetricUnion-4                    1000000              1798 ns/op
BenchmarkSerealMarshalWithoutBufferCounter-4                      200000              6866 ns/op
BenchmarkSerealMarshalWithoutBufferBatchTimer-4                   200000              7925 ns/op
BenchmarkSerealMarshalWithoutBufferMetricUnion-4                  200000              8628 ns/op
BenchmarkGencodeMarshalWithoutBufferCounter-4                   20000000                57.9 ns/op
BenchmarkGencodeMarshalWithoutBufferBatchTimer-4                20000000                83.7 ns/op
BenchmarkGencodeMarshalWithoutBufferGauge-4                     20000000                63.8 ns/op
BenchmarkGencodeMarshalWithoutBufferMetricUnion-4               20000000                97.1 ns/op
BenchmarkColferMarshalWithoutBufferCounter-4                    30000000                46.2 ns/op
BenchmarkColferMarshalWithoutBufferBatchTimer-4                 20000000                79.4 ns/op
BenchmarkColferMarshalWithoutBufferGauge-4                      30000000                53.9 ns/op
BenchmarkColferMarshalWithoutBufferMetricUnion-4                20000000                80.4 ns/op
BenchmarkHproselMarshalWithoutBufferCounter-4                     500000              2678 ns/op
BenchmarkHproselMarshalWithoutBufferBatchTimer-4                  500000              2977 ns/op
BenchmarkHproselMarshalWithoutBufferGauge-4                       500000              2694 ns/op
BenchmarkHproselMarshalWithoutBufferMetricUnion-4                 300000              3446 ns/op
```

### Unmarshal

These benchmarks look at the time it takes to `Unmarshal` a byte slice into a Metric.

```
BenchmarkVmihailencoMsgpackUnmarshalCounter-4           10000000               167 ns/op
BenchmarkVmihailencoMsgpackUnmarshalBatchTimer-4         3000000               519 ns/op
BenchmarkVmihailencoMsgpackUnmarshalGauge-4             10000000               220 ns/op
BenchmarkVmihailencoMsgpackUnmarshalMetricUnion-4        2000000               589 ns/op
BenchmarkMsgpUnmarshalCounter-4                         20000000                68.0 ns/op
BenchmarkMsgpUnmarshalBatchTimer-4                      10000000               107 ns/op
BenchmarkMsgpUnmarshalGauge-4                           20000000                71.0 ns/op
BenchmarkMsgpUnmarshalMetricUnion-4                     10000000               199 ns/op
BenchmarkCodecMsgpackUnmarshalCounter-4                 10000000               176 ns/op
BenchmarkCodecMsgpackUnmarshalBatchTimer-4              10000000               176 ns/op
BenchmarkCodecMsgpackUnmarshalGauge-4                   10000000               172 ns/op
BenchmarkCodecMsgpackUnmarshalMetricUnion-4             10000000               178 ns/op
BenchmarkFlatBuffersUnmarshalCounter-4                   3000000               591 ns/op
BenchmarkFlatBuffersUnmarshalBatchTimer-4                2000000               707 ns/op
BenchmarkFlatBuffersUnmarshalGauge-4                     2000000               585 ns/op
BenchmarkFlatBuffersUnmarshalMetricUnion-4               2000000               710 ns/op
BenchmarkProtobufUnmarshalCounter-4                      5000000               315 ns/op
BenchmarkProtobufUnmarshalBatchTimer-4                   3000000               561 ns/op
BenchmarkProtobufUnmarshalGauge-4                        5000000               331 ns/op
BenchmarkProtobufUnmarshalMetricUnion-4                  2000000               614 ns/op
BenchmarkGoGoProtobufUnmarshalCounter-4                 100000000               18.3 ns/op
BenchmarkGoGoProtobufUnmarshalGauge-4                   100000000               18.1 ns/op
BenchmarkGoGoProtobufUnmarshalMetricUnion-4             10000000               172 ns/op
BenchmarkDedisProtobufUnmarshalCounter-4                 5000000               289 ns/op
BenchmarkDedisProtobufUnmarshalBatchTimer-4              1000000              1141 ns/op
BenchmarkDedisProtobufUnmarshalGauge-4                   5000000               294 ns/op
BenchmarkDedisProtobufUnmarshalMetricUnion-4             1000000              1478 ns/op
BenchmarkCapnpUnmarshalCounter-4                         2000000               572 ns/op
BenchmarkCapnpUnmarshalBatchTimer-4                      3000000               581 ns/op
BenchmarkCapnpUnmarshalGauge-4                           3000000               571 ns/op
BenchmarkCapnpUnmarshalMetricUnion-4                     3000000               553 ns/op
BenchmarkThriftUnmarshalCounter-4                        3000000               430 ns/op
BenchmarkThriftUnmarshalGauge-4                          3000000               428 ns/op
BenchmarkThriftUnmarshalMetricUnion-4                    3000000               450 ns/op
BenchmarkGobUnmarshalCounter-4                          10000000               138 ns/op
BenchmarkGobUnmarshalBatchTimer-4                       10000000               140 ns/op
BenchmarkGobUnmarshalGauge-4                            10000000               140 ns/op
BenchmarkGobUnmarshalMetricUnion-4                      10000000               145 ns/op
BenchmarkJSONUnmarshalCounter-4                          1000000              1755 ns/op
BenchmarkJSONUnmarshalBatchTimer-4                        500000              2326 ns/op
BenchmarkJSONUnmarshalGauge-4                            1000000              1778 ns/op
BenchmarkJSONUnmarshalMetricUnion-4                       500000              3199 ns/op
BenchmarkEasyjsonUnmarshalCounter-4                      3000000               524 ns/op
BenchmarkEasyjsonUnmarshalBatchTimer-4                   2000000               889 ns/op
BenchmarkEasyjsonUnmarshalGauge-4                        2000000               572 ns/op
BenchmarkEasyjsonUnmarshalMetricUnion-4                  1000000              1306 ns/op
BenchmarkBSONUnmarshalCounter-4                          2000000               919 ns/op
BenchmarkBSONUnmarshalBatchTimer-4                       1000000              2486 ns/op
BenchmarkBSONUnmarshalGauge-4                            2000000               941 ns/op
BenchmarkBSONUnmarshalMetricUnion-4                       500000              3215 ns/op
BenchmarkXDRUnmarshalCounter-4                           5000000               330 ns/op
BenchmarkXDRUnmarshalBatchTimer-4                        3000000               577 ns/op
BenchmarkXDRUnmarshalGauge-4                             5000000               322 ns/op
BenchmarkXDRUnmarshalMetricUnion-4                       2000000               980 ns/op
BenchmarkGencodeUnmarshalCounter-4                      100000000               12.3 ns/op
BenchmarkGencodeUnmarshalGauge-4                        100000000               14.4 ns/op
BenchmarkGencodeUnmarshalMetricUnion-4                  50000000                28.6 ns/op
BenchmarkColferUnmarshalCounter-4                       30000000                49.6 ns/op
BenchmarkColferUnmarshalBatchTimer-4                    10000000               111 ns/op
BenchmarkColferUnmarshalGauge-4                         30000000                50.2 ns/op
BenchmarkColferUnmarshalMetricUnion-4                   10000000               107 ns/op
BenchmarkSerealUnmarshalCounter-4                        1000000              2512 ns/op
BenchmarkSerealUnmarshalBatchTimer-4                      500000              2791 ns/op
BenchmarkSerealUnmarshalGauge-4                           500000              2820 ns/op
BenchmarkSerealUnmarshalMetricUnion-4                     300000              3649 ns/op
BenchmarkHproselUnmarshalCounter-4                       1000000              1868 ns/op
BenchmarkHproselUnmarshalBatchTimer-4                    1000000              2226 ns/op
BenchmarkHproselUnmarshalGauge-4                         1000000              1950 ns/op
BenchmarkHproselUnmarshalMetricUnion-4                    500000              2649 ns/op
```

### Size of Serialized data

Listed below are the sizes, in bytes, of the resulting byte slices when metrics have been
marshalled.

```
Encoding Scheme                 Counter      BatchTimer      Gauge       MetricUnion
---------------                 -------      ----------      -----       -----------
VmihailencoMsgpack                40 B          85 B          48 B          96 B
TinylibMsgp                       50 B          96 B          58 B         139 B
UgorjiCodecMessagepack            51 B          97 B          59 B         140 B
Flatbuffers                       72 B         116 B          72 B         128 B
GooogleProtobuf                   41 B          81 B          48 B          94 B
GogoProtobuf                      41 B          81 B          48 B          83 B
DedisProtobuf                     42 B          81 B          48 B          94 B
Capnproto                         72 B         112 B          72 B         136 B
Thrift                            56 B          93 B          56 B         119 B
EncodingGob                       85 B         126 B          84 B         173 B
EncodingJSON                      73 B          82 B          73 B         126 B
MailruEasyjson                    73 B          82 B          73 B         126 B
BSON                              66 B         119 B          66 B         174 B
XDR                               52 B          88 B          52 B         108 B
Sereal                           106 B         155 B         111 B         202 B
Gencode                           40 B          79 B          46 B          96 B
Colfer                            42 B          82 B          49 B          84 B
Hprose                           215 B         234 B         213 B         283 B
```

