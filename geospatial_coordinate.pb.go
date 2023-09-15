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
// 	protoc        v4.24.3
// source: geospatial_coordinate.proto

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

// Represents a geospatial, or global positioning satellite (GPS), coordinate.
type GeospatialCoordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Distance north or south of the equator measured in degrees up to 90
	// degrees.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Distance measured in degrees east or west from an imaginary line (called
	// the prime meridian) that is measured from the North Pole to the South Pole
	// and that passes through Greenwich, England.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// The geospatial elevation.
	Elevation *GeospatialElevation `protobuf:"bytes,3,opt,name=elevation,proto3" json:"elevation,omitempty"`
}

func (x *GeospatialCoordinate) Reset() {
	*x = GeospatialCoordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geospatial_coordinate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeospatialCoordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeospatialCoordinate) ProtoMessage() {}

func (x *GeospatialCoordinate) ProtoReflect() protoreflect.Message {
	mi := &file_geospatial_coordinate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeospatialCoordinate.ProtoReflect.Descriptor instead.
func (*GeospatialCoordinate) Descriptor() ([]byte, []int) {
	return file_geospatial_coordinate_proto_rawDescGZIP(), []int{0}
}

func (x *GeospatialCoordinate) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *GeospatialCoordinate) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *GeospatialCoordinate) GetElevation() *GeospatialElevation {
	if x != nil {
		return x.Elevation
	}
	return nil
}

var File_geospatial_coordinate_proto protoreflect.FileDescriptor

var file_geospatial_coordinate_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x1a, 0x1a, 0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c,
	0x5f, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x43,
	0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47,
	0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x5a, 0x0a,
	0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x42, 0x18, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x45,
	0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73,
	0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_geospatial_coordinate_proto_rawDescOnce sync.Once
	file_geospatial_coordinate_proto_rawDescData = file_geospatial_coordinate_proto_rawDesc
)

func file_geospatial_coordinate_proto_rawDescGZIP() []byte {
	file_geospatial_coordinate_proto_rawDescOnce.Do(func() {
		file_geospatial_coordinate_proto_rawDescData = protoimpl.X.CompressGZIP(file_geospatial_coordinate_proto_rawDescData)
	})
	return file_geospatial_coordinate_proto_rawDescData
}

var file_geospatial_coordinate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geospatial_coordinate_proto_goTypes = []interface{}{
	(*GeospatialCoordinate)(nil), // 0: types.GeospatialCoordinate
	(*GeospatialElevation)(nil),  // 1: types.GeospatialElevation
}
var file_geospatial_coordinate_proto_depIdxs = []int32{
	1, // 0: types.GeospatialCoordinate.elevation:type_name -> types.GeospatialElevation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_geospatial_coordinate_proto_init() }
func file_geospatial_coordinate_proto_init() {
	if File_geospatial_coordinate_proto != nil {
		return
	}
	file_geospatial_elevation_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_geospatial_coordinate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeospatialCoordinate); i {
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
			RawDescriptor: file_geospatial_coordinate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geospatial_coordinate_proto_goTypes,
		DependencyIndexes: file_geospatial_coordinate_proto_depIdxs,
		MessageInfos:      file_geospatial_coordinate_proto_msgTypes,
	}.Build()
	File_geospatial_coordinate_proto = out.File
	file_geospatial_coordinate_proto_rawDesc = nil
	file_geospatial_coordinate_proto_goTypes = nil
	file_geospatial_coordinate_proto_depIdxs = nil
}
