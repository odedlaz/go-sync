// Copyright 2019 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: web_app_specifics.proto

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

type WebAppIconInfo_Purpose int32

const (
	WebAppIconInfo_UNSPECIFIED WebAppIconInfo_Purpose = 0
	// Suitable for any purpose.
	WebAppIconInfo_ANY WebAppIconInfo_Purpose = 1
	// Designed for masking.
	WebAppIconInfo_MASKABLE WebAppIconInfo_Purpose = 2 // MONOCHROME; is never serialized.
)

// Enum value maps for WebAppIconInfo_Purpose.
var (
	WebAppIconInfo_Purpose_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "ANY",
		2: "MASKABLE",
	}
	WebAppIconInfo_Purpose_value = map[string]int32{
		"UNSPECIFIED": 0,
		"ANY":         1,
		"MASKABLE":    2,
	}
)

func (x WebAppIconInfo_Purpose) Enum() *WebAppIconInfo_Purpose {
	p := new(WebAppIconInfo_Purpose)
	*p = x
	return p
}

func (x WebAppIconInfo_Purpose) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebAppIconInfo_Purpose) Descriptor() protoreflect.EnumDescriptor {
	return file_web_app_specifics_proto_enumTypes[0].Descriptor()
}

func (WebAppIconInfo_Purpose) Type() protoreflect.EnumType {
	return &file_web_app_specifics_proto_enumTypes[0]
}

func (x WebAppIconInfo_Purpose) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *WebAppIconInfo_Purpose) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = WebAppIconInfo_Purpose(num)
	return nil
}

// Deprecated: Use WebAppIconInfo_Purpose.Descriptor instead.
func (WebAppIconInfo_Purpose) EnumDescriptor() ([]byte, []int) {
	return file_web_app_specifics_proto_rawDescGZIP(), []int{0, 0}
}

// This enum should be a subset of the DisplayMode enum in
// chrome/browser/web_applications/proto/web_app.proto and
// third_party/blink/public/mojom/manifest/display_mode.mojom
type WebAppSpecifics_UserDisplayMode int32

const (
	WebAppSpecifics_UNSPECIFIED WebAppSpecifics_UserDisplayMode = 0
	WebAppSpecifics_BROWSER     WebAppSpecifics_UserDisplayMode = 1
	// MINIMAL_UI is never serialized.
	WebAppSpecifics_STANDALONE WebAppSpecifics_UserDisplayMode = 3
)

// Enum value maps for WebAppSpecifics_UserDisplayMode.
var (
	WebAppSpecifics_UserDisplayMode_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "BROWSER",
		3: "STANDALONE",
	}
	WebAppSpecifics_UserDisplayMode_value = map[string]int32{
		"UNSPECIFIED": 0,
		"BROWSER":     1,
		"STANDALONE":  3,
	}
)

func (x WebAppSpecifics_UserDisplayMode) Enum() *WebAppSpecifics_UserDisplayMode {
	p := new(WebAppSpecifics_UserDisplayMode)
	*p = x
	return p
}

func (x WebAppSpecifics_UserDisplayMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WebAppSpecifics_UserDisplayMode) Descriptor() protoreflect.EnumDescriptor {
	return file_web_app_specifics_proto_enumTypes[1].Descriptor()
}

func (WebAppSpecifics_UserDisplayMode) Type() protoreflect.EnumType {
	return &file_web_app_specifics_proto_enumTypes[1]
}

func (x WebAppSpecifics_UserDisplayMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *WebAppSpecifics_UserDisplayMode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = WebAppSpecifics_UserDisplayMode(num)
	return nil
}

// Deprecated: Use WebAppSpecifics_UserDisplayMode.Descriptor instead.
func (WebAppSpecifics_UserDisplayMode) EnumDescriptor() ([]byte, []int) {
	return file_web_app_specifics_proto_rawDescGZIP(), []int{1, 0}
}

