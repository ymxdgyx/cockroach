// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/serverpb/init.proto

package serverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

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

type BootstrapRequest struct {
}

func (m *BootstrapRequest) Reset()         { *m = BootstrapRequest{} }
func (m *BootstrapRequest) String() string { return proto.CompactTextString(m) }
func (*BootstrapRequest) ProtoMessage()    {}
func (*BootstrapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_init_119a86c40b8bcd78, []int{0}
}
func (m *BootstrapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *BootstrapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapRequest.Merge(dst, src)
}
func (m *BootstrapRequest) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapRequest proto.InternalMessageInfo

type BootstrapResponse struct {
}

func (m *BootstrapResponse) Reset()         { *m = BootstrapResponse{} }
func (m *BootstrapResponse) String() string { return proto.CompactTextString(m) }
func (*BootstrapResponse) ProtoMessage()    {}
func (*BootstrapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_init_119a86c40b8bcd78, []int{1}
}
func (m *BootstrapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BootstrapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *BootstrapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BootstrapResponse.Merge(dst, src)
}
func (m *BootstrapResponse) XXX_Size() int {
	return m.Size()
}
func (m *BootstrapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BootstrapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BootstrapResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BootstrapRequest)(nil), "cockroach.server.serverpb.BootstrapRequest")
	proto.RegisterType((*BootstrapResponse)(nil), "cockroach.server.serverpb.BootstrapResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InitClient is the client API for Init service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InitClient interface {
	// Bootstrap bootstraps an uninitialized cluster. This is primarily used by
	// cockroach init. It writes the data for a single-node cluster comprised of
	// this node (with NodeID "1") to the first store (StoreID "1"), after which
	// the node can start and allow other nodes to join the cluster.
	Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error)
}

type initClient struct {
	cc *grpc.ClientConn
}

func NewInitClient(cc *grpc.ClientConn) InitClient {
	return &initClient{cc}
}

func (c *initClient) Bootstrap(ctx context.Context, in *BootstrapRequest, opts ...grpc.CallOption) (*BootstrapResponse, error) {
	out := new(BootstrapResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.Init/Bootstrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitServer is the server API for Init service.
type InitServer interface {
	// Bootstrap bootstraps an uninitialized cluster. This is primarily used by
	// cockroach init. It writes the data for a single-node cluster comprised of
	// this node (with NodeID "1") to the first store (StoreID "1"), after which
	// the node can start and allow other nodes to join the cluster.
	Bootstrap(context.Context, *BootstrapRequest) (*BootstrapResponse, error)
}

func RegisterInitServer(s *grpc.Server, srv InitServer) {
	s.RegisterService(&_Init_serviceDesc, srv)
}

func _Init_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Init/Bootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitServer).Bootstrap(ctx, req.(*BootstrapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Init_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.Init",
	HandlerType: (*InitServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bootstrap",
			Handler:    _Init_Bootstrap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/init.proto",
}

func (m *BootstrapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *BootstrapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BootstrapResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintInit(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BootstrapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *BootstrapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovInit(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInit(x uint64) (n int) {
	return sovInit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BootstrapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInit
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
			return fmt.Errorf("proto: BootstrapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInit
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
func (m *BootstrapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInit
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
			return fmt.Errorf("proto: BootstrapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BootstrapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipInit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInit
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
func skipInit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
					return 0, ErrIntOverflowInit
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
				return 0, ErrInvalidLengthInit
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInit
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
				next, err := skipInit(dAtA[start:])
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
	ErrInvalidLengthInit = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInit   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("server/serverpb/init.proto", fileDescriptor_init_119a86c40b8bcd78) }

var fileDescriptor_init_119a86c40b8bcd78 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x05, 0x49, 0xfa, 0x99, 0x79, 0x99, 0x25, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x92, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89, 0xc9, 0x19, 0x7a, 0x10, 0x69,
	0x3d, 0x98, 0x2a, 0x25, 0x21, 0x2e, 0x01, 0xa7, 0xfc, 0xfc, 0x92, 0xe2, 0x92, 0xa2, 0xc4, 0x82,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x61, 0x2e, 0x41, 0x24, 0xb1, 0xe2, 0x82, 0xfc,
	0xbc, 0xe2, 0x54, 0xa3, 0x02, 0x2e, 0x16, 0xcf, 0xbc, 0xcc, 0x12, 0xa1, 0x0c, 0x2e, 0x4e, 0xb8,
	0xa4, 0x90, 0xb6, 0x1e, 0x4e, 0x93, 0xf5, 0xd0, 0x8d, 0x95, 0xd2, 0x21, 0x4e, 0x31, 0xc4, 0x3e,
	0x25, 0x06, 0x27, 0xad, 0x13, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e,
	0xf1, 0xc6, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58,
	0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x0e, 0x98, 0xfe, 0x24, 0x36, 0xb0, 0x47, 0x8d, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x90, 0x15, 0x22, 0x79, 0x06, 0x01, 0x00, 0x00,
}
