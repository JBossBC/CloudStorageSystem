// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: pb/fileServer.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FileServerClient is the client API for FileServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServerClient interface {
	FindOne(ctx context.Context, in *FindFileReq, opts ...grpc.CallOption) (*FileMetaInfo, error)
	QueryFiles(ctx context.Context, in *QueryFileReq, opts ...grpc.CallOption) (*QueryFileRes, error)
	InertOne(ctx context.Context, in *FileMetaInfo, opts ...grpc.CallOption) (*BaseRes, error)
	DeleteOne(ctx context.Context, in *FileMetaInfo, opts ...grpc.CallOption) (*BaseRes, error)
}

type fileServerClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServerClient(cc grpc.ClientConnInterface) FileServerClient {
	return &fileServerClient{cc}
}

func (c *fileServerClient) FindOne(ctx context.Context, in *FindFileReq, opts ...grpc.CallOption) (*FileMetaInfo, error) {
	out := new(FileMetaInfo)
	err := c.cc.Invoke(ctx, "/pb.fileServer/FindOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerClient) QueryFiles(ctx context.Context, in *QueryFileReq, opts ...grpc.CallOption) (*QueryFileRes, error) {
	out := new(QueryFileRes)
	err := c.cc.Invoke(ctx, "/pb.fileServer/QueryFiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerClient) InertOne(ctx context.Context, in *FileMetaInfo, opts ...grpc.CallOption) (*BaseRes, error) {
	out := new(BaseRes)
	err := c.cc.Invoke(ctx, "/pb.fileServer/InertOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServerClient) DeleteOne(ctx context.Context, in *FileMetaInfo, opts ...grpc.CallOption) (*BaseRes, error) {
	out := new(BaseRes)
	err := c.cc.Invoke(ctx, "/pb.fileServer/DeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServerServer is the server API for FileServer service.
// All implementations must embed UnimplementedFileServerServer
// for forward compatibility
type FileServerServer interface {
	FindOne(context.Context, *FindFileReq) (*FileMetaInfo, error)
	QueryFiles(context.Context, *QueryFileReq) (*QueryFileRes, error)
	InertOne(context.Context, *FileMetaInfo) (*BaseRes, error)
	DeleteOne(context.Context, *FileMetaInfo) (*BaseRes, error)
	mustEmbedUnimplementedFileServerServer()
}

// UnimplementedFileServerServer must be embedded to have forward compatible implementations.
type UnimplementedFileServerServer struct {
}

func (UnimplementedFileServerServer) FindOne(context.Context, *FindFileReq) (*FileMetaInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOne not implemented")
}
func (UnimplementedFileServerServer) QueryFiles(context.Context, *QueryFileReq) (*QueryFileRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFiles not implemented")
}
func (UnimplementedFileServerServer) InertOne(context.Context, *FileMetaInfo) (*BaseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InertOne not implemented")
}
func (UnimplementedFileServerServer) DeleteOne(context.Context, *FileMetaInfo) (*BaseRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOne not implemented")
}
func (UnimplementedFileServerServer) mustEmbedUnimplementedFileServerServer() {}

// UnsafeFileServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServerServer will
// result in compilation errors.
type UnsafeFileServerServer interface {
	mustEmbedUnimplementedFileServerServer()
}

func RegisterFileServerServer(s grpc.ServiceRegistrar, srv FileServerServer) {
	s.RegisterService(&FileServer_ServiceDesc, srv)
}

func _FileServer_FindOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerServer).FindOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.fileServer/FindOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerServer).FindOne(ctx, req.(*FindFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServer_QueryFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerServer).QueryFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.fileServer/QueryFiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerServer).QueryFiles(ctx, req.(*QueryFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServer_InertOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMetaInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerServer).InertOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.fileServer/InertOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerServer).InertOne(ctx, req.(*FileMetaInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileServer_DeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileMetaInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServerServer).DeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.fileServer/DeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServerServer).DeleteOne(ctx, req.(*FileMetaInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// FileServer_ServiceDesc is the grpc.ServiceDesc for FileServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.fileServer",
	HandlerType: (*FileServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindOne",
			Handler:    _FileServer_FindOne_Handler,
		},
		{
			MethodName: "QueryFiles",
			Handler:    _FileServer_QueryFiles_Handler,
		},
		{
			MethodName: "InertOne",
			Handler:    _FileServer_InertOne_Handler,
		},
		{
			MethodName: "DeleteOne",
			Handler:    _FileServer_DeleteOne_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/fileServer.proto",
}
