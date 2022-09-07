// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// This was a sync protocol datatype extension for experimental feature flags,
// also exposed via a separate ExperimentStatus API. As of M75, the datatype
// isn't used anymore, and as of M81 the ExperimentStatus API isn't used anymore
// either (see crbug.com/939819 and crbug.com/1009361). The proto definition
// needs to stay around for now so that the server can continue supporting these
// old clients .

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: experiments_specifics.proto

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

// A flag to enable support for keystore encryption.
type KeystoreEncryptionFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *KeystoreEncryptionFlags) Reset() {
	*x = KeystoreEncryptionFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeystoreEncryptionFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeystoreEncryptionFlags) ProtoMessage() {}

func (x *KeystoreEncryptionFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeystoreEncryptionFlags.ProtoReflect.Descriptor instead.
func (*KeystoreEncryptionFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *KeystoreEncryptionFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Whether history delete directives are enabled.
type HistoryDeleteDirectives struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *HistoryDeleteDirectives) Reset() {
	*x = HistoryDeleteDirectives{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryDeleteDirectives) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryDeleteDirectives) ProtoMessage() {}

func (x *HistoryDeleteDirectives) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryDeleteDirectives.ProtoReflect.Descriptor instead.
func (*HistoryDeleteDirectives) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryDeleteDirectives) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Whether this client should cull (delete) expired autofill
// entries when autofill sync is enabled.
type AutofillCullingFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *AutofillCullingFlags) Reset() {
	*x = AutofillCullingFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutofillCullingFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutofillCullingFlags) ProtoMessage() {}

