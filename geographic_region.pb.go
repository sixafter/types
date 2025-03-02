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
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: geographic_region.proto

package types

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

// The `GeographicRegion` message represents a region as defined by the ISO 3166-2 standard.
// ISO 3166-2 defines codes for identifying the principal subdivisions (e.g., states or provinces)
// of countries coded in ISO 3166-1. This message can be used in applications requiring
// geographic or administrative classification.
type GeographicRegion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The commonly used name of the region.
	// Example: "Texas".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The formal or official name of the region.
	// Example: "State of Texas".
	FormalName string `protobuf:"bytes,2,opt,name=formal_name,json=formalName,proto3" json:"formal_name,omitempty"`
	// The numeric code for the region as defined in the ISO 3166-2 standard.
	// This code is often used in databases or systems where numeric identifiers
	// are preferred for regional classification.
	// Example: 48 (numeric code for Texas in the United States, if applicable).
	NumericCode uint32 `protobuf:"varint,3,opt,name=numeric_code,json=numericCode,proto3" json:"numeric_code,omitempty"`
	// The country to which the region belongs.
	// This should use the ISO 3166-1 alpha-2 code of the country.
	// Example: "US" for the United States.
	Country string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	// The country subdivision code for the region.
	// This should use the ISO 3166-2 code for the specific subdivision.
	// Example: "US-TX" for Texas in the United States.
	CountrySubdivision string `protobuf:"bytes,5,opt,name=country_subdivision,json=countrySubdivision,proto3" json:"country_subdivision,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GeographicRegion) Reset() {
	*x = GeographicRegion{}
	mi := &file_geographic_region_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeographicRegion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeographicRegion) ProtoMessage() {}

func (x *GeographicRegion) ProtoReflect() protoreflect.Message {
	mi := &file_geographic_region_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeographicRegion.ProtoReflect.Descriptor instead.
func (*GeographicRegion) Descriptor() ([]byte, []int) {
	return file_geographic_region_proto_rawDescGZIP(), []int{0}
}

func (x *GeographicRegion) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeographicRegion) GetFormalName() string {
	if x != nil {
		return x.FormalName
	}
	return ""
}

func (x *GeographicRegion) GetNumericCode() uint32 {
	if x != nil {
		return x.NumericCode
	}
	return 0
}

func (x *GeographicRegion) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GeographicRegion) GetCountrySubdivision() string {
	if x != nil {
		return x.CountrySubdivision
	}
	return ""
}

var File_geographic_region_proto protoreflect.FileDescriptor

var file_geographic_region_proto_rawDesc = string([]byte{
	0x0a, 0x17, 0x67, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x69, 0x78, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x63, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x73, 0x75, 0x62, 0x64,
	0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x75, 0x62, 0x64, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x42, 0x77, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x15, 0x47, 0x65, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x69, 0x63, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65,
	0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x42, 0xaa, 0x02, 0x1d, 0x53, 0x69, 0x78,
	0x41, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x6c, 0x6c,
	0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_geographic_region_proto_rawDescOnce sync.Once
	file_geographic_region_proto_rawDescData []byte
)

func file_geographic_region_proto_rawDescGZIP() []byte {
	file_geographic_region_proto_rawDescOnce.Do(func() {
		file_geographic_region_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_geographic_region_proto_rawDesc), len(file_geographic_region_proto_rawDesc)))
	})
	return file_geographic_region_proto_rawDescData
}

var file_geographic_region_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geographic_region_proto_goTypes = []any{
	(*GeographicRegion)(nil), // 0: sixafter.types.GeographicRegion
}
var file_geographic_region_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_geographic_region_proto_init() }
func file_geographic_region_proto_init() {
	if File_geographic_region_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_geographic_region_proto_rawDesc), len(file_geographic_region_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geographic_region_proto_goTypes,
		DependencyIndexes: file_geographic_region_proto_depIdxs,
		MessageInfos:      file_geographic_region_proto_msgTypes,
	}.Build()
	File_geographic_region_proto = out.File
	file_geographic_region_proto_goTypes = nil
	file_geographic_region_proto_depIdxs = nil
}
