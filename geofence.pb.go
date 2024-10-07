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
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: geofence.proto

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
type Geofence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the geofence.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The set of polygons that define the geofence.
	Polygon []*Polygon `protobuf:"bytes,2,rep,name=polygon,proto3" json:"polygon,omitempty"`
}

func (x *Geofence) Reset() {
	*x = Geofence{}
	mi := &file_geofence_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Geofence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geofence) ProtoMessage() {}

func (x *Geofence) ProtoReflect() protoreflect.Message {
	mi := &file_geofence_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geofence.ProtoReflect.Descriptor instead.
func (*Geofence) Descriptor() ([]byte, []int) {
	return file_geofence_proto_rawDescGZIP(), []int{0}
}

func (x *Geofence) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Geofence) GetPolygon() []*Polygon {
	if x != nil {
		return x.Polygon
	}
	return nil
}

var File_geofence_proto protoreflect.FileDescriptor

var file_geofence_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0d, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x50, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x52, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e,
	0x42, 0x4f, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0d, 0x47, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59,
	0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geofence_proto_rawDescOnce sync.Once
	file_geofence_proto_rawDescData = file_geofence_proto_rawDesc
)

func file_geofence_proto_rawDescGZIP() []byte {
	file_geofence_proto_rawDescOnce.Do(func() {
		file_geofence_proto_rawDescData = protoimpl.X.CompressGZIP(file_geofence_proto_rawDescData)
	})
	return file_geofence_proto_rawDescData
}

var file_geofence_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geofence_proto_goTypes = []any{
	(*Geofence)(nil), // 0: types.Geofence
	(*Polygon)(nil),  // 1: types.Polygon
}
var file_geofence_proto_depIdxs = []int32{
	1, // 0: types.Geofence.polygon:type_name -> types.Polygon
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_geofence_proto_init() }
func file_geofence_proto_init() {
	if File_geofence_proto != nil {
		return
	}
	file_polygon_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geofence_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geofence_proto_goTypes,
		DependencyIndexes: file_geofence_proto_depIdxs,
		MessageInfos:      file_geofence_proto_msgTypes,
	}.Build()
	File_geofence_proto = out.File
	file_geofence_proto_rawDesc = nil
	file_geofence_proto_goTypes = nil
	file_geofence_proto_depIdxs = nil
}
