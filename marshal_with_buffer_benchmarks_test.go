package metser

import (
	"bytes"
	"testing"

	"gopkg.in/vmihailenco/msgpack.v2"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/google/flatbuffers/go"
	"github.com/ugorji/go/codec"
	"zombiezen.com/go/capnproto2"
)

// gopkg.in/vmihailenco/msgpack.v2

func BenchmarkVmihailencoMsgpackMarshalWithBufferCounter(b *testing.B) {
	c := GetCounter()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MsgpackMarshalCounter(enc, &c)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MsgpackMarshalBatchTimer(enc, &t)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithBufferGauge(b *testing.B) {
	g := GetGauge()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MsgpackMarshalGauge(enc, &g)
	}
}

func BenchmarkVmihailencoMsgpackMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MsgpackMarshalMetricUnion(enc, &u)
	}
}

// github.com/tinylib/msgp

func BenchmarkMsgpMarshalWithBufferCounter(b *testing.B) {
	c := GetCounter()
	buf := make([]byte, 0, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.MarshalMsg(buf)
	}
}

func BenchmarkMsgpMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetBatchTimer()
	buf := make([]byte, 0, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.MarshalMsg(buf)
	}
}

func BenchmarkMsgpMarshalWithBufferGauge(b *testing.B) {
	g := GetGauge()
	buf := make([]byte, 0, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.MarshalMsg(buf)
	}
}

func BenchmarkMsgpMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetMetricUnion()
	buf := make([]byte, 0, 256)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalMsg(buf)
	}
}

// github.com/ugorji/go/codec (messagepack)

func BenchmarkCodecMsgpackMarshalWithBufferCounter(b *testing.B) {
	var mh codec.MsgpackHandle
	buf := make([]byte, 0, 256)
	c := GetCounter()
	enc := codec.NewEncoderBytes(&buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(&c)
	}
}

func BenchmarkCodecMsgpackMarshalWithBufferBatchTimer(b *testing.B) {
	var mh codec.MsgpackHandle
	buf := make([]byte, 0, 256)
	t := GetBatchTimer()
	enc := codec.NewEncoderBytes(&buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(&t)
	}
}

func BenchmarkCodecMsgpackMarshalWithBufferGauge(b *testing.B) {
	var mh codec.MsgpackHandle
	buf := make([]byte, 0, 256)
	g := GetGauge()
	enc := codec.NewEncoderBytes(&buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(&g)
	}
}

func BenchmarkCodecMsgpackMarshalWithBufferMetricUnion(b *testing.B) {
	var mh codec.MsgpackHandle
	buf := make([]byte, 0, 256)
	u := GetMetricUnion()
	enc := codec.NewEncoderBytes(&buf, &mh)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc.Encode(&u)
	}
}

// github.com/google/flatbuffers/go

func BenchmarkFlatBuffersMarshalWithBufferCounter(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	c := GetCounter()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.MarshalCounter(c)
	}
}

func BenchmarkFlatBuffersMarshalWithBufferBatchTimer(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	t := GetBatchTimer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.MarshalBatchTimer(t)
	}
}

func BenchmarkFlatBuffersMarshalWithBufferGauge(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	g := GetGauge()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.MarshalGauge(g)
	}
}

func BenchmarkFlatBuffersMarshalWithBufferMetricUnion(b *testing.B) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	u := GetMetricUnion()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builder.MarshalMetricUnion(u)
	}
}

// github.com/gogo/protobuf

func BenchmarkGoGoProtobufMarshalWithBufferCounter(b *testing.B) {
	c := GetGoGoProtobufCounter()
	buf := make([]byte, c.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.MarshalTo(buf)
	}
}

func BenchmarkGoGoProtobufMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetGoGoProtobufBatchTimer()
	buf := make([]byte, t.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.MarshalTo(buf)
	}
}

