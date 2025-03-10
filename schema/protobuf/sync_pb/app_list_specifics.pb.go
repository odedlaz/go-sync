// Copyright 2013 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for the app list (aka app launcher).

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: app_list_specifics.proto

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

// What type of item this is.
type AppListSpecifics_AppListItemType int32

const (
	// An app, whether a web app, Android app, etc.
	//
	// For bookmark apps (URL shortcuts), additional information such as their
	// URLs are kept in the AppSpecifics.bookmark_app_foobar fields.
	AppListSpecifics_TYPE_APP AppListSpecifics_AppListItemType = 1
	// A request to remove any matching default installed apps.
	AppListSpecifics_TYPE_REMOVE_DEFAULT_APP AppListSpecifics_AppListItemType = 2
	// A folder containing entries whose |parent_id| matches |item_id|.
	AppListSpecifics_TYPE_FOLDER AppListSpecifics_AppListItemType = 3
	// Obsolete type, intended for URL shortcuts, that was never implemented.
	AppListSpecifics_TYPE_OBSOLETE_URL AppListSpecifics_AppListItemType = 4
	// A "page break" item (Indicate creation of a new page in app list).
	AppListSpecifics_TYPE_PAGE_BREAK AppListSpecifics_AppListItemType = 5
)

// Enum value maps for AppListSpecifics_AppListItemType.
var (
	AppListSpecifics_AppListItemType_name = map[int32]string{
		1: "TYPE_APP",
		2: "TYPE_REMOVE_DEFAULT_APP",
		3: "TYPE_FOLDER",
		4: "TYPE_OBSOLETE_URL",
		5: "TYPE_PAGE_BREAK",
	}
	AppListSpecifics_AppListItemType_value = map[string]int32{
		"TYPE_APP":                1,
		"TYPE_REMOVE_DEFAULT_APP": 2,
		"TYPE_FOLDER":             3,
		"TYPE_OBSOLETE_URL":       4,
		"TYPE_PAGE_BREAK":         5,
	}
)

func (x AppListSpecifics_AppListItemType) Enum() *AppListSpecifics_AppListItemType {
	p := new(AppListSpecifics_AppListItemType)
	*p = x
	return p
}

func (x AppListSpecifics_AppListItemType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppListSpecifics_AppListItemType) Descriptor() protoreflect.EnumDescriptor {
	return file_app_list_specifics_proto_enumTypes[0].Descriptor()
}

func (AppListSpecifics_AppListItemType) Type() protoreflect.EnumType {
	return &file_app_list_specifics_proto_enumTypes[0]
}

func (x AppListSpecifics_AppListItemType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AppListSpecifics_AppListItemType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AppListSpecifics_AppListItemType(num)
	return nil
}

// Deprecated: Use AppListSpecifics_AppListItemType.Descriptor instead.
func (AppListSpecifics_AppListItemType) EnumDescriptor() ([]byte, []int) {
	return file_app_list_specifics_proto_rawDescGZIP(), []int{0, 0}
}

// Properties of app list objects.
type AppListSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique identifier for the item:
	// * TYPE_FOLDER: Folder id (generated)
	// * TYPE_APP: App Id
	ItemId   *string                           `protobuf:"bytes,1,opt,name=item_id,json=itemId" json:"item_id,omitempty"`
	ItemType *AppListSpecifics_AppListItemType `protobuf:"varint,2,opt,name=item_type,json=itemType,enum=sync_pb.AppListSpecifics_AppListItemType" json:"item_type,omitempty"`
	// Item name (FOLDER).
	ItemName *string `protobuf:"bytes,3,opt,name=item_name,json=itemName" json:"item_name,omitempty"`
	// Id of the parent (folder) item.
	ParentId *string `protobuf:"bytes,4,opt,name=parent_id,json=parentId" json:"parent_id,omitempty"`
	// Marked OBSOLETE because this is unused for the app list.
	// Which page this item will appear on in the app list.
	//
	// Deprecated: Do not use.
	OBSOLETEPageOrdinal *string `protobuf:"bytes,5,opt,name=OBSOLETE_page_ordinal,json=OBSOLETEPageOrdinal" json:"OBSOLETE_page_ordinal,omitempty"`
	// Where on a page this item will appear.
	ItemOrdinal *string `protobuf:"bytes,6,opt,name=item_ordinal,json=itemOrdinal" json:"item_ordinal,omitempty"`
	// Where on a shelf this item will appear.
	ItemPinOrdinal *string `protobuf:"bytes,7,opt,name=item_pin_ordinal,json=itemPinOrdinal" json:"item_pin_ordinal,omitempty"`
}

