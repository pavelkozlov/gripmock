// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: wkt.proto

package well_known_types

import (
	context "context"
	api "google.golang.org/genproto/protobuf/api"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	empty "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_wkt_proto protoreflect.FileDescriptor

var file_wkt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x6b, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x77, 0x65, 0x6c,
	0x6c, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x43, 0x0a, 0x08, 0x47, 0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63,
	0x6b, 0x12, 0x37, 0x0a, 0x07, 0x41, 0x70, 0x69, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x70, 0x69, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64,
	0x69, 0x61, 0x2f, 0x67, 0x72, 0x69, 0x70, 0x6d, 0x6f, 0x63, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x77, 0x65, 0x6c,
	0x6c, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_wkt_proto_goTypes = []interface{}{
	(*empty.Empty)(nil), // 0: google.protobuf.Empty
	(*api.Api)(nil),     // 1: google.protobuf.Api
}
var file_wkt_proto_depIdxs = []int32{
	0, // 0: well_known_types.Gripmock.ApiInfo:input_type -> google.protobuf.Empty
	1, // 1: well_known_types.Gripmock.ApiInfo:output_type -> google.protobuf.Api
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_wkt_proto_init() }
func file_wkt_proto_init() {
	if File_wkt_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wkt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wkt_proto_goTypes,
		DependencyIndexes: file_wkt_proto_depIdxs,
	}.Build()
	File_wkt_proto = out.File
	file_wkt_proto_rawDesc = nil
	file_wkt_proto_goTypes = nil
	file_wkt_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripmockClient interface {
	// this shows us example on using WKT as dependency
	// api.proto in particular has go_package alias with semicolon
	// "google.golang.org/genproto/protobuf/api;api"
	ApiInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.Api, error)
}

type gripmockClient struct {
	cc grpc.ClientConnInterface
}

func NewGripmockClient(cc grpc.ClientConnInterface) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) ApiInfo(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*api.Api, error) {
	out := new(api.Api)
	err := c.cc.Invoke(ctx, "/well_known_types.Gripmock/ApiInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
type GripmockServer interface {
	// this shows us example on using WKT as dependency
	// api.proto in particular has go_package alias with semicolon
	// "google.golang.org/genproto/protobuf/api;api"
	ApiInfo(context.Context, *empty.Empty) (*api.Api, error)
}

// UnimplementedGripmockServer can be embedded to have forward compatible implementations.
type UnimplementedGripmockServer struct {
}

func (*UnimplementedGripmockServer) ApiInfo(context.Context, *empty.Empty) (*api.Api, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApiInfo not implemented")
}

func RegisterGripmockServer(s *grpc.Server, srv GripmockServer) {
	s.RegisterService(&_Gripmock_serviceDesc, srv)
}

func _Gripmock_ApiInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).ApiInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/well_known_types.Gripmock/ApiInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).ApiInfo(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "well_known_types.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApiInfo",
			Handler:    _Gripmock_ApiInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wkt.proto",
}
