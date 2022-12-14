// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: mult.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A uint64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B uint64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *MultRequest) Reset() {
	*x = MultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultRequest) ProtoMessage() {}

func (x *MultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultRequest.ProtoReflect.Descriptor instead.
func (*MultRequest) Descriptor() ([]byte, []int) {
	return file_mult_proto_rawDescGZIP(), []int{0}
}

func (x *MultRequest) GetA() uint64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *MultRequest) GetB() uint64 {
	if x != nil {
		return x.B
	}
	return 0
}

type MultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result uint64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MultResponse) Reset() {
	*x = MultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mult_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultResponse) ProtoMessage() {}

func (x *MultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mult_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultResponse.ProtoReflect.Descriptor instead.
func (*MultResponse) Descriptor() ([]byte, []int) {
	return file_mult_proto_rawDescGZIP(), []int{1}
}

func (x *MultResponse) GetResult() uint64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_mult_proto protoreflect.FileDescriptor

var file_mult_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x29, 0x0a, 0x0b, 0x4d, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a,
	0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x01, 0x62, 0x22, 0x26, 0x0a, 0x0c, 0x4d,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x32, 0x3d, 0x0a, 0x0b, 0x4d, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mult_proto_rawDescOnce sync.Once
	file_mult_proto_rawDescData = file_mult_proto_rawDesc
)

func file_mult_proto_rawDescGZIP() []byte {
	file_mult_proto_rawDescOnce.Do(func() {
		file_mult_proto_rawDescData = protoimpl.X.CompressGZIP(file_mult_proto_rawDescData)
	})
	return file_mult_proto_rawDescData
}

var file_mult_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mult_proto_goTypes = []interface{}{
	(*MultRequest)(nil),  // 0: pb.MultRequest
	(*MultResponse)(nil), // 1: pb.MultResponse
}
var file_mult_proto_depIdxs = []int32{
	0, // 0: pb.MultService.Compute:input_type -> pb.MultRequest
	1, // 1: pb.MultService.Compute:output_type -> pb.MultResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mult_proto_init() }
func file_mult_proto_init() {
	if File_mult_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_mult_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mult_proto_goTypes,
		DependencyIndexes: file_mult_proto_depIdxs,
		MessageInfos:      file_mult_proto_msgTypes,
	}.Build()
	File_mult_proto = out.File
	file_mult_proto_rawDesc = nil
	file_mult_proto_goTypes = nil
	file_mult_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MultServiceClient is the client API for MultService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MultServiceClient interface {
	Compute(ctx context.Context, in *MultRequest, opts ...grpc.CallOption) (*MultResponse, error)
}

type multServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMultServiceClient(cc grpc.ClientConnInterface) MultServiceClient {
	return &multServiceClient{cc}
}

func (c *multServiceClient) Compute(ctx context.Context, in *MultRequest, opts ...grpc.CallOption) (*MultResponse, error) {
	out := new(MultResponse)
	err := c.cc.Invoke(ctx, "/pb.MultService/Compute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MultServiceServer is the server API for MultService service.
type MultServiceServer interface {
	Compute(context.Context, *MultRequest) (*MultResponse, error)
}

// UnimplementedMultServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMultServiceServer struct {
}

func (*UnimplementedMultServiceServer) Compute(context.Context, *MultRequest) (*MultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compute not implemented")
}

func RegisterMultServiceServer(s *grpc.Server, srv MultServiceServer) {
	s.RegisterService(&_MultService_serviceDesc, srv)
}

func _MultService_Compute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultServiceServer).Compute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MultService/Compute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultServiceServer).Compute(ctx, req.(*MultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MultService",
	HandlerType: (*MultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compute",
			Handler:    _MultService_Compute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mult.proto",
}