// Sync & local data: Information about web app icon.
type WebAppIconInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The size of the square app icon, in raw pixels.
	SizeInPx *int32 `protobuf:"varint,1,opt,name=size_in_px,json=sizeInPx" json:"size_in_px,omitempty"`
	// The URL of the app icon.
	Url *string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// The purpose or context in which the icon should be used.
	Purpose *WebAppIconInfo_Purpose `protobuf:"varint,3,opt,name=purpose,enum=sync_pb.WebAppIconInfo_Purpose" json:"purpose,omitempty"`
}

func (x *WebAppIconInfo) Reset() {
	*x = WebAppIconInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_app_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebAppIconInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAppIconInfo) ProtoMessage() {}

func (x *WebAppIconInfo) ProtoReflect() protoreflect.Message {
	mi := &file_web_app_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAppIconInfo.ProtoReflect.Descriptor instead.
func (*WebAppIconInfo) Descriptor() ([]byte, []int) {
	return file_web_app_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *WebAppIconInfo) GetSizeInPx() int32 {
	if x != nil && x.SizeInPx != nil {
		return *x.SizeInPx
	}
	return 0
}

func (x *WebAppIconInfo) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *WebAppIconInfo) GetPurpose() WebAppIconInfo_Purpose {
	if x != nil && x.Purpose != nil {
		return *x.Purpose
	}
	return WebAppIconInfo_UNSPECIFIED
}

