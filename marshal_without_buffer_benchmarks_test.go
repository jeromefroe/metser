package metser

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/DeDiS/protobuf"
	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/gogo/protobuf/proto"
	"github.com/google/flatbuffers/go"
	hprose "github.com/hprose/hprose-golang/io"
	"github.com/ugorji/go/codec"

	"gopkg.in/mgo.v2/bson"
	"gopkg.in/vmihailenco/msgpack.v2"
)

var bts []byte
var buf *bytes.Buffer

// gopkg.in/vmihailenco/msgpack.v2

func BenchmarkVmihailencoMsgpackMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = bytes.NewBuffer(make([]byte, 0, 256))
		enc := msgpack.NewEncoder(buf)
		MsgpackMarshalCounter(enc, &c)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = bytes.NewBuffer(make([]byte, 0, 256))
		enc := msgpack.NewEncoder(buf)
		MsgpackMarshalBatchTimer(enc, &t)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = bytes.NewBuffer(make([]byte, 0, 256))
		enc := msgpack.NewEncoder(buf)
		MsgpackMarshalGauge(enc, &g)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf = bytes.NewBuffer(make([]byte, 0, 256))
		enc := msgpack.NewEncoder(buf)
		MsgpackMarshalMetricUnion(enc, &u)
	}
}

// github.com/tinylib/msgp

func BenchmarkMsgpMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		c.MarshalMsg(bts)
	}
}

func BenchmarkMsgpMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		t.MarshalMsg(bts)
	}
}

func BenchmarkMsgpMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		g.MarshalMsg(bts)
	}
}

func BenchmarkMsgpMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		u.MarshalMsg(bts)
	}
}

// github.com/ugorji/go/codec (messagepack)

func BenchmarkCodecMsgpackMarshalWithoutBufferCounter(b *testing.B) {
	var mh codec.MsgpackHandle
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		enc := codec.NewEncoderBytes(&bts, &mh)
		enc.Encode(&c)
	}
}

func BenchmarkCodecMsgpackMarshalWithoutBufferBatchTimer(b *testing.B) {
	var mh codec.MsgpackHandle
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		enc := codec.NewEncoderBytes(&bts, &mh)
		enc.Encode(&t)
	}
}

func BenchmarkCodecMsgpackMarshalWithoutBufferGauge(b *testing.B) {
	var mh codec.MsgpackHandle
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, 0, 256)
		enc := codec.NewEncoderBytes(&buf, &mh)
		enc.Encode(&g)
	}
}

func BenchmarkCodecMsgpackMarshalWithoutBufferMetricUnion(b *testing.B) {
	var mh codec.MsgpackHandle
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, 0, 256)
		enc := codec.NewEncoderBytes(&bts, &mh)
		enc.Encode(&u)
	}
}

// github.com/google/flatbuffers/go

func BenchmarkFlatBuffersMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
		builder.MarshalCounter(c)
	}
}

func BenchmarkFlatBuffersMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
		builder.MarshalBatchTimer(t)
	}
}

func BenchmarkFlatBuffersMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
		builder.MarshalGauge(g)
	}
}

func BenchmarkFlatBuffersMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
		builder.MarshalMetricUnion(u)
	}
}

// github.com/golang/protobuf

func BenchmarkProtobufMarshalWithoutBufferCounter(b *testing.B) {
	c := GetProtobufCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&c)
	}
}

func BenchmarkProtobufMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetProtobufBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&t)
	}
}

func BenchmarkProtobufMarshalWithoutBufferGauge(b *testing.B) {
	g := GetProtobufGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&g)
	}
}

func BenchmarkProtobufMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetProtobufMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&u)
	}
}

// github.com/gogo/protobuf

func BenchmarkGoGoProtobufMarshalWithoutBufferCounter(b *testing.B) {
	c := GetGoGoProtobufCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, c.Size())
		c.MarshalTo(bts)
	}
}

func BenchmarkGoGoProtobufMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetGoGoProtobufBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, t.Size())
		t.MarshalTo(bts)
	}
}

func BenchmarkGoGoProtobufMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGoGoProtobufGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, g.Size())
		g.MarshalTo(bts)
	}
}

func BenchmarkGoGoProtobufMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetGoGoProtobufMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, u.Size())
		u.MarshalTo(bts)
	}
}

// github.com/DeDiS/protobuf

func BenchmarkDedisProtobufMarshalWithoutBufferCounter(b *testing.B) {
	c := GetProtobufCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Encode(&c)
	}
}

func BenchmarkDedisProtobufMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetProtobufBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Encode(&t)
	}
}

func BenchmarkDedisProtobufMarshalWithoutBufferGauge(b *testing.B) {
	g := GetProtobufGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Encode(&g)
	}
}

func BenchmarkDedisProtobufMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetProtobufMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Encode(&u)
	}
}

// zombiezen.com/go/capnproto2

func BenchmarkCapnpMarshalWithoutBufferCounter(b *testing.B) {
	m := GetCapnpCounterMessage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Marshal()
	}
}

func BenchmarkCapnpfMarshalWithoutBufferBatchTimer(b *testing.B) {
	m := GetCapnpBatchTimerMessage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Marshal()
	}
}

func BenchmarkCapnpfMarshalWithoutBufferGauge(b *testing.B) {
	m := GetCapnpGaugeMessage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Marshal()
	}
}

