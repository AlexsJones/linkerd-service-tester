// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto
package proto
import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)
// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package
type SendMessageRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}
var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo
func (m *SendMessageRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}
type SendMessageResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
func (m *SendMessageResponse) Reset()         { *m = SendMessageResponse{} }
func (m *SendMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SendMessageResponse) ProtoMessage()    {}
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}
func (m *SendMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageResponse.Unmarshal(m, b)
}
func (m *SendMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageResponse.Marshal(b, m, deterministic)
}
func (m *SendMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageResponse.Merge(m, src)
}
func (m *SendMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SendMessageResponse.Size(m)
}
func (m *SendMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageResponse.DiscardUnknown(m)
}
var xxx_messageInfo_SendMessageResponse proto.InternalMessageInfo
func (m *SendMessageResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}
func init() {
	proto.RegisterType((*SendMessageRequest)(nil), "proto.SendMessageRequest")
	proto.RegisterType((*SendMessageResponse)(nil), "proto.SendMessageResponse")
}
func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }
var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 128 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x7a, 0x5c,
	0x42, 0xc1, 0xa9, 0x79, 0x29, 0xbe, 0x10, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x09, 0x2e, 0x76, 0xa8, 0x88, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0x64, 0xc8,
	0x25, 0x8c, 0xa2, 0xbe, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8a, 0x8b, 0x03, 0xc6, 0x86,
	0xea, 0x80, 0xf3, 0x8d, 0x02, 0xe1, 0x86, 0x09, 0xb9, 0x71, 0x71, 0x23, 0xe9, 0x16, 0x92, 0x84,
	0xb8, 0x45, 0x0f, 0xd3, 0x05, 0x52, 0x52, 0xd8, 0xa4, 0x20, 0x06, 0x2a, 0x31, 0x24, 0xb1, 0x81,
	0x25, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0x97, 0xfb, 0x75, 0xd4, 0x00, 0x00, 0x00,
}
// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn
// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4
// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	// Sends a greeting
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
}
type messageClient struct {
	cc *grpc.ClientConn
}
func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}
func (c *messageClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/proto.Message/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
// MessageServer is the server API for Message service.
type MessageServer interface {
	// Sends a greeting
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
}
// UnimplementedMessageServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}
func (*UnimplementedMessageServer) SendMessage(ctx context.Context, req *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}
func _Message_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Message/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _Message_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}