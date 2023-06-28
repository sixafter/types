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
// 	protoc        v4.23.3
// source: time_zone.proto

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

// A time zone is an area which observes a uniform standard time for legal,
// commercial and social purposes.
type TimeZone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation string `protobuf:"bytes,2,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	// The Standard Time (STD) offset from UTC in hours and minutes.
	UtcOffsetStd *TimeZone_TimeOffset `protobuf:"bytes,3,opt,name=utc_offset_std,json=utcOffsetStd,proto3" json:"utc_offset_std,omitempty"`
	// The Daylight Saving Time (DST) offset from UTC in hours and minutes.
	UtcOffsetDst *TimeZone_TimeOffset `protobuf:"bytes,4,opt,name=utc_offset_dst,json=utcOffsetDst,proto3" json:"utc_offset_dst,omitempty"`
}

func (x *TimeZone) Reset() {
	*x = TimeZone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeZone) ProtoMessage() {}

func (x *TimeZone) ProtoReflect() protoreflect.Message {
	mi := &file_time_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeZone.ProtoReflect.Descriptor instead.
func (*TimeZone) Descriptor() ([]byte, []int) {
	return file_time_zone_proto_rawDescGZIP(), []int{0}
}

func (x *TimeZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TimeZone) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *TimeZone) GetUtcOffsetStd() *TimeZone_TimeOffset {
	if x != nil {
		return x.UtcOffsetStd
	}
	return nil
}

func (x *TimeZone) GetUtcOffsetDst() *TimeZone_TimeOffset {
	if x != nil {
		return x.UtcOffsetDst
	}
	return nil
}

type TimeZone_TimeOffset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The offset in hours.
	Hours int32 `protobuf:"varint,1,opt,name=hours,proto3" json:"hours,omitempty"`
	// The offset in minutes.
	Minutes int32 `protobuf:"varint,2,opt,name=minutes,proto3" json:"minutes,omitempty"`
}

func (x *TimeZone_TimeOffset) Reset() {
	*x = TimeZone_TimeOffset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_time_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeZone_TimeOffset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeZone_TimeOffset) ProtoMessage() {}

func (x *TimeZone_TimeOffset) ProtoReflect() protoreflect.Message {
	mi := &file_time_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeZone_TimeOffset.ProtoReflect.Descriptor instead.
func (*TimeZone_TimeOffset) Descriptor() ([]byte, []int) {
	return file_time_zone_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TimeZone_TimeOffset) GetHours() int32 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *TimeZone_TimeOffset) GetMinutes() int32 {
	if x != nil {
		return x.Minutes
	}
	return 0
}

var File_time_zone_proto protoreflect.FileDescriptor

var file_time_zone_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x84, 0x02, 0x0a, 0x08, 0x54, 0x69, 0x6d,
	0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x62, 0x62,
	0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x61, 0x62, 0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a,
	0x0e, 0x75, 0x74, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x74, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x52, 0x0c, 0x75, 0x74, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x64, 0x12,
	0x40, 0x0a, 0x0e, 0x75, 0x74, 0x63, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x52, 0x0c, 0x75, 0x74, 0x63, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x44, 0x73,
	0x74, 0x1a, 0x3c, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x42,
	0x4f, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0d, 0x54, 0x69, 0x6d, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_time_zone_proto_rawDescOnce sync.Once
	file_time_zone_proto_rawDescData = file_time_zone_proto_rawDesc
)

func file_time_zone_proto_rawDescGZIP() []byte {
	file_time_zone_proto_rawDescOnce.Do(func() {
		file_time_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_time_zone_proto_rawDescData)
	})
	return file_time_zone_proto_rawDescData
}

var file_time_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_time_zone_proto_goTypes = []interface{}{
	(*TimeZone)(nil),            // 0: types.TimeZone
	(*TimeZone_TimeOffset)(nil), // 1: types.TimeZone.TimeOffset
}
var file_time_zone_proto_depIdxs = []int32{
	1, // 0: types.TimeZone.utc_offset_std:type_name -> types.TimeZone.TimeOffset
	1, // 1: types.TimeZone.utc_offset_dst:type_name -> types.TimeZone.TimeOffset
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_time_zone_proto_init() }
func file_time_zone_proto_init() {
	if File_time_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_time_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeZone); i {
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
		file_time_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeZone_TimeOffset); i {
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
			RawDescriptor: file_time_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_time_zone_proto_goTypes,
		DependencyIndexes: file_time_zone_proto_depIdxs,
		MessageInfos:      file_time_zone_proto_msgTypes,
	}.Build()
	File_time_zone_proto = out.File
	file_time_zone_proto_rawDesc = nil
	file_time_zone_proto_goTypes = nil
	file_time_zone_proto_depIdxs = nil
}
