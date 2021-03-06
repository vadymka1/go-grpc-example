// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/blog.proto

/*
Package blogpb is a generated protocol buffer package.

It is generated from these files:
	proto/blog.proto

It has these top-level messages:
	Blog
	CreateBlogReq
	CreateBlogRes
	UpdateBlogReq
	UpdateBlogRes
	ReadBlogReq
	ReadBlogRes
	DeleteBlogReq
	DeleteBlogRes
	ListBlogsReq
	ListBlogsRes
*/
package blogpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Blog struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId" json:"author_id,omitempty"`
	Title    string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Content  string `protobuf:"bytes,4,opt,name=content" json:"content,omitempty"`
}

func (m *Blog) Reset()                    { *m = Blog{} }
func (m *Blog) String() string            { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()               {}
func (*Blog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogReq struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *CreateBlogReq) Reset()                    { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string            { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()               {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *CreateBlogRes) Reset()                    { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string            { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()               {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogReq struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *UpdateBlogReq) Reset()                    { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()               {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *UpdateBlogRes) Reset()                    { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string            { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()               {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogReq struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadBlogReq) Reset()                    { *m = ReadBlogReq{} }
func (m *ReadBlogReq) String() string            { return proto.CompactTextString(m) }
func (*ReadBlogReq) ProtoMessage()               {}
func (*ReadBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *ReadBlogRes) Reset()                    { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string            { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()               {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogReq struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteBlogReq) Reset()                    { *m = DeleteBlogReq{} }
func (m *DeleteBlogReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlogReq) ProtoMessage()               {}
func (*DeleteBlogReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBlogRes struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *DeleteBlogRes) Reset()                    { *m = DeleteBlogRes{} }
func (m *DeleteBlogRes) String() string            { return proto.CompactTextString(m) }
func (*DeleteBlogRes) ProtoMessage()               {}
func (*DeleteBlogRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DeleteBlogRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListBlogsReq struct {
}

func (m *ListBlogsReq) Reset()                    { *m = ListBlogsReq{} }
func (m *ListBlogsReq) String() string            { return proto.CompactTextString(m) }
func (*ListBlogsReq) ProtoMessage()               {}
func (*ListBlogsReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ListBlogsRes struct {
	Blog *Blog `protobuf:"bytes,1,opt,name=blog" json:"blog,omitempty"`
}

func (m *ListBlogsRes) Reset()                    { *m = ListBlogsRes{} }
func (m *ListBlogsRes) String() string            { return proto.CompactTextString(m) }
func (*ListBlogsRes) ProtoMessage()               {}
func (*ListBlogsRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ListBlogsRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*ReadBlogReq)(nil), "blog.ReadBlogReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*DeleteBlogReq)(nil), "blog.DeleteBlogReq")
	proto.RegisterType((*DeleteBlogRes)(nil), "blog.DeleteBlogRes")
	proto.RegisterType((*ListBlogsReq)(nil), "blog.ListBlogsReq")
	proto.RegisterType((*ListBlogsRes)(nil), "blog.ListBlogsRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BlogService service

type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error)
	ListBlogs(ctx context.Context, in *ListBlogsReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error) {
	out := new(DeleteBlogRes)
	err := grpc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlogs(ctx context.Context, in *ListBlogsReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BlogService_serviceDesc.Streams[0], c.cc, "/blog.BlogService/ListBlogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogsClient interface {
	Recv() (*ListBlogsRes, error)
	grpc.ClientStream
}

type blogServiceListBlogsClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogsClient) Recv() (*ListBlogsRes, error) {
	m := new(ListBlogsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for BlogService service

type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error)
	ListBlogs(*ListBlogsReq, BlogService_ListBlogsServer) error
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlogs(m, &blogServiceListBlogsServer{stream})
}

type BlogService_ListBlogsServer interface {
	Send(*ListBlogsRes) error
	grpc.ServerStream
}

type blogServiceListBlogsServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogsServer) Send(m *ListBlogsRes) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlogs",
			Handler:       _BlogService_ListBlogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/blog.proto",
}

func init() { proto.RegisterFile("proto/blog.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xa5, 0xb1, 0xd6, 0x74, 0x6a, 0x8b, 0x8e, 0x1e, 0x42, 0xc4, 0x0f, 0x72, 0xd2, 0x83, 0x6d,
	0xa9, 0xe8, 0x0f, 0xa8, 0x5e, 0x04, 0x4f, 0x11, 0x2f, 0x5e, 0xa4, 0xcd, 0x0e, 0x75, 0x21, 0x74,
	0xd3, 0xec, 0xd6, 0x9f, 0xe3, 0x6f, 0x95, 0xdd, 0x65, 0xcd, 0x47, 0x03, 0xd1, 0xdb, 0xce, 0x9b,
	0xf7, 0xe6, 0x0d, 0x6f, 0x58, 0x38, 0xca, 0x72, 0xa1, 0xc4, 0x64, 0x99, 0x8a, 0xd5, 0xd8, 0x3c,
	0xb1, 0xab, 0xdf, 0x51, 0x02, 0xdd, 0x79, 0x2a, 0x56, 0x38, 0x02, 0x8f, 0xb3, 0xa0, 0x73, 0xd5,
	0xb9, 0xee, 0xc7, 0x1e, 0x67, 0x78, 0x06, 0xfd, 0xc5, 0x56, 0x7d, 0x8a, 0xfc, 0x83, 0xb3, 0xc0,
	0x33, 0xb0, 0x6f, 0x81, 0x67, 0x86, 0xa7, 0xb0, 0xaf, 0xb8, 0x4a, 0x29, 0xd8, 0x33, 0x0d, 0x5b,
	0x60, 0x00, 0x07, 0x89, 0x58, 0x2b, 0x5a, 0xab, 0xa0, 0x6b, 0x70, 0x57, 0x46, 0x13, 0x18, 0x3e,
	0xe6, 0xb4, 0x50, 0xa4, 0xad, 0x62, 0xda, 0xe0, 0x05, 0x18, 0x77, 0xe3, 0x37, 0x98, 0xc1, 0xd8,
	0xac, 0x65, 0x9a, 0x76, 0xab, 0x9a, 0x40, 0xfe, 0x45, 0xf0, 0x96, 0xb1, 0xff, 0x39, 0x94, 0x05,
	0xed, 0x0e, 0xe7, 0x30, 0x88, 0x69, 0xc1, 0xdc, 0xfc, 0x5a, 0x5e, 0xd1, 0x6d, 0xb9, 0xdd, 0x3e,
	0xed, 0x12, 0x86, 0x4f, 0x94, 0x52, 0xb1, 0x6f, 0x7d, 0xde, 0x4d, 0x95, 0x20, 0x75, 0xba, 0x72,
	0x9b, 0x24, 0x24, 0xa5, 0x61, 0xf9, 0xb1, 0x2b, 0xa3, 0x11, 0x1c, 0xbe, 0x70, 0xa9, 0x34, 0x51,
	0xc6, 0xb4, 0x89, 0xc6, 0x95, 0xba, 0x75, 0x97, 0xd9, 0xb7, 0x07, 0x03, 0x5d, 0xbe, 0x52, 0xfe,
	0xc5, 0x13, 0xc2, 0x07, 0x80, 0x22, 0x7c, 0x3c, 0xb1, 0xfc, 0xca, 0xfd, 0xc2, 0x06, 0x50, 0xe2,
	0x14, 0x7c, 0x17, 0x01, 0x1e, 0x5b, 0x42, 0x29, 0xb1, 0x70, 0x07, 0x92, 0xda, 0xa9, 0x38, 0x82,
	0x73, 0xaa, 0xdc, 0x31, 0x6c, 0x00, 0x8d, 0xae, 0x08, 0xc7, 0xe9, 0x2a, 0x79, 0x86, 0x0d, 0xa0,
	0xc4, 0x7b, 0xe8, 0xff, 0x26, 0x83, 0x68, 0x19, 0xe5, 0xe8, 0xc2, 0x5d, 0x4c, 0x4e, 0x3b, 0x73,
	0xff, 0xbd, 0xa7, 0xe1, 0x6c, 0xb9, 0xec, 0x99, 0xaf, 0x73, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xb3, 0xc5, 0xc4, 0xfe, 0x4e, 0x03, 0x00, 0x00,
}
