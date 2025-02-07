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
// source: map_polygon.proto

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

// The `MapPolygonPolygon` message represents a closed geometric shape on a map.
// The shape is defined by a series of points that are connected end-to-end,
// with the first and last points also connected to form a closed shape.
// This structure is commonly used in mapping applications, GIS, and geofencing.
type MapPolygon struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The set of points that define the polygon.
	// Each point is represented by a `MapPoint` message, which specifies the
	// location on the map. The points should be provided in the order they are
	// connected.
	//
	// Example: A triangle may have three points:
	// - Point 1: `x=0, y=0`
	// - Point 2: `x=10, y=0`
	// - Point 3: `x=5, y=5`
	// The first point is automatically connected to the last point to close the shape.
	Points        []*MapPoint `protobuf:"bytes,1,rep,name=points,proto3" json:"points,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MapPolygon) Reset() {
	*x = MapPolygon{}
	mi := &file_map_polygon_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapPolygon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapPolygon) ProtoMessage() {}

func (x *MapPolygon) ProtoReflect() protoreflect.Message {
	mi := &file_map_polygon_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapPolygon.ProtoReflect.Descriptor instead.
func (*MapPolygon) Descriptor() ([]byte, []int) {
	return file_map_polygon_proto_rawDescGZIP(), []int{0}
}

func (x *MapPolygon) GetPoints() []*MapPoint {
	if x != nil {
		return x.Points
	}
	return nil
}

var File_map_polygon_proto protoreflect.FileDescriptor

var file_map_polygon_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x1a, 0x0f, 0x6d, 0x61, 0x70, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x0a, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x42, 0x71, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0f, 0x4d, 0x61, 0x70, 0x50,
	0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01,
	0x01, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x42, 0xaa, 0x02, 0x1d, 0x53, 0x69, 0x78, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_map_polygon_proto_rawDescOnce sync.Once
	file_map_polygon_proto_rawDescData []byte
)

func file_map_polygon_proto_rawDescGZIP() []byte {
	file_map_polygon_proto_rawDescOnce.Do(func() {
		file_map_polygon_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_map_polygon_proto_rawDesc), len(file_map_polygon_proto_rawDesc)))
	})
	return file_map_polygon_proto_rawDescData
}

var file_map_polygon_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_map_polygon_proto_goTypes = []any{
	(*MapPolygon)(nil), // 0: sixafter.types.MapPolygon
	(*MapPoint)(nil),   // 1: sixafter.types.MapPoint
}
var file_map_polygon_proto_depIdxs = []int32{
	1, // 0: sixafter.types.MapPolygon.points:type_name -> sixafter.types.MapPoint
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_map_polygon_proto_init() }
func file_map_polygon_proto_init() {
	if File_map_polygon_proto != nil {
		return
	}
	file_map_point_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_map_polygon_proto_rawDesc), len(file_map_polygon_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_map_polygon_proto_goTypes,
		DependencyIndexes: file_map_polygon_proto_depIdxs,
		MessageInfos:      file_map_polygon_proto_msgTypes,
	}.Build()
	File_map_polygon_proto = out.File
	file_map_polygon_proto_goTypes = nil
	file_map_polygon_proto_depIdxs = nil
}
