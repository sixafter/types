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
// source: entity_metadata.proto

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

type EntityMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The timestamp at which the related entity was created.
	CreatedOn *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_on,json=createdOn,proto3" json:"created_on,omitempty"`
	// The timestamp at which the related entity was last modified.
	ModifiedOn *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=modified_on,json=modifiedOn,proto3" json:"modified_on,omitempty"`
	// A Boolean value indicating whether the related entity is deleted.
	Deleted bool `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// A Boolean value indicating the ephemeral nature of the entity.
	Ephemeral bool `protobuf:"varint,4,opt,name=ephemeral,proto3" json:"ephemeral,omitempty"`
	// A Boolean value indicating if the entity can be synchronized with external sources.
	Replicable bool `protobuf:"varint,5,opt,name=replicable,proto3" json:"replicable,omitempty"`
	// A Boolean value indicating if the entity is immutable.
	Immutable bool `protobuf:"varint,6,opt,name=immutable,proto3" json:"immutable,omitempty"`
	// The Semantic Version v2.0 compliant version of the entity.
	Version *Version `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
	// A set of tags.
	Tags []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	// A Boolean value indicating if the entity was synthesized (or is synthetic).
	Synthesized bool `protobuf:"varint,9,opt,name=synthesized,proto3" json:"synthesized,omitempty"`
}

func (x *EntityMetadata) Reset() {
	*x = EntityMetadata{}
	mi := &file_entity_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EntityMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityMetadata) ProtoMessage() {}

func (x *EntityMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_entity_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityMetadata.ProtoReflect.Descriptor instead.
func (*EntityMetadata) Descriptor() ([]byte, []int) {
	return file_entity_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *EntityMetadata) GetCreatedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedOn
	}
	return nil
}

func (x *EntityMetadata) GetModifiedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedOn
	}
	return nil
}

func (x *EntityMetadata) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

func (x *EntityMetadata) GetEphemeral() bool {
	if x != nil {
		return x.Ephemeral
	}
	return false
}

func (x *EntityMetadata) GetReplicable() bool {
	if x != nil {
		return x.Replicable
	}
	return false
}

func (x *EntityMetadata) GetImmutable() bool {
	if x != nil {
		return x.Immutable
	}
	return false
}

func (x *EntityMetadata) GetVersion() *Version {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *EntityMetadata) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *EntityMetadata) GetSynthesized() bool {
	if x != nil {
		return x.Synthesized
	}
	return false
}

var File_entity_metadata_proto protoreflect.FileDescriptor

var file_entity_metadata_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde,
	0x02, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x4f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x70, 0x68, 0x65, 0x6d, 0x65, 0x72, 0x61,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6d, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x28, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x64, 0x42,
	0x55, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x13, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entity_metadata_proto_rawDescOnce sync.Once
	file_entity_metadata_proto_rawDescData = file_entity_metadata_proto_rawDesc
)

func file_entity_metadata_proto_rawDescGZIP() []byte {
	file_entity_metadata_proto_rawDescOnce.Do(func() {
		file_entity_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_metadata_proto_rawDescData)
	})
	return file_entity_metadata_proto_rawDescData
}

var file_entity_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_metadata_proto_goTypes = []any{
	(*EntityMetadata)(nil),        // 0: types.EntityMetadata
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*Version)(nil),               // 2: types.Version
}
var file_entity_metadata_proto_depIdxs = []int32{
	1, // 0: types.EntityMetadata.created_on:type_name -> google.protobuf.Timestamp
	1, // 1: types.EntityMetadata.modified_on:type_name -> google.protobuf.Timestamp
	2, // 2: types.EntityMetadata.version:type_name -> types.Version
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_entity_metadata_proto_init() }
func file_entity_metadata_proto_init() {
	if File_entity_metadata_proto != nil {
		return
	}
	file_version_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_metadata_proto_goTypes,
		DependencyIndexes: file_entity_metadata_proto_depIdxs,
		MessageInfos:      file_entity_metadata_proto_msgTypes,
	}.Build()
	File_entity_metadata_proto = out.File
	file_entity_metadata_proto_rawDesc = nil
	file_entity_metadata_proto_goTypes = nil
	file_entity_metadata_proto_depIdxs = nil
}
