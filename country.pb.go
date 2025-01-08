// Copyright 2020-2024 SIX AFTER, INC (SIX AFTER)
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: SIX AFTER, INC (SIX AFTER)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: country.proto

package types

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

// The `Country` message represents a country as defined by the ISO-3166-1 standard.
// It is commonly used for internationalization, geolocation, and regulatory purposes.
type Country struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The commonly used name of the country.
	// Example: "United States".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The formal or official name of the country.
	// Example: "United States of America".
	FormalName string `protobuf:"bytes,2,opt,name=formal_name,json=formalName,proto3" json:"formal_name,omitempty"`
	// The ISO 3166-1 alpha-2 code for the country.
	// This is a two-letter country code commonly used in international standards and systems.
	// Example: "US" for the United States.
	AlphabeticCode_2 string `protobuf:"bytes,3,opt,name=alphabetic_code_2,json=alphabeticCode2,proto3" json:"alphabetic_code_2,omitempty"`
	// The ISO 3166-1 alpha-3 code for the country.
	// This is a three-letter country code that offers more specificity than the alpha-2 code.
	// Example: "USA" for the United States.
	AlphabeticCode_3 string `protobuf:"bytes,4,opt,name=alphabetic_code_3,json=alphabeticCode3,proto3" json:"alphabetic_code_3,omitempty"`
	// The ISO 3166-1 numeric code for the country.
	// This is a three-digit code that is language-neutral and often used in databases or systems
	// where numeric identifiers are preferred.
	// Example: 840 for the United States.
	NumericCode   uint32 `protobuf:"varint,5,opt,name=numeric_code,json=numericCode,proto3" json:"numeric_code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Country) Reset() {
	*x = Country{}
	mi := &file_country_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetFormalName() string {
	if x != nil {
		return x.FormalName
	}
	return ""
}

func (x *Country) GetAlphabeticCode_2() string {
	if x != nil {
		return x.AlphabeticCode_2
	}
	return ""
}

func (x *Country) GetAlphabeticCode_3() string {
	if x != nil {
		return x.AlphabeticCode_3
	}
	return ""
}

func (x *Country) GetNumericCode() uint32 {
	if x != nil {
		return x.NumericCode
	}
	return 0
}

var File_country_proto protoreflect.FileDescriptor

var file_country_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22,
	0xb9, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x62, 0x65, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x12, 0x2a, 0x0a, 0x11,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f,
	0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x65,
	0x74, 0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x33, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x6e, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x42, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69,
	0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70,
	0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x42, 0xaa, 0x02, 0x1d, 0x53, 0x69,
	0x78, 0x41, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x6c,
	0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_country_proto_rawDescOnce sync.Once
	file_country_proto_rawDescData = file_country_proto_rawDesc
)

func file_country_proto_rawDescGZIP() []byte {
	file_country_proto_rawDescOnce.Do(func() {
		file_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_country_proto_rawDescData)
	})
	return file_country_proto_rawDescData
}

var file_country_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_country_proto_goTypes = []any{
	(*Country)(nil), // 0: sixafter.types.Country
}
var file_country_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_country_proto_init() }
func file_country_proto_init() {
	if File_country_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_country_proto_goTypes,
		DependencyIndexes: file_country_proto_depIdxs,
		MessageInfos:      file_country_proto_msgTypes,
	}.Build()
	File_country_proto = out.File
	file_country_proto_rawDesc = nil
	file_country_proto_goTypes = nil
	file_country_proto_depIdxs = nil
}
