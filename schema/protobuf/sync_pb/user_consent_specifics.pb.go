// Copyright 2018 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Sync protocol datatype extension for user consents.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: user_consent_specifics.proto

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

// Which feature does the consent apply to?
type UserConsentSpecifics_Feature int32

const (
	UserConsentSpecifics_FEATURE_UNSPECIFIED     UserConsentSpecifics_Feature = 0
	UserConsentSpecifics_CHROME_SYNC             UserConsentSpecifics_Feature = 1
	UserConsentSpecifics_PLAY_STORE              UserConsentSpecifics_Feature = 2
	UserConsentSpecifics_BACKUP_AND_RESTORE      UserConsentSpecifics_Feature = 3
	UserConsentSpecifics_GOOGLE_LOCATION_SERVICE UserConsentSpecifics_Feature = 4
	UserConsentSpecifics_CHROME_UNIFIED_CONSENT  UserConsentSpecifics_Feature = 5
	// TODO(markusheintz): ASSISTANT_ACTIVITY_CONTROL was only added for
	// compatibility with the Feature enum in UserEventSpecifics.UserConsent.
	// Delete this value once the value is deleted from the other proto.
	UserConsentSpecifics_ASSISTANT_ACTIVITY_CONTROL UserConsentSpecifics_Feature = 6
)

// Enum value maps for UserConsentSpecifics_Feature.
var (
	UserConsentSpecifics_Feature_name = map[int32]string{
		0: "FEATURE_UNSPECIFIED",
		1: "CHROME_SYNC",
		2: "PLAY_STORE",
		3: "BACKUP_AND_RESTORE",
		4: "GOOGLE_LOCATION_SERVICE",
		5: "CHROME_UNIFIED_CONSENT",
		6: "ASSISTANT_ACTIVITY_CONTROL",
	}
	UserConsentSpecifics_Feature_value = map[string]int32{
		"FEATURE_UNSPECIFIED":        0,
		"CHROME_SYNC":                1,
		"PLAY_STORE":                 2,
		"BACKUP_AND_RESTORE":         3,
		"GOOGLE_LOCATION_SERVICE":    4,
		"CHROME_UNIFIED_CONSENT":     5,
		"ASSISTANT_ACTIVITY_CONTROL": 6,
	}
)

func (x UserConsentSpecifics_Feature) Enum() *UserConsentSpecifics_Feature {
	p := new(UserConsentSpecifics_Feature)
	*p = x
	return p
}

func (x UserConsentSpecifics_Feature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserConsentSpecifics_Feature) Descriptor() protoreflect.EnumDescriptor {
	return file_user_consent_specifics_proto_enumTypes[0].Descriptor()
}

func (UserConsentSpecifics_Feature) Type() protoreflect.EnumType {
	return &file_user_consent_specifics_proto_enumTypes[0]
}

func (x UserConsentSpecifics_Feature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics_Feature) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = UserConsentSpecifics_Feature(num)
	return nil
}

// Deprecated: Use UserConsentSpecifics_Feature.Descriptor instead.
func (UserConsentSpecifics_Feature) EnumDescriptor() ([]byte, []int) {
	return file_user_consent_specifics_proto_rawDescGZIP(), []int{0, 0}
}

