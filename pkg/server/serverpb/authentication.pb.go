// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/serverpb/authentication.proto

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

// UserLoginRequest contains credentials a user must provide to log in.
type UserLoginRequest struct {
	// A username which must correspond to a database user on the cluster.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// A password for the provided username.
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_a6dd60e2be36c9bf, []int{0}
}
func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(dst, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

// UserLoginResponse is currently empty. If a login is successful, an HTTP
// Set-Cookie header will be added to the response with a session
// cookie identifying the created session.
type UserLoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_a6dd60e2be36c9bf, []int{1}
}
func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(dst, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return m.Size()
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

// UserLogoutRequest will terminate the current session in use. The request
// is empty because the current session is identified by an HTTP cookie on the
// incoming request.
type UserLogoutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
}

func (m *UserLogoutRequest) Reset()         { *m = UserLogoutRequest{} }
func (m *UserLogoutRequest) String() string { return proto.CompactTextString(m) }
func (*UserLogoutRequest) ProtoMessage()    {}
func (*UserLogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_a6dd60e2be36c9bf, []int{2}
}
func (m *UserLogoutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *UserLogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLogoutRequest.Merge(dst, src)
}
func (m *UserLogoutRequest) XXX_Size() int {
	return m.Size()
}
func (m *UserLogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLogoutRequest proto.InternalMessageInfo

type UserLogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
}

func (m *UserLogoutResponse) Reset()         { *m = UserLogoutResponse{} }
func (m *UserLogoutResponse) String() string { return proto.CompactTextString(m) }
func (*UserLogoutResponse) ProtoMessage()    {}
func (*UserLogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_a6dd60e2be36c9bf, []int{3}
}
func (m *UserLogoutResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserLogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *UserLogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLogoutResponse.Merge(dst, src)
}
func (m *UserLogoutResponse) XXX_Size() int {
	return m.Size()
}
func (m *UserLogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLogoutResponse proto.InternalMessageInfo

// SessionCookie is a message used to encode the authentication cookie returned
// from successful login requests.
type SessionCookie struct {
	// The unique ID of the session.
	ID int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// The secret needed to verify ownership of a session.
	Secret               []byte   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
}

func (m *SessionCookie) Reset()         { *m = SessionCookie{} }
func (m *SessionCookie) String() string { return proto.CompactTextString(m) }
func (*SessionCookie) ProtoMessage()    {}
func (*SessionCookie) Descriptor() ([]byte, []int) {
	return fileDescriptor_authentication_a6dd60e2be36c9bf, []int{4}
}
func (m *SessionCookie) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SessionCookie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *SessionCookie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionCookie.Merge(dst, src)
}
func (m *SessionCookie) XXX_Size() int {
	return m.Size()
}
func (m *SessionCookie) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionCookie.DiscardUnknown(m)
}

var xxx_messageInfo_SessionCookie proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UserLoginRequest)(nil), "cockroach.server.serverpb.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "cockroach.server.serverpb.UserLoginResponse")
	proto.RegisterType((*UserLogoutRequest)(nil), "cockroach.server.serverpb.UserLogoutRequest")
	proto.RegisterType((*UserLogoutResponse)(nil), "cockroach.server.serverpb.UserLogoutResponse")
	proto.RegisterType((*SessionCookie)(nil), "cockroach.server.serverpb.SessionCookie")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogInClient is the client API for LogIn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogInClient interface {
	// UserLogin is used to create a web authentication session.
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
}

type logInClient struct {
	cc *grpc.ClientConn
}

func NewLogInClient(cc *grpc.ClientConn) LogInClient {
	return &logInClient{cc}
}

func (c *logInClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.LogIn/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogInServer is the server API for LogIn service.
type LogInServer interface {
	// UserLogin is used to create a web authentication session.
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
}

func RegisterLogInServer(s *grpc.Server, srv LogInServer) {
	s.RegisterService(&_LogIn_serviceDesc, srv)
}

func _LogIn_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogInServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.LogIn/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogInServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogIn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.LogIn",
	HandlerType: (*LogInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogin",
			Handler:    _LogIn_UserLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/authentication.proto",
}

// LogOutClient is the client API for LogOut service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogOutClient interface {
	// UserLogout terminates an active authentication session.
	UserLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error)
}

