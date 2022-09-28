// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: book.proto

package testdata

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

type Genre int32

const (
	Genre_BLUES     Genre = 0
	Genre_CLASSICAL Genre = 1
	Genre_COUNTRY   Genre = 2
	Genre_ROCK      Genre = 3
)

// Enum value maps for Genre.
var (
	Genre_name = map[int32]string{
		0: "BLUES",
		1: "CLASSICAL",
		2: "COUNTRY",
		3: "ROCK",
	}
	Genre_value = map[string]int32{
		"BLUES":     0,
		"CLASSICAL": 1,
		"COUNTRY":   2,
		"ROCK":      3,
	}
)

func (x Genre) Enum() *Genre {
	p := new(Genre)
	*p = x
	return p
}

func (x Genre) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Genre) Descriptor() protoreflect.EnumDescriptor {
	return file_book_proto_enumTypes[0].Descriptor()
}

func (Genre) Type() protoreflect.EnumType {
	return &file_book_proto_enumTypes[0]
}

func (x Genre) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Genre.Descriptor instead.
func (Genre) EnumDescriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Isbn   int64  `protobuf:"varint,1,opt,name=isbn,proto3" json:"isbn,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Genre  Genre  `protobuf:"varint,4,opt,name=genre,proto3,enum=main.Genre" json:"genre,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_book_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_book_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_book_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetIsbn() int64 {
	if x != nil {
		return x.Isbn
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetGenre() Genre {
	if x != nil {
		return x.Genre
	}
	return Genre_BLUES
}

var File_book_proto protoreflect.FileDescriptor

var file_book_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61,
	0x69, 0x6e, 0x22, 0x6b, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x62, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x69, 0x73, 0x62, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x21, 0x0a, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2a,
	0x38, 0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x55, 0x45,
	0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x43, 0x41, 0x4c,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x02, 0x12,
	0x08, 0x0a, 0x04, 0x52, 0x4f, 0x43, 0x4b, 0x10, 0x03, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_book_proto_rawDescOnce sync.Once
	file_book_proto_rawDescData = file_book_proto_rawDesc
)

func file_book_proto_rawDescGZIP() []byte {
	file_book_proto_rawDescOnce.Do(func() {
		file_book_proto_rawDescData = protoimpl.X.CompressGZIP(file_book_proto_rawDescData)
	})
	return file_book_proto_rawDescData
}

var file_book_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_book_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_book_proto_goTypes = []interface{}{
	(Genre)(0),   // 0: main.Genre
	(*Book)(nil), // 1: main.Book
}
var file_book_proto_depIdxs = []int32{
	0, // 0: main.Book.genre:type_name -> main.Genre
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_book_proto_init() }
func file_book_proto_init() {
	if File_book_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_book_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
			RawDescriptor: file_book_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_book_proto_goTypes,
		DependencyIndexes: file_book_proto_depIdxs,
		EnumInfos:         file_book_proto_enumTypes,
		MessageInfos:      file_book_proto_msgTypes,
	}.Build()
	File_book_proto = out.File
	file_book_proto_rawDesc = nil
	file_book_proto_goTypes = nil
	file_book_proto_depIdxs = nil
}
