// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol for communication between sync client and server.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: client_commands.proto

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

type CustomNudgeDelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DatatypeId *int32 `protobuf:"varint,1,opt,name=datatype_id,json=datatypeId" json:"datatype_id,omitempty"` // Datatype id.
	DelayMs    *int32 `protobuf:"varint,2,opt,name=delay_ms,json=delayMs" json:"delay_ms,omitempty"`          // Delay in milliseconds.
}

func (x *CustomNudgeDelay) Reset() {
	*x = CustomNudgeDelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_commands_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomNudgeDelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomNudgeDelay) ProtoMessage() {}

func (x *CustomNudgeDelay) ProtoReflect() protoreflect.Message {
	mi := &file_client_commands_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomNudgeDelay.ProtoReflect.Descriptor instead.
func (*CustomNudgeDelay) Descriptor() ([]byte, []int) {
	return file_client_commands_proto_rawDescGZIP(), []int{0}
}

func (x *CustomNudgeDelay) GetDatatypeId() int32 {
	if x != nil && x.DatatypeId != nil {
		return *x.DatatypeId
	}
	return 0
}

func (x *CustomNudgeDelay) GetDelayMs() int32 {
	if x != nil && x.DelayMs != nil {
		return *x.DelayMs
	}
	return 0
}

type ClientCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time to wait before sending any requests to the server.
	SetSyncPollInterval *int32 `protobuf:"varint,1,opt,name=set_sync_poll_interval,json=setSyncPollInterval" json:"set_sync_poll_interval,omitempty"` // in seconds
	// This has been deprecated since M75.
	//
	// Deprecated: Do not use.
	SetSyncLongPollInterval *int32 `protobuf:"varint,2,opt,name=set_sync_long_poll_interval,json=setSyncLongPollInterval" json:"set_sync_long_poll_interval,omitempty"`
	MaxCommitBatchSize      *int32 `protobuf:"varint,3,opt,name=max_commit_batch_size,json=maxCommitBatchSize" json:"max_commit_batch_size,omitempty"`
	// Number of seconds to delay between a sessions action and sending a commit
	// message to the server. Can be overridden by |custom_nudge_delays|.
	SessionsCommitDelaySeconds *int32 `protobuf:"varint,4,opt,name=sessions_commit_delay_seconds,json=sessionsCommitDelaySeconds" json:"sessions_commit_delay_seconds,omitempty"`
	// Number of seconds to delay before the throttled client should retry.
	ThrottleDelaySeconds *int32 `protobuf:"varint,5,opt,name=throttle_delay_seconds,json=throttleDelaySeconds" json:"throttle_delay_seconds,omitempty"`
	// Maximum number of local nudges to buffer per-type.
	ClientInvalidationHintBufferSize *int32 `protobuf:"varint,6,opt,name=client_invalidation_hint_buffer_size,json=clientInvalidationHintBufferSize" json:"client_invalidation_hint_buffer_size,omitempty"`
	// Time to wait before issuing a retry GU.
	GuRetryDelaySeconds *int32 `protobuf:"varint,7,opt,name=gu_retry_delay_seconds,json=guRetryDelaySeconds" json:"gu_retry_delay_seconds,omitempty"`
	// A dictionary of custom nudge delays.
	// Note: if a SESSIONS value is present, this will override
	// |sessions_commit_delay_seconds|
	// New in M39.
	CustomNudgeDelays []*CustomNudgeDelay `protobuf:"bytes,8,rep,name=custom_nudge_delays,json=customNudgeDelays" json:"custom_nudge_delays,omitempty"`
}

func (x *ClientCommand) Reset() {
	*x = ClientCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_commands_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCommand) ProtoMessage() {}