// WebApp data. This is a synced part of
// chrome/browser/web_applications/proto/web_app.proto data.
type WebAppSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartUrl        *string                          `protobuf:"bytes,1,opt,name=start_url,json=startUrl" json:"start_url,omitempty"`
	Name            *string                          `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	UserDisplayMode *WebAppSpecifics_UserDisplayMode `protobuf:"varint,3,opt,name=user_display_mode,json=userDisplayMode,enum=sync_pb.WebAppSpecifics_UserDisplayMode" json:"user_display_mode,omitempty"`
	ThemeColor      *uint32                          `protobuf:"varint,4,opt,name=theme_color,json=themeColor" json:"theme_color,omitempty"`
	Scope           *string                          `protobuf:"bytes,5,opt,name=scope" json:"scope,omitempty"`
	IconInfos       []*WebAppIconInfo                `protobuf:"bytes,6,rep,name=icon_infos,json=iconInfos" json:"icon_infos,omitempty"`
	// Used to store the page number that the app is displayed on in
	// chrome://apps.
	UserPageOrdinal *string `protobuf:"bytes,7,opt,name=user_page_ordinal,json=userPageOrdinal" json:"user_page_ordinal,omitempty"`
	// Used to store the in-page ranking for ordering apps in its given
	// |user_page_ordinal| page.
	UserLaunchOrdinal *string `protobuf:"bytes,8,opt,name=user_launch_ordinal,json=userLaunchOrdinal" json:"user_launch_ordinal,omitempty"`
}

func (x *WebAppSpecifics) Reset() {
	*x = WebAppSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_app_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebAppSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebAppSpecifics) ProtoMessage() {}

func (x *WebAppSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_web_app_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebAppSpecifics.ProtoReflect.Descriptor instead.
func (*WebAppSpecifics) Descriptor() ([]byte, []int) {
	return file_web_app_specifics_proto_rawDescGZIP(), []int{1}
}

func (x *WebAppSpecifics) GetStartUrl() string {
	if x != nil && x.StartUrl != nil {
		return *x.StartUrl
	}
	return ""
}

func (x *WebAppSpecifics) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *WebAppSpecifics) GetUserDisplayMode() WebAppSpecifics_UserDisplayMode {
	if x != nil && x.UserDisplayMode != nil {
		return *x.UserDisplayMode
	}
	return WebAppSpecifics_UNSPECIFIED
}

func (x *WebAppSpecifics) GetThemeColor() uint32 {
	if x != nil && x.ThemeColor != nil {
		return *x.ThemeColor
	}
	return 0
}

func (x *WebAppSpecifics) GetScope() string {
	if x != nil && x.Scope != nil {
		return *x.Scope
	}
	return ""
}

func (x *WebAppSpecifics) GetIconInfos() []*WebAppIconInfo {
	if x != nil {
		return x.IconInfos
	}
	return nil
}

func (x *WebAppSpecifics) GetUserPageOrdinal() string {
	if x != nil && x.UserPageOrdinal != nil {
		return *x.UserPageOrdinal
	}
	return ""
}

func (x *WebAppSpecifics) GetUserLaunchOrdinal() string {
	if x != nil && x.UserLaunchOrdinal != nil {
		return *x.UserLaunchOrdinal
	}
	return ""
}

var File_web_app_specifics_proto protoreflect.FileDescriptor

var file_web_app_specifics_proto_rawDesc = []byte{
	0x0a, 0x17, 0x77, 0x65, 0x62, 0x5f, 0x61, 0x70, 0x70, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x22, 0xae, 0x01, 0x0a, 0x0e, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x49, 0x63, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x69, 0x6e,
	0x5f, 0x70, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x69, 0x7a, 0x65, 0x49,
	0x6e, 0x50, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x39, 0x0a, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x49, 0x63, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x52, 0x07, 0x70, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65,
	0x22, 0x31, 0x0a, 0x07, 0x50, 0x75, 0x72, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x4e, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x41, 0x53, 0x4b, 0x41, 0x42, 0x4c,
	0x45, 0x10, 0x02, 0x22, 0xa4, 0x03, 0x0a, 0x0f, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x53, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x54, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x57, 0x65,
	0x62, 0x41, 0x70, 0x70, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0f, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x70, 0x70, 0x49, 0x63, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x69, 0x63, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x67, 0x65, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x5f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x4f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x22, 0x3f, 0x0a, 0x0f, 0x55, 0x73, 0x65,
	0x72, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x0f, 0x0a, 0x0b,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a,
	0x07, 0x42, 0x52, 0x4f, 0x57, 0x53, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x54,
	0x41, 0x4e, 0x44, 0x41, 0x4c, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72,
	0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_web_app_specifics_proto_rawDescOnce sync.Once
	file_web_app_specifics_proto_rawDescData = file_web_app_specifics_proto_rawDesc
)

func file_web_app_specifics_proto_rawDescGZIP() []byte {
	file_web_app_specifics_proto_rawDescOnce.Do(func() {
		file_web_app_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_app_specifics_proto_rawDescData)
	})
	return file_web_app_specifics_proto_rawDescData
}

var file_web_app_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_web_app_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_web_app_specifics_proto_goTypes = []interface{}{
	(WebAppIconInfo_Purpose)(0),          // 0: sync_pb.WebAppIconInfo.Purpose
	(WebAppSpecifics_UserDisplayMode)(0), // 1: sync_pb.WebAppSpecifics.UserDisplayMode
	(*WebAppIconInfo)(nil),               // 2: sync_pb.WebAppIconInfo
	(*WebAppSpecifics)(nil),              // 3: sync_pb.WebAppSpecifics
}
var file_web_app_specifics_proto_depIdxs = []int32{
	0, // 0: sync_pb.WebAppIconInfo.purpose:type_name -> sync_pb.WebAppIconInfo.Purpose
	1, // 1: sync_pb.WebAppSpecifics.user_display_mode:type_name -> sync_pb.WebAppSpecifics.UserDisplayMode
	2, // 2: sync_pb.WebAppSpecifics.icon_infos:type_name -> sync_pb.WebAppIconInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_web_app_specifics_proto_init() }
func file_web_app_specifics_proto_init() {
	if File_web_app_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_app_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebAppIconInfo); i {
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
		file_web_app_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebAppSpecifics); i {
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
			RawDescriptor: file_web_app_specifics_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_app_specifics_proto_goTypes,
		DependencyIndexes: file_web_app_specifics_proto_depIdxs,
		EnumInfos:         file_web_app_specifics_proto_enumTypes,
		MessageInfos:      file_web_app_specifics_proto_msgTypes,
	}.Build()
	File_web_app_specifics_proto = out.File
	file_web_app_specifics_proto_rawDesc = nil
	file_web_app_specifics_proto_goTypes = nil
	file_web_app_specifics_proto_depIdxs = nil
}
