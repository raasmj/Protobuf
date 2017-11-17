// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addressbook.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	addressbook.proto

It has these top-level messages:
	Message
	Person
	AddressBook
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_WORK   Person_PhoneType = 1
	Person_HOME   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "WORK",
	2: "HOME",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"WORK":   1,
	"HOME":   2,
}

func (x Person_PhoneType) String() string {
	return proto.EnumName(Person_PhoneType_name, int32(x))
}
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Message struct {
	ClientId   int32        `protobuf:"varint,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientName string       `protobuf:"bytes,2,opt,name=client_name,json=clientName" json:"client_name,omitempty"`
	Book       *AddressBook `protobuf:"bytes,3,opt,name=book" json:"book,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetClientId() int32 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *Message) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *Message) GetBook() *AddressBook {
	if m != nil {
		return m.Book
	}
	return nil
}

type Person struct {
	Name   string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id     int32                 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Email  string                `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*Person_PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type Person_PhoneNumber struct {
	Number string           `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,enum=protobuf.Person_PhoneType" json:"type,omitempty"`
}

func (m *Person_PhoneNumber) Reset()                    { *m = Person_PhoneNumber{} }
func (m *Person_PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*Person_PhoneNumber) ProtoMessage()               {}
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Person_PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *Person_PhoneNumber) GetType() Person_PhoneType {
	if m != nil {
		return m.Type
	}
	return Person_MOBILE
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "protobuf.Message")
	proto.RegisterType((*Person)(nil), "protobuf.Person")
	proto.RegisterType((*Person_PhoneNumber)(nil), "protobuf.Person.PhoneNumber")
	proto.RegisterType((*AddressBook)(nil), "protobuf.AddressBook")
	proto.RegisterEnum("protobuf.Person_PhoneType", Person_PhoneType_name, Person_PhoneType_value)
}

func init() { proto.RegisterFile("addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0x9b, 0xa6, 0xf9, 0x36, 0x13, 0x28, 0x71, 0x50, 0x09, 0x55, 0x30, 0xe4, 0x14,
	0x11, 0x72, 0xa8, 0x82, 0x67, 0x0b, 0x05, 0x8b, 0xf6, 0x07, 0x8b, 0xe2, 0x51, 0x12, 0x33, 0x6a,
	0x68, 0x93, 0x5d, 0xb2, 0xe9, 0xa1, 0xff, 0xbb, 0x07, 0xc9, 0x6e, 0xd4, 0x22, 0x9e, 0x76, 0xe6,
	0xbd, 0xe1, 0x7d, 0x76, 0x06, 0x0e, 0xd2, 0x3c, 0xaf, 0x49, 0xa9, 0x4c, 0x88, 0x75, 0x22, 0x6b,
	0xd1, 0x08, 0x1c, 0xe8, 0x27, 0xdb, 0xbe, 0x46, 0x35, 0xfc, 0x9f, 0x93, 0x52, 0xe9, 0x1b, 0xe1,
	0x09, 0xb8, 0x2f, 0x9b, 0x82, 0xaa, 0xe6, 0xb9, 0xc8, 0x03, 0x16, 0xb2, 0xb8, 0xcf, 0x07, 0x46,
	0x98, 0xe5, 0x78, 0x06, 0x5e, 0x67, 0x56, 0x69, 0x49, 0x81, 0x15, 0xb2, 0xd8, 0xe5, 0x60, 0xa4,
	0x45, 0x5a, 0x12, 0x9e, 0x83, 0xdd, 0x02, 0x82, 0x5e, 0xc8, 0x62, 0x6f, 0x7c, 0x94, 0x7c, 0x11,
	0x92, 0x1b, 0x43, 0x9f, 0x08, 0xb1, 0xe6, 0x7a, 0x24, 0xfa, 0x60, 0xe0, 0xac, 0xa8, 0x56, 0xa2,
	0x42, 0x04, 0x5b, 0xe7, 0x31, 0x9d, 0xa7, 0x6b, 0x1c, 0x82, 0x55, 0xe4, 0x9a, 0xd0, 0xe7, 0x56,
	0x91, 0xe3, 0x21, 0xf4, 0xa9, 0x4c, 0x8b, 0x8d, 0x8e, 0x76, 0xb9, 0x69, 0xf0, 0x0a, 0x1c, 0xf9,
	0x2e, 0x2a, 0x52, 0x81, 0x1d, 0xf6, 0x62, 0x6f, 0x7c, 0xfa, 0x43, 0x34, 0xd9, 0xc9, 0xaa, 0xb5,
	0x17, 0xdb, 0x32, 0xa3, 0x9a, 0x77, 0xb3, 0xa3, 0x47, 0xf0, 0xf6, 0x64, 0x3c, 0x06, 0xa7, 0xd2,
	0x55, 0xf7, 0x81, 0xae, 0xc3, 0x04, 0xec, 0x66, 0x27, 0xcd, 0x9a, 0xc3, 0xf1, 0xe8, 0xef, 0xe8,
	0x87, 0x9d, 0x24, 0xae, 0xe7, 0xa2, 0x0b, 0x70, 0xbf, 0x25, 0x04, 0x70, 0xe6, 0xcb, 0xc9, 0xec,
	0x7e, 0xea, 0xff, 0xc3, 0x01, 0xd8, 0x4f, 0x4b, 0x7e, 0xe7, 0xb3, 0xb6, 0xba, 0x5d, 0xce, 0xa7,
	0xbe, 0x15, 0x5d, 0x83, 0xb7, 0x77, 0x13, 0x8c, 0xc1, 0x91, 0x24, 0xe4, 0xa6, 0x3d, 0x42, 0xbb,
	0x88, 0xff, 0x9b, 0xc6, 0x3b, 0x3f, 0x73, 0xb4, 0x71, 0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xbc,
	0x2f, 0x8b, 0xde, 0xd1, 0x01, 0x00, 0x00,
}
