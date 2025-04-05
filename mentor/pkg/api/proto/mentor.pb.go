// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v5.28.3
// source: proto/mentor.proto

package api

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

type RatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action      string  `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	MentorEmail string  `protobuf:"bytes,2,opt,name=mentor_email,json=mentorEmail,proto3" json:"mentor_email,omitempty"`
	Rating      float32 `protobuf:"fixed32,3,opt,name=rating,proto3" json:"rating,omitempty"`
}

func (x *RatingRequest) Reset() {
	*x = RatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mentor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingRequest) ProtoMessage() {}

func (x *RatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mentor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingRequest.ProtoReflect.Descriptor instead.
func (*RatingRequest) Descriptor() ([]byte, []int) {
	return file_proto_mentor_proto_rawDescGZIP(), []int{0}
}

func (x *RatingRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *RatingRequest) GetMentorEmail() string {
	if x != nil {
		return x.MentorEmail
	}
	return ""
}

func (x *RatingRequest) GetRating() float32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

type MentorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MentorEmail string `protobuf:"bytes,1,opt,name=mentor_email,json=mentorEmail,proto3" json:"mentor_email,omitempty"`
	Contact     string `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *MentorRequest) Reset() {
	*x = MentorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mentor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MentorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MentorRequest) ProtoMessage() {}

func (x *MentorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mentor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MentorRequest.ProtoReflect.Descriptor instead.
func (*MentorRequest) Descriptor() ([]byte, []int) {
	return file_proto_mentor_proto_rawDescGZIP(), []int{1}
}

func (x *MentorRequest) GetMentorEmail() string {
	if x != nil {
		return x.MentorEmail
	}
	return ""
}

func (x *MentorRequest) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MentorEmail string `protobuf:"bytes,1,opt,name=mentor_email,json=mentorEmail,proto3" json:"mentor_email,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mentor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mentor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_proto_mentor_proto_rawDescGZIP(), []int{2}
}

func (x *CheckRequest) GetMentorEmail() string {
	if x != nil {
		return x.MentorEmail
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Exists  bool   `protobuf:"varint,2,opt,name=exists,proto3" json:"exists,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mentor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mentor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_proto_mentor_proto_rawDescGZIP(), []int{3}
}

func (x *CheckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CheckResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

func (x *CheckResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mentor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mentor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_mentor_proto_rawDescGZIP(), []int{4}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_mentor_proto protoreflect.FileDescriptor

var file_proto_mentor_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x22, 0x62, 0x0a, 0x0d,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x5f,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x22, 0x4c, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x31,
	0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x5b, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3e,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc0,
	0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3d, 0x0a, 0x12, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4d, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x15, 0x2e, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x4d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x12, 0x15, 0x2e, 0x6d,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x65,
	0x6e, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x10, 0x5a, 0x0e, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mentor_proto_rawDescOnce sync.Once
	file_proto_mentor_proto_rawDescData = file_proto_mentor_proto_rawDesc
)

func file_proto_mentor_proto_rawDescGZIP() []byte {
	file_proto_mentor_proto_rawDescOnce.Do(func() {
		file_proto_mentor_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mentor_proto_rawDescData)
	})
	return file_proto_mentor_proto_rawDescData
}

var file_proto_mentor_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_mentor_proto_goTypes = []interface{}{
	(*RatingRequest)(nil), // 0: mentor.RatingRequest
	(*MentorRequest)(nil), // 1: mentor.MentorRequest
	(*CheckRequest)(nil),  // 2: mentor.CheckRequest
	(*CheckResponse)(nil), // 3: mentor.CheckResponse
	(*Response)(nil),      // 4: mentor.Response
}
var file_proto_mentor_proto_depIdxs = []int32{
	0, // 0: mentor.MentorService.MethodMentorRating:input_type -> mentor.RatingRequest
	1, // 1: mentor.MentorService.NewMentor:input_type -> mentor.MentorRequest
	2, // 2: mentor.MentorService.CheckMentor:input_type -> mentor.CheckRequest
	4, // 3: mentor.MentorService.MethodMentorRating:output_type -> mentor.Response
	4, // 4: mentor.MentorService.NewMentor:output_type -> mentor.Response
	3, // 5: mentor.MentorService.CheckMentor:output_type -> mentor.CheckResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_mentor_proto_init() }
func file_proto_mentor_proto_init() {
	if File_proto_mentor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mentor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RatingRequest); i {
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
		file_proto_mentor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MentorRequest); i {
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
		file_proto_mentor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_proto_mentor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_proto_mentor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_mentor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_mentor_proto_goTypes,
		DependencyIndexes: file_proto_mentor_proto_depIdxs,
		MessageInfos:      file_proto_mentor_proto_msgTypes,
	}.Build()
	File_proto_mentor_proto = out.File
	file_proto_mentor_proto_rawDesc = nil
	file_proto_mentor_proto_goTypes = nil
	file_proto_mentor_proto_depIdxs = nil
}
