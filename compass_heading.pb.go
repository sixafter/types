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
// source: compass_heading.proto

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

// The heading of an object in the Compass Geodetic System (CGS).
type CompassHeading struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The heading (measured in degrees) relative to magnetic north.
	MagneticHeading float32 `protobuf:"fixed32,1,opt,name=magnetic_heading,json=magneticHeading,proto3" json:"magnetic_heading,omitempty"`
	// The heading (measured in degrees) relative to true north.
	TrueHeading float32 `protobuf:"fixed32,2,opt,name=true_heading,json=trueHeading,proto3" json:"true_heading,omitempty"`
	// The maximum deviation (measured in degrees) between the reported heading
	// and the true geomagnetic heading.
	HeadingAccuracy float32 `protobuf:"fixed32,3,opt,name=heading_accuracy,json=headingAccuracy,proto3" json:"heading_accuracy,omitempty"`
	// The time at which this heading was determined.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The geomagnetic data (measured in microteslas) for the x-axis.
	X float32 `protobuf:"fixed32,5,opt,name=x,proto3" json:"x,omitempty"`
	// The geomagnetic data (measured in microteslas) for the y-axis.
	Y float32 `protobuf:"fixed32,6,opt,name=y,proto3" json:"y,omitempty"`
	// The geomagnetic data (measured in microteslas) for the z-axis.
	Z float32 `protobuf:"fixed32,7,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *CompassHeading) Reset() {
	*x = CompassHeading{}
	mi := &file_compass_heading_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompassHeading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompassHeading) ProtoMessage() {}

func (x *CompassHeading) ProtoReflect() protoreflect.Message {
	mi := &file_compass_heading_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompassHeading.ProtoReflect.Descriptor instead.
func (*CompassHeading) Descriptor() ([]byte, []int) {
	return file_compass_heading_proto_rawDescGZIP(), []int{0}
}

func (x *CompassHeading) GetMagneticHeading() float32 {
	if x != nil {
		return x.MagneticHeading
	}
	return 0
}

func (x *CompassHeading) GetTrueHeading() float32 {
	if x != nil {
		return x.TrueHeading
	}
	return 0
}

func (x *CompassHeading) GetHeadingAccuracy() float32 {
	if x != nil {
		return x.HeadingAccuracy
	}
	return 0
}

func (x *CompassHeading) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CompassHeading) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *CompassHeading) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *CompassHeading) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

var File_compass_heading_proto protoreflect.FileDescriptor

var file_compass_heading_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xed, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x48, 0x65, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x12, 0x29, 0x0a, 0x10, 0x6d, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x5f, 0x68,
	0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x6d, 0x61,
	0x67, 0x6e, 0x65, 0x74, 0x69, 0x63, 0x48, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a,
	0x0c, 0x74, 0x72, 0x75, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x65, 0x48, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67,
	0x12, 0x29, 0x0a, 0x10, 0x68, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01,
	0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x42,
	0x55, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x73, 0x73, 0x48, 0x65,
	0x61, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compass_heading_proto_rawDescOnce sync.Once
	file_compass_heading_proto_rawDescData = file_compass_heading_proto_rawDesc
)

func file_compass_heading_proto_rawDescGZIP() []byte {
	file_compass_heading_proto_rawDescOnce.Do(func() {
		file_compass_heading_proto_rawDescData = protoimpl.X.CompressGZIP(file_compass_heading_proto_rawDescData)
	})
	return file_compass_heading_proto_rawDescData
}

var file_compass_heading_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compass_heading_proto_goTypes = []any{
	(*CompassHeading)(nil),        // 0: types.CompassHeading
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_compass_heading_proto_depIdxs = []int32{
	1, // 0: types.CompassHeading.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_compass_heading_proto_init() }
func file_compass_heading_proto_init() {
	if File_compass_heading_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_compass_heading_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compass_heading_proto_goTypes,
		DependencyIndexes: file_compass_heading_proto_depIdxs,
		MessageInfos:      file_compass_heading_proto_msgTypes,
	}.Build()
	File_compass_heading_proto = out.File
	file_compass_heading_proto_rawDesc = nil
	file_compass_heading_proto_goTypes = nil
	file_compass_heading_proto_depIdxs = nil
}
