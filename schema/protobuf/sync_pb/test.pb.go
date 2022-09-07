// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Protocol messages used only for testing.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: test.proto

package sync_pb

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

type UnknownFieldsTestA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *bool `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
}

func (x *UnknownFieldsTestA) Reset() {
	*x = UnknownFieldsTestA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnknownFieldsTestA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownFieldsTestA) ProtoMessage() {}

func (x *UnknownFieldsTestA) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownFieldsTestA.ProtoReflect.Descriptor instead.
func (*UnknownFieldsTestA) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *UnknownFieldsTestA) GetFoo() bool {
	if x != nil && x.Foo != nil {
		return *x.Foo
	}
	return false
}

type UnknownFieldsTestB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *bool `protobuf:"varint,1,req,name=foo" json:"foo,omitempty"`
	Bar *bool `protobuf:"varint,2,req,name=bar" json:"bar,omitempty"`
}

func (x *UnknownFieldsTestB) Reset() {
	*x = UnknownFieldsTestB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnknownFieldsTestB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnknownFieldsTestB) ProtoMessage() {}

func (x *UnknownFieldsTestB) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnknownFieldsTestB.ProtoReflect.Descriptor instead.
func (*UnknownFieldsTestB) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *UnknownFieldsTestB) GetFoo() bool {
	if x != nil && x.Foo != nil {
		return *x.Foo
	}
	return false
}

func (x *UnknownFieldsTestB) GetBar() bool {
	if x != nil && x.Bar != nil {
		return *x.Bar
	}
	return false
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x26, 0x0a, 0x12, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x65, 0x73, 0x74, 0x41, 0x12, 0x10, 0x0a, 0x03, 0x66,
	0x6f, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0x38, 0x0a,
	0x12, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x54, 0x65,
	0x73, 0x74, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08,
	0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x08, 0x52, 0x03, 0x62, 0x61, 0x72, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x48, 0x03, 0x50, 0x01,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*UnknownFieldsTestA)(nil), // 0: sync_pb.UnknownFieldsTestA
	(*UnknownFieldsTestB)(nil), // 1: sync_pb.UnknownFieldsTestB
}
var file_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnknownFieldsTestA); i {
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
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnknownFieldsTestB); i {
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
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}
