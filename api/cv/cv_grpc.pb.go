// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0
// source: cv.proto

package cv

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

const (
	CvService_Upload_FullMethodName = "/cv.CvService/Upload"
	CvService_GetOne_FullMethodName = "/cv.CvService/GetOne"
	CvService_GetAll_FullMethodName = "/cv.CvService/GetAll"
)

// CvServiceClient is the client API for CvService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CvServiceClient interface {
	Upload(ctx context.Context, in *CV, opts ...grpc.CallOption) (*Empty, error)
	GetOne(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CV, error)
	GetAll(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*GetAllResponse, error)
}

type cvServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCvServiceClient(cc grpc.ClientConnInterface) CvServiceClient {
	return &cvServiceClient{cc}
}

func (c *cvServiceClient) Upload(ctx context.Context, in *CV, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, CvService_Upload_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cvServiceClient) GetOne(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CV, error) {
	out := new(CV)
	err := c.cc.Invoke(ctx, CvService_GetOne_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cvServiceClient) GetAll(ctx context.Context, in *Paggination, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, CvService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CvServiceServer is the server API for CvService service.
// All implementations must embed UnimplementedCvServiceServer
// for forward compatibility
type CvServiceServer interface {
	Upload(context.Context, *CV) (*Empty, error)
	GetOne(context.Context, *Id) (*CV, error)
	GetAll(context.Context, *Paggination) (*GetAllResponse, error)
	mustEmbedUnimplementedCvServiceServer()
}

// UnimplementedCvServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCvServiceServer struct {
}

func (UnimplementedCvServiceServer) Upload(context.Context, *CV) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedCvServiceServer) GetOne(context.Context, *Id) (*CV, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOne not implemented")
}
func (UnimplementedCvServiceServer) GetAll(context.Context, *Paggination) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedCvServiceServer) mustEmbedUnimplementedCvServiceServer() {}

// UnsafeCvServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CvServiceServer will
// result in compilation errors.
type UnsafeCvServiceServer interface {
	mustEmbedUnimplementedCvServiceServer()
}

func RegisterCvServiceServer(s grpc.ServiceRegistrar, srv CvServiceServer) {
	s.RegisterService(&CvService_ServiceDesc, srv)
}

func _CvService_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CV)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CvServiceServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CvService_Upload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CvServiceServer).Upload(ctx, req.(*CV))
	}
	return interceptor(ctx, in, info, handler)
}

func _CvService_GetOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CvServiceServer).GetOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CvService_GetOne_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CvServiceServer).GetOne(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CvService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Paggination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CvServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CvService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CvServiceServer).GetAll(ctx, req.(*Paggination))
	}
	return interceptor(ctx, in, info, handler)
}

// CvService_ServiceDesc is the grpc.ServiceDesc for CvService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CvService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cv.CvService",
	HandlerType: (*CvServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upload",
			Handler:    _CvService_Upload_Handler,
		},
		{
			MethodName: "GetOne",
			Handler:    _CvService_GetOne_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _CvService_GetAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cv.proto",
}