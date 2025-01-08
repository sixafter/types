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
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: url.proto

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

// The `Url` message represents a Uniform Resource Locator (URL), which is a specific
// type of Uniform Resource Identifier (URI) used to locate resources on the web.
// A URL consists of various components like protocol, host, path, query, and more.
type Url struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The protocol (or scheme) component of the URL.
	// Example: "https", "http", or "ftp".
	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// The host component of the URL.
	// This specifies the domain name or IP address of the resource.
	// Example: "example.com" or "192.168.1.1".
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// The path component of the URL.
	// This represents the hierarchical location of the resource on the host.
	// Example: "/resources/images" in "https://example.com/resources/images".
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// The port number of the URL.
	// This specifies the communication port for the connection.
	// Example: 443 for HTTPS, 80 for HTTP.
	Port int32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The file component of the URL.
	// This is the specific file or resource being requested.
	// Example: "image.png" in "https://example.com/resources/image.png".
	File string `protobuf:"bytes,5,opt,name=file,proto3" json:"file,omitempty"`
	// The query component of the URL.
	// This contains optional parameters appended to the URL, following a "?".
	// Example: "id=123&name=test" in "https://example.com/page?id=123&name=test".
	Query string `protobuf:"bytes,6,opt,name=query,proto3" json:"query,omitempty"`
	// The authority component of the URL.
	// This typically includes the host, port, and optionally the user-information.
	// Example: "user@example.com:8080" in "https://user@example.com:8080/resource".
	Authority string `protobuf:"bytes,7,opt,name=authority,proto3" json:"authority,omitempty"`
	// The user-information component of the URL.
	// This is the optional part before the host that provides user credentials.
	// Example: "user:password" in "https://user:password@example.com".
	UserInfo string `protobuf:"bytes,8,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	// The ref (also known as anchor) component of the URL.
	// This identifies a secondary resource within the primary resource.
	// Example: "section1" in "https://example.com/page#section1".
	Ref           string `protobuf:"bytes,9,opt,name=ref,proto3" json:"ref,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Url) Reset() {
	*x = Url{}
	mi := &file_url_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_url_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_url_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Url) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Url) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Url) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Url) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *Url) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *Url) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

func (x *Url) GetUserInfo() string {
	if x != nil {
		return x.UserInfo
	}
	return ""
}

func (x *Url) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

var File_url_proto protoreflect.FileDescriptor

var file_url_proto_rawDesc = []byte{
	0x0a, 0x09, 0x75, 0x72, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x73, 0x69, 0x78,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xd4, 0x01, 0x0a, 0x03,
	0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72,
	0x65, 0x66, 0x42, 0x6a, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x08, 0x55, 0x72, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x69, 0x78, 0x61, 0x66, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b,
	0x74, 0x79, 0x70, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x42, 0xaa, 0x02,
	0x1d, 0x53, 0x69, 0x78, 0x41, 0x66, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_url_proto_rawDescOnce sync.Once
	file_url_proto_rawDescData = file_url_proto_rawDesc
)

func file_url_proto_rawDescGZIP() []byte {
	file_url_proto_rawDescOnce.Do(func() {
		file_url_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_proto_rawDescData)
	})
	return file_url_proto_rawDescData
}

var file_url_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_url_proto_goTypes = []any{
	(*Url)(nil), // 0: sixafter.types.Url
}
var file_url_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_proto_init() }
func file_url_proto_init() {
	if File_url_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_url_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_url_proto_goTypes,
		DependencyIndexes: file_url_proto_depIdxs,
		MessageInfos:      file_url_proto_msgTypes,
	}.Build()
	File_url_proto = out.File
	file_url_proto_rawDesc = nil
	file_url_proto_goTypes = nil
	file_url_proto_depIdxs = nil
}
