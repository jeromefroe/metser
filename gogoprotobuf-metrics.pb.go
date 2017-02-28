// Code generated by protoc-gen-gogo.
// source: github.com/jeromefroe/metser/gogoprotobuf-metrics.proto
// DO NOT EDIT!

/*
	Package metser is a generated protocol buffer package.

	It is generated from these files:
		github.com/jeromefroe/metser/gogoprotobuf-metrics.proto

	It has these top-level messages:
		GoGoProtobufCounter
		GoGoProtobufBatchTimer
		GoGoProtobufGauge
		GoGoProtobufMetricUnion
*/
package metser

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GoGoProtobufType int32

const (
	GoGoProtobufType_UnknownType    GoGoProtobufType = 0
	GoGoProtobufType_CounterType    GoGoProtobufType = 1
	GoGoProtobufType_BatchTimerType GoGoProtobufType = 2
	GoGoProtobufType_GaugeType      GoGoProtobufType = 3
)

var GoGoProtobufType_name = map[int32]string{
	0: "UnknownType",
	1: "CounterType",
	2: "BatchTimerType",
	3: "GaugeType",
}
var GoGoProtobufType_value = map[string]int32{
	"UnknownType":    0,
	"CounterType":    1,
	"BatchTimerType": 2,
	"GaugeType":      3,
}

func (x GoGoProtobufType) String() string {
	return proto.EnumName(GoGoProtobufType_name, int32(x))
}
func (GoGoProtobufType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorGogoprotobufMetrics, []int{0}
}

type GoGoProtobufCounter struct {
	ID    []byte `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (m *GoGoProtobufCounter) Reset()         { *m = GoGoProtobufCounter{} }
func (m *GoGoProtobufCounter) String() string { return proto.CompactTextString(m) }
func (*GoGoProtobufCounter) ProtoMessage()    {}
func (*GoGoProtobufCounter) Descriptor() ([]byte, []int) {
	return fileDescriptorGogoprotobufMetrics, []int{0}
}

func (m *GoGoProtobufCounter) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *GoGoProtobufCounter) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GoGoProtobufBatchTimer struct {
	ID     []byte    `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Values []float64 `protobuf:"fixed64,2,rep,packed,name=Values,json=values" json:"Values,omitempty"`
}

func (m *GoGoProtobufBatchTimer) Reset()         { *m = GoGoProtobufBatchTimer{} }
func (m *GoGoProtobufBatchTimer) String() string { return proto.CompactTextString(m) }
func (*GoGoProtobufBatchTimer) ProtoMessage()    {}
func (*GoGoProtobufBatchTimer) Descriptor() ([]byte, []int) {
	return fileDescriptorGogoprotobufMetrics, []int{1}
}

func (m *GoGoProtobufBatchTimer) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *GoGoProtobufBatchTimer) GetValues() []float64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type GoGoProtobufGauge struct {
	ID    []byte  `protobuf:"bytes,1,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	Value float64 `protobuf:"fixed64,2,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (m *GoGoProtobufGauge) Reset()         { *m = GoGoProtobufGauge{} }
func (m *GoGoProtobufGauge) String() string { return proto.CompactTextString(m) }
func (*GoGoProtobufGauge) ProtoMessage()    {}
func (*GoGoProtobufGauge) Descriptor() ([]byte, []int) {
	return fileDescriptorGogoprotobufMetrics, []int{2}
}

func (m *GoGoProtobufGauge) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *GoGoProtobufGauge) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type GoGoProtobufMetricUnion struct {
	Type          GoGoProtobufType `protobuf:"varint,1,opt,name=Type,json=type,proto3,enum=metser.GoGoProtobufType" json:"Type,omitempty"`
	ID            []byte           `protobuf:"bytes,2,opt,name=ID,json=iD,proto3" json:"ID,omitempty"`
	CounterVal    int64            `protobuf:"varint,3,opt,name=CounterVal,json=counterVal,proto3" json:"CounterVal,omitempty"`
	BatchTimerVal []float64        `protobuf:"fixed64,4,rep,packed,name=BatchTimerVal,json=batchTimerVal" json:"BatchTimerVal,omitempty"`
	GaugeVal      float64          `protobuf:"fixed64,5,opt,name=GaugeVal,json=gaugeVal,proto3" json:"GaugeVal,omitempty"`
}

func (m *GoGoProtobufMetricUnion) Reset()         { *m = GoGoProtobufMetricUnion{} }
func (m *GoGoProtobufMetricUnion) String() string { return proto.CompactTextString(m) }
func (*GoGoProtobufMetricUnion) ProtoMessage()    {}
func (*GoGoProtobufMetricUnion) Descriptor() ([]byte, []int) {
	return fileDescriptorGogoprotobufMetrics, []int{3}
}

