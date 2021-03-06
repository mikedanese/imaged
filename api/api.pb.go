// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ObjectMeta
	Void
	Status
	Condition
	Image
	ImageList
	Blob
	BlobList
	Bundle
	BundleList
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

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

type ObjectMeta struct {
	Name              string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CreationTimestamp *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	DeletionTimestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=deletion_timestamp,json=deletionTimestamp" json:"deletion_timestamp,omitempty"`
	Labels            map[string]string          `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Annotations       map[string]string          `protobuf:"bytes,5,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ObjectMeta) Reset()                    { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string            { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()               {}
func (*ObjectMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ObjectMeta) GetCreationTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetDeletionTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.DeletionTimestamp
	}
	return nil
}

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Status struct {
	Conditions []*Condition `protobuf:"bytes,1,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type Condition struct {
	Type           string                     `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Status         bool                       `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	Reason         string                     `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
	LastTransition *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=last_transition,json=lastTransition" json:"last_transition,omitempty"`
}

func (m *Condition) Reset()                    { *m = Condition{} }
func (m *Condition) String() string            { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()               {}
func (*Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Condition) GetLastTransition() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastTransition
	}
	return nil
}

type Image struct {
	Meta       *ObjectMeta  `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Url        string       `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Conditions []*Condition `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Image) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Image) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type ImageList struct {
	Images []*Image `protobuf:"bytes,1,rep,name=images" json:"images,omitempty"`
}

func (m *ImageList) Reset()                    { *m = ImageList{} }
func (m *ImageList) String() string            { return proto.CompactTextString(m) }
func (*ImageList) ProtoMessage()               {}
func (*ImageList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ImageList) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Blob struct {
	Meta       *ObjectMeta  `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Url        string       `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Conditions []*Condition `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *Blob) Reset()                    { *m = Blob{} }
func (m *Blob) String() string            { return proto.CompactTextString(m) }
func (*Blob) ProtoMessage()               {}
func (*Blob) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Blob) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Blob) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type BlobList struct {
	Blobs []*Blob `protobuf:"bytes,1,rep,name=blobs" json:"blobs,omitempty"`
}

func (m *BlobList) Reset()                    { *m = BlobList{} }
func (m *BlobList) String() string            { return proto.CompactTextString(m) }
func (*BlobList) ProtoMessage()               {}
func (*BlobList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *BlobList) GetBlobs() []*Blob {
	if m != nil {
		return m.Blobs
	}
	return nil
}

