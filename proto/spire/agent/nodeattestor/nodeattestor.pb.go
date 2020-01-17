// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nodeattestor.proto

package nodeattestor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
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

//* Represents an empty request
type FetchAttestationDataRequest struct {
	Challenge            []byte   `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAttestationDataRequest) Reset()         { *m = FetchAttestationDataRequest{} }
func (m *FetchAttestationDataRequest) String() string { return proto.CompactTextString(m) }
func (*FetchAttestationDataRequest) ProtoMessage()    {}
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3b2582c38c076c, []int{0}
}

func (m *FetchAttestationDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAttestationDataRequest.Unmarshal(m, b)
}
func (m *FetchAttestationDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAttestationDataRequest.Marshal(b, m, deterministic)
}
func (m *FetchAttestationDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAttestationDataRequest.Merge(m, src)
}
func (m *FetchAttestationDataRequest) XXX_Size() int {
	return xxx_messageInfo_FetchAttestationDataRequest.Size(m)
}
func (m *FetchAttestationDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAttestationDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAttestationDataRequest proto.InternalMessageInfo

func (m *FetchAttestationDataRequest) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

//* Represents the attested data and base SPIFFE ID
type FetchAttestationDataResponse struct {
	//* A type which contains attestation data for specific platform
	AttestationData *common.AttestationData `protobuf:"bytes,1,opt,name=attestation_data,json=attestationData,proto3" json:"attestation_data,omitempty"`
	//* SPIFFE ID for the agent. This field is deprecated and should no longer
	//be set by implementers. The server-side plugin is now solely in charge
	//of determining the SPIFFE ID for the agent
	DEPRECATEDSpiffeId string `protobuf:"bytes,2,opt,name=DEPRECATED_spiffe_id,json=DEPRECATEDSpiffeId,proto3" json:"DEPRECATED_spiffe_id,omitempty"`
	//* response to the challenge (if challenge was present) *
	Response             []byte   `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FetchAttestationDataResponse) Reset()         { *m = FetchAttestationDataResponse{} }
func (m *FetchAttestationDataResponse) String() string { return proto.CompactTextString(m) }
func (*FetchAttestationDataResponse) ProtoMessage()    {}
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e3b2582c38c076c, []int{1}
}

func (m *FetchAttestationDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FetchAttestationDataResponse.Unmarshal(m, b)
}
func (m *FetchAttestationDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FetchAttestationDataResponse.Marshal(b, m, deterministic)
}
func (m *FetchAttestationDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchAttestationDataResponse.Merge(m, src)
}
func (m *FetchAttestationDataResponse) XXX_Size() int {
	return xxx_messageInfo_FetchAttestationDataResponse.Size(m)
}
func (m *FetchAttestationDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchAttestationDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FetchAttestationDataResponse proto.InternalMessageInfo

func (m *FetchAttestationDataResponse) GetAttestationData() *common.AttestationData {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func (m *FetchAttestationDataResponse) GetDEPRECATEDSpiffeId() string {
	if m != nil {
		return m.DEPRECATEDSpiffeId
	}
	return ""
}

func (m *FetchAttestationDataResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchAttestationDataRequest)(nil), "spire.agent.nodeattestor.FetchAttestationDataRequest")
	proto.RegisterType((*FetchAttestationDataResponse)(nil), "spire.agent.nodeattestor.FetchAttestationDataResponse")
}

func init() { proto.RegisterFile("nodeattestor.proto", fileDescriptor_9e3b2582c38c076c) }

