package metser

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeCounter struct {
	ID    []byte
	Value int64
}

func (d *GencodeCounter) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{

		t := uint64(d.Value)
		t <<= 1
		if d.Value < 0 {
			t = ^t
		}
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	return
}
func (d *GencodeCounter) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.ID)
		i += l
	}
	{

		t := uint64(d.Value)

		t <<= 1
		if d.Value < 0 {
			t = ^t
		}

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	return buf[:i+0], nil
}

func (d *GencodeCounter) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ID)) >= l {
			d.ID = d.ID[:l]
		} else {
			d.ID = make([]byte, l)
		}
		copy(d.ID, buf[i+0:])
		i += l
	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Value = int64(t >> 1)
		if t&1 != 0 {
			d.Value = ^d.Value
		}

	}
	return i + 0, nil
}

type GencodeBatchTimer struct {
	ID     []byte
	Values []float64
}

func (d *GencodeBatchTimer) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Values))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	return
}
func (d *GencodeBatchTimer) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.ID)
		i += l
	}
	{
		l := uint64(len(d.Values))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Values {

			{

				v := *(*uint64)(unsafe.Pointer(&(d.Values[k0])))

				buf[i+0+0] = byte(v >> 0)

				buf[i+1+0] = byte(v >> 8)

				buf[i+2+0] = byte(v >> 16)

				buf[i+3+0] = byte(v >> 24)

				buf[i+4+0] = byte(v >> 32)

				buf[i+5+0] = byte(v >> 40)

				buf[i+6+0] = byte(v >> 48)

				buf[i+7+0] = byte(v >> 56)

			}

			i += 8

		}
	}
	return buf[:i+0], nil
}

func (d *GencodeBatchTimer) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ID)) >= l {
			d.ID = d.ID[:l]
		} else {
			d.ID = make([]byte, l)
		}
		copy(d.ID, buf[i+0:])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Values)) >= l {
			d.Values = d.Values[:l]
		} else {
			d.Values = make([]float64, l)
		}
		for k0 := range d.Values {

			{

				v := 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
				d.Values[k0] = *(*float64)(unsafe.Pointer(&v))

			}

			i += 8

		}
	}
	return i + 0, nil
}

type GencodeGauge struct {
	ID    []byte
	Value float64
}

func (d *GencodeGauge) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 8
	return
}
func (d *GencodeGauge) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.ID)
		i += l
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Value)))

		buf[i+0+0] = byte(v >> 0)

		buf[i+1+0] = byte(v >> 8)

		buf[i+2+0] = byte(v >> 16)

		buf[i+3+0] = byte(v >> 24)

		buf[i+4+0] = byte(v >> 32)

		buf[i+5+0] = byte(v >> 40)

		buf[i+6+0] = byte(v >> 48)

		buf[i+7+0] = byte(v >> 56)

	}
	return buf[:i+8], nil
}

func (d *GencodeGauge) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ID)) >= l {
			d.ID = d.ID[:l]
		} else {
			d.ID = make([]byte, l)
		}
		copy(d.ID, buf[i+0:])
		i += l
	}
	{

		v := 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)
		d.Value = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 8, nil
}

type GencodeMetricUnion struct {
	Type          int8
	ID            []byte
	CounterVal    int64
	BatchTimerVal []float64
	GaugeVal      float64
}

func (d *GencodeMetricUnion) Size() (s uint64) {

	{
		l := uint64(len(d.ID))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.BatchTimerVal))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	s += 17
	return
}
func (d *GencodeMetricUnion) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Type >> 0)

	}
	{
		l := uint64(len(d.ID))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+1] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+1] = byte(t)
			i++

		}
		copy(buf[i+1:], d.ID)
		i += l
	}
	{

		buf[i+0+1] = byte(d.CounterVal >> 0)

		buf[i+1+1] = byte(d.CounterVal >> 8)

		buf[i+2+1] = byte(d.CounterVal >> 16)

		buf[i+3+1] = byte(d.CounterVal >> 24)

		buf[i+4+1] = byte(d.CounterVal >> 32)

		buf[i+5+1] = byte(d.CounterVal >> 40)

		buf[i+6+1] = byte(d.CounterVal >> 48)

		buf[i+7+1] = byte(d.CounterVal >> 56)

	}
	{
		l := uint64(len(d.BatchTimerVal))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+9] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+9] = byte(t)
			i++

		}
		for k0 := range d.BatchTimerVal {

			{

				v := *(*uint64)(unsafe.Pointer(&(d.BatchTimerVal[k0])))

				buf[i+0+9] = byte(v >> 0)

				buf[i+1+9] = byte(v >> 8)

				buf[i+2+9] = byte(v >> 16)

				buf[i+3+9] = byte(v >> 24)

				buf[i+4+9] = byte(v >> 32)

				buf[i+5+9] = byte(v >> 40)

				buf[i+6+9] = byte(v >> 48)

				buf[i+7+9] = byte(v >> 56)

			}

			i += 8

		}
	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.GaugeVal)))

		buf[i+0+9] = byte(v >> 0)

		buf[i+1+9] = byte(v >> 8)

		buf[i+2+9] = byte(v >> 16)

		buf[i+3+9] = byte(v >> 24)

		buf[i+4+9] = byte(v >> 32)

		buf[i+5+9] = byte(v >> 40)

		buf[i+6+9] = byte(v >> 48)

		buf[i+7+9] = byte(v >> 56)

	}
	return buf[:i+17], nil
}

func (d *GencodeMetricUnion) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Type = 0 | (int8(buf[i+0+0]) << 0)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+1] & 0x7F)
			for buf[i+1]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+1]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ID)) >= l {
			d.ID = d.ID[:l]
		} else {
			d.ID = make([]byte, l)
		}
		copy(d.ID, buf[i+1:])
		i += l
	}
	{

		d.CounterVal = 0 | (int64(buf[i+0+1]) << 0) | (int64(buf[i+1+1]) << 8) | (int64(buf[i+2+1]) << 16) | (int64(buf[i+3+1]) << 24) | (int64(buf[i+4+1]) << 32) | (int64(buf[i+5+1]) << 40) | (int64(buf[i+6+1]) << 48) | (int64(buf[i+7+1]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+9] & 0x7F)
			for buf[i+9]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+9]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.BatchTimerVal)) >= l {
			d.BatchTimerVal = d.BatchTimerVal[:l]
		} else {
			d.BatchTimerVal = make([]float64, l)
		}
		for k0 := range d.BatchTimerVal {

			{

				v := 0 | (uint64(buf[i+0+9]) << 0) | (uint64(buf[i+1+9]) << 8) | (uint64(buf[i+2+9]) << 16) | (uint64(buf[i+3+9]) << 24) | (uint64(buf[i+4+9]) << 32) | (uint64(buf[i+5+9]) << 40) | (uint64(buf[i+6+9]) << 48) | (uint64(buf[i+7+9]) << 56)
				d.BatchTimerVal[k0] = *(*float64)(unsafe.Pointer(&v))

			}

			i += 8

		}
	}
	{

		v := 0 | (uint64(buf[i+0+9]) << 0) | (uint64(buf[i+1+9]) << 8) | (uint64(buf[i+2+9]) << 16) | (uint64(buf[i+3+9]) << 24) | (uint64(buf[i+4+9]) << 32) | (uint64(buf[i+5+9]) << 40) | (uint64(buf[i+6+9]) << 48) | (uint64(buf[i+7+9]) << 56)
		d.GaugeVal = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 17, nil
}
