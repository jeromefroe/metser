package metser

// AUTO GENERATED - DO NOT EDIT

import (
	math "math"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type CapnpType uint16

// Values of CapnpType.
const (
	CapnpType_unknownType    CapnpType = 0
	CapnpType_counterType    CapnpType = 1
	CapnpType_batchTimerType CapnpType = 2
	CapnpType_gaugeType      CapnpType = 3
)

// String returns the enum's constant name.
func (c CapnpType) String() string {
	switch c {
	case CapnpType_unknownType:
		return "unknownType"
	case CapnpType_counterType:
		return "counterType"
	case CapnpType_batchTimerType:
		return "batchTimerType"
	case CapnpType_gaugeType:
		return "gaugeType"

	default:
		return ""
	}
}

// CapnpTypeFromString returns the enum value with a name,
// or the zero value if there's no such value.
func CapnpTypeFromString(c string) CapnpType {
	switch c {
	case "unknownType":
		return CapnpType_unknownType
	case "counterType":
		return CapnpType_counterType
	case "batchTimerType":
		return CapnpType_batchTimerType
	case "gaugeType":
		return CapnpType_gaugeType

	default:
		return 0
	}
}

type CapnpType_List struct{ capnp.List }

func NewCapnpType_List(s *capnp.Segment, sz int32) (CapnpType_List, error) {
	l, err := capnp.NewUInt16List(s, sz)
	return CapnpType_List{l.List}, err
}

func (l CapnpType_List) At(i int) CapnpType {
	ul := capnp.UInt16List{List: l.List}
	return CapnpType(ul.At(i))
}

func (l CapnpType_List) Set(i int, v CapnpType) {
	ul := capnp.UInt16List{List: l.List}
	ul.Set(i, uint16(v))
}

type CapnpCounter struct{ capnp.Struct }

// CapnpCounter_TypeID is the unique identifier for the type CapnpCounter.
const CapnpCounter_TypeID = 0xd7b5a0d0acc1f200

func NewCapnpCounter(s *capnp.Segment) (CapnpCounter, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapnpCounter{st}, err
}

func NewRootCapnpCounter(s *capnp.Segment) (CapnpCounter, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapnpCounter{st}, err
}

func ReadRootCapnpCounter(msg *capnp.Message) (CapnpCounter, error) {
	root, err := msg.RootPtr()
	return CapnpCounter{root.Struct()}, err
}

func (s CapnpCounter) String() string {
	str, _ := text.Marshal(0xd7b5a0d0acc1f200, s.Struct)
	return str
}

func (s CapnpCounter) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CapnpCounter) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapnpCounter) SetId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s CapnpCounter) Value() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s CapnpCounter) SetValue(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

// CapnpCounter_List is a list of CapnpCounter.
type CapnpCounter_List struct{ capnp.List }

// NewCapnpCounter creates a new list of CapnpCounter.
func NewCapnpCounter_List(s *capnp.Segment, sz int32) (CapnpCounter_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CapnpCounter_List{l}, err
}

func (s CapnpCounter_List) At(i int) CapnpCounter { return CapnpCounter{s.List.Struct(i)} }

func (s CapnpCounter_List) Set(i int, v CapnpCounter) error { return s.List.SetStruct(i, v.Struct) }

// CapnpCounter_Promise is a wrapper for a CapnpCounter promised by a client call.
type CapnpCounter_Promise struct{ *capnp.Pipeline }

func (p CapnpCounter_Promise) Struct() (CapnpCounter, error) {
	s, err := p.Pipeline.Struct()
	return CapnpCounter{s}, err
}

type CapnpBatchTimer struct{ capnp.Struct }

// CapnpBatchTimer_TypeID is the unique identifier for the type CapnpBatchTimer.
const CapnpBatchTimer_TypeID = 0xc7e5aabb91f26222

func NewCapnpBatchTimer(s *capnp.Segment) (CapnpBatchTimer, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CapnpBatchTimer{st}, err
}

func NewRootCapnpBatchTimer(s *capnp.Segment) (CapnpBatchTimer, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return CapnpBatchTimer{st}, err
}

func ReadRootCapnpBatchTimer(msg *capnp.Message) (CapnpBatchTimer, error) {
	root, err := msg.RootPtr()
	return CapnpBatchTimer{root.Struct()}, err
}

func (s CapnpBatchTimer) String() string {
	str, _ := text.Marshal(0xc7e5aabb91f26222, s.Struct)
	return str
}

func (s CapnpBatchTimer) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CapnpBatchTimer) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapnpBatchTimer) SetId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s CapnpBatchTimer) Value() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Float64List{List: p.List()}, err
}

