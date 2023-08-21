// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: helloworld/v1/greeter.proto

package v1

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
	Greeter_SayHello_FullMethodName           = "/helloworld.v1.Greeter/SayHello"
	Greeter_UserRegister_FullMethodName       = "/helloworld.v1.Greeter/UserRegister"
	Greeter_UserList_FullMethodName           = "/helloworld.v1.Greeter/UserList"
	Greeter_ClusterList_FullMethodName        = "/helloworld.v1.Greeter/ClusterList"
	Greeter_NamespaceList_FullMethodName      = "/helloworld.v1.Greeter/NamespaceList"
	Greeter_GetInstance_FullMethodName        = "/helloworld.v1.Greeter/GetInstance"
	Greeter_CreateInstance_FullMethodName     = "/helloworld.v1.Greeter/CreateInstance"
	Greeter_ListInstances_FullMethodName      = "/helloworld.v1.Greeter/ListInstances"
	Greeter_DeleteInstanceById_FullMethodName = "/helloworld.v1.Greeter/DeleteInstanceById"
	Greeter_UpdateInstance_FullMethodName     = "/helloworld.v1.Greeter/UpdateInstance"
	Greeter_SyncFromTencent_FullMethodName    = "/helloworld.v1.Greeter/SyncFromTencent"
	Greeter_DeploymentList_FullMethodName     = "/helloworld.v1.Greeter/DeploymentList"
	Greeter_ClusterRegister_FullMethodName    = "/helloworld.v1.Greeter/ClusterRegister"
	Greeter_ClusterDelete_FullMethodName      = "/helloworld.v1.Greeter/ClusterDelete"
	Greeter_ClusterGet_FullMethodName         = "/helloworld.v1.Greeter/ClusterGet"
	Greeter_ClusterUpdate_FullMethodName      = "/helloworld.v1.Greeter/ClusterUpdate"
)

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Register a user
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
	UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error)
	ClusterList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClusterListResponse, error)
	NamespaceList(ctx context.Context, in *NamespaceReq, opts ...grpc.CallOption) (*NamespaceResp, error)
	GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceReply, error)
	CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceReply, error)
	ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesReply, error)
	DeleteInstanceById(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceReply, error)
	UpdateInstance(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*UpdateInstanceReply, error)
	SyncFromTencent(ctx context.Context, in *SyncFromTencentRequest, opts ...grpc.CallOption) (*SyncFromTencentReply, error)
	DeploymentList(ctx context.Context, in *DeploymentReq, opts ...grpc.CallOption) (*DeploymentResp, error)
	// Register a cluster
	ClusterRegister(ctx context.Context, in *ClusterKubeconfig, opts ...grpc.CallOption) (*ClusterBase, error)
	// Delete a cluster
	ClusterDelete(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*Empty, error)
	// Get a cluster
	ClusterGet(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*ClusterBase, error)
	// Update a cluster
	ClusterUpdate(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*ClusterBase, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, Greeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, Greeter_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UserList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*UserListResponse, error) {
	out := new(UserListResponse)
	err := c.cc.Invoke(ctx, Greeter_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ClusterList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := c.cc.Invoke(ctx, Greeter_ClusterList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) NamespaceList(ctx context.Context, in *NamespaceReq, opts ...grpc.CallOption) (*NamespaceResp, error) {
	out := new(NamespaceResp)
	err := c.cc.Invoke(ctx, Greeter_NamespaceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) GetInstance(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*GetInstanceReply, error) {
	out := new(GetInstanceReply)
	err := c.cc.Invoke(ctx, Greeter_GetInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) CreateInstance(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*CreateInstanceReply, error) {
	out := new(CreateInstanceReply)
	err := c.cc.Invoke(ctx, Greeter_CreateInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ListInstances(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesReply, error) {
	out := new(ListInstancesReply)
	err := c.cc.Invoke(ctx, Greeter_ListInstances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeleteInstanceById(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*DeleteInstanceReply, error) {
	out := new(DeleteInstanceReply)
	err := c.cc.Invoke(ctx, Greeter_DeleteInstanceById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) UpdateInstance(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*UpdateInstanceReply, error) {
	out := new(UpdateInstanceReply)
	err := c.cc.Invoke(ctx, Greeter_UpdateInstance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SyncFromTencent(ctx context.Context, in *SyncFromTencentRequest, opts ...grpc.CallOption) (*SyncFromTencentReply, error) {
	out := new(SyncFromTencentReply)
	err := c.cc.Invoke(ctx, Greeter_SyncFromTencent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) DeploymentList(ctx context.Context, in *DeploymentReq, opts ...grpc.CallOption) (*DeploymentResp, error) {
	out := new(DeploymentResp)
	err := c.cc.Invoke(ctx, Greeter_DeploymentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ClusterRegister(ctx context.Context, in *ClusterKubeconfig, opts ...grpc.CallOption) (*ClusterBase, error) {
	out := new(ClusterBase)
	err := c.cc.Invoke(ctx, Greeter_ClusterRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ClusterDelete(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Greeter_ClusterDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ClusterGet(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*ClusterBase, error) {
	out := new(ClusterBase)
	err := c.cc.Invoke(ctx, Greeter_ClusterGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) ClusterUpdate(ctx context.Context, in *ClusterBase, opts ...grpc.CallOption) (*ClusterBase, error) {
	out := new(ClusterBase)
	err := c.cc.Invoke(ctx, Greeter_ClusterUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	// Register a user
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error)
	UserList(context.Context, *Empty) (*UserListResponse, error)
	ClusterList(context.Context, *Empty) (*ClusterListResponse, error)
	NamespaceList(context.Context, *NamespaceReq) (*NamespaceResp, error)
	GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceReply, error)
	CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceReply, error)
	ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesReply, error)
	DeleteInstanceById(context.Context, *DeleteInstanceRequest) (*DeleteInstanceReply, error)
	UpdateInstance(context.Context, *UpdateInstanceRequest) (*UpdateInstanceReply, error)
	SyncFromTencent(context.Context, *SyncFromTencentRequest) (*SyncFromTencentReply, error)
	DeploymentList(context.Context, *DeploymentReq) (*DeploymentResp, error)
	// Register a cluster
	ClusterRegister(context.Context, *ClusterKubeconfig) (*ClusterBase, error)
	// Delete a cluster
	ClusterDelete(context.Context, *ClusterBase) (*Empty, error)
	// Get a cluster
	ClusterGet(context.Context, *ClusterBase) (*ClusterBase, error)
	// Update a cluster
	ClusterUpdate(context.Context, *ClusterBase) (*ClusterBase, error)
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedGreeterServer) UserList(context.Context, *Empty) (*UserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedGreeterServer) ClusterList(context.Context, *Empty) (*ClusterListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterList not implemented")
}
func (UnimplementedGreeterServer) NamespaceList(context.Context, *NamespaceReq) (*NamespaceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NamespaceList not implemented")
}
func (UnimplementedGreeterServer) GetInstance(context.Context, *GetInstanceRequest) (*GetInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstance not implemented")
}
func (UnimplementedGreeterServer) CreateInstance(context.Context, *CreateInstanceRequest) (*CreateInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateInstance not implemented")
}
func (UnimplementedGreeterServer) ListInstances(context.Context, *ListInstancesRequest) (*ListInstancesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedGreeterServer) DeleteInstanceById(context.Context, *DeleteInstanceRequest) (*DeleteInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstanceById not implemented")
}
func (UnimplementedGreeterServer) UpdateInstance(context.Context, *UpdateInstanceRequest) (*UpdateInstanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstance not implemented")
}
func (UnimplementedGreeterServer) SyncFromTencent(context.Context, *SyncFromTencentRequest) (*SyncFromTencentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncFromTencent not implemented")
}
func (UnimplementedGreeterServer) DeploymentList(context.Context, *DeploymentReq) (*DeploymentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeploymentList not implemented")
}
func (UnimplementedGreeterServer) ClusterRegister(context.Context, *ClusterKubeconfig) (*ClusterBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterRegister not implemented")
}
func (UnimplementedGreeterServer) ClusterDelete(context.Context, *ClusterBase) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterDelete not implemented")
}
func (UnimplementedGreeterServer) ClusterGet(context.Context, *ClusterBase) (*ClusterBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterGet not implemented")
}
func (UnimplementedGreeterServer) ClusterUpdate(context.Context, *ClusterBase) (*ClusterBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClusterUpdate not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UserList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ClusterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ClusterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ClusterList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ClusterList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_NamespaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamespaceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).NamespaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_NamespaceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).NamespaceList(ctx, req.(*NamespaceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_GetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).GetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_GetInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).GetInstance(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_CreateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).CreateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_CreateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).CreateInstance(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ListInstances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ListInstances(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeleteInstanceById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeleteInstanceById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_DeleteInstanceById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeleteInstanceById(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_UpdateInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).UpdateInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_UpdateInstance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).UpdateInstance(ctx, req.(*UpdateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SyncFromTencent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncFromTencentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SyncFromTencent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_SyncFromTencent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SyncFromTencent(ctx, req.(*SyncFromTencentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_DeploymentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeploymentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).DeploymentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_DeploymentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).DeploymentList(ctx, req.(*DeploymentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ClusterRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterKubeconfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ClusterRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ClusterRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ClusterRegister(ctx, req.(*ClusterKubeconfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ClusterDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ClusterDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ClusterDelete(ctx, req.(*ClusterBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ClusterGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ClusterGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ClusterGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ClusterGet(ctx, req.(*ClusterBase))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_ClusterUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterBase)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).ClusterUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Greeter_ClusterUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).ClusterUpdate(ctx, req.(*ClusterBase))
	}
	return interceptor(ctx, in, info, handler)
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.v1.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "UserRegister",
			Handler:    _Greeter_UserRegister_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Greeter_UserList_Handler,
		},
		{
			MethodName: "ClusterList",
			Handler:    _Greeter_ClusterList_Handler,
		},
		{
			MethodName: "NamespaceList",
			Handler:    _Greeter_NamespaceList_Handler,
		},
		{
			MethodName: "GetInstance",
			Handler:    _Greeter_GetInstance_Handler,
		},
		{
			MethodName: "CreateInstance",
			Handler:    _Greeter_CreateInstance_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _Greeter_ListInstances_Handler,
		},
		{
			MethodName: "DeleteInstanceById",
			Handler:    _Greeter_DeleteInstanceById_Handler,
		},
		{
			MethodName: "UpdateInstance",
			Handler:    _Greeter_UpdateInstance_Handler,
		},
		{
			MethodName: "SyncFromTencent",
			Handler:    _Greeter_SyncFromTencent_Handler,
		},
		{
			MethodName: "DeploymentList",
			Handler:    _Greeter_DeploymentList_Handler,
		},
		{
			MethodName: "ClusterRegister",
			Handler:    _Greeter_ClusterRegister_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Greeter_ClusterDelete_Handler,
		},
		{
			MethodName: "ClusterGet",
			Handler:    _Greeter_ClusterGet_Handler,
		},
		{
			MethodName: "ClusterUpdate",
			Handler:    _Greeter_ClusterUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld/v1/greeter.proto",
}
