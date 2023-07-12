// Copyright 2020 SIX AFTER, LLC (SIX AFTER)
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
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: geospatial_elevation.proto

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

// A reference system or an approximation of the Earth's surface against
// which positional measurements are made for computing locations.
type GeospatialElevation_GeodeticDatum int32

const (
	// The altitude above mean sea level associated with a location, measured in meters.
	GeospatialElevation_SEA_LEVEL GeospatialElevation_GeodeticDatum = 0
	// The altitude as a height above the World Geodetic System 1984 (WGS84) ellipsoid, measured in meters.
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

// The vertical direction, between a reference datum and a point or object.
type GeospatialElevation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A reference point.
	Datum GeospatialElevation_GeodeticDatum `protobuf:"varint,1,opt,name=datum,proto3,enum=types.GeospatialElevation_GeodeticDatum" json:"datum,omitempty"`
	// The positive or negative vertical elevation in relation to the datum.
	Altitude float64 `protobuf:"fixed64,2,opt,name=altitude,proto3" json:"altitude,omitempty"`
}

func (x *GeospatialElevation) Reset() {
	*x = GeospatialElevation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geospatial_elevation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeospatialElevation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeospatialElevation) ProtoMessage() {}

func (x *GeospatialElevation) ProtoReflect() protoreflect.Message {
	mi := &file_geospatial_elevation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_geospatial_elevation_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x65, 0x6c, 0x65,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69,
	0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x05, 0x64,
	0x61, 0x74, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6c, 0x65,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6f, 0x64, 0x65, 0x74, 0x69, 0x63, 0x44,
	0x61, 0x74, 0x75, 0x6d, 0x52, 0x05, 0x64, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x61,
	0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x2a, 0x0a, 0x0d, 0x47, 0x65, 0x6f, 0x64, 0x65,
	0x74, 0x69, 0x63, 0x44, 0x61, 0x74, 0x75, 0x6d, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x41, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x57, 0x47, 0x53, 0x5f, 0x38,
	0x34, 0x10, 0x01, 0x42, 0x5a, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66,
	0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x18, 0x47, 0x65, 0x6f, 0x73, 0x70,
	0x61, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geospatial_elevation_proto_rawDescOnce sync.Once
	file_geospatial_elevation_proto_rawDescData = file_geospatial_elevation_proto_rawDesc
)

func file_geospatial_elevation_proto_rawDescGZIP() []byte {
	file_geospatial_elevation_proto_rawDescOnce.Do(func() {
		file_geospatial_elevation_proto_rawDescData = protoimpl.X.CompressGZIP(file_geospatial_elevation_proto_rawDescData)
	})
	return file_geospatial_elevation_proto_rawDescData
}

var file_geospatial_elevation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_geospatial_elevation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geospatial_elevation_proto_goTypes = []interface{}{
	(GeospatialElevation_GeodeticDatum)(0), // 0: types.GeospatialElevation.GeodeticDatum
	(*GeospatialElevation)(nil),            // 1: types.GeospatialElevation
}
var file_geospatial_elevation_proto_depIdxs = []int32{
	0, // 0: types.GeospatialElevation.datum:type_name -> types.GeospatialElevation.GeodeticDatum
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
	if !protoimpl.UnsafeEnabled {
		file_geospatial_elevation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeospatialElevation); i {
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
			RawDescriptor: file_geospatial_elevation_proto_rawDesc,
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
	file_geospatial_elevation_proto_rawDesc = nil
	file_geospatial_elevation_proto_goTypes = nil
	file_geospatial_elevation_proto_depIdxs = nil
}
