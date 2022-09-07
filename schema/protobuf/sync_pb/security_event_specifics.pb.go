// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for security events.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: security_event_specifics.proto

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

type SecurityEventSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The specific security event to record.
	//
	// Types that are assignable to Event:
	//	*SecurityEventSpecifics_GaiaPasswordReuseEvent
	Event isSecurityEventSpecifics_Event `protobuf_oneof:"event"`
	// Time of event, as measured by client in microseconds since Windows epoch.
	EventTimeUsec *int64 `protobuf:"varint,2,opt,name=event_time_usec,json=eventTimeUsec" json:"event_time_usec,omitempty"`
}

func (x *SecurityEventSpecifics) Reset() {
	*x = SecurityEventSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_event_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityEventSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityEventSpecifics) ProtoMessage() {}

func (x *SecurityEventSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_security_event_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityEventSpecifics.ProtoReflect.Descriptor instead.
func (*SecurityEventSpecifics) Descriptor() ([]byte, []int) {
	return file_security_event_specifics_proto_rawDescGZIP(), []int{0}
}

func (m *SecurityEventSpecifics) GetEvent() isSecurityEventSpecifics_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *SecurityEventSpecifics) GetGaiaPasswordReuseEvent() *GaiaPasswordReuse {
	if x, ok := x.GetEvent().(*SecurityEventSpecifics_GaiaPasswordReuseEvent); ok {
		return x.GaiaPasswordReuseEvent
	}
	return nil
}

func (x *SecurityEventSpecifics) GetEventTimeUsec() int64 {
	if x != nil && x.EventTimeUsec != nil {
		return *x.EventTimeUsec
	}
	return 0
}

type isSecurityEventSpecifics_Event interface {
	isSecurityEventSpecifics_Event()
}

type SecurityEventSpecifics_GaiaPasswordReuseEvent struct {
	GaiaPasswordReuseEvent *GaiaPasswordReuse `protobuf:"bytes,1,opt,name=gaia_password_reuse_event,json=gaiaPasswordReuseEvent,oneof"`
}

func (*SecurityEventSpecifics_GaiaPasswordReuseEvent) isSecurityEventSpecifics_Event() {}

var File_security_event_specifics_proto protoreflect.FileDescriptor

var file_security_event_specifics_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x19, 0x67, 0x61, 0x69, 0x61, 0x5f,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x72, 0x65, 0x75, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x16, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12,
	0x57, 0x0a, 0x19, 0x67, 0x61, 0x69, 0x61, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x5f, 0x72, 0x65, 0x75, 0x73, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x69,
	0x61, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x75, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x16, 0x67, 0x61, 0x69, 0x61, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x75, 0x73, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63,
	0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_security_event_specifics_proto_rawDescOnce sync.Once
	file_security_event_specifics_proto_rawDescData = file_security_event_specifics_proto_rawDesc
)

func file_security_event_specifics_proto_rawDescGZIP() []byte {
	file_security_event_specifics_proto_rawDescOnce.Do(func() {
		file_security_event_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_event_specifics_proto_rawDescData)
	})
	return file_security_event_specifics_proto_rawDescData
}

var file_security_event_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_security_event_specifics_proto_goTypes = []interface{}{
	(*SecurityEventSpecifics)(nil), // 0: sync_pb.SecurityEventSpecifics
	(*GaiaPasswordReuse)(nil),      // 1: sync_pb.GaiaPasswordReuse
}
var file_security_event_specifics_proto_depIdxs = []int32{
	1, // 0: sync_pb.SecurityEventSpecifics.gaia_password_reuse_event:type_name -> sync_pb.GaiaPasswordReuse
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_security_event_specifics_proto_init() }
func file_security_event_specifics_proto_init() {
	if File_security_event_specifics_proto != nil {
		return
	}
	file_gaia_password_reuse_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_security_event_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityEventSpecifics); i {
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
	file_security_event_specifics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SecurityEventSpecifics_GaiaPasswordReuseEvent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_security_event_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_event_specifics_proto_goTypes,
		DependencyIndexes: file_security_event_specifics_proto_depIdxs,
		MessageInfos:      file_security_event_specifics_proto_msgTypes,
	}.Build()
	File_security_event_specifics_proto = out.File
	file_security_event_specifics_proto_rawDesc = nil
	file_security_event_specifics_proto_goTypes = nil
	file_security_event_specifics_proto_depIdxs = nil
}
