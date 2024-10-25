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
// 	protoc        v5.28.3
// source: geospatial_location.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Spatial or temporal scale refers to the extent of the area or the duration
// of time.
type GeospatialLocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The geographical coordinate information.
	Coordinate *GeospatialCoordinate `protobuf:"bytes,1,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	// The azimuth (orientation) of the user’s device, relative to true or
	// magnetic north.
	Heading *CompassHeading `protobuf:"bytes,2,opt,name=heading,proto3" json:"heading,omitempty"`
	// The direction in which the device is traveling, measured in degrees and
	// relative to due north.
	//
	// Course information reflects the speed and direction in which a device is
	// moving and is available only on devices with GPS hardware. Don’t confuse
	// course information with heading information. Course direction reflects the
	// direction in which the device is moving and is independent of the device’s
	// physical orientation. The most common use of course information is in
	// navigation apps.
	Course float64 `protobuf:"fixed64,3,opt,name=course,proto3" json:"course,omitempty"`
	// The instantaneous speed of the device, measured in meters per second.
	Speed float32 `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	// The time at which this location was determined.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *GeospatialLocation) Reset() {
	*x = GeospatialLocation{}
	mi := &file_geospatial_location_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeospatialLocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeospatialLocation) ProtoMessage() {}

func (x *GeospatialLocation) ProtoReflect() protoreflect.Message {
	mi := &file_geospatial_location_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeospatialLocation.ProtoReflect.Descriptor instead.
func (*GeospatialLocation) Descriptor() ([]byte, []int) {
	return file_geospatial_location_proto_rawDescGZIP(), []int{0}
}

func (x *GeospatialLocation) GetCoordinate() *GeospatialCoordinate {
	if x != nil {
		return x.Coordinate
	}
	return nil
}

func (x *GeospatialLocation) GetHeading() *CompassHeading {
	if x != nil {
		return x.Heading
	}
	return nil
}

func (x *GeospatialLocation) GetCourse() float64 {
	if x != nil {
		return x.Course
	}
	return 0
}

func (x *GeospatialLocation) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *GeospatialLocation) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

var File_geospatial_location_proto protoreflect.FileDescriptor

var file_geospatial_location_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x65, 0x6f, 0x73, 0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x68, 0x65, 0x61,
	0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x65, 0x6f, 0x73,
	0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x6f, 0x73,
	0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b,
	0x0a, 0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x6f, 0x73, 0x70,
	0x61, 0x74, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x52,
	0x0a, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x48, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x59, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61,
	0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x17, 0x47, 0x65, 0x6f, 0x73,
	0x70, 0x61, 0x74, 0x69, 0x61, 0x6c, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geospatial_location_proto_rawDescOnce sync.Once
	file_geospatial_location_proto_rawDescData = file_geospatial_location_proto_rawDesc
)

func file_geospatial_location_proto_rawDescGZIP() []byte {
	file_geospatial_location_proto_rawDescOnce.Do(func() {
		file_geospatial_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_geospatial_location_proto_rawDescData)
	})
	return file_geospatial_location_proto_rawDescData
}

var file_geospatial_location_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_geospatial_location_proto_goTypes = []any{
	(*GeospatialLocation)(nil),    // 0: types.GeospatialLocation
	(*GeospatialCoordinate)(nil),  // 1: types.GeospatialCoordinate
	(*CompassHeading)(nil),        // 2: types.CompassHeading
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_geospatial_location_proto_depIdxs = []int32{
	1, // 0: types.GeospatialLocation.coordinate:type_name -> types.GeospatialCoordinate
	2, // 1: types.GeospatialLocation.heading:type_name -> types.CompassHeading
	3, // 2: types.GeospatialLocation.timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_geospatial_location_proto_init() }
func file_geospatial_location_proto_init() {
	if File_geospatial_location_proto != nil {
		return
	}
	file_compass_heading_proto_init()
	file_geospatial_coordinate_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_geospatial_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geospatial_location_proto_goTypes,
		DependencyIndexes: file_geospatial_location_proto_depIdxs,
		MessageInfos:      file_geospatial_location_proto_msgTypes,
	}.Build()
	File_geospatial_location_proto = out.File
	file_geospatial_location_proto_rawDesc = nil
	file_geospatial_location_proto_goTypes = nil
	file_geospatial_location_proto_depIdxs = nil
}