func (s CapnpBatchTimer) HasValue() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CapnpBatchTimer) SetValue(v capnp.Float64List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewValue sets the value field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s CapnpBatchTimer) NewValue(n int32) (capnp.Float64List, error) {
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// CapnpBatchTimer_List is a list of CapnpBatchTimer.
type CapnpBatchTimer_List struct{ capnp.List }

// NewCapnpBatchTimer creates a new list of CapnpBatchTimer.
func NewCapnpBatchTimer_List(s *capnp.Segment, sz int32) (CapnpBatchTimer_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return CapnpBatchTimer_List{l}, err
}

func (s CapnpBatchTimer_List) At(i int) CapnpBatchTimer { return CapnpBatchTimer{s.List.Struct(i)} }

func (s CapnpBatchTimer_List) Set(i int, v CapnpBatchTimer) error {
	return s.List.SetStruct(i, v.Struct)
}

// CapnpBatchTimer_Promise is a wrapper for a CapnpBatchTimer promised by a client call.
type CapnpBatchTimer_Promise struct{ *capnp.Pipeline }

func (p CapnpBatchTimer_Promise) Struct() (CapnpBatchTimer, error) {
	s, err := p.Pipeline.Struct()
	return CapnpBatchTimer{s}, err
}

type CapnpGauge struct{ capnp.Struct }

// CapnpGauge_TypeID is the unique identifier for the type CapnpGauge.
const CapnpGauge_TypeID = 0xab832e531cb9aecd

func NewCapnpGauge(s *capnp.Segment) (CapnpGauge, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapnpGauge{st}, err
}

func NewRootCapnpGauge(s *capnp.Segment) (CapnpGauge, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return CapnpGauge{st}, err
}

func ReadRootCapnpGauge(msg *capnp.Message) (CapnpGauge, error) {
	root, err := msg.RootPtr()
	return CapnpGauge{root.Struct()}, err
}

func (s CapnpGauge) String() string {
	str, _ := text.Marshal(0xab832e531cb9aecd, s.Struct)
	return str
}

func (s CapnpGauge) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CapnpGauge) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapnpGauge) SetId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s CapnpGauge) Value() float64 {
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s CapnpGauge) SetValue(v float64) {
	s.Struct.SetUint64(0, math.Float64bits(v))
}

// CapnpGauge_List is a list of CapnpGauge.
type CapnpGauge_List struct{ capnp.List }

// NewCapnpGauge creates a new list of CapnpGauge.
func NewCapnpGauge_List(s *capnp.Segment, sz int32) (CapnpGauge_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return CapnpGauge_List{l}, err
}

func (s CapnpGauge_List) At(i int) CapnpGauge { return CapnpGauge{s.List.Struct(i)} }

func (s CapnpGauge_List) Set(i int, v CapnpGauge) error { return s.List.SetStruct(i, v.Struct) }

// CapnpGauge_Promise is a wrapper for a CapnpGauge promised by a client call.
type CapnpGauge_Promise struct{ *capnp.Pipeline }

func (p CapnpGauge_Promise) Struct() (CapnpGauge, error) {
	s, err := p.Pipeline.Struct()
	return CapnpGauge{s}, err
}

