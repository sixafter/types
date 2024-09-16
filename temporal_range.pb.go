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
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: temporal_range.proto

package types

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
type TemporalRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// In chronology and periodization, an epoch or reference epoch is an instant
	// in time chosen as the origin of a particular calendar era. The "epoch"
	// serves as a reference point from which time is measured.
	Epoch *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// A temporal duration.
	Duration *durationpb.Duration `protobuf:"bytes,2,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *TemporalRange) Reset() {
	*x = TemporalRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_range_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TemporalRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TemporalRange) ProtoMessage() {}

func (x *TemporalRange) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_range_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TemporalRange.ProtoReflect.Descriptor instead.
func (*TemporalRange) Descriptor() ([]byte, []int) {
	return file_temporal_range_proto_rawDescGZIP(), []int{0}
}

func (x *TemporalRange) GetEpoch() *timestamppb.Timestamp {
	if x != nil {
		return x.Epoch
	}
	return nil
}

func (x *TemporalRange) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

var File_temporal_range_proto protoreflect.FileDescriptor

var file_temporal_range_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78,
	0x0a, 0x0d, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x30, 0x0a, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63,
	0x68, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x54, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x12,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b,
	0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_range_proto_rawDescOnce sync.Once
	file_temporal_range_proto_rawDescData = file_temporal_range_proto_rawDesc
)

func file_temporal_range_proto_rawDescGZIP() []byte {
	file_temporal_range_proto_rawDescOnce.Do(func() {
		file_temporal_range_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_range_proto_rawDescData)
	})
	return file_temporal_range_proto_rawDescData
}

var file_temporal_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_range_proto_goTypes = []any{
	(*TemporalRange)(nil),         // 0: types.TemporalRange
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
}
var file_temporal_range_proto_depIdxs = []int32{
	1, // 0: types.TemporalRange.epoch:type_name -> google.protobuf.Timestamp
	2, // 1: types.TemporalRange.duration:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_temporal_range_proto_init() }
func file_temporal_range_proto_init() {
	if File_temporal_range_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_range_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TemporalRange); i {
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
			RawDescriptor: file_temporal_range_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_range_proto_goTypes,
		DependencyIndexes: file_temporal_range_proto_depIdxs,
		MessageInfos:      file_temporal_range_proto_msgTypes,
	}.Build()
	File_temporal_range_proto = out.File
	file_temporal_range_proto_rawDesc = nil
	file_temporal_range_proto_goTypes = nil
	file_temporal_range_proto_depIdxs = nil
}