type logOutClient struct {
	cc *grpc.ClientConn
}

func NewLogOutClient(cc *grpc.ClientConn) LogOutClient {
	return &logOutClient{cc}
}

func (c *logOutClient) UserLogout(ctx context.Context, in *UserLogoutRequest, opts ...grpc.CallOption) (*UserLogoutResponse, error) {
	out := new(UserLogoutResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.LogOut/UserLogout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogOutServer is the server API for LogOut service.
type LogOutServer interface {
	// UserLogout terminates an active authentication session.
	UserLogout(context.Context, *UserLogoutRequest) (*UserLogoutResponse, error)
}

func RegisterLogOutServer(s *grpc.Server, srv LogOutServer) {
	s.RegisterService(&_LogOut_serviceDesc, srv)
}

func _LogOut_UserLogout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogOutServer).UserLogout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.LogOut/UserLogout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogOutServer).UserLogout(ctx, req.(*UserLogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogOut_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.LogOut",
	HandlerType: (*LogOutServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserLogout",
			Handler:    _LogOut_UserLogout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/authentication.proto",
}

func (m *UserLoginRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	return i, nil
}

func (m *UserLoginResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLoginResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserLogoutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLogoutRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserLogoutResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserLogoutResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SessionCookie) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SessionCookie) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(m.ID))
	}
	if len(m.Secret) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintAuthentication(dAtA, i, uint64(len(m.Secret)))
		i += copy(dAtA[i:], m.Secret)
	}
	return i, nil
}

func encodeVarintAuthentication(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserLoginRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func (m *UserLoginResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UserLogoutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UserLogoutResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SessionCookie) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovAuthentication(uint64(m.ID))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovAuthentication(uint64(l))
	}
	return n
}

