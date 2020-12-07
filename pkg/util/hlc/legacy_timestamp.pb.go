// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: util/hlc/legacy_timestamp.proto

package hlc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// LegacyTimestamp is convertible to hlc.Timestamp, but uses the
// legacy encoding as it is encoded "below raft".
type LegacyTimestamp struct {
	// Holds a wall time, typically a unix epoch time expressed in
	// nanoseconds.
	WallTime int64 `protobuf:"varint,1,opt,name=wall_time,json=wallTime" json:"wall_time"`
	// The logical component captures causality for events whose wall
	// times are equal. It is effectively bounded by (maximum clock
	// skew)/(minimal ns between events) and nearly impossible to
	// overflow.
	Logical int32 `protobuf:"varint,2,opt,name=logical" json:"logical"`
	// A collection of bit flags that provide details about the timestamp
	// and its meaning. The data type is a uint32, but the number of flags
	// is limited to 8 so that the flags can be encoded into a single byte.
	//
	// Flags do not affect the sort order of Timestamps. However, they are
	// considered when performing structural equality checks (e.g. using the
	// == operator). Consider use of the EqOrdering method when testing for
	// equality.
	//
	// The field is nullable so that it is not serialized when no flags are
	// set. This ensures that the timestamp encoding does not change across
	// nodes that are and are not aware of this field.
	Flags *uint32 `protobuf:"varint,3,opt,name=flags" json:"flags,omitempty"`
}

func (m *LegacyTimestamp) Reset()      { *m = LegacyTimestamp{} }
func (*LegacyTimestamp) ProtoMessage() {}
func (*LegacyTimestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_legacy_timestamp_d72283f54eaf58e6, []int{0}
}
func (m *LegacyTimestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LegacyTimestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *LegacyTimestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LegacyTimestamp.Merge(dst, src)
}
func (m *LegacyTimestamp) XXX_Size() int {
	return m.Size()
}
func (m *LegacyTimestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_LegacyTimestamp.DiscardUnknown(m)
}

var xxx_messageInfo_LegacyTimestamp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LegacyTimestamp)(nil), "cockroach.util.hlc.LegacyTimestamp")
}
func (this *LegacyTimestamp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*LegacyTimestamp)
	if !ok {
		that2, ok := that.(LegacyTimestamp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.WallTime != that1.WallTime {
		return false
	}
	if this.Logical != that1.Logical {
		return false
	}
	if this.Flags != nil && that1.Flags != nil {
		if *this.Flags != *that1.Flags {
			return false
		}
	} else if this.Flags != nil {
		return false
	} else if that1.Flags != nil {
		return false
	}
	return true
}
func (m *LegacyTimestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LegacyTimestamp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintLegacyTimestamp(dAtA, i, uint64(m.WallTime))
	dAtA[i] = 0x10
	i++
	i = encodeVarintLegacyTimestamp(dAtA, i, uint64(m.Logical))
	if m.Flags != nil {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLegacyTimestamp(dAtA, i, uint64(*m.Flags))
	}
	return i, nil
}

func encodeVarintLegacyTimestamp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedLegacyTimestamp(r randyLegacyTimestamp, easy bool) *LegacyTimestamp {
	this := &LegacyTimestamp{}
	this.WallTime = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.WallTime *= -1
	}
	this.Logical = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Logical *= -1
	}
	if r.Intn(10) != 0 {
		v1 := uint32(r.Uint32())
		this.Flags = &v1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyLegacyTimestamp interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneLegacyTimestamp(r randyLegacyTimestamp) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringLegacyTimestamp(r randyLegacyTimestamp) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneLegacyTimestamp(r)
	}
	return string(tmps)
}
func randUnrecognizedLegacyTimestamp(r randyLegacyTimestamp, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldLegacyTimestamp(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldLegacyTimestamp(dAtA []byte, r randyLegacyTimestamp, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateLegacyTimestamp(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateLegacyTimestamp(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *LegacyTimestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovLegacyTimestamp(uint64(m.WallTime))
	n += 1 + sovLegacyTimestamp(uint64(m.Logical))
	if m.Flags != nil {
		n += 1 + sovLegacyTimestamp(uint64(*m.Flags))
	}
	return n
}

func sovLegacyTimestamp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLegacyTimestamp(x uint64) (n int) {
	return sovLegacyTimestamp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LegacyTimestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLegacyTimestamp
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
			return fmt.Errorf("proto: LegacyTimestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LegacyTimestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WallTime", wireType)
			}
			m.WallTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLegacyTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WallTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Logical", wireType)
			}
			m.Logical = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLegacyTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Logical |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flags", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLegacyTimestamp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Flags = &v
		default:
			iNdEx = preIndex
			skippy, err := skipLegacyTimestamp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLegacyTimestamp
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
func skipLegacyTimestamp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLegacyTimestamp
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
					return 0, ErrIntOverflowLegacyTimestamp
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
					return 0, ErrIntOverflowLegacyTimestamp
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
				return 0, ErrInvalidLengthLegacyTimestamp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLegacyTimestamp
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
				next, err := skipLegacyTimestamp(dAtA[start:])
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
	ErrInvalidLengthLegacyTimestamp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLegacyTimestamp   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("util/hlc/legacy_timestamp.proto", fileDescriptor_legacy_timestamp_d72283f54eaf58e6)
}

var fileDescriptor_legacy_timestamp_d72283f54eaf58e6 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2d, 0xc9, 0xcc,
	0xd1, 0xcf, 0xc8, 0x49, 0xd6, 0xcf, 0x49, 0x4d, 0x4f, 0x4c, 0xae, 0x8c, 0x2f, 0xc9, 0xcc, 0x4d,
	0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xce, 0x4f,
	0xce, 0x2e, 0xca, 0x4f, 0x4c, 0xce, 0xd0, 0x03, 0x29, 0xd5, 0xcb, 0xc8, 0x49, 0x96, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x4b, 0xeb, 0x83, 0x58, 0x10, 0x95, 0x4a, 0x15, 0x5c, 0xfc, 0x3e, 0x60,
	0x33, 0x42, 0x60, 0x46, 0x08, 0x29, 0x72, 0x71, 0x96, 0x27, 0xe6, 0xe4, 0x80, 0x0d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0x60, 0x76, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x88, 0x03, 0x24, 0x0c, 0x52,
	0x27, 0x24, 0xc7, 0xc5, 0x9e, 0x93, 0x9f, 0x9e, 0x99, 0x9c, 0x98, 0x23, 0xc1, 0xa4, 0xc0, 0xa8,
	0xc1, 0x0a, 0x55, 0x00, 0x13, 0x14, 0x12, 0xe1, 0x62, 0x4d, 0xcb, 0x49, 0x4c, 0x2f, 0x96, 0x60,
	0x56, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x70, 0xac, 0x78, 0x66, 0x2c, 0x90, 0x67, 0xd8, 0xb1, 0x40,
	0x9e, 0xf1, 0xc5, 0x02, 0x79, 0x46, 0x27, 0xd5, 0x13, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63,
	0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc6, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96,
	0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xe6, 0x8c, 0x9c, 0x64, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x5c, 0x3e, 0x65, 0xec, 0x00, 0x00, 0x00,
}
