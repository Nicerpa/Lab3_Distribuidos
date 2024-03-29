// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Lab3

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FulcrumServerClient is the client API for FulcrumServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FulcrumServerClient interface {
	// Broker
	GetClock(ctx context.Context, in *PlanetData, opts ...grpc.CallOption) (*PlanetClock, error)
	// Leia -> Broker (Leia a traves de Broker)
	GetNumber(ctx context.Context, in *CityData, opts ...grpc.CallOption) (*CityRes, error)
	// Informantes
	AddCity(ctx context.Context, in *NewCity, opts ...grpc.CallOption) (*CityStatus, error)
	DeleteCity(ctx context.Context, in *DelCity, opts ...grpc.CallOption) (*CityStatus, error)
	UpdateCity(ctx context.Context, in *UpCity, opts ...grpc.CallOption) (*CityStatus, error)
	RequestPlanetList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PlanetList, error)
	RequestLog(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*Log, error)
	UpdateFile(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*Status, error)
}

type fulcrumServerClient struct {
	cc grpc.ClientConnInterface
}

func NewFulcrumServerClient(cc grpc.ClientConnInterface) FulcrumServerClient {
	return &fulcrumServerClient{cc}
}

func (c *fulcrumServerClient) GetClock(ctx context.Context, in *PlanetData, opts ...grpc.CallOption) (*PlanetClock, error) {
	out := new(PlanetClock)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/GetClock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) GetNumber(ctx context.Context, in *CityData, opts ...grpc.CallOption) (*CityRes, error) {
	out := new(CityRes)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/GetNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) AddCity(ctx context.Context, in *NewCity, opts ...grpc.CallOption) (*CityStatus, error) {
	out := new(CityStatus)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/AddCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) DeleteCity(ctx context.Context, in *DelCity, opts ...grpc.CallOption) (*CityStatus, error) {
	out := new(CityStatus)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/DeleteCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) UpdateCity(ctx context.Context, in *UpCity, opts ...grpc.CallOption) (*CityStatus, error) {
	out := new(CityStatus)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/UpdateCity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) RequestPlanetList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PlanetList, error) {
	out := new(PlanetList)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/RequestPlanetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) RequestLog(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*Log, error) {
	out := new(Log)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/RequestLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fulcrumServerClient) UpdateFile(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/FulcrumProto.FulcrumServer/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FulcrumServerServer is the server API for FulcrumServer service.
// All implementations must embed UnimplementedFulcrumServerServer
// for forward compatibility
type FulcrumServerServer interface {
	// Broker
	GetClock(context.Context, *PlanetData) (*PlanetClock, error)
	// Leia -> Broker (Leia a traves de Broker)
	GetNumber(context.Context, *CityData) (*CityRes, error)
	// Informantes
	AddCity(context.Context, *NewCity) (*CityStatus, error)
	DeleteCity(context.Context, *DelCity) (*CityStatus, error)
	UpdateCity(context.Context, *UpCity) (*CityStatus, error)
	RequestPlanetList(context.Context, *empty.Empty) (*PlanetList, error)
	RequestLog(context.Context, *LogReq) (*Log, error)
	UpdateFile(context.Context, *NewData) (*Status, error)
	mustEmbedUnimplementedFulcrumServerServer()
}

// UnimplementedFulcrumServerServer must be embedded to have forward compatible implementations.
type UnimplementedFulcrumServerServer struct {
}

func (UnimplementedFulcrumServerServer) GetClock(context.Context, *PlanetData) (*PlanetClock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClock not implemented")
}
func (UnimplementedFulcrumServerServer) GetNumber(context.Context, *CityData) (*CityRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNumber not implemented")
}
func (UnimplementedFulcrumServerServer) AddCity(context.Context, *NewCity) (*CityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCity not implemented")
}
func (UnimplementedFulcrumServerServer) DeleteCity(context.Context, *DelCity) (*CityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCity not implemented")
}
func (UnimplementedFulcrumServerServer) UpdateCity(context.Context, *UpCity) (*CityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCity not implemented")
}
func (UnimplementedFulcrumServerServer) RequestPlanetList(context.Context, *empty.Empty) (*PlanetList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestPlanetList not implemented")
}
func (UnimplementedFulcrumServerServer) RequestLog(context.Context, *LogReq) (*Log, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLog not implemented")
}
func (UnimplementedFulcrumServerServer) UpdateFile(context.Context, *NewData) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFulcrumServerServer) mustEmbedUnimplementedFulcrumServerServer() {}

// UnsafeFulcrumServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FulcrumServerServer will
// result in compilation errors.
type UnsafeFulcrumServerServer interface {
	mustEmbedUnimplementedFulcrumServerServer()
}

func RegisterFulcrumServerServer(s grpc.ServiceRegistrar, srv FulcrumServerServer) {
	s.RegisterService(&FulcrumServer_ServiceDesc, srv)
}

func _FulcrumServer_GetClock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlanetData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).GetClock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/GetClock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).GetClock(ctx, req.(*PlanetData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_GetNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CityData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).GetNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/GetNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).GetNumber(ctx, req.(*CityData))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_AddCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).AddCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/AddCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).AddCity(ctx, req.(*NewCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_DeleteCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).DeleteCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/DeleteCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).DeleteCity(ctx, req.(*DelCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_UpdateCity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpCity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).UpdateCity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/UpdateCity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).UpdateCity(ctx, req.(*UpCity))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_RequestPlanetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).RequestPlanetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/RequestPlanetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).RequestPlanetList(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_RequestLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).RequestLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/RequestLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).RequestLog(ctx, req.(*LogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _FulcrumServer_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FulcrumServerServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FulcrumProto.FulcrumServer/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FulcrumServerServer).UpdateFile(ctx, req.(*NewData))
	}
	return interceptor(ctx, in, info, handler)
}

// FulcrumServer_ServiceDesc is the grpc.ServiceDesc for FulcrumServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FulcrumServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FulcrumProto.FulcrumServer",
	HandlerType: (*FulcrumServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClock",
			Handler:    _FulcrumServer_GetClock_Handler,
		},
		{
			MethodName: "GetNumber",
			Handler:    _FulcrumServer_GetNumber_Handler,
		},
		{
			MethodName: "AddCity",
			Handler:    _FulcrumServer_AddCity_Handler,
		},
		{
			MethodName: "DeleteCity",
			Handler:    _FulcrumServer_DeleteCity_Handler,
		},
		{
			MethodName: "UpdateCity",
			Handler:    _FulcrumServer_UpdateCity_Handler,
		},
		{
			MethodName: "RequestPlanetList",
			Handler:    _FulcrumServer_RequestPlanetList_Handler,
		},
		{
			MethodName: "RequestLog",
			Handler:    _FulcrumServer_RequestLog_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FulcrumServer_UpdateFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "FulcrumProto/FulcrumProto.proto",
}
