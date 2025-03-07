// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: coordinator/v1/coordinator.proto

// Namespace for the coordinator service
// The fn0rd.coordinator.v1 package contains definitions for the distributed computation coordination system

package coordinator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Capability enum defines the different scanning or operational capabilities
// that components in the system can support.
type Capability int32

const (
	// Default value used when capability is unknown or unspecified.
	Capability_CAPABILITY_UNKNOWN_UNSPECIFIED Capability = 0
	// Basic built-in capabilities provided by default.
	Capability_CAPABILITY_BUILTIN Capability = 1
	// Basic network mapping capability using nmap with missing scripts.
	Capability_CAPABILITY_NMAP Capability = 2
	// Full network mapping capability using nmap with all scripts.
	Capability_CAPABILITY_NMAP_FULL Capability = 3
)

// Enum value maps for Capability.
var (
	Capability_name = map[int32]string{
		0: "CAPABILITY_UNKNOWN_UNSPECIFIED",
		1: "CAPABILITY_BUILTIN",
		2: "CAPABILITY_NMAP",
		3: "CAPABILITY_NMAP_FULL",
	}
	Capability_value = map[string]int32{
		"CAPABILITY_UNKNOWN_UNSPECIFIED": 0,
		"CAPABILITY_BUILTIN":             1,
		"CAPABILITY_NMAP":                2,
		"CAPABILITY_NMAP_FULL":           3,
	}
)

func (x Capability) Enum() *Capability {
	p := new(Capability)
	*p = x
	return p
}

func (x Capability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Capability) Descriptor() protoreflect.EnumDescriptor {
	return file_coordinator_v1_coordinator_proto_enumTypes[0].Descriptor()
}

func (Capability) Type() protoreflect.EnumType {
	return &file_coordinator_v1_coordinator_proto_enumTypes[0]
}

func (x Capability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Capability.Descriptor instead.
func (Capability) EnumDescriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{0}
}

// StreamRequest encapsulates all possible client-to-coordinator messages
type StreamRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique random value to prevent replay attacks
	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Cryptographic signature proving the sender's identity
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// The actual request payload, which can be one of multiple types
	//
	// Types that are valid to be assigned to Request:
	//
	//	*StreamRequest_Register
	//	*StreamRequest_Submit
	Request       isStreamRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *StreamRequest) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *StreamRequest) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *StreamRequest) GetRequest() isStreamRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *StreamRequest) GetRegister() *RegisterRequest {
	if x != nil {
		if x, ok := x.Request.(*StreamRequest_Register); ok {
			return x.Register
		}
	}
	return nil
}

func (x *StreamRequest) GetSubmit() *SubmitRequest {
	if x != nil {
		if x, ok := x.Request.(*StreamRequest_Submit); ok {
			return x.Submit
		}
	}
	return nil
}

type isStreamRequest_Request interface {
	isStreamRequest_Request()
}

type StreamRequest_Register struct {
	// Client registration information
	Register *RegisterRequest `protobuf:"bytes,3,opt,name=register,proto3,oneof"`
}

type StreamRequest_Submit struct {
	// Computation result submission
	Submit *SubmitRequest `protobuf:"bytes,4,opt,name=submit,proto3,oneof"`
}

func (*StreamRequest_Register) isStreamRequest_Request() {}

func (*StreamRequest_Submit) isStreamRequest_Request() {}

// RegisterRequest is sent by clients to join the computation network
type RegisterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The client's public key for message verification and identity
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Number of parallel computation workers the client can provide
	Workers uint32 `protobuf:"varint,2,opt,name=workers,proto3" json:"workers,omitempty"`
	// The client's capabilities
	Capabilities  []Capability `protobuf:"varint,3,rep,packed,name=capabilities,proto3,enum=fn0rd.coordinator.v1.Capability" json:"capabilities,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterRequest) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *RegisterRequest) GetWorkers() uint32 {
	if x != nil {
		return x.Workers
	}
	return 0
}

func (x *RegisterRequest) GetCapabilities() []Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

// SubmitRequest is sent by clients when they complete a computational task
type SubmitRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The result of the completed computation
	Result        []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *SubmitRequest) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

// StreamResponse encapsulates all possible coordinator-to-client messages
type StreamResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique random value to prevent replay attacks
	Nonce []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The actual response payload, which can be one of multiple types
	//
	// Types that are valid to be assigned to Response:
	//
	//	*StreamResponse_Target
	Response      isStreamResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{3}
}

func (x *StreamResponse) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *StreamResponse) GetResponse() isStreamResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *StreamResponse) GetTarget() *TargetResponse {
	if x != nil {
		if x, ok := x.Response.(*StreamResponse_Target); ok {
			return x.Target
		}
	}
	return nil
}

type isStreamResponse_Response interface {
	isStreamResponse_Response()
}

type StreamResponse_Target struct {
	// Computational target assignment
	Target *TargetResponse `protobuf:"bytes,2,opt,name=target,proto3,oneof"`
}

func (*StreamResponse_Target) isStreamResponse_Response() {}

// TargetResponse contains a computational task assigned to a client
type TargetResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The computational targets to process
	Target []byte `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Timestamp by which the computation must be completed
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TargetResponse) Reset() {
	*x = TargetResponse{}
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TargetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetResponse) ProtoMessage() {}