type Bundle struct {
	Meta       *ObjectMeta  `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Url        string       `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Conditions []*Condition `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *Bundle) Reset()                    { *m = Bundle{} }
func (m *Bundle) String() string            { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()               {}
func (*Bundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Bundle) GetMeta() *ObjectMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *Bundle) GetConditions() []*Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type BundleList struct {
	Bundles []*Bundle `protobuf:"bytes,1,rep,name=bundles" json:"bundles,omitempty"`
}

func (m *BundleList) Reset()                    { *m = BundleList{} }
func (m *BundleList) String() string            { return proto.CompactTextString(m) }
func (*BundleList) ProtoMessage()               {}
func (*BundleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *BundleList) GetBundles() []*Bundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

func init() {
	proto.RegisterType((*ObjectMeta)(nil), "ObjectMeta")
	proto.RegisterType((*Void)(nil), "Void")
	proto.RegisterType((*Status)(nil), "Status")
	proto.RegisterType((*Condition)(nil), "Condition")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*ImageList)(nil), "ImageList")
	proto.RegisterType((*Blob)(nil), "Blob")
	proto.RegisterType((*BlobList)(nil), "BlobList")
	proto.RegisterType((*Bundle)(nil), "Bundle")
	proto.RegisterType((*BundleList)(nil), "BundleList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Images service

type ImagesClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ImageList, error)
	Put(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error)
	Delete(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error)
}

type imagesClient struct {
	cc *grpc.ClientConn
}

func NewImagesClient(cc *grpc.ClientConn) ImagesClient {
	return &imagesClient{cc}
}

func (c *imagesClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*ImageList, error) {
	out := new(ImageList)
	err := grpc.Invoke(ctx, "/Images/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Put(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/Images/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imagesClient) Delete(ctx context.Context, in *Image, opts ...grpc.CallOption) (*Image, error) {
	out := new(Image)
	err := grpc.Invoke(ctx, "/Images/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Images service

type ImagesServer interface {
	List(context.Context, *Void) (*ImageList, error)
	Put(context.Context, *Image) (*Image, error)
	Delete(context.Context, *Image) (*Image, error)
}

func RegisterImagesServer(s *grpc.Server, srv ImagesServer) {
	s.RegisterService(&_Images_serviceDesc, srv)
}

func _Images_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Images/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Images/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Put(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _Images_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImagesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Images/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImagesServer).Delete(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

var _Images_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Images",
	HandlerType: (*ImagesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Images_List_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Images_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Images_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Blobs service

type BlobsClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*BlobList, error)
	Put(ctx context.Context, in *Blob, opts ...grpc.CallOption) (*Blob, error)
	Delete(ctx context.Context, in *Blob, opts ...grpc.CallOption) (*Blob, error)
}

type blobsClient struct {
	cc *grpc.ClientConn
}

func NewBlobsClient(cc *grpc.ClientConn) BlobsClient {
	return &blobsClient{cc}
}

func (c *blobsClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*BlobList, error) {
	out := new(BlobList)
	err := grpc.Invoke(ctx, "/Blobs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobsClient) Put(ctx context.Context, in *Blob, opts ...grpc.CallOption) (*Blob, error) {
	out := new(Blob)
	err := grpc.Invoke(ctx, "/Blobs/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobsClient) Delete(ctx context.Context, in *Blob, opts ...grpc.CallOption) (*Blob, error) {
	out := new(Blob)
	err := grpc.Invoke(ctx, "/Blobs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Blobs service

type BlobsServer interface {
	List(context.Context, *Void) (*BlobList, error)
	Put(context.Context, *Blob) (*Blob, error)
	Delete(context.Context, *Blob) (*Blob, error)
}

func RegisterBlobsServer(s *grpc.Server, srv BlobsServer) {
	s.RegisterService(&_Blobs_serviceDesc, srv)
}

func _Blobs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Blobs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobsServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blobs_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobsServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Blobs/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobsServer).Put(ctx, req.(*Blob))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Blobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobsServer).Delete(ctx, req.(*Blob))
	}
	return interceptor(ctx, in, info, handler)
}

var _Blobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Blobs",
	HandlerType: (*BlobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Blobs_List_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Blobs_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Blobs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

// Client API for Bundles service

type BundlesClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*BundleList, error)
	Put(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error)
	Delete(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error)
}

type bundlesClient struct {
	cc *grpc.ClientConn
}

func NewBundlesClient(cc *grpc.ClientConn) BundlesClient {
	return &bundlesClient{cc}
}

func (c *bundlesClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*BundleList, error) {
	out := new(BundleList)
	err := grpc.Invoke(ctx, "/Bundles/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundlesClient) Put(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error) {
	out := new(Bundle)
	err := grpc.Invoke(ctx, "/Bundles/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bundlesClient) Delete(ctx context.Context, in *Bundle, opts ...grpc.CallOption) (*Bundle, error) {
	out := new(Bundle)
	err := grpc.Invoke(ctx, "/Bundles/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bundles service

type BundlesServer interface {
	List(context.Context, *Void) (*BundleList, error)
	Put(context.Context, *Bundle) (*Bundle, error)
	Delete(context.Context, *Bundle) (*Bundle, error)
}

func RegisterBundlesServer(s *grpc.Server, srv BundlesServer) {
	s.RegisterService(&_Bundles_serviceDesc, srv)
}

func _Bundles_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BundlesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bundles/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BundlesServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bundles_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BundlesServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bundles/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BundlesServer).Put(ctx, req.(*Bundle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bundles_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BundlesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bundles/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BundlesServer).Delete(ctx, req.(*Bundle))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bundles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Bundles",
	HandlerType: (*BundlesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Bundles_List_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Bundles_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Bundles_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0xf5, 0x47, 0xea, 0xb1, 0x04, 0x65, 0x05, 0x34, 0x98, 0x88, 0x16, 0x5f, 0x40, 0x20,
	0x6d, 0xa4, 0xc0, 0x01, 0x38, 0x54, 0xa2, 0x85, 0x43, 0xa5, 0x22, 0xd0, 0x12, 0x71, 0xe0, 0x52,
	0xd6, 0xc9, 0x36, 0x32, 0x38, 0x76, 0x94, 0x5d, 0x23, 0xe5, 0x7f, 0xf0, 0x0b, 0xf8, 0xa5, 0xec,
	0xce, 0x7a, 0x1d, 0x37, 0x20, 0x2a, 0x0e, 0x3d, 0x65, 0x66, 0xde, 0xbe, 0x37, 0x33, 0x6f, 0x1c,
	0x88, 0xf8, 0x32, 0xa7, 0xcb, 0x55, 0xa5, 0xaa, 0xe4, 0x60, 0x5e, 0x55, 0xf3, 0x42, 0x8c, 0x30,
	0xcb, 0xea, 0x8b, 0x91, 0xca, 0x17, 0x42, 0x2a, 0xbe, 0x58, 0xda, 0x07, 0xe9, 0x2f, 0x0f, 0xe0,
	0x43, 0xf6, 0x4d, 0x4c, 0xd5, 0x7b, 0xa1, 0x38, 0x21, 0xe0, 0x97, 0x7c, 0x21, 0x06, 0xbd, 0xc3,
	0xde, 0x93, 0x88, 0x61, 0x4c, 0x4e, 0x81, 0x4c, 0x57, 0x82, 0xab, 0xbc, 0x2a, 0xcf, 0x5b, 0xfa,
	0x60, 0x47, 0xbf, 0x88, 0xc7, 0x09, 0xb5, 0x0d, 0xa8, 0x6b, 0x40, 0x27, 0xee, 0x05, 0xbb, 0xed,
	0x58, 0x6d, 0xc9, 0x48, 0xcd, 0x44, 0x21, 0xb6, 0xa4, 0xbc, 0xab, 0xa5, 0x1c, 0x6b, 0x23, 0x35,
	0x82, 0xb0, 0xe0, 0x99, 0x28, 0xe4, 0xc0, 0x3f, 0xf4, 0x34, 0x7d, 0x9f, 0x6e, 0xd6, 0xa0, 0x67,
	0x88, 0xbc, 0x2b, 0xd5, 0x6a, 0xcd, 0x9a, 0x67, 0xe4, 0x08, 0x62, 0x5e, 0x96, 0x95, 0xc2, 0x91,
	0xe4, 0x20, 0x40, 0xd6, 0xb0, 0xcb, 0x7a, 0xb3, 0x81, 0x2d, 0xb5, 0x4b, 0x48, 0x5e, 0x41, 0xdc,
	0x91, 0x25, 0x7b, 0xe0, 0x7d, 0x17, 0xeb, 0xc6, 0x28, 0x13, 0x92, 0x3b, 0x10, 0xfc, 0xe0, 0x45,
	0x2d, 0xd0, 0x9a, 0x88, 0xd9, 0xe4, 0xf5, 0xce, 0xcb, 0x5e, 0x72, 0x04, 0x7b, 0xdb, 0xda, 0xff,
	0xc3, 0x4f, 0x43, 0xf0, 0x3f, 0x57, 0xf9, 0x2c, 0x7d, 0x01, 0xe1, 0x27, 0x2d, 0x52, 0x4b, 0xf2,
	0x14, 0x60, 0x5a, 0x95, 0xb3, 0xdc, 0xee, 0xd2, 0xc3, 0x5d, 0x80, 0x9e, 0xb8, 0x12, 0xeb, 0xa0,
	0xe9, 0xcf, 0x1e, 0x44, 0x2d, 0x62, 0x2e, 0xac, 0xd6, 0xcb, 0xf6, 0xc2, 0x26, 0x26, 0xf7, 0x20,
	0x94, 0xa8, 0x8b, 0xad, 0x77, 0x59, 0x93, 0x99, 0xba, 0x3e, 0xa1, 0xac, 0x4a, 0x3c, 0x51, 0xc4,
	0x9a, 0x8c, 0x9c, 0xc0, 0xad, 0x82, 0x4b, 0x75, 0xae, 0x56, 0xbc, 0x94, 0x28, 0xab, 0x8f, 0x70,
	0xd5, 0x0d, 0x6f, 0x1a, 0xca, 0xa4, 0x65, 0xa4, 0x17, 0x10, 0x9c, 0x2e, 0xf8, 0x5c, 0x90, 0x03,
	0xf0, 0x17, 0xda, 0x7e, 0x9c, 0x28, 0x1e, 0xc7, 0x9d, 0x8b, 0x30, 0x04, 0x8c, 0x55, 0xf5, 0xaa,
	0x68, 0x6c, 0x31, 0xe1, 0xd6, 0xfa, 0xde, 0x3f, 0xd7, 0x7f, 0x06, 0x11, 0xf6, 0x39, 0xcb, 0xa5,
	0x22, 0x0f, 0x21, 0xcc, 0x4d, 0xe2, 0x3c, 0x0b, 0x29, 0x62, 0xac, 0xa9, 0xa6, 0x02, 0xfc, 0xe3,
	0xa2, 0xca, 0xae, 0x7b, 0xa6, 0xc7, 0xb0, 0x6b, 0xda, 0xe0, 0x48, 0x0f, 0x20, 0xc8, 0x74, 0xec,
	0x26, 0x0a, 0xa8, 0x41, 0x98, 0xad, 0xa5, 0x73, 0x08, 0x8f, 0xeb, 0x72, 0x56, 0x5c, 0xbb, 0x4b,
	0x23, 0x00, 0xdb, 0x08, 0x67, 0x7a, 0x04, 0xfd, 0x0c, 0x33, 0x37, 0x55, 0x9f, 0x5a, 0x94, 0xb9,
	0xfa, 0xf8, 0x0b, 0x84, 0x68, 0x9d, 0xd4, 0x0b, 0xf8, 0x48, 0x0a, 0xa8, 0xf9, 0x48, 0x13, 0xa0,
	0xad, 0xdd, 0xe9, 0x0d, 0xb2, 0x0f, 0xde, 0xc7, 0x5a, 0x91, 0xc6, 0xe7, 0xa4, 0xf9, 0xd5, 0x40,
	0x02, 0xe1, 0x5b, 0xf3, 0xa7, 0x16, 0x7f, 0x62, 0xe3, 0x09, 0x04, 0xc6, 0x04, 0xa9, 0x1f, 0x5d,
	0x92, 0x8e, 0xa8, 0x73, 0x4d, 0x0b, 0xdc, 0xb5, 0xca, 0xd6, 0xaf, 0xc4, 0xfe, 0xe8, 0xf2, 0xa0,
	0xd5, 0xdd, 0x42, 0xc6, 0x5f, 0xa1, 0x6f, 0x97, 0x90, 0x64, 0x78, 0x59, 0x37, 0xa6, 0x9b, 0xdd,
	0xb5, 0xc4, 0x7d, 0xab, 0xec, 0x76, 0x4e, 0x5c, 0xa0, 0xa1, 0x61, 0xab, 0xfe, 0x17, 0x34, 0x0b,
	0xf1, 0xb3, 0x7f, 0xfe, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x79, 0xdc, 0xa6, 0x7b, 0x81, 0x05, 0x00,
	0x00,
}
