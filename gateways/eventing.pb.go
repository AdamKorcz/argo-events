// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateways/eventing.proto

package gateways

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

//*
// Represents an event source
type EventSource struct {
	// ID of the event source. internally generated.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The event source name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The event source configuration value.
	Data string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// Version of the event source
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSource) Reset()         { *m = EventSource{} }
func (m *EventSource) String() string { return proto.CompactTextString(m) }
func (*EventSource) ProtoMessage()    {}
func (*EventSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{0}
}

func (m *EventSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSource.Unmarshal(m, b)
}
func (m *EventSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSource.Marshal(b, m, deterministic)
}
func (m *EventSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSource.Merge(m, src)
}
func (m *EventSource) XXX_Size() int {
	return xxx_messageInfo_EventSource.Size(m)
}
func (m *EventSource) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSource.DiscardUnknown(m)
}

var xxx_messageInfo_EventSource proto.InternalMessageInfo

func (m *EventSource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventSource) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *EventSource) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

//*
// Represents an event
type Event struct {
	// The event source name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The event payload.
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

//*
// Represents if an event source is valid or not
type ValidEventSource struct {
	// whether event source is valid
	IsValid bool `protobuf:"varint,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
	// reason if an event source is invalid
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidEventSource) Reset()         { *m = ValidEventSource{} }
func (m *ValidEventSource) String() string { return proto.CompactTextString(m) }
func (*ValidEventSource) ProtoMessage()    {}
func (*ValidEventSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c25325013aefc28a, []int{2}
}

func (m *ValidEventSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidEventSource.Unmarshal(m, b)
}
func (m *ValidEventSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidEventSource.Marshal(b, m, deterministic)
}
func (m *ValidEventSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidEventSource.Merge(m, src)
}
func (m *ValidEventSource) XXX_Size() int {
	return xxx_messageInfo_ValidEventSource.Size(m)
}
func (m *ValidEventSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidEventSource.DiscardUnknown(m)
}

var xxx_messageInfo_ValidEventSource proto.InternalMessageInfo

func (m *ValidEventSource) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *ValidEventSource) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func init() {
	proto.RegisterType((*EventSource)(nil), "gateways.EventSource")
	proto.RegisterType((*Event)(nil), "gateways.Event")
	proto.RegisterType((*ValidEventSource)(nil), "gateways.ValidEventSource")
}

func init() { proto.RegisterFile("gateways/eventing.proto", fileDescriptor_c25325013aefc28a) }

var fileDescriptor_c25325013aefc28a = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xc9, 0x5a, 0xbb, 0xeb, 0x28, 0x5a, 0x46, 0xd4, 0xd0, 0x93, 0xec, 0xc9, 0xd3, 0x2a,
	0x8a, 0x37, 0x8f, 0x16, 0x3c, 0x6f, 0xc1, 0xab, 0x8c, 0x66, 0x28, 0x81, 0x9a, 0x94, 0x6c, 0xac,
	0xf4, 0x35, 0x7c, 0x62, 0xd9, 0x71, 0xa3, 0x61, 0x4f, 0xde, 0xe6, 0xff, 0x06, 0xbe, 0x99, 0x49,
	0xe0, 0x62, 0x45, 0x91, 0x3f, 0x69, 0xd7, 0x5d, 0xf3, 0x96, 0x5d, 0xb4, 0x6e, 0xd5, 0x6c, 0x82,
	0x8f, 0x1e, 0xab, 0xd4, 0xa8, 0x5f, 0xe0, 0x70, 0xd1, 0xf7, 0x96, 0xfe, 0x23, 0xbc, 0x31, 0x1e,
	0x43, 0x61, 0x8d, 0x56, 0x97, 0xea, 0xea, 0xa0, 0x2d, 0xac, 0x41, 0x84, 0x89, 0xa3, 0x77, 0xd6,
	0x85, 0x10, 0xa9, 0x7b, 0x66, 0x28, 0x92, 0xde, 0xfb, 0x61, 0x7d, 0x8d, 0x1a, 0xca, 0x2d, 0x87,
	0xce, 0x7a, 0xa7, 0x27, 0x82, 0x53, 0xac, 0xef, 0x61, 0x5f, 0x06, 0xfc, 0xaa, 0x54, 0xa6, 0xd2,
	0x50, 0x6e, 0x68, 0xb7, 0xf6, 0x64, 0x64, 0xc2, 0x51, 0x9b, 0x62, 0xfd, 0x08, 0xb3, 0x67, 0x5a,
	0x5b, 0x93, 0x2f, 0xa7, 0xa1, 0xb4, 0x9d, 0x50, 0x91, 0x54, 0x6d, 0x8a, 0x78, 0x0e, 0xd3, 0xc0,
	0xd4, 0x79, 0x37, 0x2c, 0x3a, 0xa4, 0xdb, 0x2f, 0x05, 0xd5, 0x62, 0x38, 0x1d, 0x1f, 0x60, 0xb6,
	0x8c, 0x14, 0x62, 0xae, 0x3c, 0x6b, 0xd2, 0x4b, 0x34, 0x19, 0x9e, 0x9f, 0x8c, 0xf0, 0x8d, 0xc2,
	0x27, 0x38, 0x95, 0x59, 0x14, 0xf9, 0x1f, 0x82, 0xf9, 0x1f, 0x1e, 0x9f, 0xf1, 0x3a, 0x95, 0x3f,
	0xb8, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x08, 0x55, 0x59, 0x7f, 0x9e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventingClient is the client API for Eventing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventingClient interface {
	// StartEventSource starts an event source and returns stream of events.
	StartEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (Eventing_StartEventSourceClient, error)
	// ValidateEventSource validates an event source.
	ValidateEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (*ValidEventSource, error)
}

type eventingClient struct {
	cc *grpc.ClientConn
}

func NewEventingClient(cc *grpc.ClientConn) EventingClient {
	return &eventingClient{cc}
}

func (c *eventingClient) StartEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (Eventing_StartEventSourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Eventing_serviceDesc.Streams[0], "/gateways.Eventing/StartEventSource", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventingStartEventSourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Eventing_StartEventSourceClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventingStartEventSourceClient struct {
	grpc.ClientStream
}

func (x *eventingStartEventSourceClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventingClient) ValidateEventSource(ctx context.Context, in *EventSource, opts ...grpc.CallOption) (*ValidEventSource, error) {
	out := new(ValidEventSource)
	err := c.cc.Invoke(ctx, "/gateways.Eventing/ValidateEventSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventingServer is the server API for Eventing service.
type EventingServer interface {
	// StartEventSource starts an event source and returns stream of events.
	StartEventSource(*EventSource, Eventing_StartEventSourceServer) error
	// ValidateEventSource validates an event source.
	ValidateEventSource(context.Context, *EventSource) (*ValidEventSource, error)
}

func RegisterEventingServer(s *grpc.Server, srv EventingServer) {
	s.RegisterService(&_Eventing_serviceDesc, srv)
}

func _Eventing_StartEventSource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventSource)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventingServer).StartEventSource(m, &eventingStartEventSourceServer{stream})
}

type Eventing_StartEventSourceServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventingStartEventSourceServer struct {
	grpc.ServerStream
}

func (x *eventingStartEventSourceServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Eventing_ValidateEventSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventSource)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventingServer).ValidateEventSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateways.Eventing/ValidateEventSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventingServer).ValidateEventSource(ctx, req.(*EventSource))
	}
	return interceptor(ctx, in, info, handler)
}

var _Eventing_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateways.Eventing",
	HandlerType: (*EventingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateEventSource",
			Handler:    _Eventing_ValidateEventSource_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartEventSource",
			Handler:       _Eventing_StartEventSource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gateways/eventing.proto",
}