func (m *GoGoProtobufMetricUnion) GetType() GoGoProtobufType {
	if m != nil {
		return m.Type
	}
	return GoGoProtobufType_UnknownType
}

func (m *GoGoProtobufMetricUnion) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *GoGoProtobufMetricUnion) GetCounterVal() int64 {
	if m != nil {
		return m.CounterVal
	}
	return 0
}

func (m *GoGoProtobufMetricUnion) GetBatchTimerVal() []float64 {
	if m != nil {
		return m.BatchTimerVal
	}
	return nil
}

func (m *GoGoProtobufMetricUnion) GetGaugeVal() float64 {
	if m != nil {
		return m.GaugeVal
	}
	return 0
}

func init() {
	proto.RegisterType((*GoGoProtobufCounter)(nil), "metser.GoGoProtobufCounter")
	proto.RegisterType((*GoGoProtobufBatchTimer)(nil), "metser.GoGoProtobufBatchTimer")
	proto.RegisterType((*GoGoProtobufGauge)(nil), "metser.GoGoProtobufGauge")
	proto.RegisterType((*GoGoProtobufMetricUnion)(nil), "metser.GoGoProtobufMetricUnion")
	proto.RegisterEnum("metser.GoGoProtobufType", GoGoProtobufType_name, GoGoProtobufType_value)
}
func (m *GoGoProtobufCounter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoProtobufCounter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Value != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *GoGoProtobufBatchTimer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoProtobufBatchTimer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if len(m.Values) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.Values)*8))
		for _, num := range m.Values {
			f1 := math.Float64bits(float64(num))
			dAtA[i] = uint8(f1)
			i++
			dAtA[i] = uint8(f1 >> 8)
			i++
			dAtA[i] = uint8(f1 >> 16)
			i++
			dAtA[i] = uint8(f1 >> 24)
			i++
			dAtA[i] = uint8(f1 >> 32)
			i++
			dAtA[i] = uint8(f1 >> 40)
			i++
			dAtA[i] = uint8(f1 >> 48)
			i++
			dAtA[i] = uint8(f1 >> 56)
			i++
		}
	}
	return i, nil
}

func (m *GoGoProtobufGauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoProtobufGauge) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.Value != 0 {
		dAtA[i] = 0x11
		i++
		i = encodeFixed64GogoprotobufMetrics(dAtA, i, uint64(math.Float64bits(float64(m.Value))))
	}
	return i, nil
}

func (m *GoGoProtobufMetricUnion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GoGoProtobufMetricUnion) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(m.Type))
	}
	if len(m.ID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.CounterVal != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(m.CounterVal))
	}
	if len(m.BatchTimerVal) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintGogoprotobufMetrics(dAtA, i, uint64(len(m.BatchTimerVal)*8))
		for _, num := range m.BatchTimerVal {
			f2 := math.Float64bits(float64(num))
			dAtA[i] = uint8(f2)
			i++
			dAtA[i] = uint8(f2 >> 8)
			i++
			dAtA[i] = uint8(f2 >> 16)
			i++
			dAtA[i] = uint8(f2 >> 24)
			i++
			dAtA[i] = uint8(f2 >> 32)
			i++
			dAtA[i] = uint8(f2 >> 40)
			i++
			dAtA[i] = uint8(f2 >> 48)
			i++
			dAtA[i] = uint8(f2 >> 56)
			i++
		}
	}
	if m.GaugeVal != 0 {
		dAtA[i] = 0x29
		i++
		i = encodeFixed64GogoprotobufMetrics(dAtA, i, uint64(math.Float64bits(float64(m.GaugeVal))))
	}
	return i, nil
}

func encodeFixed64GogoprotobufMetrics(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32GogoprotobufMetrics(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGogoprotobufMetrics(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *GoGoProtobufCounter) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGogoprotobufMetrics(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovGogoprotobufMetrics(uint64(m.Value))
	}
	return n
}

func (m *GoGoProtobufBatchTimer) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGogoprotobufMetrics(uint64(l))
	}
	if len(m.Values) > 0 {
		n += 1 + sovGogoprotobufMetrics(uint64(len(m.Values)*8)) + len(m.Values)*8
	}
	return n
}

func (m *GoGoProtobufGauge) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGogoprotobufMetrics(uint64(l))
	}
	if m.Value != 0 {
		n += 9
	}
	return n
}