var fileDescriptor_9e3b2582c38c076c = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xdd, 0x4e, 0xf2, 0x30,
	0x18, 0xc7, 0x53, 0xde, 0xe4, 0x8d, 0x54, 0x8c, 0xa6, 0xe1, 0x60, 0x4e, 0x4c, 0x08, 0x89, 0x06,
	0x3d, 0xd8, 0x08, 0x46, 0x62, 0xe2, 0x11, 0x02, 0x2a, 0x27, 0x86, 0x4c, 0x8f, 0x38, 0x21, 0x85,
	0x3d, 0x1b, 0x4d, 0xa0, 0x9d, 0x6b, 0x77, 0x11, 0xde, 0x90, 0x97, 0xe3, 0xb5, 0x18, 0xdb, 0xf2,
	0x31, 0x33, 0xbf, 0x8e, 0xba, 0xf5, 0xf7, 0x7f, 0x3e, 0xfa, 0x7f, 0x1e, 0x4c, 0xb8, 0x08, 0x81,
	0x2a, 0x05, 0x52, 0x89, 0xd4, 0x4b, 0x52, 0xa1, 0x04, 0x71, 0x64, 0xc2, 0x52, 0xf0, 0x68, 0x0c,
	0x5c, 0x79, 0xdb, 0xdc, 0x3d, 0xd4, 0xc4, 0x9f, 0x89, 0xe5, 0x52, 0x70, 0x7b, 0x98, 0x20, 0xb7,
	0x9e, 0x43, 0xc9, 0x22, 0x8b, 0xd9, 0xea, 0x30, 0x8a, 0xc6, 0x35, 0x3e, 0xba, 0x05, 0x35, 0x9b,
	0x77, 0x75, 0x36, 0xaa, 0x98, 0xe0, 0x7d, 0xaa, 0x68, 0x00, 0xcf, 0x19, 0x48, 0x45, 0x6a, 0xb8,
	0x3c, 0x9b, 0xd3, 0xc5, 0x02, 0x78, 0x0c, 0x0e, 0xaa, 0xa3, 0x66, 0x25, 0xd8, 0x5c, 0x34, 0x5e,
	0x11, 0xae, 0x15, 0x47, 0xcb, 0x44, 0x70, 0x09, 0xe4, 0x1e, 0x1f, 0xd0, 0x0d, 0x9a, 0x84, 0x54,
	0x51, 0x9d, 0x65, 0xb7, 0x7d, 0xec, 0x99, 0xf7, 0xd8, 0x76, 0x3f, 0x27, 0xd8, 0xa7, 0xf9, 0x0b,
	0xd2, 0xc2, 0xd5, 0xfe, 0x60, 0x14, 0x0c, 0x7a, 0xdd, 0xa7, 0x41, 0x7f, 0x22, 0x13, 0x16, 0x45,
	0x30, 0x61, 0xa1, 0x53, 0xaa, 0xa3, 0x66, 0x39, 0x20, 0x1b, 0xf6, 0xa8, 0xd1, 0x30, 0x24, 0x2e,
	0xde, 0x49, 0x6d, 0x1f, 0xce, 0x3f, 0xdd, 0xf9, 0xfa, 0xbf, 0xfd, 0x56, 0xc2, 0x95, 0x07, 0x11,
	0x42, 0xd7, 0x7a, 0x48, 0x5e, 0x10, 0xae, 0x16, 0xbd, 0x84, 0x5c, 0x7a, 0x5f, 0xf9, 0xee, 0x7d,
	0xe3, 0x9b, 0xdb, 0xf9, 0x6b, 0x98, 0x69, 0xac, 0x89, 0x5a, 0x88, 0x8c, 0x71, 0xb9, 0x27, 0x78,
	0xc4, 0xe2, 0x2c, 0x05, 0x72, 0x92, 0xf7, 0xc9, 0xce, 0x6e, 0xcd, 0x57, 0xf5, 0x4e, 0x7f, 0x92,
	0xd9, 0x81, 0x44, 0x78, 0xef, 0x0e, 0xd4, 0x48, 0xe3, 0x21, 0x8f, 0x04, 0x39, 0x2b, 0x0c, 0xcc,
	0x69, 0x56, 0x35, 0xce, 0x7f, 0x23, 0x35, 0x75, 0x6e, 0xae, 0xc6, 0x9d, 0x98, 0xa9, 0x79, 0x36,
	0xfd, 0x50, 0xfb, 0x66, 0x5c, 0xbe, 0x59, 0x46, 0xbd, 0x77, 0xf6, 0x5b, 0xdb, 0xe3, 0x6f, 0xdb,
	0x33, 0xfd, 0xaf, 0xf9, 0xc5, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x19, 0x4b, 0x4e, 0xed, 0x04,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeAttestorClient is the client API for NodeAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAttestorClient interface {
	//* Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation
	FetchAttestationData(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_FetchAttestationDataClient, error)
	//* Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc *grpc.ClientConn
}

func NewNodeAttestorClient(cc *grpc.ClientConn) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) FetchAttestationData(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_FetchAttestationDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAttestor_serviceDesc.Streams[0], "/spire.agent.nodeattestor.NodeAttestor/FetchAttestationData", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttestorFetchAttestationDataClient{stream}
	return x, nil
}

type NodeAttestor_FetchAttestationDataClient interface {
	Send(*FetchAttestationDataRequest) error
	Recv() (*FetchAttestationDataResponse, error)
	grpc.ClientStream
}

type nodeAttestorFetchAttestationDataClient struct {
	grpc.ClientStream
}

func (x *nodeAttestorFetchAttestationDataClient) Send(m *FetchAttestationDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeAttestorFetchAttestationDataClient) Recv() (*FetchAttestationDataResponse, error) {
	m := new(FetchAttestationDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAttestorServer is the server API for NodeAttestor service.
type NodeAttestorServer interface {
	//* Returns the node attestation data for specific platform and the generated Base SPIFFE ID for CSR formation
	FetchAttestationData(NodeAttestor_FetchAttestationDataServer) error
	//* Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

// UnimplementedNodeAttestorServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAttestorServer struct {
}

func (*UnimplementedNodeAttestorServer) FetchAttestationData(srv NodeAttestor_FetchAttestationDataServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchAttestationData not implemented")
}
func (*UnimplementedNodeAttestorServer) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedNodeAttestorServer) GetPluginInfo(ctx context.Context, req *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_FetchAttestationData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeAttestorServer).FetchAttestationData(&nodeAttestorFetchAttestationDataServer{stream})
}

type NodeAttestor_FetchAttestationDataServer interface {
	Send(*FetchAttestationDataResponse) error
	Recv() (*FetchAttestationDataRequest, error)
	grpc.ServerStream
}

type nodeAttestorFetchAttestationDataServer struct {
	grpc.ServerStream
}

func (x *nodeAttestorFetchAttestationDataServer) Send(m *FetchAttestationDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeAttestorFetchAttestationDataServer) Recv() (*FetchAttestationDataRequest, error) {
	m := new(FetchAttestationDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchAttestationData",
			Handler:       _NodeAttestor_FetchAttestationData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "nodeattestor.proto",
}