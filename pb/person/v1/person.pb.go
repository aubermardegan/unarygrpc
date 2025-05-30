// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/v1/person.proto

package personv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Person struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_proto_v1_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type AddPersonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age           int32                  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPersonRequest) Reset() {
	*x = AddPersonRequest{}
	mi := &file_proto_v1_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonRequest) ProtoMessage() {}

func (x *AddPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonRequest.ProtoReflect.Descriptor instead.
func (*AddPersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{1}
}

func (x *AddPersonRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddPersonRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type AddPersonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Person        *Person                `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddPersonResponse) Reset() {
	*x = AddPersonResponse{}
	mi := &file_proto_v1_person_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPersonResponse) ProtoMessage() {}

func (x *AddPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPersonResponse.ProtoReflect.Descriptor instead.
func (*AddPersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{2}
}

func (x *AddPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type GetPersonRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	mi := &file_proto_v1_person_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type GetPersonResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Person        *Person                `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	mi := &file_proto_v1_person_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{4}
}

func (x *GetPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type ListPeopleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPeopleRequest) Reset() {
	*x = ListPeopleRequest{}
	mi := &file_proto_v1_person_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeopleRequest) ProtoMessage() {}

func (x *ListPeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeopleRequest.ProtoReflect.Descriptor instead.
func (*ListPeopleRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{5}
}

type ListPeopleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	People        []*Person              `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPeopleResponse) Reset() {
	*x = ListPeopleResponse{}
	mi := &file_proto_v1_person_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPeopleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPeopleResponse) ProtoMessage() {}

func (x *ListPeopleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_person_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPeopleResponse.ProtoReflect.Descriptor instead.
func (*ListPeopleResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_person_proto_rawDescGZIP(), []int{6}
}

func (x *ListPeopleResponse) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

var File_proto_v1_person_proto protoreflect.FileDescriptor

const file_proto_v1_person_proto_rawDesc = "" +
	"\n" +
	"\x15proto/v1/person.proto\x12\tperson.v1\"B\n" +
	"\x06Person\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x03 \x01(\x05R\x03age\"8\n" +
	"\x10AddPersonRequest\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03age\x18\x03 \x01(\x05R\x03age\">\n" +
	"\x11AddPersonResponse\x12)\n" +
	"\x06person\x18\x01 \x01(\v2\x11.person.v1.PersonR\x06person\"&\n" +
	"\x10GetPersonRequest\x12\x12\n" +
	"\x04uuid\x18\x01 \x01(\tR\x04uuid\">\n" +
	"\x11GetPersonResponse\x12)\n" +
	"\x06person\x18\x01 \x01(\v2\x11.person.v1.PersonR\x06person\"\x13\n" +
	"\x11ListPeopleRequest\"?\n" +
	"\x12ListPeopleResponse\x12)\n" +
	"\x06people\x18\x01 \x03(\v2\x11.person.v1.PersonR\x06people2\xea\x01\n" +
	"\rPersonService\x12F\n" +
	"\tAddPerson\x12\x1b.person.v1.AddPersonRequest\x1a\x1c.person.v1.AddPersonResponse\x12F\n" +
	"\tGetPerson\x12\x1b.person.v1.GetPersonRequest\x1a\x1c.person.v1.GetPersonResponse\x12I\n" +
	"\n" +
	"ListPeople\x12\x1c.person.v1.ListPeopleRequest\x1a\x1d.person.v1.ListPeopleResponseB\x19Z\x17./pb/person/v1;personv1b\x06proto3"

var (
	file_proto_v1_person_proto_rawDescOnce sync.Once
	file_proto_v1_person_proto_rawDescData []byte
)

func file_proto_v1_person_proto_rawDescGZIP() []byte {
	file_proto_v1_person_proto_rawDescOnce.Do(func() {
		file_proto_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_v1_person_proto_rawDesc), len(file_proto_v1_person_proto_rawDesc)))
	})
	return file_proto_v1_person_proto_rawDescData
}

var file_proto_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_v1_person_proto_goTypes = []any{
	(*Person)(nil),             // 0: person.v1.Person
	(*AddPersonRequest)(nil),   // 1: person.v1.AddPersonRequest
	(*AddPersonResponse)(nil),  // 2: person.v1.AddPersonResponse
	(*GetPersonRequest)(nil),   // 3: person.v1.GetPersonRequest
	(*GetPersonResponse)(nil),  // 4: person.v1.GetPersonResponse
	(*ListPeopleRequest)(nil),  // 5: person.v1.ListPeopleRequest
	(*ListPeopleResponse)(nil), // 6: person.v1.ListPeopleResponse
}
var file_proto_v1_person_proto_depIdxs = []int32{
	0, // 0: person.v1.AddPersonResponse.person:type_name -> person.v1.Person
	0, // 1: person.v1.GetPersonResponse.person:type_name -> person.v1.Person
	0, // 2: person.v1.ListPeopleResponse.people:type_name -> person.v1.Person
	1, // 3: person.v1.PersonService.AddPerson:input_type -> person.v1.AddPersonRequest
	3, // 4: person.v1.PersonService.GetPerson:input_type -> person.v1.GetPersonRequest
	5, // 5: person.v1.PersonService.ListPeople:input_type -> person.v1.ListPeopleRequest
	2, // 6: person.v1.PersonService.AddPerson:output_type -> person.v1.AddPersonResponse
	4, // 7: person.v1.PersonService.GetPerson:output_type -> person.v1.GetPersonResponse
	6, // 8: person.v1.PersonService.ListPeople:output_type -> person.v1.ListPeopleResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1_person_proto_init() }
func file_proto_v1_person_proto_init() {
	if File_proto_v1_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_v1_person_proto_rawDesc), len(file_proto_v1_person_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_person_proto_goTypes,
		DependencyIndexes: file_proto_v1_person_proto_depIdxs,
		MessageInfos:      file_proto_v1_person_proto_msgTypes,
	}.Build()
	File_proto_v1_person_proto = out.File
	file_proto_v1_person_proto_goTypes = nil
	file_proto_v1_person_proto_depIdxs = nil
}
