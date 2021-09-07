// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ova_meeting_api

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

// MeetingsClient is the client API for Meetings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MeetingsClient interface {
	// AddMeetingRequestV1V1 create new Meeting
	CreateMeetingV1(ctx context.Context, in *AddMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error)
	// Создает несколько новых опросов
	MultiCreateMeetingV1(ctx context.Context, in *MultiCreateMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error)
	// AddMeetingRequestV1V1 create new Meeting
	UpdateMeetingV1(ctx context.Context, in *UpdateMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error)
	// DescribeMeetingV1 get Meeting Info by ID
	DescribeMeetingV1(ctx context.Context, in *MeetingIDRequestV1, opts ...grpc.CallOption) (*MeetingResponseV1, error)
	// ListMeetingsV1 get all Meetings
	ListMeetingsV1(ctx context.Context, in *ListMeetingsRequestV1, opts ...grpc.CallOption) (*ListMeetingsResponseV1, error)
	// RemoveMeetingV1 remove Meeting by ID
	RemoveMeetingV1(ctx context.Context, in *MeetingIDRequestV1, opts ...grpc.CallOption) (*empty.Empty, error)
}

type meetingsClient struct {
	cc grpc.ClientConnInterface
}

func NewMeetingsClient(cc grpc.ClientConnInterface) MeetingsClient {
	return &meetingsClient{cc}
}

func (c *meetingsClient) CreateMeetingV1(ctx context.Context, in *AddMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/CreateMeetingV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsClient) MultiCreateMeetingV1(ctx context.Context, in *MultiCreateMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/MultiCreateMeetingV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsClient) UpdateMeetingV1(ctx context.Context, in *UpdateMeetingRequestV1, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/UpdateMeetingV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsClient) DescribeMeetingV1(ctx context.Context, in *MeetingIDRequestV1, opts ...grpc.CallOption) (*MeetingResponseV1, error) {
	out := new(MeetingResponseV1)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/DescribeMeetingV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsClient) ListMeetingsV1(ctx context.Context, in *ListMeetingsRequestV1, opts ...grpc.CallOption) (*ListMeetingsResponseV1, error) {
	out := new(ListMeetingsResponseV1)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/ListMeetingsV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meetingsClient) RemoveMeetingV1(ctx context.Context, in *MeetingIDRequestV1, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/ova.meeting.api.Meetings/RemoveMeetingV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MeetingsServer is the server API for Meetings service.
// All implementations must embed UnimplementedMeetingsServer
// for forward compatibility
type MeetingsServer interface {
	// AddMeetingRequestV1V1 create new Meeting
	CreateMeetingV1(context.Context, *AddMeetingRequestV1) (*empty.Empty, error)
	// Создает несколько новых опросов
	MultiCreateMeetingV1(context.Context, *MultiCreateMeetingRequestV1) (*empty.Empty, error)
	// AddMeetingRequestV1V1 create new Meeting
	UpdateMeetingV1(context.Context, *UpdateMeetingRequestV1) (*empty.Empty, error)
	// DescribeMeetingV1 get Meeting Info by ID
	DescribeMeetingV1(context.Context, *MeetingIDRequestV1) (*MeetingResponseV1, error)
	// ListMeetingsV1 get all Meetings
	ListMeetingsV1(context.Context, *ListMeetingsRequestV1) (*ListMeetingsResponseV1, error)
	// RemoveMeetingV1 remove Meeting by ID
	RemoveMeetingV1(context.Context, *MeetingIDRequestV1) (*empty.Empty, error)
	mustEmbedUnimplementedMeetingsServer()
}

// UnimplementedMeetingsServer must be embedded to have forward compatible implementations.
type UnimplementedMeetingsServer struct {
}

func (UnimplementedMeetingsServer) CreateMeetingV1(context.Context, *AddMeetingRequestV1) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMeetingV1 not implemented")
}
func (UnimplementedMeetingsServer) MultiCreateMeetingV1(context.Context, *MultiCreateMeetingRequestV1) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateMeetingV1 not implemented")
}
func (UnimplementedMeetingsServer) UpdateMeetingV1(context.Context, *UpdateMeetingRequestV1) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMeetingV1 not implemented")
}
func (UnimplementedMeetingsServer) DescribeMeetingV1(context.Context, *MeetingIDRequestV1) (*MeetingResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeMeetingV1 not implemented")
}
func (UnimplementedMeetingsServer) ListMeetingsV1(context.Context, *ListMeetingsRequestV1) (*ListMeetingsResponseV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMeetingsV1 not implemented")
}
func (UnimplementedMeetingsServer) RemoveMeetingV1(context.Context, *MeetingIDRequestV1) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMeetingV1 not implemented")
}
func (UnimplementedMeetingsServer) mustEmbedUnimplementedMeetingsServer() {}

// UnsafeMeetingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeetingsServer will
// result in compilation errors.
type UnsafeMeetingsServer interface {
	mustEmbedUnimplementedMeetingsServer()
}

func RegisterMeetingsServer(s grpc.ServiceRegistrar, srv MeetingsServer) {
	s.RegisterService(&Meetings_ServiceDesc, srv)
}

func _Meetings_CreateMeetingV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMeetingRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).CreateMeetingV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/CreateMeetingV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).CreateMeetingV1(ctx, req.(*AddMeetingRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meetings_MultiCreateMeetingV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateMeetingRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).MultiCreateMeetingV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/MultiCreateMeetingV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).MultiCreateMeetingV1(ctx, req.(*MultiCreateMeetingRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meetings_UpdateMeetingV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMeetingRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).UpdateMeetingV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/UpdateMeetingV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).UpdateMeetingV1(ctx, req.(*UpdateMeetingRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meetings_DescribeMeetingV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeetingIDRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).DescribeMeetingV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/DescribeMeetingV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).DescribeMeetingV1(ctx, req.(*MeetingIDRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meetings_ListMeetingsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMeetingsRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).ListMeetingsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/ListMeetingsV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).ListMeetingsV1(ctx, req.(*ListMeetingsRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _Meetings_RemoveMeetingV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeetingIDRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeetingsServer).RemoveMeetingV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ova.meeting.api.Meetings/RemoveMeetingV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeetingsServer).RemoveMeetingV1(ctx, req.(*MeetingIDRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

// Meetings_ServiceDesc is the grpc.ServiceDesc for Meetings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Meetings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ova.meeting.api.Meetings",
	HandlerType: (*MeetingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMeetingV1",
			Handler:    _Meetings_CreateMeetingV1_Handler,
		},
		{
			MethodName: "MultiCreateMeetingV1",
			Handler:    _Meetings_MultiCreateMeetingV1_Handler,
		},
		{
			MethodName: "UpdateMeetingV1",
			Handler:    _Meetings_UpdateMeetingV1_Handler,
		},
		{
			MethodName: "DescribeMeetingV1",
			Handler:    _Meetings_DescribeMeetingV1_Handler,
		},
		{
			MethodName: "ListMeetingsV1",
			Handler:    _Meetings_ListMeetingsV1_Handler,
		},
		{
			MethodName: "RemoveMeetingV1",
			Handler:    _Meetings_RemoveMeetingV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ova-meeting-api/api.proto",
}