func (x *AppListSpecifics) Reset() {
	*x = AppListSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_list_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListSpecifics) ProtoMessage() {}

func (x *AppListSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_app_list_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListSpecifics.ProtoReflect.Descriptor instead.
func (*AppListSpecifics) Descriptor() ([]byte, []int) {
	return file_app_list_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *AppListSpecifics) GetItemId() string {
	if x != nil && x.ItemId != nil {
		return *x.ItemId
	}
	return ""
}

func (x *AppListSpecifics) GetItemType() AppListSpecifics_AppListItemType {
	if x != nil && x.ItemType != nil {
		return *x.ItemType
	}
	return AppListSpecifics_TYPE_APP
}

func (x *AppListSpecifics) GetItemName() string {
	if x != nil && x.ItemName != nil {
		return *x.ItemName
	}
	return ""
}

func (x *AppListSpecifics) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

// Deprecated: Do not use.
func (x *AppListSpecifics) GetOBSOLETEPageOrdinal() string {
	if x != nil && x.OBSOLETEPageOrdinal != nil {
		return *x.OBSOLETEPageOrdinal
	}
	return ""
}

func (x *AppListSpecifics) GetItemOrdinal() string {
	if x != nil && x.ItemOrdinal != nil {
		return *x.ItemOrdinal
	}
	return ""
}

func (x *AppListSpecifics) GetItemPinOrdinal() string {
	if x != nil && x.ItemPinOrdinal != nil {
		return *x.ItemPinOrdinal
	}
	return ""
}

var File_app_list_specifics_proto protoreflect.FileDescriptor

var file_app_list_specifics_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x22, 0xad, 0x03, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x46, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x41,
	0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e,
	0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65,
	0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74,
	0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x15, 0x4f, 0x42, 0x53, 0x4f, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x13, 0x4f, 0x42, 0x53, 0x4f, 0x4c, 0x45, 0x54, 0x45,
	0x50, 0x61, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x28,
	0x0a, 0x10, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x69, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69, 0x74, 0x65, 0x6d, 0x50, 0x69,
	0x6e, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x79, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54,
	0x5f, 0x41, 0x50, 0x50, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46,
	0x4f, 0x4c, 0x44, 0x45, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x42, 0x53, 0x4f, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x41, 0x47, 0x45, 0x5f, 0x42, 0x52, 0x45, 0x41,
	0x4b, 0x10, 0x05, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d,
	0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_app_list_specifics_proto_rawDescOnce sync.Once
	file_app_list_specifics_proto_rawDescData = file_app_list_specifics_proto_rawDesc
)

func file_app_list_specifics_proto_rawDescGZIP() []byte {
	file_app_list_specifics_proto_rawDescOnce.Do(func() {
		file_app_list_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_list_specifics_proto_rawDescData)
	})
	return file_app_list_specifics_proto_rawDescData
}

var file_app_list_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_app_list_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_app_list_specifics_proto_goTypes = []interface{}{
	(AppListSpecifics_AppListItemType)(0), // 0: sync_pb.AppListSpecifics.AppListItemType
	(*AppListSpecifics)(nil),              // 1: sync_pb.AppListSpecifics
}
var file_app_list_specifics_proto_depIdxs = []int32{
	0, // 0: sync_pb.AppListSpecifics.item_type:type_name -> sync_pb.AppListSpecifics.AppListItemType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_list_specifics_proto_init() }
func file_app_list_specifics_proto_init() {
	if File_app_list_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_list_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListSpecifics); i {
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
			RawDescriptor: file_app_list_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_list_specifics_proto_goTypes,
		DependencyIndexes: file_app_list_specifics_proto_depIdxs,
		EnumInfos:         file_app_list_specifics_proto_enumTypes,
		MessageInfos:      file_app_list_specifics_proto_msgTypes,
	}.Build()
	File_app_list_specifics_proto = out.File
	file_app_list_specifics_proto_rawDesc = nil
	file_app_list_specifics_proto_goTypes = nil
	file_app_list_specifics_proto_depIdxs = nil
}
