// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.1
// source: proto/benchmark/benchmark.proto

package benchmark

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{0}
}

type HelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ComputeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Iterations    int64                  `protobuf:"varint,1,opt,name=iterations,proto3" json:"iterations,omitempty"` // Number of iterations for computation
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComputeRequest) Reset() {
	*x = ComputeRequest{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComputeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRequest) ProtoMessage() {}

func (x *ComputeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRequest.ProtoReflect.Descriptor instead.
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeRequest) GetIterations() int64 {
	if x != nil {
		return x.Iterations
	}
	return 0
}

type ComputeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ComputeResponse) Reset() {
	*x = ComputeResponse{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ComputeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResponse) ProtoMessage() {}

func (x *ComputeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResponse.ProtoReflect.Descriptor instead.
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{3}
}

func (x *ComputeResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type MemoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SizeMb        int32                  `protobuf:"varint,1,opt,name=size_mb,json=sizeMb,proto3" json:"size_mb,omitempty"` // Memory size in MB to allocate
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemoryRequest) Reset() {
	*x = MemoryRequest{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryRequest) ProtoMessage() {}

func (x *MemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryRequest.ProtoReflect.Descriptor instead.
func (*MemoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{4}
}

func (x *MemoryRequest) GetSizeMb() int32 {
	if x != nil {
		return x.SizeMb
	}
	return 0
}

type MemoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Hash          []byte                 `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	AccessTimeNs  int64                  `protobuf:"varint,2,opt,name=access_time_ns,json=accessTimeNs,proto3" json:"access_time_ns,omitempty"`  // Time taken for random access
	PagesAccessed int32                  `protobuf:"varint,3,opt,name=pages_accessed,json=pagesAccessed,proto3" json:"pages_accessed,omitempty"` // Number of pages accessed
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemoryResponse) Reset() {
	*x = MemoryResponse{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoryResponse) ProtoMessage() {}

func (x *MemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoryResponse.ProtoReflect.Descriptor instead.
func (*MemoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{5}
}

func (x *MemoryResponse) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *MemoryResponse) GetAccessTimeNs() int64 {
	if x != nil {
		return x.AccessTimeNs
	}
	return 0
}

func (x *MemoryResponse) GetPagesAccessed() int32 {
	if x != nil {
		return x.PagesAccessed
	}
	return 0
}

type IORequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileSizeMb    int32                  `protobuf:"varint,1,opt,name=file_size_mb,json=fileSizeMb,proto3" json:"file_size_mb,omitempty"` // Size of each file in MB
	NumFiles      int32                  `protobuf:"varint,2,opt,name=num_files,json=numFiles,proto3" json:"num_files,omitempty"`         // Number of files to create
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IORequest) Reset() {
	*x = IORequest{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IORequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IORequest) ProtoMessage() {}

func (x *IORequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IORequest.ProtoReflect.Descriptor instead.
func (*IORequest) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{6}
}

func (x *IORequest) GetFileSizeMb() int32 {
	if x != nil {
		return x.FileSizeMb
	}
	return 0
}

func (x *IORequest) GetNumFiles() int32 {
	if x != nil {
		return x.NumFiles
	}
	return 0
}

type IOResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IOResponse) Reset() {
	*x = IOResponse{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IOResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOResponse) ProtoMessage() {}

func (x *IOResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOResponse.ProtoReflect.Descriptor instead.
func (*IOResponse) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{7}
}

func (x *IOResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MixedRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MixedRequest) Reset() {
	*x = MixedRequest{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MixedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MixedRequest) ProtoMessage() {}

func (x *MixedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MixedRequest.ProtoReflect.Descriptor instead.
func (*MixedRequest) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{8}
}

type MixedResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MixedResponse) Reset() {
	*x = MixedResponse{}
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MixedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MixedResponse) ProtoMessage() {}

func (x *MixedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_benchmark_benchmark_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MixedResponse.ProtoReflect.Descriptor instead.
func (*MixedResponse) Descriptor() ([]byte, []int) {
	return file_proto_benchmark_benchmark_proto_rawDescGZIP(), []int{9}
}

func (x *MixedResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_benchmark_benchmark_proto protoreflect.FileDescriptor

var file_proto_benchmark_benchmark_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x0e, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0d,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x30, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x74, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69,
	0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x28, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x22, 0x71,
	0x0a, 0x0e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61,
	0x67, 0x65, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x73, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65,
	0x64, 0x22, 0x4a, 0x0a, 0x09, 0x49, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x62,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x75, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x6e, 0x75, 0x6d, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x26, 0x0a,
	0x0a, 0x49, 0x4f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x69, 0x78, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x0d, 0x4d, 0x69, 0x78, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xd0, 0x02, 0x0a, 0x10, 0x42, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x17,
	0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d,
	0x61, 0x72, 0x6b, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x45, 0x0a, 0x0c, 0x43, 0x50, 0x55, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x76,
	0x65, 0x12, 0x19, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x62,
	0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x76, 0x65, 0x12, 0x18, 0x2e, 0x62, 0x65,
	0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72,
	0x6b, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x4f, 0x12, 0x14, 0x2e, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x49, 0x4f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x49, 0x4f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x4d, 0x69, 0x78, 0x65, 0x64,
	0x12, 0x17, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x4d, 0x69, 0x78,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x65, 0x6e, 0x63,
	0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x4d, 0x69, 0x78, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x30, 0x78, 0x77, 0x6f, 0x6e, 0x6a, 0x2f, 0x74, 0x64, 0x78, 0x2d, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x65, 0x6e,
	0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_benchmark_benchmark_proto_rawDescOnce sync.Once
	file_proto_benchmark_benchmark_proto_rawDescData = file_proto_benchmark_benchmark_proto_rawDesc
)

func file_proto_benchmark_benchmark_proto_rawDescGZIP() []byte {
	file_proto_benchmark_benchmark_proto_rawDescOnce.Do(func() {
		file_proto_benchmark_benchmark_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_benchmark_benchmark_proto_rawDescData)
	})
	return file_proto_benchmark_benchmark_proto_rawDescData
}

var file_proto_benchmark_benchmark_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_benchmark_benchmark_proto_goTypes = []any{
	(*HelloRequest)(nil),    // 0: benchmark.HelloRequest
	(*HelloResponse)(nil),   // 1: benchmark.HelloResponse
	(*ComputeRequest)(nil),  // 2: benchmark.ComputeRequest
	(*ComputeResponse)(nil), // 3: benchmark.ComputeResponse
	(*MemoryRequest)(nil),   // 4: benchmark.MemoryRequest
	(*MemoryResponse)(nil),  // 5: benchmark.MemoryResponse
	(*IORequest)(nil),       // 6: benchmark.IORequest
	(*IOResponse)(nil),      // 7: benchmark.IOResponse
	(*MixedRequest)(nil),    // 8: benchmark.MixedRequest
	(*MixedResponse)(nil),   // 9: benchmark.MixedResponse
}
var file_proto_benchmark_benchmark_proto_depIdxs = []int32{
	0, // 0: benchmark.BenchmarkService.Hello:input_type -> benchmark.HelloRequest
	2, // 1: benchmark.BenchmarkService.CPUIntensive:input_type -> benchmark.ComputeRequest
	4, // 2: benchmark.BenchmarkService.MemoryIntensive:input_type -> benchmark.MemoryRequest
	6, // 3: benchmark.BenchmarkService.DiskIO:input_type -> benchmark.IORequest
	8, // 4: benchmark.BenchmarkService.Mixed:input_type -> benchmark.MixedRequest
	1, // 5: benchmark.BenchmarkService.Hello:output_type -> benchmark.HelloResponse
	3, // 6: benchmark.BenchmarkService.CPUIntensive:output_type -> benchmark.ComputeResponse
	5, // 7: benchmark.BenchmarkService.MemoryIntensive:output_type -> benchmark.MemoryResponse
	7, // 8: benchmark.BenchmarkService.DiskIO:output_type -> benchmark.IOResponse
	9, // 9: benchmark.BenchmarkService.Mixed:output_type -> benchmark.MixedResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_benchmark_benchmark_proto_init() }
func file_proto_benchmark_benchmark_proto_init() {
	if File_proto_benchmark_benchmark_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_benchmark_benchmark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_benchmark_benchmark_proto_goTypes,
		DependencyIndexes: file_proto_benchmark_benchmark_proto_depIdxs,
		MessageInfos:      file_proto_benchmark_benchmark_proto_msgTypes,
	}.Build()
	File_proto_benchmark_benchmark_proto = out.File
	file_proto_benchmark_benchmark_proto_rawDesc = nil
	file_proto_benchmark_benchmark_proto_goTypes = nil
	file_proto_benchmark_benchmark_proto_depIdxs = nil
}