func BenchmarkCapnpfMarshalWithoutBufferMetricUnion(b *testing.B) {
	m := GetCapnpCounterMessage()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Marshal()
	}
}

// git.apache.org/thrift.git/lib/go/thrift

func BenchmarkThriftMarshalWithoutBufferCounter(b *testing.B) {
	c := GetThriftCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transport := thrift.NewTMemoryBufferLen(256)
		proto := thrift.NewTBinaryProtocolTransport(transport)
		c.Write(proto)
	}
}

func BenchmarkThriftMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetThriftBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transport := thrift.NewTMemoryBufferLen(256)
		proto := thrift.NewTBinaryProtocolTransport(transport)
		t.Write(proto)
	}
}

func BenchmarkThriftMarshalWithoutBufferGauge(b *testing.B) {
	g := GetThriftGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transport := thrift.NewTMemoryBufferLen(256)
		proto := thrift.NewTBinaryProtocolTransport(transport)
		g.Write(proto)
	}
}

func BenchmarkThriftMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetThriftMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transport := thrift.NewTMemoryBufferLen(256)
		proto := thrift.NewTBinaryProtocolTransport(transport)
		u.Write(proto)
	}
}

// encoding/gob

func BenchmarkGobMarshalWithoutBufferCounter(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := gob.NewEncoder(buf)
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(&c)
	}
}

func BenchmarkGobMarshalWithoutBufferBatchTimer(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := gob.NewEncoder(buf)
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(&t)
	}
}

func BenchmarkGobMarshalWithoutBufferGauge(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := gob.NewEncoder(buf)
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(&g)
	}
}

func BenchmarkGobMarshalWithoutBufferMetricUnion(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := gob.NewEncoder(buf)
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(&u)
	}
}

// encoding/json

func BenchmarkJSONMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(c)
	}
}

func BenchmarkJSONMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(t)
	}
}

func BenchmarkJSONMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(g)
	}
}

func BenchmarkJSONMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Marshal(u)
	}
}

// github.com/mailru/easyjson

func BenchmarkEasyjsonMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.MarshalJSON()
	}
}

func BenchmarkEasyjsonMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.MarshalJSON()
	}
}

func BenchmarkEasyjsonMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.MarshalJSON()
	}
}

func BenchmarkEasyjsonMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalJSON()
	}
}

// labix.org/v2/mgo/bson

func BenchmarkBSONMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Marshal(c)
	}
}

func BenchmarkBSONMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Marshal(t)
	}
}

func BenchmarkBSONMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Marshal(g)
	}
}

func BenchmarkBSONMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Marshal(u)
	}
}

// github.com/davecgh/go-xdr/xdr2

func BenchmarkXDRMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Marshal(c)
	}
}

func BenchmarkXDRMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Marshal(t)
	}
}

func BenchmarkXDRMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Marshal(g)
	}
}

func BenchmarkXDRMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Marshal(u)
	}
}

// github.com/Sereal/Sereal/Go/sereal

func BenchmarkSerealMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Marshal(c)
	}
}

func BenchmarkSerealMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Marshal(t)
	}
}

func BenchmarkSerealWithoutBufferMarshalGauge(b *testing.B) {
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Marshal(g)
	}
}

func BenchmarkSerealMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Marshal(u)
	}
}

// github.com/andyleap/gencode

func BenchmarkGencodeMarshalWithoutBufferCounter(b *testing.B) {
	c := GetGencodeCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, c.Size())
		c.Marshal(bts)
	}
}

func BenchmarkGencodeMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetGencodeBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, t.Size())
		t.Marshal(bts)
	}
}

func BenchmarkGencodeMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGencodeGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, g.Size())
		g.Marshal(bts)
	}
}

func BenchmarkGencodeMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetGencodeMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = make([]byte, u.Size())
		u.Marshal(bts)
	}
}

// github.com/pascaldekloe/colfer

func BenchmarkColferMarshalWithoutBufferCounter(b *testing.B) {
	c := GetColferCounter()
	n, _ := c.MarshalLen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, n)
		c.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetColferBatchTimer()
	n, _ := t.MarshalLen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, n)
		t.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithoutBufferGauge(b *testing.B) {
	g := GetColferGauge()
	n, _ := g.MarshalLen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, n)
		g.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetColferMetricUnion()
	n, _ := u.MarshalLen()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := make([]byte, n)
		u.MarshalTo(buf)
	}
}

// github.com/hprose/hprose-golang

func BenchmarkHproselMarshalWithoutBufferCounter(b *testing.B) {
	c := GetCounter()
	buf := make([]byte, 0, 256)
	w := hprose.NewWriter(true, buf...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Reset()
		w.Serialize(c)
	}
}

func BenchmarkHproselMarshalWithoutBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	buf := make([]byte, 0, 256)
	w := hprose.NewWriter(true, buf...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Reset()
		w.Serialize(t)
	}
}

func BenchmarkHproselMarshalWithoutBufferGauge(b *testing.B) {
	g := GetGauge()
	buf := make([]byte, 0, 256)
	w := hprose.NewWriter(true, buf...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Reset()
		w.Serialize(g)
	}
}

func BenchmarkHproselMarshalWithoutBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	buf := make([]byte, 0, 256)
	w := hprose.NewWriter(true, buf...)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Reset()
		w.Serialize(u)
	}
}
