// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: autorization.proto

package proto

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

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName  string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName string `protobuf:"bytes,2,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Birthdate  string `protobuf:"bytes,3,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Biography  string `protobuf:"bytes,4,opt,name=biography,proto3" json:"biography,omitempty"`
	City       string `protobuf:"bytes,5,opt,name=city,proto3" json:"city,omitempty"`
	Password   string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_autorization_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserRequest) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *CreateUserRequest) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

func (x *CreateUserRequest) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *CreateUserRequest) GetCity() string {
	if x != nil {
		return x.City
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_autorization_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CreateUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreateUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_autorization_proto_rawDescGZIP(), []int{2}
}

func (x *UserLoginRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_autorization_proto_rawDescGZIP(), []int{3}
}

func (x *UserLoginResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserLoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserIdRequest) Reset() {
	*x = UserIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdRequest) ProtoMessage() {}

func (x *UserIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdRequest.ProtoReflect.Descriptor instead.
func (*UserIdRequest) Descriptor() ([]byte, []int) {
	return file_autorization_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UserGetByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  uint32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	User    *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserGetByIdResponse) Reset() {
	*x = UserGetByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserGetByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserGetByIdResponse) ProtoMessage() {}

func (x *UserGetByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserGetByIdResponse.ProtoReflect.Descriptor instead.
func (*UserGetByIdResponse) Descriptor() ([]byte, []int) {
	return file_autorization_proto_rawDescGZIP(), []int{5}
}

func (x *UserGetByIdResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UserGetByIdResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UserGetByIdResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName  string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	SecondName string `protobuf:"bytes,3,opt,name=second_name,json=secondName,proto3" json:"second_name,omitempty"`
	Birthdate  string `protobuf:"bytes,4,opt,name=birthdate,proto3" json:"birthdate,omitempty"`
	Biography  string `protobuf:"bytes,5,opt,name=biography,proto3" json:"biography,omitempty"`
	City       string `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autorization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_autorization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_autorization_proto_rawDescGZIP(), []int{6}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetSecondName() string {
	if x != nil {
		return x.SecondName
	}
	return ""
}

func (x *User) GetBirthdate() string {
	if x != nil {
		return x.Birthdate
	}
	return ""
}

func (x *User) GetBiography() string {
	if x != nil {
		return x.Biography
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

var File_autorization_proto protoreflect.FileDescriptor

var file_autorization_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x11,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5f, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5b, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x28, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x68,
	0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa6, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x69, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x32, 0xdc, 0x01, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_autorization_proto_rawDescOnce sync.Once
	file_autorization_proto_rawDescData = file_autorization_proto_rawDesc
)

func file_autorization_proto_rawDescGZIP() []byte {
	file_autorization_proto_rawDescOnce.Do(func() {
		file_autorization_proto_rawDescData = protoimpl.X.CompressGZIP(file_autorization_proto_rawDescData)
	})
	return file_autorization_proto_rawDescData
}

var file_autorization_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_autorization_proto_goTypes = []interface{}{
	(*CreateUserRequest)(nil),   // 0: proto.CreateUserRequest
	(*CreateUserResponse)(nil),  // 1: proto.CreateUserResponse
	(*UserLoginRequest)(nil),    // 2: proto.UserLoginRequest
	(*UserLoginResponse)(nil),   // 3: proto.UserLoginResponse
	(*UserIdRequest)(nil),       // 4: proto.UserIdRequest
	(*UserGetByIdResponse)(nil), // 5: proto.UserGetByIdResponse
	(*User)(nil),                // 6: proto.User
}
var file_autorization_proto_depIdxs = []int32{
	6, // 0: proto.UserGetByIdResponse.user:type_name -> proto.User
	0, // 1: proto.AuthorizationService.UserRegister:input_type -> proto.CreateUserRequest
	2, // 2: proto.AuthorizationService.UserLogin:input_type -> proto.UserLoginRequest
	4, // 3: proto.AuthorizationService.UserGetById:input_type -> proto.UserIdRequest
	1, // 4: proto.AuthorizationService.UserRegister:output_type -> proto.CreateUserResponse
	3, // 5: proto.AuthorizationService.UserLogin:output_type -> proto.UserLoginResponse
	5, // 6: proto.AuthorizationService.UserGetById:output_type -> proto.UserGetByIdResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autorization_proto_init() }
func file_autorization_proto_init() {
	if File_autorization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autorization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_autorization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_autorization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginRequest); i {
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
		file_autorization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginResponse); i {
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
		file_autorization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIdRequest); i {
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
		file_autorization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserGetByIdResponse); i {
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
		file_autorization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_autorization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autorization_proto_goTypes,
		DependencyIndexes: file_autorization_proto_depIdxs,
		MessageInfos:      file_autorization_proto_msgTypes,
	}.Build()
	File_autorization_proto = out.File
	file_autorization_proto_rawDesc = nil
	file_autorization_proto_goTypes = nil
	file_autorization_proto_depIdxs = nil
}
