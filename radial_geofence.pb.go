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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: radial_geofence.proto

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

// A geofence is a virtual perimeter for a real-world geographic area.
type RadialGeofence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this geofence.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The geospatial coordinate representing the center of a circle.
	Center *GeospatialCoordinate `protobuf:"bytes,2,opt,name=center,proto3" json:"center,omitempty"`
	// A distance measurement (measured in meters) from the coordinate.
	Radius float64 `protobuf:"fixed64,3,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *RadialGeofence) Reset() {
	*x = RadialGeofence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_radial_geofence_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadialGeofence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadialGeofence) ProtoMessage() {}

func (x *RadialGeofence) ProtoReflect() protoreflect.Message {
	mi := &file_radial_geofence_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadialGeofence.ProtoReflect.Descriptor instead.
func (*RadialGeofence) Descriptor() ([]byte, []int) {
	return file_radial_geofence_proto_rawDescGZIP(), []int{0}
}

func (x *RadialGeofence) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RadialGeofence) GetCenter() *GeospatialCoordinate {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *RadialGeofence) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

var File_radial_geofence_proto protoreflect.FileDescriptor

var file_radial_geofence_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1b,
	0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0e, 0x52,
	0x61, 0x64, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x33, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6f, 0x73, 0x70, 0x61,
	0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52, 0x06,
	0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x42, 0x55,
	0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x42, 0x13, 0x52, 0x61, 0x64, 0x69, 0x61, 0x6c, 0x47, 0x65, 0x6f, 0x66,
	0x65, 0x6e, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2,
	0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_radial_geofence_proto_rawDescOnce sync.Once
	file_radial_geofence_proto_rawDescData = file_radial_geofence_proto_rawDesc
)

func file_radial_geofence_proto_rawDescGZIP() []byte {
	file_radial_geofence_proto_rawDescOnce.Do(func() {
		file_radial_geofence_proto_rawDescData = protoimpl.X.CompressGZIP(file_radial_geofence_proto_rawDescData)
	})
	return file_radial_geofence_proto_rawDescData
}

var file_radial_geofence_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_radial_geofence_proto_goTypes = []interface{}{
	(*RadialGeofence)(nil),       // 0: types.RadialGeofence
	(*GeospatialCoordinate)(nil), // 1: types.GeospatialCoordinate
}
var file_radial_geofence_proto_depIdxs = []int32{
	1, // 0: types.RadialGeofence.center:type_name -> types.GeospatialCoordinate
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_radial_geofence_proto_init() }
func file_radial_geofence_proto_init() {
	if File_radial_geofence_proto != nil {
		return
	}
	file_geospatial_coordinate_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_radial_geofence_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadialGeofence); i {
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
			RawDescriptor: file_radial_geofence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_radial_geofence_proto_goTypes,
		DependencyIndexes: file_radial_geofence_proto_depIdxs,
		MessageInfos:      file_radial_geofence_proto_msgTypes,
	}.Build()
	File_radial_geofence_proto = out.File
	file_radial_geofence_proto_rawDesc = nil
	file_radial_geofence_proto_goTypes = nil
	file_radial_geofence_proto_depIdxs = nil
}
