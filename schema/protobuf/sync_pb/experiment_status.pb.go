// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file

// NOTE: This API is not used in Chromium anymore, but the server still needs to
// support it for the benefit of older clients (Chrome versions up to and
// including M80 rely on it, see crbug.com/1009361).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: experiment_status.proto

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

// This request allows an unauthenticated client to check the status of the
// experiments which do not require user authentication.  The status of an
// experiment for a specific client can be determined based on the user agent
// string and/or the client id sent in the HTTPS POST request.
type ExperimentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enumerates the experiments for which the status should be returned.  An
	// experiment name must be the same as one of the field names specified in
	// ExperimentsSpecifics.  See sync/protocol/experiments_specifics.proto.
	ExperimentName []string `protobuf:"bytes,1,rep,name=experiment_name,json=experimentName" json:"experiment_name,omitempty"`
}

func (x *ExperimentStatusRequest) Reset() {
	*x = ExperimentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentStatusRequest) ProtoMessage() {}

func (x *ExperimentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentStatusRequest.ProtoReflect.Descriptor instead.
func (*ExperimentStatusRequest) Descriptor() ([]byte, []int) {
	return file_experiment_status_proto_rawDescGZIP(), []int{0}
}

func (x *ExperimentStatusRequest) GetExperimentName() []string {
	if x != nil {
		return x.ExperimentName
	}
	return nil
}

// Response to an experiment status request.
type ExperimentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Minimal time to wait before issuing another request.
	PollIntervalSeconds *int32 `protobuf:"varint,1,opt,name=poll_interval_seconds,json=pollIntervalSeconds,def=3600" json:"poll_interval_seconds,omitempty"`
	// The experiments that the client has asked for, with each experiment
	// containing exactly one experiment flag.  The client can inspect the
	// embedded flag to obtain the experiment status.  Note that the number of
	// experiments should be less than or equal to the number of experiment_name
	// sent in the request since it is possible that there is no experiment
	// matching an experiment_name.
	Experiment []*ExperimentsSpecifics `protobuf:"bytes,2,rep,name=experiment" json:"experiment,omitempty"`
}

// Default values for ExperimentStatusResponse fields.
const (
	Default_ExperimentStatusResponse_PollIntervalSeconds = int32(3600)
)

func (x *ExperimentStatusResponse) Reset() {
	*x = ExperimentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiment_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentStatusResponse) ProtoMessage() {}

func (x *ExperimentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_experiment_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentStatusResponse.ProtoReflect.Descriptor instead.
func (*ExperimentStatusResponse) Descriptor() ([]byte, []int) {
	return file_experiment_status_proto_rawDescGZIP(), []int{1}
}

func (x *ExperimentStatusResponse) GetPollIntervalSeconds() int32 {
	if x != nil && x.PollIntervalSeconds != nil {
		return *x.PollIntervalSeconds
	}
	return Default_ExperimentStatusResponse_PollIntervalSeconds
}

func (x *ExperimentStatusResponse) GetExperiment() []*ExperimentsSpecifics {
	if x != nil {
		return x.Experiment
	}
	return nil
}

var File_experiment_status_proto protoreflect.FileDescriptor

var file_experiment_status_proto_rawDesc = []byte{
	0x0a, 0x17, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x1a, 0x1b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x42, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x15, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61,
	0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x3a,
	0x04, 0x33, 0x36, 0x30, 0x30, 0x52, 0x13, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x65, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x52, 0x0a, 0x65,
	0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67,
	0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_experiment_status_proto_rawDescOnce sync.Once
	file_experiment_status_proto_rawDescData = file_experiment_status_proto_rawDesc
)

func file_experiment_status_proto_rawDescGZIP() []byte {
	file_experiment_status_proto_rawDescOnce.Do(func() {
		file_experiment_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_experiment_status_proto_rawDescData)
	})
	return file_experiment_status_proto_rawDescData
}

var file_experiment_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_experiment_status_proto_goTypes = []interface{}{
	(*ExperimentStatusRequest)(nil),  // 0: sync_pb.ExperimentStatusRequest
	(*ExperimentStatusResponse)(nil), // 1: sync_pb.ExperimentStatusResponse
	(*ExperimentsSpecifics)(nil),     // 2: sync_pb.ExperimentsSpecifics
}
var file_experiment_status_proto_depIdxs = []int32{
	2, // 0: sync_pb.ExperimentStatusResponse.experiment:type_name -> sync_pb.ExperimentsSpecifics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_experiment_status_proto_init() }
func file_experiment_status_proto_init() {
	if File_experiment_status_proto != nil {
		return
	}
	file_experiments_specifics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_experiment_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentStatusRequest); i {
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
		file_experiment_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentStatusResponse); i {
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
			RawDescriptor: file_experiment_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_experiment_status_proto_goTypes,
		DependencyIndexes: file_experiment_status_proto_depIdxs,
		MessageInfos:      file_experiment_status_proto_msgTypes,
	}.Build()
	File_experiment_status_proto = out.File
	file_experiment_status_proto_rawDesc = nil
	file_experiment_status_proto_goTypes = nil
	file_experiment_status_proto_depIdxs = nil
}
