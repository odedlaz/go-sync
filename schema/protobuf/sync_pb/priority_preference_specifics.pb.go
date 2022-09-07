// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for priority preferences.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: priority_preference_specifics.proto

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

// Properties of a synced priority preference.
type PriorityPreferenceSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Preference *PreferenceSpecifics `protobuf:"bytes,1,opt,name=preference" json:"preference,omitempty"`
}

func (x *PriorityPreferenceSpecifics) Reset() {
	*x = PriorityPreferenceSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_priority_preference_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriorityPreferenceSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriorityPreferenceSpecifics) ProtoMessage() {}

func (x *PriorityPreferenceSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_priority_preference_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriorityPreferenceSpecifics.ProtoReflect.Descriptor instead.
func (*PriorityPreferenceSpecifics) Descriptor() ([]byte, []int) {
	return file_priority_preference_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *PriorityPreferenceSpecifics) GetPreference() *PreferenceSpecifics {
	if x != nil {
		return x.Preference
	}
	return nil
}

var File_priority_preference_specifics_proto protoreflect.FileDescriptor

var file_priority_preference_specifics_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x1a,
	0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x1b, 0x50, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x70, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63,
	0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x48, 0x03, 0x50, 0x01,
}

var (
	file_priority_preference_specifics_proto_rawDescOnce sync.Once
	file_priority_preference_specifics_proto_rawDescData = file_priority_preference_specifics_proto_rawDesc
)

func file_priority_preference_specifics_proto_rawDescGZIP() []byte {
	file_priority_preference_specifics_proto_rawDescOnce.Do(func() {
		file_priority_preference_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_priority_preference_specifics_proto_rawDescData)
	})
	return file_priority_preference_specifics_proto_rawDescData
}

var file_priority_preference_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_priority_preference_specifics_proto_goTypes = []interface{}{
	(*PriorityPreferenceSpecifics)(nil), // 0: sync_pb.PriorityPreferenceSpecifics
	(*PreferenceSpecifics)(nil),         // 1: sync_pb.PreferenceSpecifics
}
var file_priority_preference_specifics_proto_depIdxs = []int32{
	1, // 0: sync_pb.PriorityPreferenceSpecifics.preference:type_name -> sync_pb.PreferenceSpecifics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_priority_preference_specifics_proto_init() }
func file_priority_preference_specifics_proto_init() {
	if File_priority_preference_specifics_proto != nil {
		return
	}
	file_preference_specifics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_priority_preference_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriorityPreferenceSpecifics); i {
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
			RawDescriptor: file_priority_preference_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_priority_preference_specifics_proto_goTypes,
		DependencyIndexes: file_priority_preference_specifics_proto_depIdxs,
		MessageInfos:      file_priority_preference_specifics_proto_msgTypes,
	}.Build()
	File_priority_preference_specifics_proto = out.File
	file_priority_preference_specifics_proto_rawDesc = nil
	file_priority_preference_specifics_proto_goTypes = nil
	file_priority_preference_specifics_proto_depIdxs = nil
}