func (x *ClientCommand) ProtoReflect() protoreflect.Message {
	mi := &file_client_commands_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCommand.ProtoReflect.Descriptor instead.
func (*ClientCommand) Descriptor() ([]byte, []int) {
	return file_client_commands_proto_rawDescGZIP(), []int{1}
}

func (x *ClientCommand) GetSetSyncPollInterval() int32 {
	if x != nil && x.SetSyncPollInterval != nil {
		return *x.SetSyncPollInterval
	}
	return 0
}

// Deprecated: Do not use.
func (x *ClientCommand) GetSetSyncLongPollInterval() int32 {
	if x != nil && x.SetSyncLongPollInterval != nil {
		return *x.SetSyncLongPollInterval
	}
	return 0
}

func (x *ClientCommand) GetMaxCommitBatchSize() int32 {
	if x != nil && x.MaxCommitBatchSize != nil {
		return *x.MaxCommitBatchSize
	}
	return 0
}

func (x *ClientCommand) GetSessionsCommitDelaySeconds() int32 {
	if x != nil && x.SessionsCommitDelaySeconds != nil {
		return *x.SessionsCommitDelaySeconds
	}
	return 0
}

func (x *ClientCommand) GetThrottleDelaySeconds() int32 {
	if x != nil && x.ThrottleDelaySeconds != nil {
		return *x.ThrottleDelaySeconds
	}
	return 0
}

func (x *ClientCommand) GetClientInvalidationHintBufferSize() int32 {
	if x != nil && x.ClientInvalidationHintBufferSize != nil {
		return *x.ClientInvalidationHintBufferSize
	}
	return 0
}

func (x *ClientCommand) GetGuRetryDelaySeconds() int32 {
	if x != nil && x.GuRetryDelaySeconds != nil {
		return *x.GuRetryDelaySeconds
	}
	return 0
}

func (x *ClientCommand) GetCustomNudgeDelays() []*CustomNudgeDelay {
	if x != nil {
		return x.CustomNudgeDelays
	}
	return nil
}

var File_client_commands_proto protoreflect.FileDescriptor

var file_client_commands_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x22, 0x4e, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x75, 0x64, 0x67, 0x65, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x74,
	0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73,
	0x22, 0x82, 0x04, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x33, 0x0a, 0x16, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x13, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6f, 0x6c, 0x6c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x1b, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01,
	0x52, 0x17, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x6f, 0x6e, 0x67, 0x50, 0x6f, 0x6c,
	0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x78,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x41, 0x0a, 0x1d,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x1a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x34, 0x0a, 0x16, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x14, 0x74, 0x68, 0x72, 0x6f, 0x74, 0x74, 0x6c, 0x65, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x4e, 0x0a, 0x24, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x69, 0x6e,
	0x74, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x69, 0x6e, 0x74, 0x42, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x67, 0x75, 0x5f, 0x72, 0x65, 0x74, 0x72,
	0x79, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x67, 0x75, 0x52, 0x65, 0x74, 0x72, 0x79, 0x44, 0x65,
	0x6c, 0x61, 0x79, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x49, 0x0a, 0x13, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x75, 0x64, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70,
	0x62, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x75, 0x64, 0x67, 0x65, 0x44, 0x65, 0x6c,
	0x61, 0x79, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x75, 0x64, 0x67, 0x65, 0x44,
	0x65, 0x6c, 0x61, 0x79, 0x73, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72,
	0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03,
	0x50, 0x01,
}

var (
	file_client_commands_proto_rawDescOnce sync.Once
	file_client_commands_proto_rawDescData = file_client_commands_proto_rawDesc
)

func file_client_commands_proto_rawDescGZIP() []byte {
	file_client_commands_proto_rawDescOnce.Do(func() {
		file_client_commands_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_commands_proto_rawDescData)
	})
	return file_client_commands_proto_rawDescData
}

var file_client_commands_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_commands_proto_goTypes = []interface{}{
	(*CustomNudgeDelay)(nil), // 0: sync_pb.CustomNudgeDelay
	(*ClientCommand)(nil),    // 1: sync_pb.ClientCommand
}
var file_client_commands_proto_depIdxs = []int32{
	0, // 0: sync_pb.ClientCommand.custom_nudge_delays:type_name -> sync_pb.CustomNudgeDelay
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_commands_proto_init() }
func file_client_commands_proto_init() {
	if File_client_commands_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_commands_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomNudgeDelay); i {
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
		file_client_commands_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCommand); i {
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
			RawDescriptor: file_client_commands_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_commands_proto_goTypes,
		DependencyIndexes: file_client_commands_proto_depIdxs,
		MessageInfos:      file_client_commands_proto_msgTypes,
	}.Build()
	File_client_commands_proto = out.File
	file_client_commands_proto_rawDesc = nil
	file_client_commands_proto_goTypes = nil
	file_client_commands_proto_depIdxs = nil
}
