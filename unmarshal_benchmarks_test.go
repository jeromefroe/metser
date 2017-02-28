package metser

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io"
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

	"zombiezen.com/go/capnproto2"
)

// gopkg.in/vmihailenco/msgpack.v2

func BenchmarkVmihailencoMsgpackUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalCounter(enc, &c)
	bts := buf.Bytes()
	r := bytes.NewReader(bts)
	dec := msgpack.NewDecoder(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, io.SeekStart)
		MsgpackUnmarshalCounter(dec, &c)
	}
}

func BenchmarkVmihailencoMsgpackUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalBatchTimer(enc, &t)
	bts := buf.Bytes()
	r := bytes.NewReader(bts)
	dec := msgpack.NewDecoder(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, io.SeekStart)
		MsgpackUnmarshalBatchTimer(dec, &t)
	}
}

func BenchmarkVmihailencoMsgpackUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalGauge(enc, &g)
	bts := buf.Bytes()
	r := bytes.NewReader(bts)
	dec := msgpack.NewDecoder(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, io.SeekStart)
		MsgpackUnmarshalGauge(dec, &g)
	}
}

func BenchmarkVmihailencoMsgpackUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalMetricUnion(enc, &u)
	bts := buf.Bytes()
	r := bytes.NewReader(bts)
	dec := msgpack.NewDecoder(r)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Seek(0, io.SeekStart)
		MsgpackUnmarshalMetricUnion(dec, &u)
	}
}

// github.com/tinylib/msgp

func BenchmarkMsgpUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := c.MarshalMsg(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.UnmarshalMsg(bts)
	}
}

func BenchmarkMsgpUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := t.MarshalMsg(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.UnmarshalMsg(bts)
	}
}

func BenchmarkMsgpUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := g.MarshalMsg(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.UnmarshalMsg(bts)
	}
}

func BenchmarkMsgpUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := u.MarshalMsg(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.UnmarshalMsg(bts)
	}
}

// github.com/ugorji/go/codec (messagepack)

func BenchmarkCodecMsgpackUnmarshalCounter(b *testing.B) {
	var mh codec.MsgpackHandle
	c := GetCounter()
	buf := make([]byte, 0, 256)
	enc := codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&c)
	dec := codec.NewDecoderBytes(buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&c)
	}
}

func BenchmarkCodecMsgpackUnmarshalBatchTimer(b *testing.B) {
	var mh codec.MsgpackHandle
	t := GetBatchTimer()
	buf := make([]byte, 0, 256)
	enc := codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&t)
	dec := codec.NewDecoderBytes(buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&t)
	}
}

func BenchmarkCodecMsgpackUnmarshalGauge(b *testing.B) {
	var mh codec.MsgpackHandle
	g := GetGauge()
	buf := make([]byte, 0, 256)
	enc := codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&g)
	dec := codec.NewDecoderBytes(buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&g)
	}
}

func BenchmarkCodecMsgpackUnmarshalMetricUnion(b *testing.B) {
	var mh codec.MsgpackHandle
	u := GetMetricUnion()
	buf := make([]byte, 0, 256)
	enc := codec.NewEncoderBytes(&buf, &mh)
	enc.Encode(&u)
	dec := codec.NewDecoderBytes(buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&u)
	}
}

// github.com/google/flatbuffers/go

func BenchmarkFlatBuffersUnmarshalCounter(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	c := GetCounter()
	bts := builder.MarshalCounter(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.UnmarshalCounter(bts, &c)
	}
}

func BenchmarkFlatBuffersUnmarshalBatchTimer(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	t := GetBatchTimer()
	bts := builder.MarshalBatchTimer(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.UnmarshalBatchTimer(bts, &t)
	}
}

func BenchmarkFlatBuffersUnmarshalGauge(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	g := GetGauge()
	bts := builder.MarshalGauge(g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.UnmarshalGauge(bts, &g)
	}
}

func BenchmarkFlatBuffersUnmarshalMetricUnion(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	u := GetMetricUnion()
	bts := builder.MarshalMetricUnion(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.UnmarshalMetricUnion(bts, &u)
	}
}

// github.com/golang/protobuf

func BenchmarkProtobufUnmarshalCounter(b *testing.B) {
	c := GetProtobufCounter()
	bts, _ := proto.Marshal(&c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bts, &c)
	}
}

func BenchmarkProtobufUnmarshalBatchTimer(b *testing.B) {
	t := GetProtobufBatchTimer()
	bts, _ := proto.Marshal(&t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bts, &t)
	}
}

func BenchmarkProtobufUnmarshalGauge(b *testing.B) {
	g := GetProtobufGauge()
	bts, _ := proto.Marshal(&g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bts, &g)
	}
}

func BenchmarkProtobufUnmarshalMetricUnion(b *testing.B) {
	u := GetProtobufMetricUnion()
	bts, _ := proto.Marshal(&u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bts, &u)
	}
}

// github.com/gogo/protobuf

