// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: someuser.proto

package someuser

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_someuser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	mi := &file_someuser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateUserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetAllUsersResponse struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SchemaSomeUsers []*CreateUserResponse  `protobuf:"bytes,1,rep,name=SchemaSomeUsers,proto3" json:"SchemaSomeUsers,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetAllUsersResponse) Reset() {
	*x = GetAllUsersResponse{}
	mi := &file_someuser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersResponse) ProtoMessage() {}

func (x *GetAllUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUsersResponse) GetSchemaSomeUsers() []*CreateUserResponse {
	if x != nil {
		return x.SchemaSomeUsers
	}
	return nil
}

type UserIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserIDRequest) Reset() {
	*x = UserIDRequest{}
	mi := &file_someuser_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDRequest) ProtoMessage() {}

func (x *UserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDRequest.ProtoReflect.Descriptor instead.
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{3}
}

func (x *UserIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageResponse) Reset() {
	*x = MessageResponse{}
	mi := &file_someuser_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageResponse) ProtoMessage() {}

func (x *MessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageResponse.ProtoReflect.Descriptor instead.
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{4}
}

func (x *MessageResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdateBody    *CreateUserRequest     `protobuf:"bytes,2,opt,name=updateBody,proto3" json:"updateBody,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_someuser_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_someuser_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_someuser_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserRequest) GetUpdateBody() *CreateUserRequest {
	if x != nil {
		return x.UpdateBody
	}
	return nil
}

var File_someuser_proto protoreflect.FileDescriptor

const file_someuser_proto_rawDesc = "" +
	"\n" +
	"\x0esomeuser.proto\x12\x05proto\x1a\x1bgoogle/protobuf/empty.proto\"a\n" +
	"\x11CreateUserRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x03 \x01(\tR\bpassword\"r\n" +
	"\x12CreateUserResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bpassword\x18\x04 \x01(\tR\bpassword\"Z\n" +
	"\x13GetAllUsersResponse\x12C\n" +
	"\x0fSchemaSomeUsers\x18\x01 \x03(\v2\x19.proto.CreateUserResponseR\x0fSchemaSomeUsers\"\x1f\n" +
	"\rUserIDRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"+\n" +
	"\x0fMessageResponse\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"]\n" +
	"\x11UpdateUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x128\n" +
	"\n" +
	"updateBody\x18\x02 \x01(\v2\x18.proto.CreateUserRequestR\n" +
	"updateBody2\xdd\x02\n" +
	"\bSomeUser\x12C\n" +
	"\n" +
	"CreateUser\x12\x18.proto.CreateUserRequest\x1a\x19.proto.CreateUserResponse\"\x00\x12C\n" +
	"\vGetAllUsers\x12\x16.google.protobuf.Empty\x1a\x1a.proto.GetAllUsersResponse\"\x00\x12@\n" +
	"\vGetUserByID\x12\x14.proto.UserIDRequest\x1a\x19.proto.CreateUserResponse\"\x00\x12@\n" +
	"\x0eDeleteUserByID\x12\x14.proto.UserIDRequest\x1a\x16.proto.MessageResponse\"\x00\x12C\n" +
	"\n" +
	"UpdateUser\x12\x18.proto.UpdateUserRequest\x1a\x19.proto.CreateUserResponse\"\x00B\x1dZ\x1bgithub.com/mellgit/someuserb\x06proto3"

var (
	file_someuser_proto_rawDescOnce sync.Once
	file_someuser_proto_rawDescData []byte
)

func file_someuser_proto_rawDescGZIP() []byte {
	file_someuser_proto_rawDescOnce.Do(func() {
		file_someuser_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_someuser_proto_rawDesc), len(file_someuser_proto_rawDesc)))
	})
	return file_someuser_proto_rawDescData
}

var file_someuser_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_someuser_proto_goTypes = []any{
	(*CreateUserRequest)(nil),   // 0: proto.CreateUserRequest
	(*CreateUserResponse)(nil),  // 1: proto.CreateUserResponse
	(*GetAllUsersResponse)(nil), // 2: proto.GetAllUsersResponse
	(*UserIDRequest)(nil),       // 3: proto.UserIDRequest
	(*MessageResponse)(nil),     // 4: proto.MessageResponse
	(*UpdateUserRequest)(nil),   // 5: proto.UpdateUserRequest
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_someuser_proto_depIdxs = []int32{
	1, // 0: proto.GetAllUsersResponse.SchemaSomeUsers:type_name -> proto.CreateUserResponse
	0, // 1: proto.UpdateUserRequest.updateBody:type_name -> proto.CreateUserRequest
	0, // 2: proto.SomeUser.CreateUser:input_type -> proto.CreateUserRequest
	6, // 3: proto.SomeUser.GetAllUsers:input_type -> google.protobuf.Empty
	3, // 4: proto.SomeUser.GetUserByID:input_type -> proto.UserIDRequest
	3, // 5: proto.SomeUser.DeleteUserByID:input_type -> proto.UserIDRequest
	5, // 6: proto.SomeUser.UpdateUser:input_type -> proto.UpdateUserRequest
	1, // 7: proto.SomeUser.CreateUser:output_type -> proto.CreateUserResponse
	2, // 8: proto.SomeUser.GetAllUsers:output_type -> proto.GetAllUsersResponse
	1, // 9: proto.SomeUser.GetUserByID:output_type -> proto.CreateUserResponse
	4, // 10: proto.SomeUser.DeleteUserByID:output_type -> proto.MessageResponse
	1, // 11: proto.SomeUser.UpdateUser:output_type -> proto.CreateUserResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_someuser_proto_init() }
func file_someuser_proto_init() {
	if File_someuser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_someuser_proto_rawDesc), len(file_someuser_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_someuser_proto_goTypes,
		DependencyIndexes: file_someuser_proto_depIdxs,
		MessageInfos:      file_someuser_proto_msgTypes,
	}.Build()
	File_someuser_proto = out.File
	file_someuser_proto_goTypes = nil
	file_someuser_proto_depIdxs = nil
}
