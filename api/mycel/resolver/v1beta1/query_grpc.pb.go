// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package resolverv1beta1

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of QueryWalletRecord items.
	WalletRecord(ctx context.Context, in *QueryWalletRecordRequest, opts ...grpc.CallOption) (*QueryWalletRecordResponse, error)
	// Queries a list of DnsRecord items.
	DnsRecord(ctx context.Context, in *QueryDnsRecordRequest, opts ...grpc.CallOption) (*QueryDnsRecordResponse, error)
	// Queries a list of AllRecord items.
	AllRecords(ctx context.Context, in *QueryAllRecordsRequest, opts ...grpc.CallOption) (*QueryAllRecordsResponse, error)
	// Queries a list of TextRecord items.
	TextRecord(ctx context.Context, in *QueryTextRecordRequest, opts ...grpc.CallOption) (*QueryTextRecordResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/mycel.resolver.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) WalletRecord(ctx context.Context, in *QueryWalletRecordRequest, opts ...grpc.CallOption) (*QueryWalletRecordResponse, error) {
	out := new(QueryWalletRecordResponse)
	err := c.cc.Invoke(ctx, "/mycel.resolver.v1beta1.Query/WalletRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DnsRecord(ctx context.Context, in *QueryDnsRecordRequest, opts ...grpc.CallOption) (*QueryDnsRecordResponse, error) {
	out := new(QueryDnsRecordResponse)
	err := c.cc.Invoke(ctx, "/mycel.resolver.v1beta1.Query/DnsRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllRecords(ctx context.Context, in *QueryAllRecordsRequest, opts ...grpc.CallOption) (*QueryAllRecordsResponse, error) {
	out := new(QueryAllRecordsResponse)
	err := c.cc.Invoke(ctx, "/mycel.resolver.v1beta1.Query/AllRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TextRecord(ctx context.Context, in *QueryTextRecordRequest, opts ...grpc.CallOption) (*QueryTextRecordResponse, error) {
	out := new(QueryTextRecordResponse)
	err := c.cc.Invoke(ctx, "/mycel.resolver.v1beta1.Query/TextRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of QueryWalletRecord items.
	WalletRecord(context.Context, *QueryWalletRecordRequest) (*QueryWalletRecordResponse, error)
	// Queries a list of DnsRecord items.
	DnsRecord(context.Context, *QueryDnsRecordRequest) (*QueryDnsRecordResponse, error)
	// Queries a list of AllRecord items.
	AllRecords(context.Context, *QueryAllRecordsRequest) (*QueryAllRecordsResponse, error)
	// Queries a list of TextRecord items.
	TextRecord(context.Context, *QueryTextRecordRequest) (*QueryTextRecordResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) WalletRecord(context.Context, *QueryWalletRecordRequest) (*QueryWalletRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletRecord not implemented")
}
func (UnimplementedQueryServer) DnsRecord(context.Context, *QueryDnsRecordRequest) (*QueryDnsRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DnsRecord not implemented")
}
func (UnimplementedQueryServer) AllRecords(context.Context, *QueryAllRecordsRequest) (*QueryAllRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllRecords not implemented")
}
func (UnimplementedQueryServer) TextRecord(context.Context, *QueryTextRecordRequest) (*QueryTextRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextRecord not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.resolver.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_WalletRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryWalletRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).WalletRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.resolver.v1beta1.Query/WalletRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).WalletRecord(ctx, req.(*QueryWalletRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DnsRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDnsRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DnsRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.resolver.v1beta1.Query/DnsRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DnsRecord(ctx, req.(*QueryDnsRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.resolver.v1beta1.Query/AllRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllRecords(ctx, req.(*QueryAllRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TextRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTextRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TextRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mycel.resolver.v1beta1.Query/TextRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TextRecord(ctx, req.(*QueryTextRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mycel.resolver.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "WalletRecord",
			Handler:    _Query_WalletRecord_Handler,
		},
		{
			MethodName: "DnsRecord",
			Handler:    _Query_DnsRecord_Handler,
		},
		{
			MethodName: "AllRecords",
			Handler:    _Query_AllRecords_Handler,
		},
		{
			MethodName: "TextRecord",
			Handler:    _Query_TextRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mycel/resolver/v1beta1/query.proto",
}