func BenchmarkGoGoProtobufUnmarshalCounter(b *testing.B) {
	c := GetGoGoProtobufCounter()
	buf := make([]byte, c.Size())
	_, _ = c.MarshalTo(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Unmarshal(buf)
	}
}

func BenchmarGoGoProtobufUnmarshalBatchTimer(b *testing.B) {
	t := GetGoGoProtobufBatchTimer()
	buf := make([]byte, t.Size())
	_, _ = t.MarshalTo(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Unmarshal(buf)
	}
}

func BenchmarkGoGoProtobufUnmarshalGauge(b *testing.B) {
	g := GetGoGoProtobufGauge()
	buf := make([]byte, g.Size())
	_, _ = g.MarshalTo(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Unmarshal(buf)
	}
}

func BenchmarkGoGoProtobufUnmarshalMetricUnion(b *testing.B) {
	u := GetGoGoProtobufMetricUnion()
	buf := make([]byte, u.Size())
	_, _ = u.MarshalTo(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Unmarshal(buf)
	}
}

// github.com/DeDiS/protobuf

func BenchmarkDedisProtobufUnmarshalCounter(b *testing.B) {
	c := GetProtobufCounter()
	bts, _ := protobuf.Encode(&c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Decode(bts, &c)
	}
}

func BenchmarkDedisProtobufUnmarshalBatchTimer(b *testing.B) {
	t := GetProtobufBatchTimer()
	bts, _ := protobuf.Encode(&t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Decode(bts, &t)
	}
}

func BenchmarkDedisProtobufUnmarshalGauge(b *testing.B) {
	g := GetProtobufGauge()
	bts, _ := protobuf.Encode(&g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Decode(bts, &g)
	}
}

func BenchmarkDedisProtobufUnmarshalMetricUnion(b *testing.B) {
	u := GetProtobufMetricUnion()
	bts, _ := protobuf.Encode(&u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		protobuf.Decode(bts, &u)
	}
}

// zombiezen.com/go/capnproto2

func BenchmarkCapnpUnmarshalCounter(b *testing.B) {
	m := GetCapnpCounterMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	enc.Encode(m)
	bts := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, _ := capnp.Unmarshal(bts)
		ReadRootCapnpCounter(m)
	}
}

func BenchmarkCapnpUnmarshalBatchTimer(b *testing.B) {
	m := GetCapnpBatchTimerMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	enc.Encode(m)
	bts := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, _ := capnp.Unmarshal(bts)
		ReadRootCapnpBatchTimer(m)
	}
}

func BenchmarkCapnpUnmarshalGauge(b *testing.B) {
	m := GetCapnpGaugeMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	enc.Encode(m)
	bts := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, _ := capnp.Unmarshal(bts)
		ReadRootCapnpGauge(m)
	}
}

func BenchmarkCapnpUnmarshalMetricUnion(b *testing.B) {
	m := GetCapnpMetricUnionMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	enc.Encode(m)
	bts := buf.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m, _ := capnp.Unmarshal(bts)
		ReadRootCapnpMetricUnion(m)
	}
}

// git.apache.org/thrift.git/lib/go/thrift

func BenchmarkThriftUnmarshalCounter(b *testing.B) {
	c := GetThriftCounter()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	c.Write(proto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Read(proto)
	}
}

func BenchmarkThriftUnarshalBatchTimer(b *testing.B) {
	t := GetThriftBatchTimer()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	t.Write(proto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Read(proto)
	}
}

func BenchmarkThriftUnmarshalGauge(b *testing.B) {
	g := GetThriftGauge()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	g.Write(proto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Read(proto)
	}
}

func BenchmarkThriftUnmarshalMetricUnion(b *testing.B) {
	u := GetThriftMetricUnion()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	u.Write(proto)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Read(proto)
	}
}

// encoding/gob

func BenchmarkGobUnmarshalCounter(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	enc := gob.NewEncoder(buf)
	c := GetCounter()
	enc.Encode(&c)
	dec := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&c)
	}
}

func BenchmarkGobUnmarshalBatchTimer(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	enc := gob.NewEncoder(buf)
	t := GetBatchTimer()
	enc.Encode(&t)
	dec := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&t)
	}
}

func BenchmarkGobUnmarshalGauge(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	enc := gob.NewEncoder(buf)
	g := GetGauge()
	enc.Encode(&g)
	dec := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&g)
	}
}

func BenchmarkGobUnmarshalMetricUnion(b *testing.B) {
	buf := bytes.NewBuffer(make([]byte, 0, 128))
	enc := gob.NewEncoder(buf)
	u := GetMetricUnion()
	enc.Encode(&u)
	dec := gob.NewDecoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec.Decode(&u)
	}
}

// encoding/json

func BenchmarkJSONUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := json.Marshal(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bts, &c)
	}
}

func BenchmarkJSONUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := json.Marshal(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bts, &t)
	}
}

func BenchmarkJSONUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := json.Marshal(g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bts, &g)
	}
}

func BenchmarkJSONUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := json.Marshal(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bts, &u)
	}
}

