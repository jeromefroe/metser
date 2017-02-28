package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v2"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/DeDiS/protobuf"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/gogo/protobuf/proto"
	"github.com/google/flatbuffers/go"
	hprose "github.com/hprose/hprose-golang/io"
	"github.com/jeromefroe/metser"
	"github.com/ugorji/go/codec"
)

// gopkg.in/vmihailenco/msgpack.v2

func vmihailencoMsgpackSizes() {
	c := metser.GetCounter()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	metser.MsgpackMarshalCounter(enc, &c)
	cn := len(buf.Bytes())

	t := metser.GetBatchTimer()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = msgpack.NewEncoder(buf)
	metser.MsgpackMarshalBatchTimer(enc, &t)
	tn := len(buf.Bytes())

	g := metser.GetGauge()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = msgpack.NewEncoder(buf)
	metser.MsgpackMarshalGauge(enc, &g)
	gn := len(buf.Bytes())

	u := metser.GetMetricUnion()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = msgpack.NewEncoder(buf)
	metser.MsgpackMarshalMetricUnion(enc, &u)
	un := len(buf.Bytes())

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "VmihailencoMsgpack", cn, tn, gn, un)
}

// github.com/tinylib/msgp

func msgpSizes() {
	c := metser.GetCounter()
	buf := make([]byte, 0, 256)
	buf, _ = c.MarshalMsg(buf)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf = make([]byte, 0, 256)
	buf, _ = t.MarshalMsg(buf)
	tn := len(buf)

	g := metser.GetGauge()
	buf = make([]byte, 0, 256)
	buf, _ = g.MarshalMsg(buf)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf = make([]byte, 0, 256)
	buf, _ = u.MarshalMsg(buf)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "TinylibMsgp", cn, tn, gn, un)
}

// github.com/ugorji/go/codec (messagepack)

func codecMsgpackSizes() {
	var mh codec.MsgpackHandle

	c := metser.GetCounter()
	buf := make([]byte, 0, 256)
	enc := codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf = make([]byte, 0, 256)
	enc = codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&t)
	tn := len(buf)

	g := metser.GetGauge()
	buf = make([]byte, 0, 256)
	enc = codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf = make([]byte, 0, 256)
	enc = codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "UgorjiCodecMessagepack", cn, tn, gn, un)
}

// github.com/google/flatbuffers/go

