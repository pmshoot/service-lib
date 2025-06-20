// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package hash

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// Sha3Client is the client API for Sha3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Sha3Client interface {
	Hash(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error)
	Hash_224(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error)
	Hash_256(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error)
	Hash_384(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error)
	Hash_512(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error)
}

type sha3Client struct {
	cc grpc.ClientConnInterface
}

func NewSha3Client(cc grpc.ClientConnInterface) Sha3Client {
	return &sha3Client{cc}
}

func (c *sha3Client) Hash(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error) {
	out := new(HashArrayResponse)
	err := c.cc.Invoke(ctx, "/hash.Sha3/Hash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sha3Client) Hash_224(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error) {
	out := new(HashArrayResponse)
	err := c.cc.Invoke(ctx, "/hash.Sha3/Hash_224", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sha3Client) Hash_256(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error) {
	out := new(HashArrayResponse)
	err := c.cc.Invoke(ctx, "/hash.Sha3/Hash_256", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sha3Client) Hash_384(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error) {
	out := new(HashArrayResponse)
	err := c.cc.Invoke(ctx, "/hash.Sha3/Hash_384", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sha3Client) Hash_512(ctx context.Context, in *StringArrayRequest, opts ...grpc.CallOption) (*HashArrayResponse, error) {
	out := new(HashArrayResponse)
	err := c.cc.Invoke(ctx, "/hash.Sha3/Hash_512", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Sha3Server is the server API for Sha3 service.
// All implementations must embed UnimplementedSha3Server
// for forward compatibility
type Sha3Server interface {
	Hash(context.Context, *StringArrayRequest) (*HashArrayResponse, error)
	Hash_224(context.Context, *StringArrayRequest) (*HashArrayResponse, error)
	Hash_256(context.Context, *StringArrayRequest) (*HashArrayResponse, error)
	Hash_384(context.Context, *StringArrayRequest) (*HashArrayResponse, error)
	Hash_512(context.Context, *StringArrayRequest) (*HashArrayResponse, error)
	mustEmbedUnimplementedSha3Server()
}

// UnimplementedSha3Server must be embedded to have forward compatible implementations.
type UnimplementedSha3Server struct {
}

func (UnimplementedSha3Server) Hash(context.Context, *StringArrayRequest) (*HashArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash not implemented")
}
func (UnimplementedSha3Server) Hash_224(context.Context, *StringArrayRequest) (*HashArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash_224 not implemented")
}
func (UnimplementedSha3Server) Hash_256(context.Context, *StringArrayRequest) (*HashArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash_256 not implemented")
}
func (UnimplementedSha3Server) Hash_384(context.Context, *StringArrayRequest) (*HashArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash_384 not implemented")
}
func (UnimplementedSha3Server) Hash_512(context.Context, *StringArrayRequest) (*HashArrayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hash_512 not implemented")
}
func (UnimplementedSha3Server) mustEmbedUnimplementedSha3Server() {}

// UnsafeSha3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Sha3Server will
// result in compilation errors.
type UnsafeSha3Server interface {
	mustEmbedUnimplementedSha3Server()
}

func RegisterSha3Server(s *grpc.Server, srv Sha3Server) {
	s.RegisterService(&_Sha3_serviceDesc, srv)
}

func _Sha3_Hash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Sha3Server).Hash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Sha3/Hash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Sha3Server).Hash(ctx, req.(*StringArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sha3_Hash_224_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Sha3Server).Hash_224(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Sha3/Hash_224",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Sha3Server).Hash_224(ctx, req.(*StringArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sha3_Hash_256_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Sha3Server).Hash_256(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Sha3/Hash_256",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Sha3Server).Hash_256(ctx, req.(*StringArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sha3_Hash_384_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Sha3Server).Hash_384(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Sha3/Hash_384",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Sha3Server).Hash_384(ctx, req.(*StringArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sha3_Hash_512_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringArrayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Sha3Server).Hash_512(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hash.Sha3/Hash_512",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Sha3Server).Hash_512(ctx, req.(*StringArrayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sha3_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hash.Sha3",
	HandlerType: (*Sha3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hash",
			Handler:    _Sha3_Hash_Handler,
		},
		{
			MethodName: "Hash_224",
			Handler:    _Sha3_Hash_224_Handler,
		},
		{
			MethodName: "Hash_256",
			Handler:    _Sha3_Hash_256_Handler,
		},
		{
			MethodName: "Hash_384",
			Handler:    _Sha3_Hash_384_Handler,
		},
		{
			MethodName: "Hash_512",
			Handler:    _Sha3_Hash_512_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hash.proto",
}