// github.com/mailru/easyjson

func BenchmarkEasyjsonUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := c.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.UnmarshalJSON(bts)
	}
}

func BenchmarkEasyjsonUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := t.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.UnmarshalJSON(bts)
	}
}

func BenchmarkEasyjsonUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := g.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.UnmarshalJSON(bts)
	}
}

func BenchmarkEasyjsonUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := u.MarshalJSON()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.UnmarshalJSON(bts)
	}
}

// labix.org/v2/mgo/bson

func BenchmarkBSONUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := bson.Marshal(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Unmarshal(bts, &c)
	}
}

func BenchmarkBSONUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := bson.Marshal(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Unmarshal(bts, &t)
	}
}

func BenchmarkBSONUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := bson.Marshal(g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Unmarshal(bts, &g)
	}
}

func BenchmarkBSONUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := bson.Marshal(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bson.Unmarshal(bts, &u)
	}
}

// github.com/davecgh/go-xdr/xdr2

func BenchmarkXDRUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := xdr.Marshal(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Unmarshal(bts, &c)
	}
}

func BenchmarkXDRUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := xdr.Marshal(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Unmarshal(bts, &t)
	}
}

func BenchmarkXDRUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := xdr.Marshal(g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Unmarshal(bts, &g)
	}
}

func BenchmarkXDRUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := xdr.Marshal(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		xdr.Unmarshal(bts, &u)
	}
}

// github.com/andyleap/gencode

func BenchmarkGencodeUnmarshalCounter(b *testing.B) {
	c := GetGencodeCounter()
	buf := make([]byte, c.Size())
	_, _ = c.Marshal(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Unmarshal(buf)
	}
}

func BenchmarGencodeUnmarshalBatchTimer(b *testing.B) {
	t := GetGencodeBatchTimer()
	buf := make([]byte, t.Size())
	_, _ = t.Marshal(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Unmarshal(buf)
	}
}

func BenchmarkGencodeUnmarshalGauge(b *testing.B) {
	g := GetGencodeGauge()
	buf := make([]byte, g.Size())
	_, _ = g.Marshal(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Unmarshal(buf)
	}
}

func BenchmarkGencodeUnmarshalMetricUnion(b *testing.B) {
	u := GetGencodeMetricUnion()
	buf := make([]byte, u.Size())
	_, _ = u.Marshal(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Unmarshal(buf)
	}
}

// github.com/pascaldekloe/colfer

func BenchmarkColferUnmarshalCounter(b *testing.B) {
	c := GetColferCounter()
	bts, _ := c.MarshalBinary()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Unmarshal(bts)
	}
}

func BenchmarkColferUnmarshalBatchTimer(b *testing.B) {
	t := GetColferBatchTimer()
	bts, _ := t.MarshalBinary()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Unmarshal(bts)
	}
}

func BenchmarkColferUnmarshalGauge(b *testing.B) {
	g := GetColferGauge()
	bts, _ := g.MarshalBinary()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Unmarshal(bts)
	}
}

func BenchmarkColferUnmarshalMetricUnion(b *testing.B) {
	u := GetColferMetricUnion()
	bts, _ := u.MarshalBinary()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Unmarshal(bts)
	}
}

// github.com/Sereal/Sereal/Go/sereal

func BenchmarkSerealUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	bts, _ := sereal.Marshal(c)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Unmarshal(bts, &c)
	}
}

func BenchmarkSerealUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	bts, _ := sereal.Marshal(t)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Unmarshal(bts, &t)
	}
}

func BenchmarkSerealUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	bts, _ := sereal.Marshal(g)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Unmarshal(bts, &g)
	}
}

func BenchmarkSerealUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	bts, _ := sereal.Marshal(u)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sereal.Unmarshal(bts, &u)
	}
}

// github.com/hprose/hprose-golang

func BenchmarkHproselUnmarshalCounter(b *testing.B) {
	c := GetCounter()
	buf := make([]byte, 0, 128)
	w := hprose.NewWriter(true, buf...)
	w.Serialize(c)
	buf = w.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := hprose.NewReader(buf, true)
		r.Unserialize(&c)
	}
}

func BenchmarkHproselUnmarshalBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	buf := make([]byte, 0, 128)
	w := hprose.NewWriter(true, buf...)
	w.Serialize(t)
	buf = w.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := hprose.NewReader(buf, true)
		r.Unserialize(&t)
	}
}

func BenchmarkHproselUnmarshalGauge(b *testing.B) {
	g := GetGauge()
	buf := make([]byte, 0, 128)
	w := hprose.NewWriter(true, buf...)
	w.Serialize(g)
	buf = w.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := hprose.NewReader(buf, true)
		r.Unserialize(&g)
	}
}

func BenchmarkHproselUnmarshalMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	buf := make([]byte, 0, 128)
	w := hprose.NewWriter(true, buf...)
	w.Serialize(u)
	buf = w.Bytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r := hprose.NewReader(buf, true)
		r.Unserialize(&u)
	}
}