type CapnpMetricUnion struct{ capnp.Struct }

// CapnpMetricUnion_TypeID is the unique identifier for the type CapnpMetricUnion.
const CapnpMetricUnion_TypeID = 0x89aa81e7fcc64df3

func NewCapnpMetricUnion(s *capnp.Segment) (CapnpMetricUnion, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return CapnpMetricUnion{st}, err
}

func NewRootCapnpMetricUnion(s *capnp.Segment) (CapnpMetricUnion, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return CapnpMetricUnion{st}, err
}

func ReadRootCapnpMetricUnion(msg *capnp.Message) (CapnpMetricUnion, error) {
	root, err := msg.RootPtr()
	return CapnpMetricUnion{root.Struct()}, err
}

func (s CapnpMetricUnion) String() string {
	str, _ := text.Marshal(0x89aa81e7fcc64df3, s.Struct)
	return str
}

func (s CapnpMetricUnion) Type() CapnpType {
	return CapnpType(s.Struct.Uint16(0))
}

func (s CapnpMetricUnion) SetType(v CapnpType) {
	s.Struct.SetUint16(0, uint16(v))
}

func (s CapnpMetricUnion) Id() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s CapnpMetricUnion) HasId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s CapnpMetricUnion) SetId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s CapnpMetricUnion) CounterVal() int64 {
	return int64(s.Struct.Uint64(8))
}

func (s CapnpMetricUnion) SetCounterVal(v int64) {
	s.Struct.SetUint64(8, uint64(v))
}

func (s CapnpMetricUnion) BatchTimerVal() (capnp.Float64List, error) {
	p, err := s.Struct.Ptr(1)
	return capnp.Float64List{List: p.List()}, err
}