// Next id: 14
type UserConsentSpecifics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The UI language Chrome is using, represented as the IETF language tag
	// defined in BCP 47. The region subtag is not included when it adds no
	// distinguishing information to the language tag (e.g. both "en-US"
	// and "fr" are correct here).
	Locale *string `protobuf:"bytes,4,opt,name=locale" json:"locale,omitempty"`
	// The local time on the client when the user consent was recorded. The time
	// as measured by client is given in microseconds since Windows epoch. This
	// is needed since user consent recording may happen when a client is
	// offline.
	ClientConsentTimeUsec *int64 `protobuf:"varint,12,opt,name=client_consent_time_usec,json=clientConsentTimeUsec" json:"client_consent_time_usec,omitempty"`
	// Types that are assignable to Consent:
	//	*UserConsentSpecifics_SyncConsent
	//	*UserConsentSpecifics_ArcBackupAndRestoreConsent
	//	*UserConsentSpecifics_ArcLocationServiceConsent
	//	*UserConsentSpecifics_ArcPlayTermsOfServiceConsent
	//	*UserConsentSpecifics_UnifiedConsent
	//	*UserConsentSpecifics_AssistantActivityControlConsent
	//	*UserConsentSpecifics_AccountPasswordsConsent
	Consent isUserConsentSpecifics_Consent `protobuf_oneof:"consent"`
	// The account ID of the user who gave the consent. This field is used
	// by UserEventService to distinguish consents from different users,
	// as UserConsent does not get deleted when a user signs out. However,
	// it should be cleared before being sent over the wire, as the UserEvent
	// is sent over an authenticated channel, so this information would be
	// redundant.
	//
	// For semantics and usage of the |account_id| in the signin codebase,
	// see IdentityManager::GetPrimaryAccountId() or CoreAccountId.
	AccountId *string `protobuf:"bytes,6,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	// Deprecated: Do not use.
	Feature *UserConsentSpecifics_Feature `protobuf:"varint,1,opt,name=feature,enum=sync_pb.UserConsentSpecifics_Feature" json:"feature,omitempty"`
	// Ids of the strings of the consent text presented to the user.
	//
	// Deprecated: Do not use.
	DescriptionGrdIds []int32 `protobuf:"varint,2,rep,name=description_grd_ids,json=descriptionGrdIds" json:"description_grd_ids,omitempty"`
	// Id of the string of the UI element the user clicked when consenting.
	//
	// Deprecated: Do not use.
	ConfirmationGrdId *int32 `protobuf:"varint,3,opt,name=confirmation_grd_id,json=confirmationGrdId" json:"confirmation_grd_id,omitempty"`
	// Was the consent for |feature| given or not given (denied/revoked)?
	//
	// Deprecated: Do not use.
	Status *UserConsentTypes_ConsentStatus `protobuf:"varint,5,opt,name=status,enum=sync_pb.UserConsentTypes_ConsentStatus" json:"status,omitempty"`
}

func (x *UserConsentSpecifics) Reset() {
	*x = UserConsentSpecifics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_consent_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConsentSpecifics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConsentSpecifics) ProtoMessage() {}

func (x *UserConsentSpecifics) ProtoReflect() protoreflect.Message {
	mi := &file_user_consent_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConsentSpecifics.ProtoReflect.Descriptor instead.
func (*UserConsentSpecifics) Descriptor() ([]byte, []int) {
	return file_user_consent_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *UserConsentSpecifics) GetLocale() string {
	if x != nil && x.Locale != nil {
		return *x.Locale
	}
	return ""
}

func (x *UserConsentSpecifics) GetClientConsentTimeUsec() int64 {
	if x != nil && x.ClientConsentTimeUsec != nil {
		return *x.ClientConsentTimeUsec
	}
	return 0
}

func (m *UserConsentSpecifics) GetConsent() isUserConsentSpecifics_Consent {
	if m != nil {
		return m.Consent
	}
	return nil
}

func (x *UserConsentSpecifics) GetSyncConsent() *UserConsentTypes_SyncConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_SyncConsent); ok {
		return x.SyncConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetArcBackupAndRestoreConsent() *UserConsentTypes_ArcBackupAndRestoreConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_ArcBackupAndRestoreConsent); ok {
		return x.ArcBackupAndRestoreConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetArcLocationServiceConsent() *UserConsentTypes_ArcGoogleLocationServiceConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_ArcLocationServiceConsent); ok {
		return x.ArcLocationServiceConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetArcPlayTermsOfServiceConsent() *UserConsentTypes_ArcPlayTermsOfServiceConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_ArcPlayTermsOfServiceConsent); ok {
		return x.ArcPlayTermsOfServiceConsent
	}
	return nil
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics) GetUnifiedConsent() *UserConsentTypes_UnifiedConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_UnifiedConsent); ok {
		return x.UnifiedConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetAssistantActivityControlConsent() *UserConsentTypes_AssistantActivityControlConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_AssistantActivityControlConsent); ok {
		return x.AssistantActivityControlConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetAccountPasswordsConsent() *UserConsentTypes_AccountPasswordsConsent {
	if x, ok := x.GetConsent().(*UserConsentSpecifics_AccountPasswordsConsent); ok {
		return x.AccountPasswordsConsent
	}
	return nil
}

func (x *UserConsentSpecifics) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics) GetFeature() UserConsentSpecifics_Feature {
	if x != nil && x.Feature != nil {
		return *x.Feature
	}
	return UserConsentSpecifics_FEATURE_UNSPECIFIED
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics) GetDescriptionGrdIds() []int32 {
	if x != nil {
		return x.DescriptionGrdIds
	}
	return nil
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics) GetConfirmationGrdId() int32 {
	if x != nil && x.ConfirmationGrdId != nil {
		return *x.ConfirmationGrdId
	}
	return 0
}

// Deprecated: Do not use.
func (x *UserConsentSpecifics) GetStatus() UserConsentTypes_ConsentStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return UserConsentTypes_CONSENT_STATUS_UNSPECIFIED
}

type isUserConsentSpecifics_Consent interface {
	isUserConsentSpecifics_Consent()
}

type UserConsentSpecifics_SyncConsent struct {
	SyncConsent *UserConsentTypes_SyncConsent `protobuf:"bytes,7,opt,name=sync_consent,json=syncConsent,oneof"`
}

type UserConsentSpecifics_ArcBackupAndRestoreConsent struct {
	ArcBackupAndRestoreConsent *UserConsentTypes_ArcBackupAndRestoreConsent `protobuf:"bytes,8,opt,name=arc_backup_and_restore_consent,json=arcBackupAndRestoreConsent,oneof"`
}

type UserConsentSpecifics_ArcLocationServiceConsent struct {
	ArcLocationServiceConsent *UserConsentTypes_ArcGoogleLocationServiceConsent `protobuf:"bytes,9,opt,name=arc_location_service_consent,json=arcLocationServiceConsent,oneof"`
}

type UserConsentSpecifics_ArcPlayTermsOfServiceConsent struct {
	ArcPlayTermsOfServiceConsent *UserConsentTypes_ArcPlayTermsOfServiceConsent `protobuf:"bytes,10,opt,name=arc_play_terms_of_service_consent,json=arcPlayTermsOfServiceConsent,oneof"`
}

type UserConsentSpecifics_UnifiedConsent struct {
	// Deprecated: Do not use.
	UnifiedConsent *UserConsentTypes_UnifiedConsent `protobuf:"bytes,13,opt,name=unified_consent,json=unifiedConsent,oneof"`
}

type UserConsentSpecifics_AssistantActivityControlConsent struct {
	AssistantActivityControlConsent *UserConsentTypes_AssistantActivityControlConsent `protobuf:"bytes,14,opt,name=assistant_activity_control_consent,json=assistantActivityControlConsent,oneof"`
}

type UserConsentSpecifics_AccountPasswordsConsent struct {
	AccountPasswordsConsent *UserConsentTypes_AccountPasswordsConsent `protobuf:"bytes,15,opt,name=account_passwords_consent,json=accountPasswordsConsent,oneof"`
}

func (*UserConsentSpecifics_SyncConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_ArcBackupAndRestoreConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_ArcLocationServiceConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_ArcPlayTermsOfServiceConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_UnifiedConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_AssistantActivityControlConsent) isUserConsentSpecifics_Consent() {}

func (*UserConsentSpecifics_AccountPasswordsConsent) isUserConsentSpecifics_Consent() {}

var File_user_consent_specifics_proto protoreflect.FileDescriptor

var file_user_consent_specifics_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x1a, 0x18, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xfe, 0x0a, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61,
	0x6c, 0x65, 0x12, 0x37, 0x0a, 0x18, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x63, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x63, 0x12, 0x4a, 0x0a, 0x0c, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e,
	0x63, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63,
	0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x7a, 0x0a, 0x1e, 0x61, 0x72, 0x63, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x63, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x1a, 0x61, 0x72, 0x63, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x12, 0x7c, 0x0a, 0x1c, 0x61, 0x72, 0x63, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x63, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x19, 0x61, 0x72, 0x63, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x12, 0x81, 0x01, 0x0a, 0x21, 0x61, 0x72, 0x63, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74,
	0x65, 0x72, 0x6d, 0x73, 0x5f, 0x6f, 0x66, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x72, 0x63, 0x50, 0x6c, 0x61, 0x79,
	0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x1c, 0x61, 0x72, 0x63, 0x50, 0x6c, 0x61, 0x79,
	0x54, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64,
	0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x42, 0x02, 0x18, 0x01, 0x48, 0x00, 0x52, 0x0e,
	0x75, 0x6e, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x88,
	0x01, 0x0a, 0x22, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x1f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x6f, 0x0a, 0x19, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x48,
	0x00, 0x52, 0x17, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x07, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x73, 0x79, 0x6e,
	0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74,
	0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x32,
	0x0a, 0x13, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72,
	0x64, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x18, 0x01, 0x52,
	0x11, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x64, 0x49,
	0x64, 0x73, 0x12, 0x32, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x72, 0x64, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x07,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x45, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x10,
	0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4c, 0x41, 0x59, 0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10,
	0x02, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x5f, 0x41, 0x4e, 0x44, 0x5f,
	0x52, 0x45, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x4f, 0x4f,
	0x47, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x49, 0x43, 0x45, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x43, 0x48, 0x52, 0x4f, 0x4d, 0x45,
	0x5f, 0x55, 0x4e, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x4e, 0x54,
	0x10, 0x05, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x53, 0x53, 0x49, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x4f, 0x4c,
	0x10, 0x06, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08,
	0x0b, 0x10, 0x0c, 0x52, 0x1d, 0x61, 0x72, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x74, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69,
	0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79,
	0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_user_consent_specifics_proto_rawDescOnce sync.Once
	file_user_consent_specifics_proto_rawDescData = file_user_consent_specifics_proto_rawDesc
)

func file_user_consent_specifics_proto_rawDescGZIP() []byte {
	file_user_consent_specifics_proto_rawDescOnce.Do(func() {
		file_user_consent_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_consent_specifics_proto_rawDescData)
	})
	return file_user_consent_specifics_proto_rawDescData
}

var file_user_consent_specifics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_consent_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_user_consent_specifics_proto_goTypes = []interface{}{
	(UserConsentSpecifics_Feature)(0),                        // 0: sync_pb.UserConsentSpecifics.Feature
	(*UserConsentSpecifics)(nil),                             // 1: sync_pb.UserConsentSpecifics
	(*UserConsentTypes_SyncConsent)(nil),                     // 2: sync_pb.UserConsentTypes.SyncConsent
	(*UserConsentTypes_ArcBackupAndRestoreConsent)(nil),      // 3: sync_pb.UserConsentTypes.ArcBackupAndRestoreConsent
	(*UserConsentTypes_ArcGoogleLocationServiceConsent)(nil), // 4: sync_pb.UserConsentTypes.ArcGoogleLocationServiceConsent
	(*UserConsentTypes_ArcPlayTermsOfServiceConsent)(nil),    // 5: sync_pb.UserConsentTypes.ArcPlayTermsOfServiceConsent
	(*UserConsentTypes_UnifiedConsent)(nil),                  // 6: sync_pb.UserConsentTypes.UnifiedConsent
	(*UserConsentTypes_AssistantActivityControlConsent)(nil), // 7: sync_pb.UserConsentTypes.AssistantActivityControlConsent
	(*UserConsentTypes_AccountPasswordsConsent)(nil),         // 8: sync_pb.UserConsentTypes.AccountPasswordsConsent
	(UserConsentTypes_ConsentStatus)(0),                      // 9: sync_pb.UserConsentTypes.ConsentStatus
}
var file_user_consent_specifics_proto_depIdxs = []int32{
	2, // 0: sync_pb.UserConsentSpecifics.sync_consent:type_name -> sync_pb.UserConsentTypes.SyncConsent
	3, // 1: sync_pb.UserConsentSpecifics.arc_backup_and_restore_consent:type_name -> sync_pb.UserConsentTypes.ArcBackupAndRestoreConsent
	4, // 2: sync_pb.UserConsentSpecifics.arc_location_service_consent:type_name -> sync_pb.UserConsentTypes.ArcGoogleLocationServiceConsent
	5, // 3: sync_pb.UserConsentSpecifics.arc_play_terms_of_service_consent:type_name -> sync_pb.UserConsentTypes.ArcPlayTermsOfServiceConsent
	6, // 4: sync_pb.UserConsentSpecifics.unified_consent:type_name -> sync_pb.UserConsentTypes.UnifiedConsent
	7, // 5: sync_pb.UserConsentSpecifics.assistant_activity_control_consent:type_name -> sync_pb.UserConsentTypes.AssistantActivityControlConsent
	8, // 6: sync_pb.UserConsentSpecifics.account_passwords_consent:type_name -> sync_pb.UserConsentTypes.AccountPasswordsConsent
	0, // 7: sync_pb.UserConsentSpecifics.feature:type_name -> sync_pb.UserConsentSpecifics.Feature
	9, // 8: sync_pb.UserConsentSpecifics.status:type_name -> sync_pb.UserConsentTypes.ConsentStatus
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_user_consent_specifics_proto_init() }
func file_user_consent_specifics_proto_init() {
	if File_user_consent_specifics_proto != nil {
		return
	}
	file_user_consent_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_consent_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConsentSpecifics); i {
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
	file_user_consent_specifics_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*UserConsentSpecifics_SyncConsent)(nil),
		(*UserConsentSpecifics_ArcBackupAndRestoreConsent)(nil),
		(*UserConsentSpecifics_ArcLocationServiceConsent)(nil),
		(*UserConsentSpecifics_ArcPlayTermsOfServiceConsent)(nil),
		(*UserConsentSpecifics_UnifiedConsent)(nil),
		(*UserConsentSpecifics_AssistantActivityControlConsent)(nil),
		(*UserConsentSpecifics_AccountPasswordsConsent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_consent_specifics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_consent_specifics_proto_goTypes,
		DependencyIndexes: file_user_consent_specifics_proto_depIdxs,
		EnumInfos:         file_user_consent_specifics_proto_enumTypes,
		MessageInfos:      file_user_consent_specifics_proto_msgTypes,
	}.Build()
	File_user_consent_specifics_proto = out.File
	file_user_consent_specifics_proto_rawDesc = nil
	file_user_consent_specifics_proto_goTypes = nil
	file_user_consent_specifics_proto_depIdxs = nil
}
