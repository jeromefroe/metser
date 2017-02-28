package metser

import (
	"bytes"
	"testing"

	"gopkg.in/vmihailenco/msgpack.v2"

	"github.com/google/flatbuffers/go"
	"github.com/stretchr/testify/assert"
)

// msgpack

func TestMsgpackCounter(t *testing.T) {
	c := GetCounter()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalCounter(enc, &c)
	c2 := Counter{}
	dec := msgpack.NewDecoder(buf)
	err := MsgpackUnmarshalCounter(dec, &c2)
	assert.NoError(t, err)
	assert.Equal(t, c, c2)
}

func TestMsgpackBatchTimer(t *testing.T) {
	bt := GetBatchTimer()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalBatchTimer(enc, &bt)
	bt2 := BatchTimer{}
	dec := msgpack.NewDecoder(buf)
	err := MsgpackUnmarshalBatchTimer(dec, &bt2)
	assert.NoError(t, err)
	assert.Equal(t, bt, bt2)
}

func TestMsgpackGauge(t *testing.T) {
	g := GetGauge()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalGauge(enc, &g)
	g2 := Gauge{}
	dec := msgpack.NewDecoder(buf)
	err := MsgpackUnmarshalGauge(dec, &g2)
	assert.NoError(t, err)
	assert.Equal(t, g, g2)
}

func TestMsgpackMetricUnion(t *testing.T) {
	u := GetMetricUnion()
	buf := bytes.NewBuffer(make([]byte, 0, 256))
	enc := msgpack.NewEncoder(buf)
	MsgpackMarshalMetricUnion(enc, &u)
	u2 := MetricUnion{}
	dec := msgpack.NewDecoder(buf)
	err := MsgpackUnmarshalMetricUnion(dec, &u2)
	assert.NoError(t, err)
	assert.Equal(t, u, u2)
}

// flatbuffers

func TestFlatBuffersCounter(t *testing.T) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	c := GetCounter()
	buf := builder.MarshalCounter(c)
	c2 := Counter{}
	builder.UnmarshalCounter(buf, &c2)
	assert.Equal(t, c, c2)
}

func TestFlatBuffersBatchTimer(t *testing.T) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	bt := GetBatchTimer()
	buf := builder.MarshalBatchTimer(bt)
	bt2 := BatchTimer{}
	builder.UnmarshalBatchTimer(buf, &bt2)
	assert.Equal(t, bt, bt2)
}

func TestFlatBuffersGauge(t *testing.T) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	g := GetGauge()
	buf := builder.MarshalGauge(g)
	g2 := Gauge{}
	builder.UnmarshalGauge(buf, &g2)
	assert.Equal(t, g, g2)
}

func TestFlatBuffersMetricUnion(t *testing.T) {
	builder := FlatBuffersSerializer{flatbuffers.NewBuilder(256)}
	u := GetMetricUnion()
	buf := builder.MarshalMetricUnion(u)
	u2 := MetricUnion{}
	builder.UnmarshalMetricUnion(buf, &u2)
	assert.Equal(t, u, u2)
}
