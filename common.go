package metser

import (
	"github.com/google/flatbuffers/go"
	"github.com/jeromefroe/metser/colfer"
	fb "github.com/jeromefroe/metser/flatbuffers"
	th "github.com/jeromefroe/metser/gen-go/thrift"
	"gopkg.in/vmihailenco/msgpack.v2"
	"zombiezen.com/go/capnproto2"
)

func GetID() ID {
	return []byte("name=bulbasaur,trainer=ash,type=grass")
}

func GetCounter() Counter {
	id := GetID()
	return Counter{
		ID:    id,
		Value: 100,
	}
}

func GetColferCounter() colfer.Counter {
	id := GetID()
	return colfer.Counter{
		ID:    id,
		Value: 100,
	}
}

func GetProtobufCounter() ProtobufCounter {
	id := GetID()
	return ProtobufCounter{
		ID:    id,
		Value: 100,
	}
}

func GetGoGoProtobufCounter() GoGoProtobufCounter {
	id := GetID()
	return GoGoProtobufCounter{
		ID:    id,
		Value: 100,
	}
}

func GetGencodeCounter() GencodeCounter {
	id := GetID()
	return GencodeCounter{
		ID:    id,
		Value: 100,
	}
}

func GetThriftCounter() th.ThriftCounter {
	id := GetID()
	return th.ThriftCounter{
		ID:    id,
		Value: 100,
	}
}

func GetCapnpCounterMessage() *capnp.Message {
	id := GetID()
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	c, err := NewRootCapnpCounter(seg)
	if err != nil {
		panic(err)
	}

	c.SetId(id)
	c.SetValue(100)

	return msg
}

