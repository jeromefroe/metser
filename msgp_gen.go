package metser

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// MarshalMsg implements msgp.Marshaler
func (z *BatchTimer) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, []byte(z.ID))
	// string "Values"
	o = append(o, 0xa6, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Values)))
	for zxvk := range z.Values {
		o = msgp.AppendFloat64(o, z.Values[zxvk])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *BatchTimer) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			{
				var zbai []byte
				zbai, bts, err = msgp.ReadBytesBytes(bts, []byte(z.ID))
				z.ID = ID(zbai)
			}
			if err != nil {
				return
			}
		case "Values":
			var zcmr uint32
			zcmr, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Values) >= int(zcmr) {
				z.Values = (z.Values)[:zcmr]
			} else {
				z.Values = make([]float64, zcmr)
			}
			for zxvk := range z.Values {
				z.Values[zxvk], bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *BatchTimer) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len([]byte(z.ID)) + 7 + msgp.ArrayHeaderSize + (len(z.Values) * (msgp.Float64Size))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Counter) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, []byte(z.ID))
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendInt64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Counter) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			{
				var zwht []byte
				zwht, bts, err = msgp.ReadBytesBytes(bts, []byte(z.ID))
				z.ID = ID(zwht)
			}
			if err != nil {
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Counter) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len([]byte(z.ID)) + 6 + msgp.Int64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Gauge) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "ID"
	o = append(o, 0x82, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, []byte(z.ID))
	// string "Value"
	o = append(o, 0xa5, 0x56, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendFloat64(o, z.Value)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Gauge) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zhct uint32
	zhct, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zhct > 0 {
		zhct--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ID":
			{
				var zcua []byte
				zcua, bts, err = msgp.ReadBytesBytes(bts, []byte(z.ID))
				z.ID = ID(zcua)
			}
			if err != nil {
				return
			}
		case "Value":
			z.Value, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Gauge) Msgsize() (s int) {
	s = 1 + 3 + msgp.BytesPrefixSize + len([]byte(z.ID)) + 6 + msgp.Float64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendBytes(o, []byte(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zxhx []byte
		zxhx, bts, err = msgp.ReadBytesBytes(bts, []byte((*z)))
		(*z) = ID(zxhx)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ID) Msgsize() (s int) {
	s = msgp.BytesPrefixSize + len([]byte(z))
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *MetricUnion) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 5
	// string "Type"
	o = append(o, 0x85, 0xa4, 0x54, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, int8(z.Type))
	// string "ID"
	o = append(o, 0xa2, 0x49, 0x44)
	o = msgp.AppendBytes(o, []byte(z.ID))
	// string "CounterVal"
	o = append(o, 0xaa, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c)
	o = msgp.AppendInt64(o, z.CounterVal)
	// string "BatchTimerVal"
	o = append(o, 0xad, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x72, 0x56, 0x61, 0x6c)
	o = msgp.AppendArrayHeader(o, uint32(len(z.BatchTimerVal)))
	for zlqf := range z.BatchTimerVal {
		o = msgp.AppendFloat64(o, z.BatchTimerVal[zlqf])
	}
	// string "GaugeVal"
	o = append(o, 0xa8, 0x47, 0x61, 0x75, 0x67, 0x65, 0x56, 0x61, 0x6c)
	o = msgp.AppendFloat64(o, z.GaugeVal)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *MetricUnion) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Type":
			{
				var zpks int8
				zpks, bts, err = msgp.ReadInt8Bytes(bts)
				z.Type = Type(zpks)
			}
			if err != nil {
				return
			}
		case "ID":
			{
				var zjfb []byte
				zjfb, bts, err = msgp.ReadBytesBytes(bts, []byte(z.ID))
				z.ID = ID(zjfb)
			}
			if err != nil {
				return
			}
		case "CounterVal":
			z.CounterVal, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "BatchTimerVal":
			var zcxo uint32
			zcxo, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.BatchTimerVal) >= int(zcxo) {
				z.BatchTimerVal = (z.BatchTimerVal)[:zcxo]
			} else {
				z.BatchTimerVal = make([]float64, zcxo)
			}
			for zlqf := range z.BatchTimerVal {
				z.BatchTimerVal[zlqf], bts, err = msgp.ReadFloat64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "GaugeVal":
			z.GaugeVal, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *MetricUnion) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int8Size + 3 + msgp.BytesPrefixSize + len([]byte(z.ID)) + 11 + msgp.Int64Size + 14 + msgp.ArrayHeaderSize + (len(z.BatchTimerVal) * (msgp.Float64Size)) + 9 + msgp.Float64Size
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Type) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Type) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zeff int8
		zeff, bts, err = msgp.ReadInt8Bytes(bts)
		(*z) = Type(zeff)
	}
	if err != nil {
		return
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Type) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}