func (s CapnpMetricUnion) HasBatchTimerVal() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s CapnpMetricUnion) SetBatchTimerVal(v capnp.Float64List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewBatchTimerVal sets the batchTimerVal field to a newly
// allocated capnp.Float64List, preferring placement in s's segment.
func (s CapnpMetricUnion) NewBatchTimerVal(n int32) (capnp.Float64List, error) {
	l, err := capnp.NewFloat64List(s.Struct.Segment(), n)
	if err != nil {
		return capnp.Float64List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

func (s CapnpMetricUnion) GaugeVal() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s CapnpMetricUnion) SetGaugeVal(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// CapnpMetricUnion_List is a list of CapnpMetricUnion.
type CapnpMetricUnion_List struct{ capnp.List }

// NewCapnpMetricUnion creates a new list of CapnpMetricUnion.
func NewCapnpMetricUnion_List(s *capnp.Segment, sz int32) (CapnpMetricUnion_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	return CapnpMetricUnion_List{l}, err
}

func (s CapnpMetricUnion_List) At(i int) CapnpMetricUnion { return CapnpMetricUnion{s.List.Struct(i)} }

func (s CapnpMetricUnion_List) Set(i int, v CapnpMetricUnion) error {
	return s.List.SetStruct(i, v.Struct)
}

// CapnpMetricUnion_Promise is a wrapper for a CapnpMetricUnion promised by a client call.
type CapnpMetricUnion_Promise struct{ *capnp.Pipeline }

func (p CapnpMetricUnion_Promise) Struct() (CapnpMetricUnion, error) {
	s, err := p.Pipeline.Struct()
	return CapnpMetricUnion{s}, err
}

const schema_c347c7e4235b72dc = "x\xda\xa4\x91OHT_\x14\xc7\xbf\xdf{f~\xfe" +
	"Hl\xba\xbd\xb7\x8ad\xd0\x0cRP,[\xb512" +
	"\x10\"\xc1[\x93\x8bZ=\xa7\x87\x0d\x8d\xcf\xc7k&" +
	"\x9b\xa8\x85\xa4\x90P\x90\x8b\xf6-\x0a\xc2\xfeP\x0b\x17" +
	"\xd2*\xa2\\\x09\x16D\x8bZ\x16\xedZ\xd8.\xa2\x17" +
	"\xf7\xcd\xe4\xc8\x8c\xb6i3s\xef\xe1\xf3\xce9\x9f\xef" +
	"\xed?\xc9\xa3\xea`:+\x80\xe9J\xff\x17\x7f\x1fy" +
	"\xf3\xf3\xeb\xcc\xe2<\x8cK\x89?E\xe7\xf6}^\x19" +
	"~\x85\xb4j\x01\x06\xca\xec\xa13G{\x9ca\x96`" +
	"\xbc\xfaty\xef\xe9\xbe\x1b\x8f,\xceM\xb8e\x9c'" +
	"\xea\x87\xb3l?t\x96\xd44\x18w\x8e\xaf/\xbcX" +
	"\xfc\xb2\x02\xed\xb2\xb1\xb5\x96N:\x1db\xe1v\xb1\xf0" +
	"\xee\x8f\xcf\xbf\xad\xbe[\\\x83vU\x1d\x06\x9dyY" +
	"w\xee&\xe0\x82L\x80\xbf\xd6_>^\xbb\xb7\xf4a" +
	"\x8b\x05\x06^\xcb\x0e:\xef\x13\xf6\xadL\xa37\xce{" +
	"a\x10\xf6N\xfaR\x8a\x0a\xf9K}\xc9\xf5\xc8\x90\xfd" +
	"\x1d\xf1m\xe9L&(L\x05\xa3\xa4q%\x05\xa4\x08" +
	"\xe8\xeb=\x80\xb9\"4\xb3\x8a\xa4K[\x9b\xd9\x03\x98" +
	"kBsSQ+\xbaT\x80\x9e;\x0b\x98Y\xa1\xb9" +
	"\xa3\xa8\x85.\x05\xd0\xb7#\xc0\xdc\x12\x9a\x87\x8a:\xa5" +
	"\\\xa6\x00\xfd\xe0\x04`\xee\x0b\xcd3\xc5L\xa9\x12\xfa" +
	"\xcc\xd4}Af@)\x9cg\x1b\x14\xdb\xc08?U" +
	"\x0eJ~4\x06\xf1\x8aLC1\x0d\xc6\xe3^)\x7f" +
	"!W\x98D\xd6\x8f\xc6\xbc\"w\x82\xa3B\xb6B\xd9" +
	"c<\xe1\x95'\xfc1\xaf\x08 \xa9\xb5\xda>5{" +
	"\xd5h?\x18\x0e[\xdcz\xff\xbf\xe1\xddm\x1d\xbb\x84" +
	"\xa6\xbf\xee\xdd{\x080\x07\x84\xe6\xb0\xda\xbca\xf6\xb2" +
	"W,\xfbMs\x9aS>V]\xbae\xd2\x8f\xb6\x1f" +
	"\xa6\x9b\xa6\x1d\xdfjZ\x83\xf0\xb6r\xd90W\x09\x13" +
	"\xb7]\xc93u\x8f\xdb\x88\xf5~\xfb\xa7t\xc7U\x80" +
	"\xa2\xdbO\x01q9\xb8\x18LM\x079\xb4TB\xff" +
	"O\xea\xb5\xdbF\xdc\x83~d\xfbU\x03\xceUB\xd0" +
	"\xff\x8b\xf0P\xb5\x09\xf0\xaf\xd9\xd6\x9e\xfdw\x00\x00\x00" +
	"\xff\xffC\xfe\xe7\xb7"

func init() {
	schemas.Register(schema_c347c7e4235b72dc,
		0x89aa81e7fcc64df3,
		0xab832e531cb9aecd,
		0xc7e5aabb91f26222,
		0xd0aad2cdeeb1db12,
		0xd7b5a0d0acc1f200)
}