func flatbuffersSizes() {
	c := metser.GetCounter()
	builder := metser.FlatBuffersSerializer{Builder: flatbuffers.NewBuilder(256)}
	buf := builder.MarshalCounter(c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	builder = metser.FlatBuffersSerializer{Builder: flatbuffers.NewBuilder(256)}
	buf = builder.MarshalBatchTimer(t)
	tn := len(buf)

	g := metser.GetGauge()
	builder = metser.FlatBuffersSerializer{Builder: flatbuffers.NewBuilder(256)}
	buf = builder.MarshalGauge(g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	builder = metser.FlatBuffersSerializer{Builder: flatbuffers.NewBuilder(256)}
	buf = builder.MarshalMetricUnion(u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Flatbuffers", cn, tn, gn, un)
}

// github.com/golang/protobuf

func protobufSizes() {
	c := metser.GetProtobufCounter()
	buf, _ := proto.Marshal(&c)
	cn := len(buf)

	t := metser.GetProtobufBatchTimer()
	buf, _ = proto.Marshal(&t)
	tn := len(buf)

	g := metser.GetProtobufGauge()
	buf, _ = proto.Marshal(&g)
	gn := len(buf)

	u := metser.GetProtobufMetricUnion()
	buf, _ = proto.Marshal(&u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "GooogleProtobuf", cn, tn, gn, un)
}

// github.com/gogo/protobuf

func gogoProtobufSizes() {
	c := metser.GetGoGoProtobufCounter()
	buf := make([]byte, c.Size())
	cn, _ := c.MarshalTo(buf)

	t := metser.GetGoGoProtobufBatchTimer()
	buf = make([]byte, t.Size())
	tn, _ := t.MarshalTo(buf)

	g := metser.GetGoGoProtobufGauge()
	buf = make([]byte, g.Size())
	gn, _ := g.MarshalTo(buf)

	u := metser.GetGoGoProtobufMetricUnion()
	buf = make([]byte, u.Size())
	un, _ := u.MarshalTo(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "GogoProtobuf", cn, tn, gn, un)
}

// github.com/DeDiS/protobuf

func dedisProtobufSizes() {
	c := metser.GetProtobufCounter()
	buf, _ := protobuf.Encode(&c)
	cn := len(buf)

	t := metser.GetProtobufBatchTimer()
	buf, _ = protobuf.Encode(&t)
	tn := len(buf)

	g := metser.GetProtobufGauge()
	buf, _ = protobuf.Encode(&g)
	gn := len(buf)

	u := metser.GetProtobufMetricUnion()
	buf, _ = protobuf.Encode(&u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "DedisProtobuf", cn, tn, gn, un)
}

// zombiezen.com/go/capnproto2

func capnpSizes() {
	c := metser.GetCapnpCounterMessage()
	buf, _ := c.Marshal()
	cn := len(buf)

	t := metser.GetCapnpBatchTimerMessage()
	buf, _ = t.Marshal()
	tn := len(buf)

	g := metser.GetCapnpGaugeMessage()
	buf, _ = g.Marshal()
	gn := len(buf)

	u := metser.GetCapnpMetricUnionMessage()
	buf, _ = u.Marshal()
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Capnproto", cn, tn, gn, un)
}

// git.apache.org/thrift.git/lib/go/thrift

func thriftSizes() {
	c := metser.GetThriftCounter()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	c.Write(proto)
	cn := transport.RemainingBytes()

	t := metser.GetThriftBatchTimer()
	transport = thrift.NewTMemoryBufferLen(256)
	proto = thrift.NewTBinaryProtocolTransport(transport)
	t.Write(proto)
	tn := transport.RemainingBytes()

	g := metser.GetThriftGauge()
	transport = thrift.NewTMemoryBufferLen(256)
	proto = thrift.NewTBinaryProtocolTransport(transport)
	g.Write(proto)
	gn := transport.RemainingBytes()

	u := metser.GetThriftMetricUnion()
	transport = thrift.NewTMemoryBufferLen(256)
	proto = thrift.NewTBinaryProtocolTransport(transport)
	u.Write(proto)
	un := transport.RemainingBytes()

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Thrift", cn, tn, gn, un)
}

// encoding/gob

func gobSizes() {
	c := metser.GetCounter()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := gob.NewEncoder(buf)
	enc.Encode(&c)
	cn := len(buf.Bytes())

	t := metser.GetBatchTimer()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = gob.NewEncoder(buf)
	enc.Encode(&t)
	tn := len(buf.Bytes())

	g := metser.GetGauge()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = gob.NewEncoder(buf)
	enc.Encode(&g)
	gn := len(buf.Bytes())

	u := metser.GetMetricUnion()
	buf = bytes.NewBuffer(make([]byte, 0, 256))
	enc = gob.NewEncoder(buf)
	enc.Encode(&u)
	un := len(buf.Bytes())

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "EncodingGob", cn, tn, gn, un)
}

// encoding/json

func jsonSizes() {
	c := metser.GetCounter()
	buf, _ := json.Marshal(c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf, _ = json.Marshal(t)
	tn := len(buf)

	g := metser.GetGauge()
	buf, _ = json.Marshal(g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf, _ = json.Marshal(u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "EncodingJSON", cn, tn, gn, un)
}

// github.com/mailru/easyjson

func easyjsonSizes() {
	c := metser.GetCounter()
	buf, _ := c.MarshalJSON()
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf, _ = t.MarshalJSON()
	tn := len(buf)

	g := metser.GetGauge()
	buf, _ = g.MarshalJSON()
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf, _ = u.MarshalJSON()
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "MailruEasyjson", cn, tn, gn, un)
}

// labix.org/v2/mgo/bson

func bsonSizes() {
	c := metser.GetCounter()
	buf, _ := bson.Marshal(c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf, _ = bson.Marshal(t)
	tn := len(buf)

	g := metser.GetGauge()
	buf, _ = bson.Marshal(g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf, _ = bson.Marshal(u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "BSON", cn, tn, gn, un)
}

// github.com/davecgh/go-xdr/xdr2

func xdrSizes() {
	c := metser.GetCounter()
	buf, _ := xdr.Marshal(c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf, _ = xdr.Marshal(t)
	tn := len(buf)

	g := metser.GetGauge()
	buf, _ = xdr.Marshal(g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf, _ = xdr.Marshal(u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "XDR", cn, tn, gn, un)
}

// github.com/Sereal/Sereal/Go/sereal

func serealSizes() {
	c := metser.GetCounter()
	buf, _ := sereal.Marshal(c)
	cn := len(buf)

	t := metser.GetBatchTimer()
	buf, _ = sereal.Marshal(t)
	tn := len(buf)

	g := metser.GetGauge()
	buf, _ = sereal.Marshal(g)
	gn := len(buf)

	u := metser.GetMetricUnion()
	buf, _ = sereal.Marshal(u)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Sereal", cn, tn, gn, un)
}

// github.com/andyleap/gencode

func gencodeSizes() {
	c := metser.GetGencodeCounter()
	bts := make([]byte, c.Size())
	buf, _ := c.Marshal(bts)
	cn := len(buf)

	t := metser.GetGencodeBatchTimer()
	bts = make([]byte, t.Size())
	buf, _ = t.Marshal(bts)
	tn := len(buf)

	g := metser.GetGencodeGauge()
	bts = make([]byte, g.Size())
	buf, _ = g.Marshal(bts)
	gn := len(buf)

	u := metser.GetGencodeMetricUnion()
	bts = make([]byte, u.Size())
	buf, _ = u.Marshal(bts)
	un := len(buf)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Gencode", cn, tn, gn, un)
}

// github.com/pascaldekloe/colfer

func colferSizes() {
	c := metser.GetColferCounter()
	n, _ := c.MarshalLen()
	bts := make([]byte, n)
	cn := c.MarshalTo(bts)

	t := metser.GetColferBatchTimer()
	n, _ = t.MarshalLen()
	bts = make([]byte, n)
	tn := t.MarshalTo(bts)

	g := metser.GetColferGauge()
	n, _ = g.MarshalLen()
	bts = make([]byte, n)
	gn := g.MarshalTo(bts)

	u := metser.GetColferMetricUnion()
	n, _ = u.MarshalLen()
	bts = make([]byte, n)
	un := u.MarshalTo(bts)

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Colfer", cn, tn, gn, un)
}

// github.com/hprose/hprose-golang

func hproseSizes() {
	c := metser.GetCounter()
	buf := make([]byte, 0, 256)
	w := hprose.NewWriter(true, buf...)
	w.Serialize(c)
	cn := len(w.Bytes())

	t := metser.GetBatchTimer()
	buf = make([]byte, 0, 256)
	w = hprose.NewWriter(true, buf...)
	w.Serialize(t)
	tn := len(w.Bytes())

	g := metser.GetGauge()
	buf = make([]byte, 0, 256)
	w = hprose.NewWriter(true, buf...)
	w.Serialize(g)
	gn := len(w.Bytes())

	u := metser.GetMetricUnion()
	buf = make([]byte, 0, 256)
	w = hprose.NewWriter(true, buf...)
	w.Serialize(u)
	un := len(w.Bytes())

	fmt.Printf("%-24s%12d B%12d B%12d B%12d B\n", "Hprose", cn, tn, gn, un)
}

func main() {
	fmt.Println("Encoding Scheme                 Counter      BatchTimer      Gauge       MetricUnion")
	fmt.Println("---------------                 -------      ----------      -----       -----------")
	vmihailencoMsgpackSizes()
	msgpSizes()
	codecMsgpackSizes()
	flatbuffersSizes()
	protobufSizes()
	gogoProtobufSizes()
	dedisProtobufSizes()
	capnpSizes()
	thriftSizes()
	gobSizes()
	jsonSizes()
	easyjsonSizes()
	bsonSizes()
	xdrSizes()
	serealSizes()
	gencodeSizes()
	colferSizes()
	hproseSizes()
}