func (x *AutofillCullingFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutofillCullingFlags.ProtoReflect.Descriptor instead.
func (*AutofillCullingFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{2}
}

func (x *AutofillCullingFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Whether the favicon sync datatypes are enabled, and what parameters
// they should operate under.
type FaviconSyncFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled          *bool  `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	FaviconSyncLimit *int32 `protobuf:"varint,2,opt,name=favicon_sync_limit,json=faviconSyncLimit,def=200" json:"favicon_sync_limit,omitempty"`
}

// Default values for FaviconSyncFlags fields.
const (
	Default_FaviconSyncFlags_FaviconSyncLimit = int32(200)
)

func (x *FaviconSyncFlags) Reset() {
	*x = FaviconSyncFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaviconSyncFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaviconSyncFlags) ProtoMessage() {}

func (x *FaviconSyncFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaviconSyncFlags.ProtoReflect.Descriptor instead.
func (*FaviconSyncFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{3}
}

func (x *FaviconSyncFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *FaviconSyncFlags) GetFaviconSyncLimit() int32 {
	if x != nil && x.FaviconSyncLimit != nil {
		return *x.FaviconSyncLimit
	}
	return Default_FaviconSyncFlags_FaviconSyncLimit
}

// Flags for enabling the experimental no-precommit GU feature.
type PreCommitUpdateAvoidanceFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *PreCommitUpdateAvoidanceFlags) Reset() {
	*x = PreCommitUpdateAvoidanceFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreCommitUpdateAvoidanceFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreCommitUpdateAvoidanceFlags) ProtoMessage() {}

func (x *PreCommitUpdateAvoidanceFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreCommitUpdateAvoidanceFlags.ProtoReflect.Descriptor instead.
func (*PreCommitUpdateAvoidanceFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{4}
}

func (x *PreCommitUpdateAvoidanceFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Flags for enabling the GCM feature.
type GcmChannelFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *GcmChannelFlags) Reset() {
	*x = GcmChannelFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcmChannelFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcmChannelFlags) ProtoMessage() {}

func (x *GcmChannelFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcmChannelFlags.ProtoReflect.Descriptor instead.
func (*GcmChannelFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{5}
}

func (x *GcmChannelFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Flags for enabling the experimental enhanced bookmarks feature.
type EnhancedBookmarksFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled     *bool   `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	ExtensionId *string `protobuf:"bytes,2,opt,name=extension_id,json=extensionId" json:"extension_id,omitempty"`
}

func (x *EnhancedBookmarksFlags) Reset() {
	*x = EnhancedBookmarksFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnhancedBookmarksFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnhancedBookmarksFlags) ProtoMessage() {}

func (x *EnhancedBookmarksFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnhancedBookmarksFlags.ProtoReflect.Descriptor instead.
func (*EnhancedBookmarksFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{6}
}

func (x *EnhancedBookmarksFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

func (x *EnhancedBookmarksFlags) GetExtensionId() string {
	if x != nil && x.ExtensionId != nil {
		return *x.ExtensionId
	}
	return ""
}

// Flags for enabling GCM channel for invalidations.
type GcmInvalidationsFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *GcmInvalidationsFlags) Reset() {
	*x = GcmInvalidationsFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcmInvalidationsFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcmInvalidationsFlags) ProtoMessage() {}

func (x *GcmInvalidationsFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcmInvalidationsFlags.ProtoReflect.Descriptor instead.
func (*GcmInvalidationsFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{7}
}

func (x *GcmInvalidationsFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Flags for enabling wallet data syncing.
type WalletSyncFlags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (x *WalletSyncFlags) Reset() {
	*x = WalletSyncFlags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletSyncFlags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletSyncFlags) ProtoMessage() {}

func (x *WalletSyncFlags) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletSyncFlags.ProtoReflect.Descriptor instead.
func (*WalletSyncFlags) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{8}
}

func (x *WalletSyncFlags) GetEnabled() bool {
	if x != nil && x.Enabled != nil {
		return *x.Enabled
	}
	return false
}

// Contains one flag or set of related flags.  Each node of the experiments type
// will have a unique_client_tag identifying which flags it contains.  By
// convention, the tag name should match the sub-message name.
type ExperimentsSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeystoreEncryption       *KeystoreEncryptionFlags       `protobuf:"bytes,1,opt,name=keystore_encryption,json=keystoreEncryption" json:"keystore_encryption,omitempty"`
	HistoryDeleteDirectives  *HistoryDeleteDirectives       `protobuf:"bytes,2,opt,name=history_delete_directives,json=historyDeleteDirectives" json:"history_delete_directives,omitempty"`
	AutofillCulling          *AutofillCullingFlags          `protobuf:"bytes,3,opt,name=autofill_culling,json=autofillCulling" json:"autofill_culling,omitempty"`
	FaviconSync              *FaviconSyncFlags              `protobuf:"bytes,4,opt,name=favicon_sync,json=faviconSync" json:"favicon_sync,omitempty"`
	PreCommitUpdateAvoidance *PreCommitUpdateAvoidanceFlags `protobuf:"bytes,5,opt,name=pre_commit_update_avoidance,json=preCommitUpdateAvoidance" json:"pre_commit_update_avoidance,omitempty"`
	GcmChannel               *GcmChannelFlags               `protobuf:"bytes,6,opt,name=gcm_channel,json=gcmChannel" json:"gcm_channel,omitempty"`
	// No longer used as of M43.
	ObsoleteEnhancedBookmarks *EnhancedBookmarksFlags `protobuf:"bytes,7,opt,name=obsolete_enhanced_bookmarks,json=obsoleteEnhancedBookmarks" json:"obsolete_enhanced_bookmarks,omitempty"`
	// No longer used as of M72.
	GcmInvalidations *GcmInvalidationsFlags `protobuf:"bytes,8,opt,name=gcm_invalidations,json=gcmInvalidations" json:"gcm_invalidations,omitempty"`
	// No longer used as of M51.
	ObsoleteWalletSync *WalletSyncFlags `protobuf:"bytes,9,opt,name=obsolete_wallet_sync,json=obsoleteWalletSync" json:"obsolete_wallet_sync,omitempty"`
}

func (x *ExperimentsSpecifics) Reset() {
	*x = ExperimentsSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_experiments_specifics_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExperimentsSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExperimentsSpecifics) ProtoMessage() {}

func (x *ExperimentsSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_experiments_specifics_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExperimentsSpecifics.ProtoReflect.Descriptor instead.
func (*ExperimentsSpecifics) Descriptor() ([]byte, []int) {
	return file_experiments_specifics_proto_rawDescGZIP(), []int{9}
}

func (x *ExperimentsSpecifics) GetKeystoreEncryption() *KeystoreEncryptionFlags {
	if x != nil {
		return x.KeystoreEncryption
	}
	return nil
}

func (x *ExperimentsSpecifics) GetHistoryDeleteDirectives() *HistoryDeleteDirectives {
	if x != nil {
		return x.HistoryDeleteDirectives
	}
	return nil
}

func (x *ExperimentsSpecifics) GetAutofillCulling() *AutofillCullingFlags {
	if x != nil {
		return x.AutofillCulling
	}
	return nil
}

func (x *ExperimentsSpecifics) GetFaviconSync() *FaviconSyncFlags {
	if x != nil {
		return x.FaviconSync
	}
	return nil
}

func (x *ExperimentsSpecifics) GetPreCommitUpdateAvoidance() *PreCommitUpdateAvoidanceFlags {
	if x != nil {
		return x.PreCommitUpdateAvoidance
	}
	return nil
}

func (x *ExperimentsSpecifics) GetGcmChannel() *GcmChannelFlags {
	if x != nil {
		return x.GcmChannel
	}
	return nil
}

func (x *ExperimentsSpecifics) GetObsoleteEnhancedBookmarks() *EnhancedBookmarksFlags {
	if x != nil {
		return x.ObsoleteEnhancedBookmarks
	}
	return nil
}

func (x *ExperimentsSpecifics) GetGcmInvalidations() *GcmInvalidationsFlags {
	if x != nil {
		return x.GcmInvalidations
	}
	return nil
}

func (x *ExperimentsSpecifics) GetObsoleteWalletSync() *WalletSyncFlags {
	if x != nil {
		return x.ObsoleteWalletSync
	}
	return nil
}

var File_experiments_specifics_proto protoreflect.FileDescriptor

var file_experiments_specifics_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x33, 0x0a, 0x17, 0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x33, 0x0a, 0x17, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x22, 0x30, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x43, 0x75, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x22, 0x5f, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x53, 0x79, 0x6e,
	0x63, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x31, 0x0a, 0x12, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x3a, 0x03, 0x32, 0x30,
	0x30, 0x52, 0x10, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x53, 0x79, 0x6e, 0x63, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x39, 0x0a, 0x1d, 0x50, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x76, 0x6f, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x2b,
	0x0a, 0x0f, 0x47, 0x63, 0x6d, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x55, 0x0a, 0x16, 0x45,
	0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x47, 0x63, 0x6d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x22, 0xeb, 0x05, 0x0a, 0x14, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x51, 0x0a, 0x13, 0x6b,
	0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f,
	0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x12, 0x6b, 0x65, 0x79, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c,
	0x0a, 0x19, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x73, 0x52, 0x17, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x48, 0x0a, 0x10,
	0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x63, 0x75, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x41, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x43, 0x75, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x0f, 0x61, 0x75, 0x74, 0x6f, 0x66, 0x69, 0x6c, 0x6c, 0x43,
	0x75, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x0c, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f,
	0x6e, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x46, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x53, 0x79,
	0x6e, 0x63, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x0b, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e,
	0x53, 0x79, 0x6e, 0x63, 0x12, 0x65, 0x0a, 0x1b, 0x70, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x76, 0x6f, 0x69, 0x64, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x76, 0x6f, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x52, 0x18, 0x70, 0x72, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x76, 0x6f, 0x69, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0b, 0x67,
	0x63, 0x6d, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x63, 0x6d, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x0a, 0x67, 0x63, 0x6d, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x5f, 0x0a, 0x1b, 0x6f, 0x62, 0x73, 0x6f, 0x6c, 0x65,
	0x74, 0x65, 0x5f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6b,
	0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x19, 0x6f, 0x62,
	0x73, 0x6f, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x12, 0x4b, 0x0a, 0x11, 0x67, 0x63, 0x6d, 0x5f, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x63, 0x6d,
	0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x52, 0x10, 0x67, 0x63, 0x6d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x14, 0x6f, 0x62, 0x73, 0x6f, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x12, 0x6f, 0x62,
	0x73, 0x6f, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_experiments_specifics_proto_rawDescOnce sync.Once
	file_experiments_specifics_proto_rawDescData = file_experiments_specifics_proto_rawDesc
)

func file_experiments_specifics_proto_rawDescGZIP() []byte {
	file_experiments_specifics_proto_rawDescOnce.Do(func() {
		file_experiments_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_experiments_specifics_proto_rawDescData)
	})
	return file_experiments_specifics_proto_rawDescData
}

var file_experiments_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_experiments_specifics_proto_goTypes = []interface{}{
	(*KeystoreEncryptionFlags)(nil),       // 0: sync_pb.KeystoreEncryptionFlags
	(*HistoryDeleteDirectives)(nil),       // 1: sync_pb.HistoryDeleteDirectives
	(*AutofillCullingFlags)(nil),          // 2: sync_pb.AutofillCullingFlags
	(*FaviconSyncFlags)(nil),              // 3: sync_pb.FaviconSyncFlags
	(*PreCommitUpdateAvoidanceFlags)(nil), // 4: sync_pb.PreCommitUpdateAvoidanceFlags
	(*GcmChannelFlags)(nil),               // 5: sync_pb.GcmChannelFlags
	(*EnhancedBookmarksFlags)(nil),        // 6: sync_pb.EnhancedBookmarksFlags
	(*GcmInvalidationsFlags)(nil),         // 7: sync_pb.GcmInvalidationsFlags
	(*WalletSyncFlags)(nil),               // 8: sync_pb.WalletSyncFlags
	(*ExperimentsSpecifics)(nil),          // 9: sync_pb.ExperimentsSpecifics
}
var file_experiments_specifics_proto_depIdxs = []int32{
	0, // 0: sync_pb.ExperimentsSpecifics.keystore_encryption:type_name -> sync_pb.KeystoreEncryptionFlags
	1, // 1: sync_pb.ExperimentsSpecifics.history_delete_directives:type_name -> sync_pb.HistoryDeleteDirectives
	2, // 2: sync_pb.ExperimentsSpecifics.autofill_culling:type_name -> sync_pb.AutofillCullingFlags
	3, // 3: sync_pb.ExperimentsSpecifics.favicon_sync:type_name -> sync_pb.FaviconSyncFlags
	4, // 4: sync_pb.ExperimentsSpecifics.pre_commit_update_avoidance:type_name -> sync_pb.PreCommitUpdateAvoidanceFlags
	5, // 5: sync_pb.ExperimentsSpecifics.gcm_channel:type_name -> sync_pb.GcmChannelFlags
	6, // 6: sync_pb.ExperimentsSpecifics.obsolete_enhanced_bookmarks:type_name -> sync_pb.EnhancedBookmarksFlags
	7, // 7: sync_pb.ExperimentsSpecifics.gcm_invalidations:type_name -> sync_pb.GcmInvalidationsFlags
	8, // 8: sync_pb.ExperimentsSpecifics.obsolete_wallet_sync:type_name -> sync_pb.WalletSyncFlags
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_experiments_specifics_proto_init() }
func file_experiments_specifics_proto_init() {
	if File_experiments_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_experiments_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeystoreEncryptionFlags); i {
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
		file_experiments_specifics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryDeleteDirectives); i {
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
		file_experiments_specifics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutofillCullingFlags); i {
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
		file_experiments_specifics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaviconSyncFlags); i {
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
		file_experiments_specifics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreCommitUpdateAvoidanceFlags); i {
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
		file_experiments_specifics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcmChannelFlags); i {
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
		file_experiments_specifics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnhancedBookmarksFlags); i {
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
		file_experiments_specifics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcmInvalidationsFlags); i {
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
		file_experiments_specifics_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletSyncFlags); i {
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
		file_experiments_specifics_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExperimentsSpecifics); i {
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
			RawDescriptor: file_experiments_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_experiments_specifics_proto_goTypes,
		DependencyIndexes: file_experiments_specifics_proto_depIdxs,
		MessageInfos:      file_experiments_specifics_proto_msgTypes,
	}.Build()
	File_experiments_specifics_proto = out.File
	file_experiments_specifics_proto_rawDesc = nil
	file_experiments_specifics_proto_goTypes = nil
	file_experiments_specifics_proto_depIdxs = nil
}