func BenchmarkGoGoProtobufMarshalWithBufferGauge(b *testing.B) {
	g := GetGoGoProtobufGauge()
	buf := make([]byte, g.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.MarshalTo(buf)
	}
}

func BenchmarkGoGoProtobufMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetGoGoProtobufMetricUnion()
	buf := make([]byte, u.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.MarshalTo(buf)
	}
}

// zombiezen.com/go/capnproto2

func BenchmarkCapnpMarshalWithBufferCounter(b *testing.B) {
	m := GetCapnpCounterMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(m)
	}
}

func BenchmarkCapnpMarshalWithBufferBatchTimer(b *testing.B) {
	m := GetCapnpBatchTimerMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(m)
	}
}

func BenchmarkCapnpMarshalWithBufferGauge(b *testing.B) {
	m := GetCapnpGaugeMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(m)
	}
}

func BenchmarkCapnpfMarshalWithBufferMetricUnion(b *testing.B) {
	m := GetCapnpCounterMessage()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := capnp.NewEncoder(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		enc.Encode(m)
	}
}

// git.apache.org/thrift.git/lib/go/thrift

func BenchmarkThriftMarshalWithBufferCounter(b *testing.B) {
	c := GetThriftCounter()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Write(proto)
	}
}

func BenchmarkThriftMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetThriftBatchTimer()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Write(proto)
	}
}

func BenchmarkThriftMarshalWithBufferGauge(b *testing.B) {
	g := GetThriftGauge()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Write(proto)
	}
}

func BenchmarkThriftMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetThriftMetricUnion()
	transport := thrift.NewTMemoryBufferLen(256)
	proto := thrift.NewTBinaryProtocolTransport(transport)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Write(proto)
	}
}

// github.com/andyleap/gencode

func BenchmarkGencodeMarshalWithBufferCounter(b *testing.B) {
	c := GetGencodeCounter()
	buf := make([]byte, c.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Marshal(buf)
	}
}

func BenchmarkGencodeMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetGencodeBatchTimer()
	buf := make([]byte, t.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		t.Marshal(buf)
	}
}

func BenchmarkGencodeMarshalWithBufferGauge(b *testing.B) {
	g := GetGencodeGauge()
	buf := make([]byte, g.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Marshal(buf)
	}
}

func BenchmarkGencodeMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetGencodeMetricUnion()
	buf := make([]byte, u.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		u.Marshal(buf)
	}
}

// github.com/pascaldekloe/colfer
//
// NB(jeromefroe): here we check the length of the buffer to simulate real world use
// because MarshalTo will panic if it is not long enough

func BenchmarkColferMarshalWithBufferCounter(b *testing.B) {
	c := GetColferCounter()
	n, err := c.MarshalLen()
	if err != nil {
		b.Fatal(err)
	}
	buf := make([]byte, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n, err := c.MarshalLen()
		if err != nil || n < len(buf) {
			b.Fatal(err)
		}
		c.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithBufferBatchTimer(b *testing.B) {
	t := GetColferBatchTimer()
	n, err := t.MarshalLen()
	if err != nil {
		b.Fatal(err)
	}
	buf := make([]byte, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n, err := t.MarshalLen()
		if err != nil || n < len(buf) {
			b.Fatal(err)
		}
		t.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithBufferGauge(b *testing.B) {
	g := GetColferGauge()
	n, err := g.MarshalLen()
	if err != nil {
		b.Fatal(err)
	}
	buf := make([]byte, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n, err := g.MarshalLen()
		if err != nil || n < len(buf) {
			b.Fatal(err)
		}
		g.MarshalTo(buf)
	}
}

func BenchmarkColferMarshalWithBufferMetricUnion(b *testing.B) {
	u := GetColferMetricUnion()
	n, err := u.MarshalLen()
	if err != nil {
		b.Fatal(err)
	}
	buf := make([]byte, n)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		n, err := u.MarshalLen()
		if err != nil || n < len(buf) {
			b.Fatal(err)
		}
		u.MarshalTo(buf)
	}
}