func sovAuthentication(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozAuthentication(x uint64) (n int) {
	return sovAuthentication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserLoginRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLoginRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLoginResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLoginResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLoginResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLogoutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLogoutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLogoutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *UserLogoutResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: UserLogoutResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserLogoutResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func (m *SessionCookie) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthentication
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
			return fmt.Errorf("proto: SessionCookie: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SessionCookie: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthentication
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
				return ErrInvalidLengthAuthentication
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = append(m.Secret[:0], dAtA[iNdEx:postIndex]...)
			if m.Secret == nil {
				m.Secret = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthentication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuthentication
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
func skipAuthentication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthentication
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
					return 0, ErrIntOverflowAuthentication
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
					return 0, ErrIntOverflowAuthentication
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
				return 0, ErrInvalidLengthAuthentication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowAuthentication
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
				next, err := skipAuthentication(dAtA[start:])
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
	ErrInvalidLengthAuthentication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthentication   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/serverpb/authentication.proto", fileDescriptor_authentication_a6dd60e2be36c9bf)
}

var fileDescriptor_authentication_a6dd60e2be36c9bf = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0x86, 0x3b, 0xb9, 0xdc, 0xdc, 0x76, 0xae, 0xa2, 0x1d, 0x4b, 0xa9, 0x41, 0xa2, 0x04, 0x11,
	0xf1, 0x23, 0x81, 0xba, 0x73, 0x23, 0x54, 0x37, 0x95, 0x82, 0x10, 0x71, 0xe3, 0x2e, 0x4d, 0x87,
	0x74, 0x68, 0x9d, 0x13, 0x67, 0x26, 0x8a, 0x2e, 0x8b, 0x3f, 0x40, 0xf0, 0x4f, 0x75, 0x59, 0x70,
	0xe3, 0x4a, 0x34, 0xf5, 0x87, 0x48, 0x3e, 0xfa, 0x81, 0x20, 0x74, 0x95, 0x9c, 0xf3, 0xbe, 0xf3,
	0xce, 0x73, 0x0e, 0x83, 0xb7, 0x25, 0x15, 0x77, 0x54, 0x38, 0xd9, 0x27, 0x6c, 0x3b, 0x5e, 0xa4,
	0xba, 0x94, 0x2b, 0xe6, 0x7b, 0x8a, 0x01, 0xb7, 0x43, 0x01, 0x0a, 0xc8, 0xba, 0x0f, 0x7e, 0x4f,
	0x80, 0xe7, 0x77, 0xed, 0xcc, 0x68, 0x4f, 0xfc, 0x46, 0x25, 0x80, 0x00, 0x52, 0x97, 0x93, 0xfc,
	0x65, 0x07, 0x8c, 0x8d, 0x00, 0x20, 0xe8, 0x53, 0xc7, 0x0b, 0x99, 0xe3, 0x71, 0x0e, 0x2a, 0x4d,
	0x93, 0x99, 0x6a, 0x9d, 0xe3, 0xd5, 0x2b, 0x49, 0x45, 0x0b, 0x02, 0xc6, 0x5d, 0x7a, 0x1b, 0x51,
	0xa9, 0x88, 0x81, 0x8b, 0x91, 0xa4, 0x82, 0x7b, 0x37, 0xb4, 0x86, 0xb6, 0xd0, 0x6e, 0xc9, 0x9d,
	0xd6, 0x89, 0x16, 0x7a, 0x52, 0xde, 0x83, 0xe8, 0xd4, 0xb4, 0x4c, 0x9b, 0xd4, 0xd6, 0x1a, 0x2e,
	0xcf, 0x65, 0xc9, 0x10, 0xb8, 0xa4, 0x73, 0x4d, 0x88, 0x54, 0x7e, 0x83, 0x55, 0xc1, 0x64, 0xbe,
	0x99, 0x5b, 0x4f, 0xf0, 0xf2, 0x25, 0x95, 0x92, 0x01, 0x3f, 0x05, 0xe8, 0x31, 0x4a, 0xaa, 0x58,
	0x63, 0x9d, 0x14, 0xe1, 0x4f, 0x43, 0x8f, 0xdf, 0x37, 0xb5, 0xe6, 0x99, 0xab, 0xb1, 0x0e, 0xa9,
	0x62, 0x5d, 0x52, 0x5f, 0x50, 0x95, 0x22, 0x2c, 0xb9, 0x79, 0x55, 0x1f, 0x20, 0xfc, 0xb7, 0x05,
	0x41, 0x93, 0x93, 0x07, 0x5c, 0x9a, 0xa2, 0x90, 0x7d, 0xfb, 0xd7, 0x9d, 0xd9, 0x3f, 0x87, 0x37,
	0x0e, 0x16, 0x33, 0xe7, 0xc8, 0xe5, 0xc1, 0xeb, 0xd7, 0x8b, 0xf6, 0xdf, 0xd2, 0x9d, 0x7e, 0xd2,
	0x3f, 0x46, 0x7b, 0xf5, 0x27, 0x84, 0xf5, 0x16, 0x04, 0x17, 0x91, 0x22, 0x8f, 0x18, 0xcf, 0xc6,
	0x24, 0x0b, 0x24, 0xcf, 0x56, 0x64, 0x1c, 0x2e, 0xe8, 0xce, 0x41, 0x56, 0x52, 0x90, 0x12, 0xf9,
	0x97, 0x80, 0x40, 0xa4, 0x1a, 0x3b, 0xc3, 0x4f, 0xb3, 0x30, 0x8c, 0x4d, 0x34, 0x8a, 0x4d, 0xf4,
	0x16, 0x9b, 0xe8, 0x23, 0x36, 0xd1, 0xf3, 0xd8, 0x2c, 0x8c, 0xc6, 0x66, 0xe1, 0xba, 0x38, 0xc9,
	0x6a, 0xeb, 0xe9, 0x3b, 0x38, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x92, 0x6a, 0xf5, 0xe6, 0x7e,
	0x02, 0x00, 0x00,
}