func GetBatchTimer() BatchTimer {
	id := GetID()
	return BatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetColferBatchTimer() colfer.BatchTimer {
	id := GetID()
	return colfer.BatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetProtobufBatchTimer() ProtobufBatchTimer {
	id := GetID()
	return ProtobufBatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetGoGoProtobufBatchTimer() GoGoProtobufBatchTimer {
	id := GetID()
	return GoGoProtobufBatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetGencodeBatchTimer() GencodeBatchTimer {
	id := GetID()
	return GencodeBatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetThriftBatchTimer() th.ThriftBatchTimer {
	id := GetID()
	return th.ThriftBatchTimer{
		ID:     id,
		Values: []float64{1, 2, 3, 4, 5},
	}
}

func GetCapnpBatchTimerMessage() *capnp.Message {
	id := GetID()
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	t, err := NewRootCapnpBatchTimer(seg)
	if err != nil {
		panic(err)
	}

	t.SetId(id)
	values, err := t.NewValue(5)
	if err != nil {
		panic(err)
	}

	values.Set(0, 1)
	values.Set(1, 2)
	values.Set(2, 3)
	values.Set(3, 4)
	values.Set(4, 5)

	return msg
}

func GetGauge() Gauge {
	id := GetID()
	return Gauge{
		ID:    id,
		Value: 100,
	}
}

func GetColferGauge() colfer.Gauge {
	id := GetID()
	return colfer.Gauge{
		ID:    id,
		Value: 100,
	}
}

func GetProtobufGauge() ProtobufGauge {
	id := GetID()
	return ProtobufGauge{
		ID:    id,
		Value: 100,
	}
}

func GetGoGoProtobufGauge() GoGoProtobufGauge {
	id := GetID()
	return GoGoProtobufGauge{
		ID:    id,
		Value: 100,
	}
}

func GetGencodeGauge() GencodeGauge {
	id := GetID()
	return GencodeGauge{
		ID:    id,
		Value: 100,
	}
}

func GetThriftGauge() th.ThriftGauge {
	id := GetID()
	return th.ThriftGauge{
		ID:    id,
		Value: 100,
	}
}

func GetCapnpGaugeMessage() *capnp.Message {
	id := GetID()
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	g, err := NewRootCapnpGauge(seg)
	if err != nil {
		panic(err)
	}

	g.SetId(id)
	g.SetValue(100)

	return msg
}

func GetMetricUnion() MetricUnion {
	id := GetID()
	return MetricUnion{
		Type:          BatchTimerType,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetColferMetricUnion() colfer.MetricUnion {
	id := GetID()
	return colfer.MetricUnion{
		Type:          2,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetProtobufMetricUnion() ProtobufMetricUnion {
	id := GetID()
	return ProtobufMetricUnion{
		Type:          ProtobufType_BatchTimerType,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetGoGoProtobufMetricUnion() GoGoProtobufMetricUnion {
	id := GetID()
	return GoGoProtobufMetricUnion{
		Type:          GoGoProtobufType_BatchTimerType,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetGencodeMetricUnion() GencodeMetricUnion {
	id := GetID()
	return GencodeMetricUnion{
		Type:          2,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetThriftMetricUnion() th.ThriftMetricUnion {
	id := GetID()
	return th.ThriftMetricUnion{
		Type:          2,
		ID:            id,
		CounterVal:    0,
		BatchTimerVal: []float64{1, 2, 3, 4, 5},
		GaugeVal:      0,
	}
}

func GetCapnpMetricUnionMessage() *capnp.Message {
	id := GetID()
	msg, seg, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		panic(err)
	}

	u, err := NewRootCapnpMetricUnion(seg)
	if err != nil {
		panic(err)
	}

	u.SetType(CapnpType_batchTimerType)
	u.SetId(id)
	u.SetCounterVal(0)
	values, err := u.NewBatchTimerVal(5)
	if err != nil {
		panic(err)
	}

	values.Set(0, 1)
	values.Set(1, 2)
	values.Set(2, 3)
	values.Set(3, 4)
	values.Set(4, 5)

	u.SetGaugeVal(0)

	return msg
}

// gopkg.in/vmihailenco/msgpack.v2 marshaling and unmarshaling methods

func MsgpackMarshalCounter(enc *msgpack.Encoder, c *Counter) {
	enc.EncodeBytes(c.ID)
	enc.EncodeInt64(c.Value)
}

func MsgpackUnmarshalCounter(dec *msgpack.Decoder, c *Counter) error {
	id, err := dec.DecodeBytes()
	if err != nil {
		return err
	}

	v, err := dec.DecodeInt64()
	if err != nil {
		return err
	}

	c.ID = id
	c.Value = v
	return nil
}

func MsgpackMarshalBatchTimer(enc *msgpack.Encoder, t *BatchTimer) {
	enc.EncodeBytes(t.ID)
	enc.EncodeArrayLen(len(t.Values))
	for _, v := range t.Values {
		enc.EncodeFloat64(v)
	}
}

func MsgpackUnmarshalBatchTimer(dec *msgpack.Decoder, t *BatchTimer) error {
	id, err := dec.DecodeBytes()
	if err != nil {
		return err
	}
	t.ID = id

	l, err := dec.DecodeArrayLen()
	if err != nil {
		return err
	}

	if cap(t.Values) < l {
		t.Values = make([]float64, 0, l)
	} else {
		t.Values = t.Values[:0]
	}

	for i := 0; i < l; i++ {
		v, err := dec.DecodeFloat64()
		if err != nil {
			return err
		}
		t.Values = append(t.Values, v)
	}

	return nil
}

func MsgpackMarshalGauge(enc *msgpack.Encoder, g *Gauge) {
	enc.EncodeBytes(g.ID)
	enc.EncodeFloat64(g.Value)
}

func MsgpackUnmarshalGauge(dec *msgpack.Decoder, g *Gauge) error {
	id, err := dec.DecodeBytes()
	if err != nil {
		return err
	}

	v, err := dec.DecodeFloat64()
	if err != nil {
		return err
	}

	g.ID = id
	g.Value = v
	return nil
}

func MsgpackMarshalMetricUnion(enc *msgpack.Encoder, u *MetricUnion) {
	enc.EncodeInt8(int8(u.Type))
	enc.EncodeBytes(u.ID)
	enc.EncodeInt64(u.CounterVal)
	enc.EncodeArrayLen(len(u.BatchTimerVal))
	for _, v := range u.BatchTimerVal {
		enc.EncodeFloat64(v)
	}
	enc.EncodeFloat64(u.GaugeVal)
}

func MsgpackUnmarshalMetricUnion(dec *msgpack.Decoder, u *MetricUnion) error {
	t, err := dec.DecodeInt8()
	if err != nil {
		return err
	}
	u.Type = Type(t)

	id, err := dec.DecodeBytes()
	if err != nil {
		return err
	}
	u.ID = id

	cv, err := dec.DecodeInt64()
	if err != nil {
		return err
	}
	u.CounterVal = cv

	l, err := dec.DecodeArrayLen()
	if err != nil {
		return err
	}

	if cap(u.BatchTimerVal) < l {
		u.BatchTimerVal = make([]float64, 0, l)
	} else {
		u.BatchTimerVal = u.BatchTimerVal[:0]
	}

	for i := 0; i < l; i++ {
		v, err := dec.DecodeFloat64()
		if err != nil {
			return err
		}
		u.BatchTimerVal = append(u.BatchTimerVal, v)
	}

	gv, err := dec.DecodeFloat64()
	if err != nil {
		return err
	}
	u.GaugeVal = gv

	return nil
}

// flatbuffers marshaling and umarshaling methods

type FlatBuffersSerializer struct {
	Builder *flatbuffers.Builder
}

func (s *FlatBuffersSerializer) MarshalCounter(c Counter) []byte {
	s.Builder.Reset()

	fb.FlatBuffersCounterStartIDVector(s.Builder, len(c.ID))
	for i := len(c.ID) - 1; i >= 0; i-- {
		s.Builder.PrependByte(c.ID[i])
	}
	id := s.Builder.EndVector(len(c.ID))

	fb.FlatBuffersCounterStart(s.Builder)
	fb.FlatBuffersCounterAddID(s.Builder, id)
	fb.FlatBuffersCounterAddValue(s.Builder, c.Value)

	s.Builder.Finish(fb.FlatBuffersCounterEnd(s.Builder))
	return s.Builder.FinishedBytes()
}

func (s *FlatBuffersSerializer) UnmarshalCounter(d []byte, c *Counter) error {
	o := fb.FlatBuffersCounter{}
	o.Init(d, flatbuffers.GetUOffsetT(d))

	c.ID = c.ID[:0]
	l := o.IDLength()
	if cap(c.ID) < l {
		c.ID = make([]byte, 0, l)
	}
	for i := 0; i < l; i++ {
		c.ID = append(c.ID, o.ID(i))
	}
	c.Value = o.Value()

	return nil
}

func (s *FlatBuffersSerializer) MarshalBatchTimer(t BatchTimer) []byte {
	s.Builder.Reset()

	fb.FlatBuffersBatchTimerStartIDVector(s.Builder, len(t.ID))
	for i := len(t.ID) - 1; i >= 0; i-- {
		s.Builder.PrependByte(t.ID[i])
	}
	id := s.Builder.EndVector(len(t.ID))

	fb.FlatBuffersBatchTimerStartValuesVector(s.Builder, len(t.Values))
	for i := len(t.Values) - 1; i >= 0; i-- {
		s.Builder.PrependFloat64(t.Values[i])
	}
	values := s.Builder.EndVector(len(t.Values))

	fb.FlatBuffersBatchTimerStart(s.Builder)
	fb.FlatBuffersBatchTimerAddID(s.Builder, id)
	fb.FlatBuffersBatchTimerAddValues(s.Builder, values)

	s.Builder.Finish(fb.FlatBuffersBatchTimerEnd(s.Builder))
	return s.Builder.FinishedBytes()
}

func (s *FlatBuffersSerializer) UnmarshalBatchTimer(d []byte, t *BatchTimer) error {
	o := fb.FlatBuffersBatchTimer{}
	o.Init(d, flatbuffers.GetUOffsetT(d))

	t.ID = t.ID[:0]
	l := o.IDLength()
	if cap(t.ID) < l {
		t.ID = make([]byte, 0, l)
	}
	for i := 0; i < l; i++ {
		t.ID = append(t.ID, o.ID(i))
	}

	t.Values = t.Values[:0]
	l = o.ValuesLength()
	if cap(t.Values) < l {
		t.Values = make([]float64, 0, l)
	}
	for i := 0; i < l; i++ {
		t.Values = append(t.Values, o.Values(i))
	}

	return nil
}

func (s *FlatBuffersSerializer) MarshalGauge(g Gauge) []byte {
	s.Builder.Reset()

	fb.FlatBuffersCounterStartIDVector(s.Builder, len(g.ID))
	for i := len(g.ID) - 1; i >= 0; i-- {
		s.Builder.PrependByte(g.ID[i])
	}
	id := s.Builder.EndVector(len(g.ID))

	fb.FlatBuffersGaugeStart(s.Builder)
	fb.FlatBuffersGaugeAddID(s.Builder, id)
	fb.FlatBuffersGaugeAddValue(s.Builder, g.Value)

	s.Builder.Finish(fb.FlatBuffersGaugeEnd(s.Builder))
	return s.Builder.FinishedBytes()
}

func (s *FlatBuffersSerializer) UnmarshalGauge(d []byte, g *Gauge) error {
	o := fb.FlatBuffersGauge{}
	o.Init(d, flatbuffers.GetUOffsetT(d))

	g.ID = g.ID[:0]
	l := o.IDLength()
	if cap(g.ID) < l {
		g.ID = make([]byte, 0, l)
	}
	for i := 0; i < l; i++ {
		g.ID = append(g.ID, o.ID(i))
	}
	g.Value = o.Value()

	return nil
}

func (s *FlatBuffersSerializer) UnmarshalMetricUnion(d []byte, u *MetricUnion) error {
	o := fb.FlatBuffersMetricUnion{}
	o.Init(d, flatbuffers.GetUOffsetT(d))

	u.Type = Type(o.Type())

	u.ID = u.ID[:0]
	l := o.IDLength()
	if cap(u.ID) < l {
		u.ID = make([]byte, 0, l)
	}
	for i := 0; i < l; i++ {
		u.ID = append(u.ID, o.ID(i))
	}

	u.CounterVal = o.CounterVal()

	u.BatchTimerVal = u.BatchTimerVal[:0]
	l = o.BatchTimerValLength()
	if cap(u.BatchTimerVal) < l {
		u.BatchTimerVal = make([]float64, 0, l)
	}
	for i := 0; i < l; i++ {
		u.BatchTimerVal = append(u.BatchTimerVal, o.BatchTimerVal(i))
	}

	u.GaugeVal = o.GaugeVal()

	return nil
}

func (s *FlatBuffersSerializer) MarshalMetricUnion(u MetricUnion) []byte {
	s.Builder.Reset()

	fb.FlatBuffersMetricUnionStartIDVector(s.Builder, len(u.ID))
	for i := len(u.ID) - 1; i >= 0; i-- {
		s.Builder.PrependByte(u.ID[i])
	}
	id := s.Builder.EndVector(len(u.ID))

	fb.FlatBuffersMetricUnionStartBatchTimerValVector(s.Builder, len(u.BatchTimerVal))
	for i := len(u.BatchTimerVal) - 1; i >= 0; i-- {
		s.Builder.PrependFloat64(u.BatchTimerVal[i])
	}
	batchTimerVal := s.Builder.EndVector(len(u.BatchTimerVal))

	fb.FlatBuffersMetricUnionStart(s.Builder)
	fb.FlatBuffersMetricUnionAddType(s.Builder, int8(u.Type))
	fb.FlatBuffersMetricUnionAddID(s.Builder, id)
	fb.FlatBuffersMetricUnionAddCounterVal(s.Builder, u.CounterVal)
	fb.FlatBuffersMetricUnionAddBatchTimerVal(s.Builder, batchTimerVal)
	fb.FlatBuffersMetricUnionAddGaugeVal(s.Builder, u.GaugeVal)

	s.Builder.Finish(fb.FlatBuffersMetricUnionEnd(s.Builder))
	return s.Builder.FinishedBytes()
}
