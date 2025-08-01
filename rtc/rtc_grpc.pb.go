// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.11.4
// source: rtc/rtc.proto

package rtc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RtcService_SignalMessageAssemble_FullMethodName           = "/openim.rtc.RtcService/SignalMessageAssemble"
	RtcService_SignalGetRoomByGroupID_FullMethodName          = "/openim.rtc.RtcService/SignalGetRoomByGroupID"
	RtcService_SignalGetTokenByRoomID_FullMethodName          = "/openim.rtc.RtcService/SignalGetTokenByRoomID"
	RtcService_SignalGetRooms_FullMethodName                  = "/openim.rtc.RtcService/SignalGetRooms"
	RtcService_GetSignalInvitationInfo_FullMethodName         = "/openim.rtc.RtcService/GetSignalInvitationInfo"
	RtcService_GetSignalInvitationInfoStartApp_FullMethodName = "/openim.rtc.RtcService/GetSignalInvitationInfoStartApp"
	RtcService_SignalSendCustomSignal_FullMethodName          = "/openim.rtc.RtcService/SignalSendCustomSignal"
	RtcService_GetSignalInvitationRecords_FullMethodName      = "/openim.rtc.RtcService/GetSignalInvitationRecords"
	RtcService_DeleteSignalRecords_FullMethodName             = "/openim.rtc.RtcService/DeleteSignalRecords"
)

// RtcServiceClient is the client API for RtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RtcServiceClient interface {
	SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error)
	SignalGetRoomByGroupID(ctx context.Context, in *SignalGetRoomByGroupIDReq, opts ...grpc.CallOption) (*SignalGetRoomByGroupIDResp, error)
	SignalGetTokenByRoomID(ctx context.Context, in *SignalGetTokenByRoomIDReq, opts ...grpc.CallOption) (*SignalGetTokenByRoomIDResp, error)
	SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error)
	GetSignalInvitationInfo(ctx context.Context, in *GetSignalInvitationInfoReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoResp, error)
	GetSignalInvitationInfoStartApp(ctx context.Context, in *GetSignalInvitationInfoStartAppReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoStartAppResp, error)
	// custom signal
	SignalSendCustomSignal(ctx context.Context, in *SignalSendCustomSignalReq, opts ...grpc.CallOption) (*SignalSendCustomSignalResp, error)
	// rtc cms
	GetSignalInvitationRecords(ctx context.Context, in *GetSignalInvitationRecordsReq, opts ...grpc.CallOption) (*GetSignalInvitationRecordsResp, error)
	DeleteSignalRecords(ctx context.Context, in *DeleteSignalRecordsReq, opts ...grpc.CallOption) (*DeleteSignalRecordsResp, error)
}

type rtcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRtcServiceClient(cc grpc.ClientConnInterface) RtcServiceClient {
	return &rtcServiceClient{cc}
}

