// Copyright 2021 Security Scorecard Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: policy.proto

package policy

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

// Mode definition.
type CheckPolicy_Mode int32

const (
	CheckPolicy_DISABLED CheckPolicy_Mode = 0
	CheckPolicy_ENFORCED CheckPolicy_Mode = 1
)

// Enum value maps for CheckPolicy_Mode.
var (
	CheckPolicy_Mode_name = map[int32]string{
		0: "DISABLED",
		1: "ENFORCED",
	}
	CheckPolicy_Mode_value = map[string]int32{
		"DISABLED": 0,
		"ENFORCED": 1,
	}
)

func (x CheckPolicy_Mode) Enum() *CheckPolicy_Mode {
	p := new(CheckPolicy_Mode)
	*p = x
	return p
}

func (x CheckPolicy_Mode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckPolicy_Mode) Descriptor() protoreflect.EnumDescriptor {
	return file_policy_proto_enumTypes[0].Descriptor()
}

func (CheckPolicy_Mode) Type() protoreflect.EnumType {
	return &file_policy_proto_enumTypes[0]
}

func (x CheckPolicy_Mode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckPolicy_Mode.Descriptor instead.
func (CheckPolicy_Mode) EnumDescriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0, 0}
}

type CheckPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode  CheckPolicy_Mode `protobuf:"varint,1,opt,name=mode,proto3,enum=ossf.scorecard.policy.CheckPolicy_Mode" json:"mode,omitempty"`
	Score int32            `protobuf:"zigzag32,2,opt,name=score,proto3" json:"score,omitempty"` // TODO: add Risk.
}

func (x *CheckPolicy) Reset() {
	*x = CheckPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPolicy) ProtoMessage() {}

func (x *CheckPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPolicy.ProtoReflect.Descriptor instead.
func (*CheckPolicy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPolicy) GetMode() CheckPolicy_Mode {
	if x != nil {
		return x.Mode
	}
	return CheckPolicy_DISABLED
}

func (x *CheckPolicy) GetScore() int32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type ScorecardPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  int32                   `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Policies map[string]*CheckPolicy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ScorecardPolicy) Reset() {
	*x = ScorecardPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_policy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScorecardPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScorecardPolicy) ProtoMessage() {}

func (x *ScorecardPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_policy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScorecardPolicy.ProtoReflect.Descriptor instead.
func (*ScorecardPolicy) Descriptor() ([]byte, []int) {
	return file_policy_proto_rawDescGZIP(), []int{1}
}

func (x *ScorecardPolicy) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ScorecardPolicy) GetPolicies() map[string]*CheckPolicy {
	if x != nil {
		return x.Policies
	}
	return nil
}

var File_policy_proto protoreflect.FileDescriptor

var file_policy_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x6f, 0x73, 0x73, 0x66, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x84, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3b, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6f, 0x73, 0x73, 0x66, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x11, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x22, 0x0a, 0x04, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x45, 0x4e, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x44, 0x10, 0x01, 0x22, 0xde, 0x01, 0x0a,
	0x0f, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x08, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6f,
	0x73, 0x73, 0x66, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x1a, 0x5f, 0x0a, 0x0d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6f, 0x73, 0x73, 0x66, 0x2e, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x73, 0x73, 0x66,
	0x2f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x63, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_policy_proto_rawDescOnce sync.Once
	file_policy_proto_rawDescData = file_policy_proto_rawDesc
)

func file_policy_proto_rawDescGZIP() []byte {
	file_policy_proto_rawDescOnce.Do(func() {
		file_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_policy_proto_rawDescData)
	})
	return file_policy_proto_rawDescData
}

var file_policy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_policy_proto_goTypes = []interface{}{
	(CheckPolicy_Mode)(0),   // 0: ossf.scorecard.policy.CheckPolicy.Mode
	(*CheckPolicy)(nil),     // 1: ossf.scorecard.policy.CheckPolicy
	(*ScorecardPolicy)(nil), // 2: ossf.scorecard.policy.ScorecardPolicy
	nil,                     // 3: ossf.scorecard.policy.ScorecardPolicy.PoliciesEntry
}
var file_policy_proto_depIdxs = []int32{
	0, // 0: ossf.scorecard.policy.CheckPolicy.mode:type_name -> ossf.scorecard.policy.CheckPolicy.Mode
	3, // 1: ossf.scorecard.policy.ScorecardPolicy.policies:type_name -> ossf.scorecard.policy.ScorecardPolicy.PoliciesEntry
	1, // 2: ossf.scorecard.policy.ScorecardPolicy.PoliciesEntry.value:type_name -> ossf.scorecard.policy.CheckPolicy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_policy_proto_init() }
func file_policy_proto_init() {
	if File_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPolicy); i {
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
		file_policy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScorecardPolicy); i {
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
			RawDescriptor: file_policy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_policy_proto_goTypes,
		DependencyIndexes: file_policy_proto_depIdxs,
		EnumInfos:         file_policy_proto_enumTypes,
		MessageInfos:      file_policy_proto_msgTypes,
	}.Build()
	File_policy_proto = out.File
	file_policy_proto_rawDesc = nil
	file_policy_proto_goTypes = nil
	file_policy_proto_depIdxs = nil
}
