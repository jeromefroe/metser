// automatically generated by the FlatBuffers compiler, do not modify

package flatbuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatBuffersMetricUnion struct {
	_tab flatbuffers.Table
}

func GetRootAsFlatBuffersMetricUnion(buf []byte, offset flatbuffers.UOffsetT) *FlatBuffersMetricUnion {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlatBuffersMetricUnion{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FlatBuffersMetricUnion) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlatBuffersMetricUnion) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FlatBuffersMetricUnion) Type() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) MutateType(n int8) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *FlatBuffersMetricUnion) ID(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) IDLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) IDBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatBuffersMetricUnion) CounterVal() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) MutateCounterVal(n int64) bool {
	return rcv._tab.MutateInt64Slot(8, n)
}

func (rcv *FlatBuffersMetricUnion) BatchTimerVal(j int) float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetFloat64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) BatchTimerValLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatBuffersMetricUnion) GaugeVal() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *FlatBuffersMetricUnion) MutateGaugeVal(n float64) bool {
	return rcv._tab.MutateFloat64Slot(12, n)
}

func FlatBuffersMetricUnionStart(builder *flatbuffers.Builder) {
	builder.StartObject(5)
}
func FlatBuffersMetricUnionAddType(builder *flatbuffers.Builder, Type int8) {
	builder.PrependInt8Slot(0, Type, 0)
}
func FlatBuffersMetricUnionAddID(builder *flatbuffers.Builder, ID flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ID), 0)
}
func FlatBuffersMetricUnionStartIDVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func FlatBuffersMetricUnionAddCounterVal(builder *flatbuffers.Builder, CounterVal int64) {
	builder.PrependInt64Slot(2, CounterVal, 0)
}
func FlatBuffersMetricUnionAddBatchTimerVal(builder *flatbuffers.Builder, BatchTimerVal flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(BatchTimerVal), 0)
}
func FlatBuffersMetricUnionStartBatchTimerValVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func FlatBuffersMetricUnionAddGaugeVal(builder *flatbuffers.Builder, GaugeVal float64) {
	builder.PrependFloat64Slot(4, GaugeVal, 0.0)
}
func FlatBuffersMetricUnionEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
