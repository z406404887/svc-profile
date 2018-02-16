// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/profile.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/profile.proto

It has these top-level messages:
	EmptyMessage
	VersionResponse
	ServiceStatus
	ServicesStatusList
	EchoRequest
	EchoResponse
*/
package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/gomeet/go-proto-gomeetfaker"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ServiceStatus_Status int32

const (
	ServiceStatus_OK          ServiceStatus_Status = 0
	ServiceStatus_UNAVAILABLE ServiceStatus_Status = 1
)

var ServiceStatus_Status_name = map[int32]string{
	0: "OK",
	1: "UNAVAILABLE",
}
var ServiceStatus_Status_value = map[string]int32{
	"OK":          0,
	"UNAVAILABLE": 1,
}

func (x ServiceStatus_Status) String() string {
	return proto.EnumName(ServiceStatus_Status_name, int32(x))
}
func (ServiceStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorProfile, []int{2, 0}
}

type EmptyMessage struct {
}

func (m *EmptyMessage) Reset()                    { *m = EmptyMessage{} }
func (m *EmptyMessage) String() string            { return proto.CompactTextString(m) }
func (*EmptyMessage) ProtoMessage()               {}
func (*EmptyMessage) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{0} }

// VersionMessage represents a version message
type VersionResponse struct {
	// Id represents the message identifier.
	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{1} }

