// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: api/notes/service.proto

package notes

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
	NotesService_SaveNotesStream_FullMethodName = "/github.com.moguchev.microservices.grpc_stream.NotesService/SaveNotesStream"
	NotesService_ListNotesStream_FullMethodName = "/github.com.moguchev.microservices.grpc_stream.NotesService/ListNotesStream"
)

// NotesServiceClient is the client API for NotesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotesServiceClient interface {
	// SaveNotesStream - save notes
	SaveNotesStream(ctx context.Context, opts ...grpc.CallOption) (NotesService_SaveNotesStreamClient, error)
	// ListNotesStream - list all notes in stream
	ListNotesStream(ctx context.Context, in *ListNotesStreamRequest, opts ...grpc.CallOption) (NotesService_ListNotesStreamClient, error)
}

type notesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotesServiceClient(cc grpc.ClientConnInterface) NotesServiceClient {
	return &notesServiceClient{cc}
}

func (c *notesServiceClient) SaveNotesStream(ctx context.Context, opts ...grpc.CallOption) (NotesService_SaveNotesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotesService_ServiceDesc.Streams[0], NotesService_SaveNotesStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &notesServiceSaveNotesStreamClient{stream}
	return x, nil
}

type NotesService_SaveNotesStreamClient interface {
	Send(*NoteInfo) error
	Recv() (*Note, error)
	grpc.ClientStream
}

type notesServiceSaveNotesStreamClient struct {
	grpc.ClientStream
}

func (x *notesServiceSaveNotesStreamClient) Send(m *NoteInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *notesServiceSaveNotesStreamClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *notesServiceClient) ListNotesStream(ctx context.Context, in *ListNotesStreamRequest, opts ...grpc.CallOption) (NotesService_ListNotesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &NotesService_ServiceDesc.Streams[1], NotesService_ListNotesStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &notesServiceListNotesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NotesService_ListNotesStreamClient interface {
	Recv() (*Note, error)
	grpc.ClientStream
}

type notesServiceListNotesStreamClient struct {
	grpc.ClientStream
}

func (x *notesServiceListNotesStreamClient) Recv() (*Note, error) {
	m := new(Note)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotesServiceServer is the server API for NotesService service.
// All implementations must embed UnimplementedNotesServiceServer
// for forward compatibility
type NotesServiceServer interface {
	// SaveNotesStream - save notes
	SaveNotesStream(NotesService_SaveNotesStreamServer) error
	// ListNotesStream - list all notes in stream
	ListNotesStream(*ListNotesStreamRequest, NotesService_ListNotesStreamServer) error
	mustEmbedUnimplementedNotesServiceServer()
}

// UnimplementedNotesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotesServiceServer struct {
}

func (UnimplementedNotesServiceServer) SaveNotesStream(NotesService_SaveNotesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveNotesStream not implemented")
}
func (UnimplementedNotesServiceServer) ListNotesStream(*ListNotesStreamRequest, NotesService_ListNotesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ListNotesStream not implemented")
}
func (UnimplementedNotesServiceServer) mustEmbedUnimplementedNotesServiceServer() {}

// UnsafeNotesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotesServiceServer will
// result in compilation errors.
type UnsafeNotesServiceServer interface {
	mustEmbedUnimplementedNotesServiceServer()
}

func RegisterNotesServiceServer(s grpc.ServiceRegistrar, srv NotesServiceServer) {
	s.RegisterService(&NotesService_ServiceDesc, srv)
}

func _NotesService_SaveNotesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NotesServiceServer).SaveNotesStream(&notesServiceSaveNotesStreamServer{stream})
}

type NotesService_SaveNotesStreamServer interface {
	Send(*Note) error
	Recv() (*NoteInfo, error)
	grpc.ServerStream
}

type notesServiceSaveNotesStreamServer struct {
	grpc.ServerStream
}

func (x *notesServiceSaveNotesStreamServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

func (x *notesServiceSaveNotesStreamServer) Recv() (*NoteInfo, error) {
	m := new(NoteInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NotesService_ListNotesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNotesStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotesServiceServer).ListNotesStream(m, &notesServiceListNotesStreamServer{stream})
}

type NotesService_ListNotesStreamServer interface {
	Send(*Note) error
	grpc.ServerStream
}

type notesServiceListNotesStreamServer struct {
	grpc.ServerStream
}

func (x *notesServiceListNotesStreamServer) Send(m *Note) error {
	return x.ServerStream.SendMsg(m)
}

// NotesService_ServiceDesc is the grpc.ServiceDesc for NotesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "github.com.moguchev.microservices.grpc_stream.NotesService",
	HandlerType: (*NotesServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveNotesStream",
			Handler:       _NotesService_SaveNotesStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ListNotesStream",
			Handler:       _NotesService_ListNotesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/notes/service.proto",
}
