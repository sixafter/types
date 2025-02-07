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
// source: geospatial_elevation.proto

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

// The `GeodeticDatum` enum defines the reference system or surface
// against which the elevation is measured. These reference systems are
// approximations of the Earth's surface.
type GeospatialElevation_GeodeticDatum int32

const (
	// The altitude above mean sea level (MSL), measured in meters.
	// This is a commonly used datum for everyday elevation measurements.
	GeospatialElevation_SEA_LEVEL GeospatialElevation_GeodeticDatum = 0
	// The altitude as a height above the World Geodetic System 1984 (WGS84)
	// ellipsoid, measured in meters. This is commonly used in GPS systems.
	GeospatialElevation_WGS_84 GeospatialElevation_GeodeticDatum = 1
)

// Enum value maps for GeospatialElevation_GeodeticDatum.
var (
	GeospatialElevation_GeodeticDatum_name = map[int32]string{
		0: "SEA_LEVEL",
		1: "WGS_84",
	}
	GeospatialElevation_GeodeticDatum_value = map[string]int32{
		"SEA_LEVEL": 0,
		"WGS_84":    1,
	}
)

func (x GeospatialElevation_GeodeticDatum) Enum() *GeospatialElevation_GeodeticDatum {
	p := new(GeospatialElevation_GeodeticDatum)
	*p = x
	return p
}

func (x GeospatialElevation_GeodeticDatum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GeospatialElevation_GeodeticDatum) Descriptor() protoreflect.EnumDescriptor {
	return file_geospatial_elevation_proto_enumTypes[0].Descriptor()
}

func (GeospatialElevation_GeodeticDatum) Type() protoreflect.EnumType {
	return &file_geospatial_elevation_proto_enumTypes[0]
}

func (x GeospatialElevation_GeodeticDatum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GeospatialElevation_GeodeticDatum.Descriptor instead.
func (GeospatialElevation_GeodeticDatum) EnumDescriptor() ([]byte, []int) {
	return file_geospatial_elevation_proto_rawDescGZIP(), []int{0, 0}
}

// The `GeospatialElevation` message represents the vertical direction
// or height of a point relative to a specified reference datum. This
// is commonly used in geolocation, mapping, and geographic information systems (GIS).
type GeospatialElevation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The geodetic datum used as the reference point for the elevation measurement.
	// Example: `SEA_LEVEL` for altitude above mean sea level.
	Datum GeospatialElevation_GeodeticDatum `protobuf:"varint,1,opt,name=datum,proto3,enum=sixafter.types.GeospatialElevation_GeodeticDatum" json:"datum,omitempty"`
	// The altitude or vertical elevation of the point, measured in meters.
	// This value can be positive (above the datum) or negative (below the datum).
	// Example: An altitude of 212 meters above sea level for Westlake, Texas.
	Altitude      float64 `protobuf:"fixed64,2,opt,name=altitude,proto3" json:"altitude,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GeospatialElevation) Reset() {
	*x = GeospatialElevation{}
	mi := &file_geospatial_elevation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeospatialElevation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeospatialElevation) ProtoMessage() {}

func (x *GeospatialElevation) ProtoReflect() protoreflect.Message {
	mi := &file_geospatial_elevation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeospatialElevation.ProtoReflect.Descriptor instead.
func (*GeospatialElevation) Descriptor() ([]byte, []int) {
	return file_geospatial_elevation_proto_rawDescGZIP(), []int{0}
}

func (x *GeospatialElevation) GetDatum() GeospatialElevation_GeodeticDatum {
	if x != nil {
		return x.Datum
	}
	return GeospatialElevation_SEA_LEVEL
}

func (x *GeospatialElevation) GetAltitude() float64 {
	if x != nil {
		return x.Altitude
	}
	return 0
}

var File_geospatial_elevation_proto protoreflect.FileDescriptor

var file_geospatial_elevation_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6c, 0x65,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x69,
	0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xa6, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45,
	0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x74, 0x69,
	0x63, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x52, 0x05, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x6f,
	0x64, 0x65, 0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45,
	0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x47, 0x53,
	0x5f, 0x38, 0x34, 0x10, 0x01, 0x42, 0x7a, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x18, 0x47, 0x65, 0x6f,
	0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x50,
	0x42, 0xaa, 0x02, 0x1d, 0x53, 0x69, 0x78, 0x41, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_geospatial_elevation_proto_rawDescOnce sync.Once
	file_geospatial_elevation_proto_rawDescData []byte
)

func file_geospatial_elevation_proto_rawDescGZIP() []byte {
	file_geospatial_elevation_proto_rawDescOnce.Do(func() {
		file_geospatial_elevation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_geospatial_elevation_proto_rawDesc), len(file_geospatial_elevation_proto_rawDesc)))
	})
	return file_geospatial_elevation_proto_rawDescData
}

var file_geospatial_elevation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_geospatial_elevation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geospatial_elevation_proto_goTypes = []any{
	(GeospatialElevation_GeodeticDatum)(0), // 0: sixafter.types.GeospatialElevation.GeodeticDatum
	(*GeospatialElevation)(nil),            // 1: sixafter.types.GeospatialElevation
}
var file_geospatial_elevation_proto_depIdxs = []int32{
	0, // 0: sixafter.types.GeospatialElevation.datum:type_name -> sixafter.types.GeospatialElevation.GeodeticDatum
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_geospatial_elevation_proto_init() }
func file_geospatial_elevation_proto_init() {
	if File_geospatial_elevation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_geospatial_elevation_proto_rawDesc), len(file_geospatial_elevation_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geospatial_elevation_proto_goTypes,
		DependencyIndexes: file_geospatial_elevation_proto_depIdxs,
		EnumInfos:         file_geospatial_elevation_proto_enumTypes,
		MessageInfos:      file_geospatial_elevation_proto_msgTypes,
	}.Build()
	File_geospatial_elevation_proto = out.File
	file_geospatial_elevation_proto_goTypes = nil
	file_geospatial_elevation_proto_depIdxs = nil
}
