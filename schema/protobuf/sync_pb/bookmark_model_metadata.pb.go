// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: bookmark_model_metadata.proto

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

// Corresponds to a single bookmark id/metadata pair.
type BookmarkMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bookmark local id.
	Id *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Bookmarks sync metadata.
	Metadata *EntityMetadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (x *BookmarkMetadata) Reset() {
	*x = BookmarkMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmark_model_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookmarkMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkMetadata) ProtoMessage() {}

func (x *BookmarkMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_bookmark_model_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkMetadata.ProtoReflect.Descriptor instead.
func (*BookmarkMetadata) Descriptor() ([]byte, []int) {
	return file_bookmark_model_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *BookmarkMetadata) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *BookmarkMetadata) GetMetadata() *EntityMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Sync proto to carry the sync metadata for the bookmarks model. It is used for
// persisting and loading sync metadata from disk.
type BookmarkModelMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bookmark global metadata.
	ModelTypeState *ModelTypeState `protobuf:"bytes,1,opt,name=model_type_state,json=modelTypeState" json:"model_type_state,omitempty"`
	// A set of all bookmarks metadata.
	BookmarksMetadata []*BookmarkMetadata `protobuf:"bytes,2,rep,name=bookmarks_metadata,json=bookmarksMetadata" json:"bookmarks_metadata,omitempty"`
	// Indicates whether the reupload of bookmarks has been triggered such that
	// they include the full title, which means that their sequence number has
	// been increased (independently of whether the commit has succeeded or even
	// started).
	// TODO(crbug.com/1066962): remove this code when most of bookmarks are
	// reuploaded.
	BookmarksFullTitleReuploaded *bool `protobuf:"varint,3,opt,name=bookmarks_full_title_reuploaded,json=bookmarksFullTitleReuploaded" json:"bookmarks_full_title_reuploaded,omitempty"`
	// The local timestamp corresponding to the last time remote updates were
	// received, in milliseconds since Unix epoch. Introduced in M86.
	// TODO(crbug.com/1032052): Remove this code once all local sync metadata
	// is required to populate the client tag (and be considered invalid
	// otherwise).
	LastSyncTime *int64 `protobuf:"varint,4,opt,name=last_sync_time,json=lastSyncTime" json:"last_sync_time,omitempty"`
}

func (x *BookmarkModelMetadata) Reset() {
	*x = BookmarkModelMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookmark_model_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookmarkModelMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookmarkModelMetadata) ProtoMessage() {}

func (x *BookmarkModelMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_bookmark_model_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookmarkModelMetadata.ProtoReflect.Descriptor instead.
func (*BookmarkModelMetadata) Descriptor() ([]byte, []int) {
	return file_bookmark_model_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *BookmarkModelMetadata) GetModelTypeState() *ModelTypeState {
	if x != nil {
		return x.ModelTypeState
	}
	return nil
}

func (x *BookmarkModelMetadata) GetBookmarksMetadata() []*BookmarkMetadata {
	if x != nil {
		return x.BookmarksMetadata
	}
	return nil
}

func (x *BookmarkModelMetadata) GetBookmarksFullTitleReuploaded() bool {
	if x != nil && x.BookmarksFullTitleReuploaded != nil {
		return *x.BookmarksFullTitleReuploaded
	}
	return false
}

func (x *BookmarkModelMetadata) GetLastSyncTime() int64 {
	if x != nil && x.LastSyncTime != nil {
		return *x.LastSyncTime
	}
	return 0
}

var File_bookmark_model_metadata_proto protoreflect.FileDescriptor

var file_bookmark_model_metadata_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x16, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x10, 0x42, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x91, 0x02, 0x0a, 0x15, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x41, 0x0a, 0x10, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x48, 0x0a,
	0x12, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x11, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x45, 0x0a, 0x1f, 0x62, 0x6f, 0x6f, 0x6b, 0x6d,
	0x61, 0x72, 0x6b, 0x73, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f,
	0x72, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x1c, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x46, 0x75, 0x6c, 0x6c, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x24,
	0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x54, 0x69, 0x6d, 0x65, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f,
	0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50,
	0x01,
}

var (
	file_bookmark_model_metadata_proto_rawDescOnce sync.Once
	file_bookmark_model_metadata_proto_rawDescData = file_bookmark_model_metadata_proto_rawDesc
)

func file_bookmark_model_metadata_proto_rawDescGZIP() []byte {
	file_bookmark_model_metadata_proto_rawDescOnce.Do(func() {
		file_bookmark_model_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookmark_model_metadata_proto_rawDescData)
	})
	return file_bookmark_model_metadata_proto_rawDescData
}

var file_bookmark_model_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bookmark_model_metadata_proto_goTypes = []interface{}{
	(*BookmarkMetadata)(nil),      // 0: sync_pb.BookmarkMetadata
	(*BookmarkModelMetadata)(nil), // 1: sync_pb.BookmarkModelMetadata
	(*EntityMetadata)(nil),        // 2: sync_pb.EntityMetadata
	(*ModelTypeState)(nil),        // 3: sync_pb.ModelTypeState
}
var file_bookmark_model_metadata_proto_depIdxs = []int32{
	2, // 0: sync_pb.BookmarkMetadata.metadata:type_name -> sync_pb.EntityMetadata
	3, // 1: sync_pb.BookmarkModelMetadata.model_type_state:type_name -> sync_pb.ModelTypeState
	0, // 2: sync_pb.BookmarkModelMetadata.bookmarks_metadata:type_name -> sync_pb.BookmarkMetadata
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bookmark_model_metadata_proto_init() }
func file_bookmark_model_metadata_proto_init() {
	if File_bookmark_model_metadata_proto != nil {
		return
	}
	file_model_type_state_proto_init()
	file_entity_metadata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_bookmark_model_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookmarkMetadata); i {
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
		file_bookmark_model_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookmarkModelMetadata); i {
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
			RawDescriptor: file_bookmark_model_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_bookmark_model_metadata_proto_goTypes,
		DependencyIndexes: file_bookmark_model_metadata_proto_depIdxs,
		MessageInfos:      file_bookmark_model_metadata_proto_msgTypes,
	}.Build()
	File_bookmark_model_metadata_proto = out.File
	file_bookmark_model_metadata_proto_rawDesc = nil
	file_bookmark_model_metadata_proto_goTypes = nil
	file_bookmark_model_metadata_proto_depIdxs = nil
}
