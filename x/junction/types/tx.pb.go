// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: junction/junction/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgInitStation struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	VerificationKey []byte `protobuf:"bytes,2,opt,name=verificationKey,proto3" json:"verificationKey,omitempty"`
	StationId       string `protobuf:"bytes,3,opt,name=stationId,proto3" json:"stationId,omitempty"`
	StationInfo     string `protobuf:"bytes,4,opt,name=stationInfo,proto3" json:"stationInfo,omitempty"`
}

func (m *MsgInitStation) Reset()         { *m = MsgInitStation{} }
func (m *MsgInitStation) String() string { return proto.CompactTextString(m) }
func (*MsgInitStation) ProtoMessage()    {}
func (*MsgInitStation) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{0}
}
func (m *MsgInitStation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitStation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitStation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitStation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitStation.Merge(m, src)
}
func (m *MsgInitStation) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitStation) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitStation.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitStation proto.InternalMessageInfo

func (m *MsgInitStation) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgInitStation) GetVerificationKey() []byte {
	if m != nil {
		return m.VerificationKey
	}
	return nil
}

func (m *MsgInitStation) GetStationId() string {
	if m != nil {
		return m.StationId
	}
	return ""
}

func (m *MsgInitStation) GetStationInfo() string {
	if m != nil {
		return m.StationInfo
	}
	return ""
}

type MsgInitStationResponse struct {
	Status    bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	StationId string `protobuf:"bytes,2,opt,name=stationId,proto3" json:"stationId,omitempty"`
}

func (m *MsgInitStationResponse) Reset()         { *m = MsgInitStationResponse{} }
func (m *MsgInitStationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitStationResponse) ProtoMessage()    {}
func (*MsgInitStationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_04349ac28bbdc1dc, []int{1}
}
func (m *MsgInitStationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitStationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitStationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitStationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitStationResponse.Merge(m, src)
}
func (m *MsgInitStationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitStationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitStationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitStationResponse proto.InternalMessageInfo

func (m *MsgInitStationResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *MsgInitStationResponse) GetStationId() string {
	if m != nil {
		return m.StationId
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgInitStation)(nil), "junction.junction.MsgInitStation")
	proto.RegisterType((*MsgInitStationResponse)(nil), "junction.junction.MsgInitStationResponse")
}

func init() { proto.RegisterFile("junction/junction/tx.proto", fileDescriptor_04349ac28bbdc1dc) }

var fileDescriptor_04349ac28bbdc1dc = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x2a, 0xcd, 0x4b,
	0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x87, 0x33, 0x4a, 0x2a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x04, 0x61, 0x42, 0x7a, 0x30, 0x86, 0xd2, 0x14, 0x46, 0x2e, 0x3e, 0xdf, 0xe2, 0x74, 0xcf, 0xbc,
	0xcc, 0x92, 0xe0, 0x92, 0x44, 0x90, 0x90, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51, 0x6a, 0x62, 0x49,
	0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0xa4, 0xc1, 0xc5, 0x5f, 0x96,
	0x5a, 0x94, 0x99, 0x96, 0x99, 0x0c, 0x56, 0xe9, 0x9d, 0x5a, 0x29, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1,
	0x13, 0x84, 0x2e, 0x2c, 0x24, 0xc3, 0xc5, 0x59, 0x0c, 0x31, 0xce, 0x33, 0x45, 0x82, 0x19, 0x6c,
	0x0a, 0x42, 0x40, 0x48, 0x81, 0x8b, 0x1b, 0xc6, 0xc9, 0x4b, 0xcb, 0x97, 0x60, 0x01, 0xcb, 0x23,
	0x0b, 0x29, 0xf9, 0x71, 0x89, 0xa1, 0xba, 0x2a, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x8c, 0x8b, 0x0d, 0xa4, 0xb0, 0xb4, 0x18, 0xec, 0x38, 0x8e, 0x20, 0x28, 0x0f, 0xd5, 0x46,
	0x26, 0x34, 0x1b, 0x8d, 0x92, 0xb8, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xa2, 0xb9, 0xb8, 0x91, 0x7d,
	0xaa, 0xa8, 0x87, 0x11, 0x20, 0x7a, 0xa8, 0xd6, 0x4a, 0x69, 0x12, 0x54, 0x02, 0x73, 0x99, 0x53,
	0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1,
	0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0xa5, 0x67, 0x96, 0x64,
	0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x66, 0x16, 0x25, 0x67, 0x24, 0x66, 0xe6, 0x15,
	0xeb, 0xe6, 0xa5, 0x96, 0x94, 0xe7, 0x17, 0x65, 0x23, 0xe2, 0xa9, 0x02, 0x29, 0xca, 0x2a, 0x0b,
	0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xd1, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7f, 0xe4, 0x28,
	0x4a, 0xd4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) InitStation(ctx context.Context, in *MsgInitStation, opts ...grpc.CallOption) (*MsgInitStationResponse, error) {
	out := new(MsgInitStationResponse)
	err := c.cc.Invoke(ctx, "/junction.junction.Msg/InitStation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	InitStation(context.Context, *MsgInitStation) (*MsgInitStationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) InitStation(ctx context.Context, req *MsgInitStation) (*MsgInitStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitStation not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_InitStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInitStation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InitStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/junction.junction.Msg/InitStation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InitStation(ctx, req.(*MsgInitStation))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "junction.junction.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InitStation",
			Handler:    _Msg_InitStation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "junction/junction/tx.proto",
}

func (m *MsgInitStation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitStation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitStation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StationInfo) > 0 {
		i -= len(m.StationInfo)
		copy(dAtA[i:], m.StationInfo)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationInfo)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.StationId) > 0 {
		i -= len(m.StationId)
		copy(dAtA[i:], m.StationId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VerificationKey) > 0 {
		i -= len(m.VerificationKey)
		copy(dAtA[i:], m.VerificationKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VerificationKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitStationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitStationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitStationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StationId) > 0 {
		i -= len(m.StationId)
		copy(dAtA[i:], m.StationId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.StationId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Status {
		i--
		if m.Status {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInitStation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VerificationKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StationId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.StationInfo)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgInitStationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status {
		n += 2
	}
	l = len(m.StationId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitStation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitStation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitStation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationKey = append(m.VerificationKey[:0], dAtA[iNdEx:postIndex]...)
			if m.VerificationKey == nil {
				m.VerificationKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationInfo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationInfo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitStationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgInitStationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitStationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Status = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StationId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StationId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
