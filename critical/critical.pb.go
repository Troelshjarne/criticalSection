// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: critical/critical.proto

package criticalpackage

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

// node sends request for entering the critical section.
//the node attach its lamport time and node id.
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId int64 `protobuf:"varint,1,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_critical_critical_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_critical_critical_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_critical_critical_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetNodeId() int64 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

//
//a node sends a reply - accepting the request
type Reply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access bool `protobuf:"varint,1,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *Reply) Reset() {
	*x = Reply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_critical_critical_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reply) ProtoMessage() {}

func (x *Reply) ProtoReflect() protoreflect.Message {
	mi := &file_critical_critical_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reply.ProtoReflect.Descriptor instead.
func (*Reply) Descriptor() ([]byte, []int) {
	return file_critical_critical_proto_rawDescGZIP(), []int{1}
}

func (x *Reply) GetAccess() bool {
	if x != nil {
		return x.Access
	}
	return false
}

//
//acknowledge for sending a request
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_critical_critical_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_critical_critical_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_critical_critical_proto_rawDescGZIP(), []int{2}
}

func (x *Ack) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_critical_critical_proto protoreflect.FileDescriptor

var file_critical_critical_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x1f, 0x0a,
	0x05, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x1d,
	0x0a, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x97, 0x01,
	0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x43, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x16,
	0x2e, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x1a, 0x18, 0x2e, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e,
	0x61, 0x63, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x42, 0x41, 0x5a, 0x3f, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72,
	0x6f, 0x65, 0x6c, 0x73, 0x68, 0x6a, 0x61, 0x72, 0x6e, 0x65, 0x2f, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x72, 0x69, 0x74, 0x69,
	0x63, 0x61, 0x6c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_critical_critical_proto_rawDescOnce sync.Once
	file_critical_critical_proto_rawDescData = file_critical_critical_proto_rawDesc
)

func file_critical_critical_proto_rawDescGZIP() []byte {
	file_critical_critical_proto_rawDescOnce.Do(func() {
		file_critical_critical_proto_rawDescData = protoimpl.X.CompressGZIP(file_critical_critical_proto_rawDescData)
	})
	return file_critical_critical_proto_rawDescData
}

var file_critical_critical_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_critical_critical_proto_goTypes = []interface{}{
	(*Request)(nil), // 0: criticalpackage.request
	(*Reply)(nil),   // 1: criticalpackage.reply
	(*Ack)(nil),     // 2: criticalpackage.ack
}
var file_critical_critical_proto_depIdxs = []int32{
	1, // 0: criticalpackage.Communication.serverReply:input_type -> criticalpackage.reply
	0, // 1: criticalpackage.Communication.sendRequest:input_type -> criticalpackage.request
	0, // 2: criticalpackage.Communication.serverReply:output_type -> criticalpackage.request
	2, // 3: criticalpackage.Communication.sendRequest:output_type -> criticalpackage.ack
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_critical_critical_proto_init() }
func file_critical_critical_proto_init() {
	if File_critical_critical_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_critical_critical_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_critical_critical_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_critical_critical_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_critical_critical_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_critical_critical_proto_goTypes,
		DependencyIndexes: file_critical_critical_proto_depIdxs,
		MessageInfos:      file_critical_critical_proto_msgTypes,
	}.Build()
	File_critical_critical_proto = out.File
	file_critical_critical_proto_rawDesc = nil
	file_critical_critical_proto_goTypes = nil
	file_critical_critical_proto_depIdxs = nil
}
