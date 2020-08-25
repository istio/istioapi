// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: meta/v1alpha1/status.proto

// $title: Istio Status
// $description: Common status field for all istio collections.
// $location: https://istio.io/docs/reference/config/meta/v1beta1/istio-status.html

package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1alpha1 "istio.io/api/analysis/v1alpha1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IstioStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current service state of pod.
	// More info: https://istio.io/docs/reference/config/config-status/
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	Conditions []*IstioCondition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
	// Includes any errors or warnings detected by Istio's analyzers.
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	ValidationMessages []*v1alpha1.AnalysisMessageBase `protobuf:"bytes,2,rep,name=validation_messages,json=validationMessages,proto3" json:"validation_messages,omitempty"`
}

func (x *IstioStatus) Reset() {
	*x = IstioStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_v1alpha1_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioStatus) ProtoMessage() {}

func (x *IstioStatus) ProtoReflect() protoreflect.Message {
	mi := &file_meta_v1alpha1_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioStatus.ProtoReflect.Descriptor instead.
func (*IstioStatus) Descriptor() ([]byte, []int) {
	return file_meta_v1alpha1_status_proto_rawDescGZIP(), []int{0}
}

func (x *IstioStatus) GetConditions() []*IstioCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *IstioStatus) GetValidationMessages() []*v1alpha1.AnalysisMessageBase {
	if x != nil {
		return x.ValidationMessages
	}
	return nil
}

type IstioCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type is the type of the condition.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=last_probe_time,json=lastProbeTime,proto3" json:"last_probe_time,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_transition_time,json=lastTransitionTime,proto3" json:"last_transition_time,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `protobuf:"bytes,5,opt,name=reason,proto3" json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IstioCondition) Reset() {
	*x = IstioCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_v1alpha1_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IstioCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IstioCondition) ProtoMessage() {}

func (x *IstioCondition) ProtoReflect() protoreflect.Message {
	mi := &file_meta_v1alpha1_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IstioCondition.ProtoReflect.Descriptor instead.
func (*IstioCondition) Descriptor() ([]byte, []int) {
	return file_meta_v1alpha1_status_proto_rawDescGZIP(), []int{1}
}

func (x *IstioCondition) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *IstioCondition) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *IstioCondition) GetLastProbeTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastProbeTime
	}
	return nil
}

func (x *IstioCondition) GetLastTransitionTime() *timestamp.Timestamp {
	if x != nil {
		return x.LastTransitionTime
	}
	return nil
}

func (x *IstioCondition) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *IstioCondition) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_meta_v1alpha1_status_proto protoreflect.FileDescriptor

var file_meta_v1alpha1_status_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f,
	0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5d, 0x0a, 0x13, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x80, 0x02, 0x0a, 0x0e, 0x49, 0x73, 0x74,
	0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x62, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x69,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_meta_v1alpha1_status_proto_rawDescOnce sync.Once
	file_meta_v1alpha1_status_proto_rawDescData = file_meta_v1alpha1_status_proto_rawDesc
)

func file_meta_v1alpha1_status_proto_rawDescGZIP() []byte {
	file_meta_v1alpha1_status_proto_rawDescOnce.Do(func() {
		file_meta_v1alpha1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_v1alpha1_status_proto_rawDescData)
	})
	return file_meta_v1alpha1_status_proto_rawDescData
}

var file_meta_v1alpha1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_meta_v1alpha1_status_proto_goTypes = []interface{}{
	(*IstioStatus)(nil),                  // 0: istio.meta.v1alpha1.IstioStatus
	(*IstioCondition)(nil),               // 1: istio.meta.v1alpha1.IstioCondition
	(*v1alpha1.AnalysisMessageBase)(nil), // 2: istio.analysis.v1alpha1.AnalysisMessageBase
	(*timestamp.Timestamp)(nil),          // 3: google.protobuf.Timestamp
}
var file_meta_v1alpha1_status_proto_depIdxs = []int32{
	1, // 0: istio.meta.v1alpha1.IstioStatus.conditions:type_name -> istio.meta.v1alpha1.IstioCondition
	2, // 1: istio.meta.v1alpha1.IstioStatus.validation_messages:type_name -> istio.analysis.v1alpha1.AnalysisMessageBase
	3, // 2: istio.meta.v1alpha1.IstioCondition.last_probe_time:type_name -> google.protobuf.Timestamp
	3, // 3: istio.meta.v1alpha1.IstioCondition.last_transition_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_meta_v1alpha1_status_proto_init() }
func file_meta_v1alpha1_status_proto_init() {
	if File_meta_v1alpha1_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_v1alpha1_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioStatus); i {
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
		file_meta_v1alpha1_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IstioCondition); i {
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
			RawDescriptor: file_meta_v1alpha1_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_v1alpha1_status_proto_goTypes,
		DependencyIndexes: file_meta_v1alpha1_status_proto_depIdxs,
		MessageInfos:      file_meta_v1alpha1_status_proto_msgTypes,
	}.Build()
	File_meta_v1alpha1_status_proto = out.File
	file_meta_v1alpha1_status_proto_rawDesc = nil
	file_meta_v1alpha1_status_proto_goTypes = nil
	file_meta_v1alpha1_status_proto_depIdxs = nil
}