func (x *TargetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coordinator_v1_coordinator_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetResponse.ProtoReflect.Descriptor instead.
func (*TargetResponse) Descriptor() ([]byte, []int) {
	return file_coordinator_v1_coordinator_proto_rawDescGZIP(), []int{4}
}

func (x *TargetResponse) GetTarget() []byte {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *TargetResponse) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

var File_coordinator_v1_coordinator_proto protoreflect.FileDescriptor

var file_coordinator_v1_coordinator_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0d, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x43, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x90,
	0x01, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x44, 0x0a, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x27, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x72, 0x0a, 0x0e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x60,
	0x0a, 0x0e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65,
	0x2a, 0x77, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x22,
	0x0a, 0x1e, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x42, 0x55, 0x49, 0x4c, 0x54, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x41,
	0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4d, 0x41, 0x50, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x4d,
	0x41, 0x50, 0x5f, 0x46, 0x55, 0x4c, 0x4c, 0x10, 0x03, 0x32, 0x6f, 0x0a, 0x12, 0x43, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x59, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x66, 0x6e, 0x30, 0x72,
	0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2e, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6e, 0x30, 0x72, 0x64, 0x2d, 0x69,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69,
	0x6e, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_coordinator_v1_coordinator_proto_rawDescOnce sync.Once
	file_coordinator_v1_coordinator_proto_rawDescData []byte
)

func file_coordinator_v1_coordinator_proto_rawDescGZIP() []byte {
	file_coordinator_v1_coordinator_proto_rawDescOnce.Do(func() {
		file_coordinator_v1_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_coordinator_v1_coordinator_proto_rawDesc), len(file_coordinator_v1_coordinator_proto_rawDesc)))
	})
	return file_coordinator_v1_coordinator_proto_rawDescData
}

var file_coordinator_v1_coordinator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_coordinator_v1_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_coordinator_v1_coordinator_proto_goTypes = []any{
	(Capability)(0),               // 0: fn0rd.coordinator.v1.Capability
	(*StreamRequest)(nil),         // 1: fn0rd.coordinator.v1.StreamRequest
	(*RegisterRequest)(nil),       // 2: fn0rd.coordinator.v1.RegisterRequest
	(*SubmitRequest)(nil),         // 3: fn0rd.coordinator.v1.SubmitRequest
	(*StreamResponse)(nil),        // 4: fn0rd.coordinator.v1.StreamResponse
	(*TargetResponse)(nil),        // 5: fn0rd.coordinator.v1.TargetResponse
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_coordinator_v1_coordinator_proto_depIdxs = []int32{
	2, // 0: fn0rd.coordinator.v1.StreamRequest.register:type_name -> fn0rd.coordinator.v1.RegisterRequest
	3, // 1: fn0rd.coordinator.v1.StreamRequest.submit:type_name -> fn0rd.coordinator.v1.SubmitRequest
	0, // 2: fn0rd.coordinator.v1.RegisterRequest.capabilities:type_name -> fn0rd.coordinator.v1.Capability
	5, // 3: fn0rd.coordinator.v1.StreamResponse.target:type_name -> fn0rd.coordinator.v1.TargetResponse
	6, // 4: fn0rd.coordinator.v1.TargetResponse.deadline:type_name -> google.protobuf.Timestamp
	1, // 5: fn0rd.coordinator.v1.CoordinatorService.Stream:input_type -> fn0rd.coordinator.v1.StreamRequest
	4, // 6: fn0rd.coordinator.v1.CoordinatorService.Stream:output_type -> fn0rd.coordinator.v1.StreamResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_coordinator_v1_coordinator_proto_init() }
func file_coordinator_v1_coordinator_proto_init() {
	if File_coordinator_v1_coordinator_proto != nil {
		return
	}
	file_coordinator_v1_coordinator_proto_msgTypes[0].OneofWrappers = []any{
		(*StreamRequest_Register)(nil),
		(*StreamRequest_Submit)(nil),
	}
	file_coordinator_v1_coordinator_proto_msgTypes[3].OneofWrappers = []any{
		(*StreamResponse_Target)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_coordinator_v1_coordinator_proto_rawDesc), len(file_coordinator_v1_coordinator_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coordinator_v1_coordinator_proto_goTypes,
		DependencyIndexes: file_coordinator_v1_coordinator_proto_depIdxs,
		EnumInfos:         file_coordinator_v1_coordinator_proto_enumTypes,
		MessageInfos:      file_coordinator_v1_coordinator_proto_msgTypes,
	}.Build()
	File_coordinator_v1_coordinator_proto = out.File
	file_coordinator_v1_coordinator_proto_goTypes = nil
	file_coordinator_v1_coordinator_proto_depIdxs = nil
}