func (m *VersionResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// SeviceStatus represents a sub services status message
type ServiceStatus struct {
	Name    string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string               `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Status  ServiceStatus_Status `protobuf:"varint,3,opt,name=status,proto3,enum=grpc.gomeetexamples.profile.ServiceStatus_Status" json:"status,omitempty"`
	EMsg    string               `protobuf:"bytes,4,opt,name=e_msg,json=eMsg,proto3" json:"e_msg,omitempty"`
}

func (m *ServiceStatus) Reset()                    { *m = ServiceStatus{} }
func (m *ServiceStatus) String() string            { return proto.CompactTextString(m) }
func (*ServiceStatus) ProtoMessage()               {}
func (*ServiceStatus) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{2} }

func (m *ServiceStatus) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceStatus) GetStatus() ServiceStatus_Status {
	if m != nil {
		return m.Status
	}
	return ServiceStatus_OK
}

func (m *ServiceStatus) GetEMsg() string {
	if m != nil {
		return m.EMsg
	}
	return ""
}

// ServicesStatusList is the sub services status list
type ServicesStatusList struct {
	Services []*ServiceStatus `protobuf:"bytes,1,rep,name=services" json:"services,omitempty"`
}

func (m *ServicesStatusList) Reset()                    { *m = ServicesStatusList{} }
func (m *ServicesStatusList) String() string            { return proto.CompactTextString(m) }
func (*ServicesStatusList) ProtoMessage()               {}
func (*ServicesStatusList) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{3} }

func (m *ServicesStatusList) GetServices() []*ServiceStatus {
	if m != nil {
		return m.Services
	}
	return nil
}

// EchoRequest represents a simple message sent to the Echo service.
type EchoRequest struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *EchoRequest) Reset()                    { *m = EchoRequest{} }
func (m *EchoRequest) String() string            { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()               {}
func (*EchoRequest) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{4} }

func (m *EchoRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *EchoRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// EchoResponse represents a simple message that the Echo service return.
type EchoResponse struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *EchoResponse) Reset()                    { *m = EchoResponse{} }
func (m *EchoResponse) String() string            { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()               {}
func (*EchoResponse) Descriptor() ([]byte, []int) { return fileDescriptorProfile, []int{5} }

func (m *EchoResponse) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *EchoResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyMessage)(nil), "grpc.gomeetexamples.profile.EmptyMessage")
	proto.RegisterType((*VersionResponse)(nil), "grpc.gomeetexamples.profile.VersionResponse")
	proto.RegisterType((*ServiceStatus)(nil), "grpc.gomeetexamples.profile.ServiceStatus")
	proto.RegisterType((*ServicesStatusList)(nil), "grpc.gomeetexamples.profile.ServicesStatusList")
	proto.RegisterType((*EchoRequest)(nil), "grpc.gomeetexamples.profile.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "grpc.gomeetexamples.profile.EchoResponse")
	proto.RegisterEnum("grpc.gomeetexamples.profile.ServiceStatus_Status", ServiceStatus_Status_name, ServiceStatus_Status_value)
}
func (this *EmptyMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EmptyMessage)
	if !ok {
		that2, ok := that.(EmptyMessage)
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
	return true
}
func (this *VersionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VersionResponse)
	if !ok {
		that2, ok := that.(VersionResponse)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	return true
}
func (this *ServiceStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServiceStatus)
	if !ok {
		that2, ok := that.(ServiceStatus)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.EMsg != that1.EMsg {
		return false
	}
	return true
}
func (this *ServicesStatusList) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ServicesStatusList)
	if !ok {
		that2, ok := that.(ServicesStatusList)
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
	if len(this.Services) != len(that1.Services) {
		return false
	}
	for i := range this.Services {
		if !this.Services[i].Equal(that1.Services[i]) {
			return false
		}
	}
	return true
}
func (this *EchoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EchoRequest)
	if !ok {
		that2, ok := that.(EchoRequest)
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
	if this.Uuid != that1.Uuid {
		return false
	}
	if this.Content != that1.Content {
		return false
	}
	return true
}
func (this *EchoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EchoResponse)
	if !ok {
		that2, ok := that.(EchoResponse)
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
	if this.Uuid != that1.Uuid {
		return false
	}
	if this.Content != that1.Content {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Profile service

type ProfileClient interface {
	// Version method receives no paramaters and returns a version message.
	Version(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*VersionResponse, error)
	// ServicesStatus method receives no paramaters and returns all services status message
	ServicesStatus(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ServicesStatusList, error)
	// Echo method receives a simple message and returns it.
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type profileClient struct {
	cc *grpc.ClientConn
}

func NewProfileClient(cc *grpc.ClientConn) ProfileClient {
	return &profileClient{cc}
}

func (c *profileClient) Version(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/grpc.gomeetexamples.profile.Profile/Version", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) ServicesStatus(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*ServicesStatusList, error) {
	out := new(ServicesStatusList)
	err := grpc.Invoke(ctx, "/grpc.gomeetexamples.profile.Profile/ServicesStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := grpc.Invoke(ctx, "/grpc.gomeetexamples.profile.Profile/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profile service

type ProfileServer interface {
	// Version method receives no paramaters and returns a version message.
	Version(context.Context, *EmptyMessage) (*VersionResponse, error)
	// ServicesStatus method receives no paramaters and returns all services status message
	ServicesStatus(context.Context, *EmptyMessage) (*ServicesStatusList, error)
	// Echo method receives a simple message and returns it.
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterProfileServer(s *grpc.Server, srv ProfileServer) {
	s.RegisterService(&_Profile_serviceDesc, srv)
}

func _Profile_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gomeetexamples.profile.Profile/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Version(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_ServicesStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).ServicesStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gomeetexamples.profile.Profile/ServicesStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).ServicesStatus(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profile_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gomeetexamples.profile.Profile/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profile_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gomeetexamples.profile.Profile",
	HandlerType: (*ProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _Profile_Version_Handler,
		},
		{
			MethodName: "ServicesStatus",
			Handler:    _Profile_ServicesStatus_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _Profile_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/profile.proto",
}

func NewPopulatedEmptyMessage(r randyProfile, easy bool) *EmptyMessage {
	this := &EmptyMessage{}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedVersionResponse(r randyProfile, easy bool) *VersionResponse {
	this := &VersionResponse{}
	this.Name = string(randStringProfile(r))
	this.Version = string(randStringProfile(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedServiceStatus(r randyProfile, easy bool) *ServiceStatus {
	this := &ServiceStatus{}
	this.Name = string(randStringProfile(r))
	this.Version = string(randStringProfile(r))
	this.Status = ServiceStatus_Status([]int32{0, 1}[r.Intn(2)])
	this.EMsg = string(randStringProfile(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedServicesStatusList(r randyProfile, easy bool) *ServicesStatusList {
	this := &ServicesStatusList{}
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Services = make([]*ServiceStatus, v1)
		for i := 0; i < v1; i++ {
			this.Services[i] = NewPopulatedServiceStatus(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedEchoRequest(r randyProfile, easy bool) *EchoRequest {
	this := &EchoRequest{}
	this.Uuid = string(randStringProfile(r))
	this.Content = string(randStringProfile(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedEchoResponse(r randyProfile, easy bool) *EchoResponse {
	this := &EchoResponse{}
	this.Uuid = string(randStringProfile(r))
	this.Content = string(randStringProfile(r))
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyProfile interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneProfile(r randyProfile) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringProfile(r randyProfile) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneProfile(r)
	}
	return string(tmps)
}
func randUnrecognizedProfile(r randyProfile, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldProfile(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldProfile(dAtA []byte, r randyProfile, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateProfile(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateProfile(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

func init() { proto.RegisterFile("pb/profile.proto", fileDescriptorProfile) }

var fileDescriptorProfile = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xd3, 0x3c,
	0x18, 0xfe, 0x9c, 0x66, 0xe9, 0xe6, 0xed, 0xdb, 0x8a, 0x2f, 0x0b, 0x65, 0x94, 0x11, 0x2e, 0xdd,
	0xb4, 0xc6, 0x5a, 0x41, 0x08, 0xc1, 0xa9, 0x95, 0x0a, 0x9a, 0x68, 0xa1, 0x4a, 0x45, 0x0f, 0x08,
	0x09, 0x92, 0xcc, 0x4d, 0xa3, 0x35, 0x71, 0x88, 0x9d, 0x0c, 0xc4, 0x6d, 0x67, 0x38, 0xf1, 0x0f,
	0x38, 0x71, 0xe8, 0x0f, 0xd8, 0x6f, 0xe8, 0x95, 0x33, 0x07, 0x26, 0x7e, 0x04, 0x47, 0x84, 0xe3,
	0x6c, 0x2d, 0x88, 0xaa, 0x08, 0x4e, 0xb1, 0x5f, 0x3f, 0xcf, 0xfb, 0x3e, 0x7e, 0x9f, 0x37, 0x86,
	0xa5, 0xc8, 0xc1, 0x51, 0x4c, 0x07, 0xfe, 0x88, 0x98, 0x51, 0x4c, 0x39, 0x45, 0x57, 0xbc, 0x38,
	0x72, 0x4d, 0x8f, 0x06, 0x84, 0x70, 0xf2, 0xca, 0x0e, 0xa2, 0x11, 0x61, 0xa6, 0x84, 0x94, 0xb7,
	0x3c, 0x4a, 0xbd, 0x11, 0xc1, 0x76, 0xe4, 0x63, 0x3b, 0x0c, 0x29, 0xb7, 0xb9, 0x4f, 0x43, 0x96,
	0x51, 0xcb, 0xb7, 0x3d, 0x9f, 0x0f, 0x13, 0xc7, 0x74, 0x69, 0x80, 0x83, 0x63, 0x9f, 0x1f, 0xd1,
	0x63, 0xec, 0xd1, 0x9a, 0x38, 0xac, 0xa5, 0xf6, 0xc8, 0x3f, 0xb4, 0x39, 0x8d, 0x19, 0x3e, 0x5f,
	0x4a, 0xde, 0x9d, 0x29, 0x5e, 0x56, 0xf8, 0x82, 0x96, 0xed, 0x07, 0xf6, 0x11, 0x89, 0xf1, 0xd4,
	0x5a, 0x32, 0x6b, 0x33, 0x4c, 0x8f, 0x62, 0x11, 0x76, 0x92, 0x81, 0xd8, 0x89, 0x8d, 0x58, 0x65,
	0x70, 0x63, 0x1d, 0xae, 0xb5, 0x82, 0x88, 0xbf, 0xee, 0x10, 0xc6, 0x6c, 0x8f, 0x18, 0x3d, 0xb8,
	0xd1, 0x27, 0x31, 0xf3, 0x69, 0x68, 0x11, 0x16, 0xd1, 0x90, 0x11, 0xb4, 0x05, 0xd5, 0xd0, 0x0e,
	0x88, 0x0e, 0xb6, 0x41, 0x75, 0xa5, 0xb9, 0x3c, 0x19, 0x6b, 0xea, 0xae, 0x52, 0x02, 0x96, 0x88,
	0x22, 0x03, 0x16, 0xd3, 0x8c, 0xa0, 0x2b, 0xd3, 0x00, 0x1d, 0x58, 0xf9, 0x81, 0xf1, 0x19, 0xc0,
	0xff, 0x7b, 0x24, 0x4e, 0x7d, 0x97, 0xf4, 0xb8, 0xcd, 0x13, 0xf6, 0xf7, 0x39, 0xd1, 0x01, 0xd4,
	0x98, 0xc8, 0xa5, 0x17, 0xb6, 0x41, 0x75, 0xbd, 0xbe, 0x6f, 0xce, 0x71, 0xc9, 0x9c, 0xa9, 0x6e,
	0x66, 0x1f, 0x4b, 0x26, 0x40, 0x5b, 0x70, 0x89, 0x3c, 0x0f, 0x98, 0xa7, 0xab, 0xa2, 0x58, 0x71,
	0x32, 0xd6, 0x0a, 0xa7, 0x00, 0x58, 0x2a, 0xe9, 0x30, 0xcf, 0xb8, 0x0e, 0x35, 0x29, 0x5a, 0x83,
	0xca, 0xe3, 0x87, 0xa5, 0xff, 0xd0, 0x06, 0x5c, 0x7d, 0xf2, 0xa8, 0xd1, 0x6f, 0x1c, 0xb4, 0x1b,
	0xcd, 0x76, 0xab, 0x04, 0x8c, 0x67, 0x10, 0xc9, 0x02, 0x2c, 0x83, 0xb6, 0x7d, 0xc6, 0xd1, 0x7d,
	0xb8, 0xcc, 0x64, 0x54, 0x07, 0xdb, 0x85, 0xea, 0x6a, 0x7d, 0x77, 0x71, 0x8d, 0xd6, 0x39, 0xd7,
	0xb0, 0xe0, 0x6a, 0xcb, 0x1d, 0x52, 0x8b, 0xbc, 0x4c, 0x08, 0xe3, 0xa8, 0x02, 0xd5, 0x24, 0xf1,
	0x0f, 0x65, 0xeb, 0xe0, 0x64, 0xac, 0x69, 0x48, 0x45, 0x4a, 0xff, 0x96, 0x25, 0xe2, 0xe8, 0x06,
	0x2c, 0xba, 0x34, 0xe4, 0x24, 0xe4, 0xb2, 0x79, 0x2b, 0x93, 0xb1, 0xb6, 0x74, 0x02, 0x94, 0x2e,
	0xb0, 0xf2, 0x13, 0xa3, 0x07, 0xd7, 0xb2, 0x9c, 0xd2, 0xe3, 0x7f, 0x91, 0xb4, 0xfe, 0xb6, 0x00,
	0x8b, 0xdd, 0xec, 0x32, 0xe8, 0x0d, 0x2c, 0xca, 0x39, 0x42, 0x3b, 0x73, 0x6f, 0x3d, 0x3d, 0x7d,
	0xe5, 0xbd, 0xb9, 0xd0, 0x9f, 0x06, 0xd3, 0xd8, 0x3c, 0xf9, 0xf4, 0xf5, 0xbd, 0x72, 0x09, 0x6d,
	0x88, 0x9f, 0x2f, 0xdd, 0xc7, 0xf9, 0x6c, 0xbc, 0x03, 0x70, 0x7d, 0xd6, 0x90, 0x3f, 0x11, 0x81,
	0x17, 0x71, 0x69, 0xca, 0x68, 0xe3, 0x9a, 0xd0, 0x71, 0x19, 0x6d, 0xe6, 0x3a, 0x72, 0xeb, 0xb0,
	0x1c, 0x30, 0x0e, 0xd5, 0x1f, 0xdd, 0x46, 0xd5, 0xf9, 0x22, 0x2e, 0x4c, 0x2e, 0xef, 0x2c, 0x80,
	0x9c, 0xed, 0x82, 0xb1, 0x96, 0x57, 0x27, 0xee, 0x90, 0xde, 0x05, 0xbb, 0xcd, 0x17, 0xdf, 0xbe,
	0x54, 0xc0, 0xc7, 0xb3, 0x0a, 0x38, 0x3d, 0xab, 0x00, 0x78, 0xd5, 0xa5, 0xc1, 0xef, 0xf2, 0xa5,
	0xfb, 0x4f, 0xf7, 0x7e, 0x79, 0x70, 0x6a, 0x39, 0x0c, 0xb3, 0xd4, 0xad, 0x49, 0x28, 0x8e, 0x9c,
	0x7b, 0x91, 0xf3, 0x41, 0x51, 0x1f, 0x74, 0xba, 0x4d, 0x47, 0x13, 0x6f, 0xc8, 0xcd, 0xef, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xb9, 0x83, 0x24, 0xd9, 0x33, 0x05, 0x00, 0x00,
}
