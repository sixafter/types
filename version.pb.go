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
// source: version.proto

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

// A Semantic Versioning 2.0.0 compliant version number.
// See https://semver.org
// See the Backus-Naur form at https://semver.org/#backusnaur-form-grammar-for-valid-semver-versions
type Version struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Major version zero (0.y.z) is for initial development. Anything MAY change
	// at any time. The public API SHOULD NOT be considered stable.
	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	// Minor version Y (x.Y.z | x > 0) MUST be incremented if new, backwards compatible
	// functionality is introduced to the public API. It MUST be incremented if any
	// public API functionality is marked as deprecated. It MAY be incremented if
	// substantial new functionality or improvements are introduced within the private
	// code. It MAY include patch level changes. Patch version MUST be reset to 0 when
	// minor version is incremented.
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	// Patch version Z (x.y.Z | x > 0) MUST be incremented if only backwards compatible
	// bug fixes are introduced. A bug fix is defined as an internal change that fixes
	// incorrect behavior.
	Patch uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	// A pre-release version MAY be denoted by appending a hyphen and a series of dot
	// separated identifiers immediately following the patch version. Identifiers MUST
	// comprise only ASCII alphanumerics and hyphens [0-9A-Za-z-]. Identifiers MUST NOT
	// be empty. Numeric identifiers MUST NOT include leading zeroes. Pre-release
	// versions have a lower precedence than the associated normal version.
	//
	// A pre-release version indicates that the version is unstable and might not satisfy
	// the intended compatibility requirements as denoted by its associated normal version.
	//
	// Examples: 1.0.0-alpha, 1.0.0-alpha.1, 1.0.0-0.3.7, 1.0.0-x.7.z.92, 1.0.0-x-y-z.–.
	Prerelease string `protobuf:"bytes,4,opt,name=prerelease,proto3" json:"prerelease,omitempty"`
	// Build metadata MAY be denoted by appending a plus sign and a series of dot separated
	// identifiers immediately following the patch or pre-release version. Identifiers MUST
	// comprise only ASCII alphanumerics and hyphens [0-9A-Za-z-]. Identifiers MUST NOT be
	// empty.
	//
	// Build metadata MUST be ignored when determining version precedence. Thus two
	// versions that differ only in the build metadata, have the same precedence.
	//
	// Examples: 1.0.0-alpha+001, 1.0.0+20130313144700, 1.0.0-beta+exp.sha.5114f85,
	// 1.0.0+21AF26D3—-117B344092BD.
	BuildMetadata string `protobuf:"bytes,5,opt,name=build_metadata,json=buildMetadata,proto3" json:"build_metadata,omitempty"`
}

func (x *Version) Reset() {
	*x = Version{}
	if protoimpl.UnsafeEnabled {
		mi := &file_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Version) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Version) ProtoMessage() {}

func (x *Version) ProtoReflect() protoreflect.Message {
	mi := &file_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Version.ProtoReflect.Descriptor instead.
func (*Version) Descriptor() ([]byte, []int) {
	return file_version_proto_rawDescGZIP(), []int{0}
}

func (x *Version) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *Version) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *Version) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

func (x *Version) GetPrerelease() string {
	if x != nil {
		return x.Prerelease
	}
	return ""
}

func (x *Version) GetBuildMetadata() string {
	if x != nil {
		return x.BuildMetadata
	}
	return ""
}

var File_version_proto protoreflect.FileDescriptor

var file_version_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x72, 0x65, 0x6c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x4e, 0x0a, 0x12, 0x63,
	0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x42, 0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x69,
	0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74, 0x79, 0x70,
	0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x59, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_version_proto_rawDescOnce sync.Once
	file_version_proto_rawDescData = file_version_proto_rawDesc
)

func file_version_proto_rawDescGZIP() []byte {
	file_version_proto_rawDescOnce.Do(func() {
		file_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_version_proto_rawDescData)
	})
	return file_version_proto_rawDescData
}

var file_version_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_version_proto_goTypes = []interface{}{
	(*Version)(nil), // 0: types.Version
}
var file_version_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_version_proto_init() }
func file_version_proto_init() {
	if File_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Version); i {
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
			RawDescriptor: file_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_version_proto_goTypes,
		DependencyIndexes: file_version_proto_depIdxs,
		MessageInfos:      file_version_proto_msgTypes,
	}.Build()
	File_version_proto = out.File
	file_version_proto_rawDesc = nil
	file_version_proto_goTypes = nil
	file_version_proto_depIdxs = nil
}