func (c *rtcServiceClient) SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalMessageAssembleResp)
	err := c.cc.Invoke(ctx, RtcService_SignalMessageAssemble_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetRoomByGroupID(ctx context.Context, in *SignalGetRoomByGroupIDReq, opts ...grpc.CallOption) (*SignalGetRoomByGroupIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalGetRoomByGroupIDResp)
	err := c.cc.Invoke(ctx, RtcService_SignalGetRoomByGroupID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetTokenByRoomID(ctx context.Context, in *SignalGetTokenByRoomIDReq, opts ...grpc.CallOption) (*SignalGetTokenByRoomIDResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalGetTokenByRoomIDResp)
	err := c.cc.Invoke(ctx, RtcService_SignalGetTokenByRoomID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalGetRoomsResp)
	err := c.cc.Invoke(ctx, RtcService_SignalGetRooms_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationInfo(ctx context.Context, in *GetSignalInvitationInfoReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSignalInvitationInfoResp)
	err := c.cc.Invoke(ctx, RtcService_GetSignalInvitationInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationInfoStartApp(ctx context.Context, in *GetSignalInvitationInfoStartAppReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoStartAppResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSignalInvitationInfoStartAppResp)
	err := c.cc.Invoke(ctx, RtcService_GetSignalInvitationInfoStartApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalSendCustomSignal(ctx context.Context, in *SignalSendCustomSignalReq, opts ...grpc.CallOption) (*SignalSendCustomSignalResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignalSendCustomSignalResp)
	err := c.cc.Invoke(ctx, RtcService_SignalSendCustomSignal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationRecords(ctx context.Context, in *GetSignalInvitationRecordsReq, opts ...grpc.CallOption) (*GetSignalInvitationRecordsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSignalInvitationRecordsResp)
	err := c.cc.Invoke(ctx, RtcService_GetSignalInvitationRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) DeleteSignalRecords(ctx context.Context, in *DeleteSignalRecordsReq, opts ...grpc.CallOption) (*DeleteSignalRecordsResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteSignalRecordsResp)
	err := c.cc.Invoke(ctx, RtcService_DeleteSignalRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RtcServiceServer is the server API for RtcService service.
// All implementations must embed UnimplementedRtcServiceServer
// for forward compatibility.
type RtcServiceServer interface {
	SignalMessageAssemble(context.Context, *SignalMessageAssembleReq) (*SignalMessageAssembleResp, error)
	SignalGetRoomByGroupID(context.Context, *SignalGetRoomByGroupIDReq) (*SignalGetRoomByGroupIDResp, error)
	SignalGetTokenByRoomID(context.Context, *SignalGetTokenByRoomIDReq) (*SignalGetTokenByRoomIDResp, error)
	SignalGetRooms(context.Context, *SignalGetRoomsReq) (*SignalGetRoomsResp, error)
	GetSignalInvitationInfo(context.Context, *GetSignalInvitationInfoReq) (*GetSignalInvitationInfoResp, error)
	GetSignalInvitationInfoStartApp(context.Context, *GetSignalInvitationInfoStartAppReq) (*GetSignalInvitationInfoStartAppResp, error)
	// custom signal
	SignalSendCustomSignal(context.Context, *SignalSendCustomSignalReq) (*SignalSendCustomSignalResp, error)
	// rtc cms
	GetSignalInvitationRecords(context.Context, *GetSignalInvitationRecordsReq) (*GetSignalInvitationRecordsResp, error)
	DeleteSignalRecords(context.Context, *DeleteSignalRecordsReq) (*DeleteSignalRecordsResp, error)
	mustEmbedUnimplementedRtcServiceServer()
}

// UnimplementedRtcServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRtcServiceServer struct{}

func (UnimplementedRtcServiceServer) SignalMessageAssemble(context.Context, *SignalMessageAssembleReq) (*SignalMessageAssembleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalMessageAssemble not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetRoomByGroupID(context.Context, *SignalGetRoomByGroupIDReq) (*SignalGetRoomByGroupIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetRoomByGroupID not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetTokenByRoomID(context.Context, *SignalGetTokenByRoomIDReq) (*SignalGetTokenByRoomIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetTokenByRoomID not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetRooms(context.Context, *SignalGetRoomsReq) (*SignalGetRoomsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetRooms not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationInfo(context.Context, *GetSignalInvitationInfoReq) (*GetSignalInvitationInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationInfo not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationInfoStartApp(context.Context, *GetSignalInvitationInfoStartAppReq) (*GetSignalInvitationInfoStartAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationInfoStartApp not implemented")
}
func (UnimplementedRtcServiceServer) SignalSendCustomSignal(context.Context, *SignalSendCustomSignalReq) (*SignalSendCustomSignalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalSendCustomSignal not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationRecords(context.Context, *GetSignalInvitationRecordsReq) (*GetSignalInvitationRecordsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationRecords not implemented")
}
func (UnimplementedRtcServiceServer) DeleteSignalRecords(context.Context, *DeleteSignalRecordsReq) (*DeleteSignalRecordsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSignalRecords not implemented")
}
func (UnimplementedRtcServiceServer) mustEmbedUnimplementedRtcServiceServer() {}
func (UnimplementedRtcServiceServer) testEmbeddedByValue()                    {}

// UnsafeRtcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RtcServiceServer will
// result in compilation errors.
type UnsafeRtcServiceServer interface {
	mustEmbedUnimplementedRtcServiceServer()
}

func RegisterRtcServiceServer(s grpc.ServiceRegistrar, srv RtcServiceServer) {
	// If the following call pancis, it indicates UnimplementedRtcServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RtcService_ServiceDesc, srv)
}

func _RtcService_SignalMessageAssemble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalMessageAssembleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_SignalMessageAssemble_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, req.(*SignalMessageAssembleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetRoomByGroupID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetRoomByGroupIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetRoomByGroupID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_SignalGetRoomByGroupID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetRoomByGroupID(ctx, req.(*SignalGetRoomByGroupIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetTokenByRoomID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetTokenByRoomIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetTokenByRoomID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_SignalGetTokenByRoomID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetTokenByRoomID(ctx, req.(*SignalGetTokenByRoomIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetRoomsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_SignalGetRooms_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, req.(*SignalGetRoomsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_GetSignalInvitationInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationInfo(ctx, req.(*GetSignalInvitationInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationInfoStartApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationInfoStartAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationInfoStartApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_GetSignalInvitationInfoStartApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationInfoStartApp(ctx, req.(*GetSignalInvitationInfoStartAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalSendCustomSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalSendCustomSignalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalSendCustomSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_SignalSendCustomSignal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalSendCustomSignal(ctx, req.(*SignalSendCustomSignalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_GetSignalInvitationRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationRecords(ctx, req.(*GetSignalInvitationRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_DeleteSignalRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSignalRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).DeleteSignalRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RtcService_DeleteSignalRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).DeleteSignalRecords(ctx, req.(*DeleteSignalRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RtcService_ServiceDesc is the grpc.ServiceDesc for RtcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RtcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.rtc.RtcService",
	HandlerType: (*RtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignalMessageAssemble",
			Handler:    _RtcService_SignalMessageAssemble_Handler,
		},
		{
			MethodName: "SignalGetRoomByGroupID",
			Handler:    _RtcService_SignalGetRoomByGroupID_Handler,
		},
		{
			MethodName: "SignalGetTokenByRoomID",
			Handler:    _RtcService_SignalGetTokenByRoomID_Handler,
		},
		{
			MethodName: "SignalGetRooms",
			Handler:    _RtcService_SignalGetRooms_Handler,
		},
		{
			MethodName: "GetSignalInvitationInfo",
			Handler:    _RtcService_GetSignalInvitationInfo_Handler,
		},
		{
			MethodName: "GetSignalInvitationInfoStartApp",
			Handler:    _RtcService_GetSignalInvitationInfoStartApp_Handler,
		},
		{
			MethodName: "SignalSendCustomSignal",
			Handler:    _RtcService_SignalSendCustomSignal_Handler,
		},
		{
			MethodName: "GetSignalInvitationRecords",
			Handler:    _RtcService_GetSignalInvitationRecords_Handler,
		},
		{
			MethodName: "DeleteSignalRecords",
			Handler:    _RtcService_DeleteSignalRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rtc/rtc.proto",
}