func (m *GoGoProtobufMetricUnion) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovGogoprotobufMetrics(uint64(m.Type))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovGogoprotobufMetrics(uint64(l))
	}
	if m.CounterVal != 0 {
		n += 1 + sovGogoprotobufMetrics(uint64(m.CounterVal))
	}
	if len(m.BatchTimerVal) > 0 {
		n += 1 + sovGogoprotobufMetrics(uint64(len(m.BatchTimerVal)*8)) + len(m.BatchTimerVal)*8
	}
	if m.GaugeVal != 0 {
		n += 9
	}
	return n
}

func sovGogoprotobufMetrics(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGogoprotobufMetrics(x uint64) (n int) {
	return sovGogoprotobufMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GoGoProtobufCounter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogoprotobufMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoGoProtobufCounter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoProtobufCounter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGogoprotobufMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoGoProtobufBatchTimer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogoprotobufMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoGoProtobufBatchTimer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoProtobufBatchTimer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				v2 := float64(math.Float64frombits(v))
				m.Values = append(m.Values, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGogoprotobufMetrics
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGogoprotobufMetrics
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					v2 := float64(math.Float64frombits(v))
					m.Values = append(m.Values, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGogoprotobufMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoGoProtobufGauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogoprotobufMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoGoProtobufGauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoProtobufGauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.Value = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipGogoprotobufMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GoGoProtobufMetricUnion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGogoprotobufMetrics
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GoGoProtobufMetricUnion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GoGoProtobufMetricUnion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (GoGoProtobufType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = append(m.ID[:0], dAtA[iNdEx:postIndex]...)
			if m.ID == nil {
				m.ID = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterVal", wireType)
			}
			m.CounterVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CounterVal |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 1 {
				var v uint64
				if (iNdEx + 8) > l {
					return io.ErrUnexpectedEOF
				}
				iNdEx += 8
				v = uint64(dAtA[iNdEx-8])
				v |= uint64(dAtA[iNdEx-7]) << 8
				v |= uint64(dAtA[iNdEx-6]) << 16
				v |= uint64(dAtA[iNdEx-5]) << 24
				v |= uint64(dAtA[iNdEx-4]) << 32
				v |= uint64(dAtA[iNdEx-3]) << 40
				v |= uint64(dAtA[iNdEx-2]) << 48
				v |= uint64(dAtA[iNdEx-1]) << 56
				v2 := float64(math.Float64frombits(v))
				m.BatchTimerVal = append(m.BatchTimerVal, v2)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGogoprotobufMetrics
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGogoprotobufMetrics
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					if (iNdEx + 8) > l {
						return io.ErrUnexpectedEOF
					}
					iNdEx += 8
					v = uint64(dAtA[iNdEx-8])
					v |= uint64(dAtA[iNdEx-7]) << 8
					v |= uint64(dAtA[iNdEx-6]) << 16
					v |= uint64(dAtA[iNdEx-5]) << 24
					v |= uint64(dAtA[iNdEx-4]) << 32
					v |= uint64(dAtA[iNdEx-3]) << 40
					v |= uint64(dAtA[iNdEx-2]) << 48
					v |= uint64(dAtA[iNdEx-1]) << 56
					v2 := float64(math.Float64frombits(v))
					m.BatchTimerVal = append(m.BatchTimerVal, v2)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchTimerVal", wireType)
			}
		case 5:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field GaugeVal", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += 8
			v = uint64(dAtA[iNdEx-8])
			v |= uint64(dAtA[iNdEx-7]) << 8
			v |= uint64(dAtA[iNdEx-6]) << 16
			v |= uint64(dAtA[iNdEx-5]) << 24
			v |= uint64(dAtA[iNdEx-4]) << 32
			v |= uint64(dAtA[iNdEx-3]) << 40
			v |= uint64(dAtA[iNdEx-2]) << 48
			v |= uint64(dAtA[iNdEx-1]) << 56
			m.GaugeVal = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipGogoprotobufMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGogoprotobufMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGogoprotobufMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGogoprotobufMetrics
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGogoprotobufMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGogoprotobufMetrics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGogoprotobufMetrics
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGogoprotobufMetrics(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGogoprotobufMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGogoprotobufMetrics   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/jeromefroe/metser/gogoprotobuf-metrics.proto", fileDescriptorGogoprotobufMetrics)
}

var fileDescriptorGogoprotobufMetrics = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4e, 0xc2, 0x40,
	0x14, 0x66, 0x5a, 0x68, 0xf0, 0x29, 0x88, 0xa3, 0xc1, 0x86, 0x45, 0x43, 0x88, 0x0b, 0x62, 0xa4,
	0x4d, 0x74, 0x61, 0x8c, 0x1b, 0x83, 0x24, 0xc4, 0x85, 0x89, 0x69, 0x00, 0xd7, 0x6d, 0x33, 0x94,
	0x2a, 0xed, 0x90, 0x76, 0xaa, 0xe1, 0x26, 0x5e, 0xc5, 0x1b, 0xb8, 0xf4, 0x08, 0x06, 0x2f, 0x62,
	0xe6, 0x51, 0x68, 0x63, 0x8c, 0xbb, 0xbe, 0xaf, 0xdf, 0xcf, 0xfb, 0xde, 0xc0, 0xa5, 0x1f, 0x88,
	0x59, 0xea, 0x9a, 0x1e, 0x0f, 0xad, 0x27, 0x16, 0xf3, 0x90, 0x4d, 0x63, 0xce, 0xac, 0x90, 0x89,
	0x84, 0xc5, 0x96, 0xcf, 0x7d, 0xbe, 0x88, 0xb9, 0xe0, 0x6e, 0x3a, 0xed, 0x85, 0x4c, 0xc4, 0x81,
	0x97, 0x98, 0x08, 0x50, 0x6d, 0x4d, 0x69, 0xf5, 0x0a, 0x06, 0x92, 0x6e, 0x6d, 0xf8, 0xb9, 0x18,
	0xbf, 0xd6, 0xb2, 0xce, 0x35, 0x1c, 0x0e, 0xf9, 0x90, 0x3f, 0x64, 0xa4, 0x5b, 0x9e, 0x46, 0x82,
	0xc5, 0xb4, 0x0e, 0xca, 0xdd, 0x40, 0x27, 0x6d, 0xd2, 0xdd, 0xb3, 0x95, 0x60, 0x40, 0x8f, 0xa0,
	0x32, 0x71, 0xe6, 0x29, 0xd3, 0x95, 0x36, 0xe9, 0xaa, 0x76, 0xe5, 0x45, 0x0e, 0x9d, 0x1b, 0x68,
	0x16, 0xc5, 0x7d, 0x47, 0x78, 0xb3, 0x51, 0x10, 0xfe, 0xa1, 0x6f, 0x82, 0x86, 0xfa, 0x44, 0x57,
	0xda, 0x6a, 0x97, 0xd8, 0x1a, 0x1a, 0x24, 0x9d, 0x2b, 0x38, 0x28, 0x3a, 0x0c, 0x9d, 0xd4, 0x67,
	0xff, 0x87, 0x93, 0x4d, 0xf8, 0x3b, 0x81, 0xe3, 0xa2, 0xf6, 0x1e, 0xcf, 0x31, 0x8e, 0x02, 0x1e,
	0xd1, 0x33, 0x28, 0x8f, 0x96, 0x0b, 0x86, 0x1e, 0xf5, 0x73, 0xdd, 0x5c, 0xdf, 0xc6, 0x2c, 0xd2,
	0xe5, 0x7f, 0xbb, 0x2c, 0x96, 0x8b, 0x4d, 0x9e, 0xb2, 0xcd, 0x33, 0x00, 0xb2, 0x3b, 0x4c, 0x9c,
	0xb9, 0xae, 0x62, 0x63, 0xf0, 0xb6, 0x08, 0x3d, 0x81, 0x5a, 0x5e, 0x55, 0x52, 0xca, 0xd8, 0xa9,
	0xe6, 0x16, 0x41, 0xda, 0x82, 0x2a, 0xd6, 0x91, 0x84, 0x0a, 0x2e, 0x5e, 0xf5, 0xb3, 0xf9, 0xf4,
	0x11, 0x1a, 0xbf, 0x77, 0xa1, 0xfb, 0xb0, 0x3b, 0x8e, 0x9e, 0x23, 0xfe, 0x1a, 0xc9, 0xb1, 0x51,
	0x92, 0x40, 0xb6, 0x06, 0x02, 0x84, 0x52, 0xa8, 0xe7, 0xb9, 0x88, 0x29, 0xb4, 0x06, 0x3b, 0x98,
	0x82, 0xa3, 0xda, 0x6f, 0x7c, 0xac, 0x0c, 0xf2, 0xb9, 0x32, 0xc8, 0xd7, 0xca, 0x20, 0x6f, 0xdf,
	0x46, 0xc9, 0xd5, 0xf0, 0x9d, 0x2f, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xad, 0x90, 0x67,
	0x59, 0x02, 0x00, 0x00,
}