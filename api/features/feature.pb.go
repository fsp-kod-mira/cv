// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0
// source: feature.proto

package features

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{0}
}

type FeatureStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PriorityId uint64 `protobuf:"varint,3,opt,name=priority_id,json=priorityId,proto3" json:"priority_id,omitempty"`
}

func (x *FeatureStruct) Reset() {
	*x = FeatureStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureStruct) ProtoMessage() {}

func (x *FeatureStruct) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureStruct.ProtoReflect.Descriptor instead.
func (*FeatureStruct) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{1}
}

func (x *FeatureStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FeatureStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureStruct) GetPriorityId() uint64 {
	if x != nil {
		return x.PriorityId
	}
	return 0
}

type PriorityStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Coefficient float32 `protobuf:"fixed32,3,opt,name=coefficient,proto3" json:"coefficient,omitempty"`
}

func (x *PriorityStruct) Reset() {
	*x = PriorityStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PriorityStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PriorityStruct) ProtoMessage() {}

func (x *PriorityStruct) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PriorityStruct.ProtoReflect.Descriptor instead.
func (*PriorityStruct) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{2}
}

func (x *PriorityStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PriorityStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PriorityStruct) GetCoefficient() float32 {
	if x != nil {
		return x.Coefficient
	}
	return 0
}

type HibrydFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureId    uint64  `protobuf:"varint,1,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	PriorityId   uint64  `protobuf:"varint,2,opt,name=priority_id,json=priorityId,proto3" json:"priority_id,omitempty"`
	FeatureName  string  `protobuf:"bytes,3,opt,name=feature_name,json=featureName,proto3" json:"feature_name,omitempty"`
	PriorityName string  `protobuf:"bytes,4,opt,name=priority_name,json=priorityName,proto3" json:"priority_name,omitempty"`
	Coefficient  float32 `protobuf:"fixed32,5,opt,name=coefficient,proto3" json:"coefficient,omitempty"`
}

func (x *HibrydFeature) Reset() {
	*x = HibrydFeature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HibrydFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HibrydFeature) ProtoMessage() {}

func (x *HibrydFeature) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HibrydFeature.ProtoReflect.Descriptor instead.
func (*HibrydFeature) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{3}
}

func (x *HibrydFeature) GetFeatureId() uint64 {
	if x != nil {
		return x.FeatureId
	}
	return 0
}

func (x *HibrydFeature) GetPriorityId() uint64 {
	if x != nil {
		return x.PriorityId
	}
	return 0
}

func (x *HibrydFeature) GetFeatureName() string {
	if x != nil {
		return x.FeatureName
	}
	return ""
}

func (x *HibrydFeature) GetPriorityName() string {
	if x != nil {
		return x.PriorityName
	}
	return ""
}

func (x *HibrydFeature) GetCoefficient() float32 {
	if x != nil {
		return x.Coefficient
	}
	return 0
}

type HibrydFeatureList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*HibrydFeature `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *HibrydFeatureList) Reset() {
	*x = HibrydFeatureList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HibrydFeatureList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HibrydFeatureList) ProtoMessage() {}

func (x *HibrydFeatureList) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HibrydFeatureList.ProtoReflect.Descriptor instead.
func (*HibrydFeatureList) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{4}
}

func (x *HibrydFeatureList) GetItems() []*HibrydFeature {
	if x != nil {
		return x.Items
	}
	return nil
}

type IdStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdStruct) Reset() {
	*x = IdStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdStruct) ProtoMessage() {}

func (x *IdStruct) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdStruct.ProtoReflect.Descriptor instead.
func (*IdStruct) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{5}
}

func (x *IdStruct) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type NameStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NameStruct) Reset() {
	*x = NameStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_feature_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameStruct) ProtoMessage() {}

func (x *NameStruct) ProtoReflect() protoreflect.Message {
	mi := &file_feature_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameStruct.ProtoReflect.Descriptor instead.
func (*NameStruct) Descriptor() ([]byte, []int) {
	return file_feature_proto_rawDescGZIP(), []int{6}
}

func (x *NameStruct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_feature_proto protoreflect.FileDescriptor

var file_feature_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x54, 0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x22, 0x56,
	0x0a, 0x0e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66,
	0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x22, 0xb9, 0x01, 0x0a, 0x0d, 0x48, 0x69, 0x62, 0x72, 0x79,
	0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72,
	0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x65, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0x48, 0x0a, 0x11, 0x48, 0x69, 0x62, 0x72, 0x79, 0x64, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x62, 0x72, 0x79, 0x64, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x1a, 0x0a, 0x08,
	0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x20, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x89, 0x05, 0x0a, 0x07, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x50, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x18, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x1d, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x1d,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12,
	0x40, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x15, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x18, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x15, 0x2e, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x18, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x15, 0x2e, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x18, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x15,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x15,
	0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x62, 0x72, 0x79, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49,
	0x64, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x1d, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x48, 0x69, 0x62, 0x72, 0x79, 0x64, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x4e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x2e, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x1a, 0x1d, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x0e, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_feature_proto_rawDescOnce sync.Once
	file_feature_proto_rawDescData = file_feature_proto_rawDesc
)

func file_feature_proto_rawDescGZIP() []byte {
	file_feature_proto_rawDescOnce.Do(func() {
		file_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_feature_proto_rawDescData)
	})
	return file_feature_proto_rawDescData
}

var file_feature_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_feature_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: FeatureService.Empty
	(*FeatureStruct)(nil),     // 1: FeatureService.FeatureStruct
	(*PriorityStruct)(nil),    // 2: FeatureService.PriorityStruct
	(*HibrydFeature)(nil),     // 3: FeatureService.HibrydFeature
	(*HibrydFeatureList)(nil), // 4: FeatureService.HibrydFeatureList
	(*IdStruct)(nil),          // 5: FeatureService.IdStruct
	(*NameStruct)(nil),        // 6: FeatureService.NameStruct
}
var file_feature_proto_depIdxs = []int32{
	3,  // 0: FeatureService.HibrydFeatureList.items:type_name -> FeatureService.HibrydFeature
	2,  // 1: FeatureService.Feature.AddPriority:input_type -> FeatureService.PriorityStruct
	1,  // 2: FeatureService.Feature.AddFeature:input_type -> FeatureService.FeatureStruct
	5,  // 3: FeatureService.Feature.DelPriority:input_type -> FeatureService.IdStruct
	5,  // 4: FeatureService.Feature.DelFeature:input_type -> FeatureService.IdStruct
	5,  // 5: FeatureService.Feature.EditPriority:input_type -> FeatureService.IdStruct
	5,  // 6: FeatureService.Feature.EditFeature:input_type -> FeatureService.IdStruct
	0,  // 7: FeatureService.Feature.Get:input_type -> FeatureService.Empty
	5,  // 8: FeatureService.Feature.GetFeaturesById:input_type -> FeatureService.IdStruct
	6,  // 9: FeatureService.Feature.GetFeaturesByName:input_type -> FeatureService.NameStruct
	5,  // 10: FeatureService.Feature.AddPriority:output_type -> FeatureService.IdStruct
	1,  // 11: FeatureService.Feature.AddFeature:output_type -> FeatureService.FeatureStruct
	0,  // 12: FeatureService.Feature.DelPriority:output_type -> FeatureService.Empty
	0,  // 13: FeatureService.Feature.DelFeature:output_type -> FeatureService.Empty
	0,  // 14: FeatureService.Feature.EditPriority:output_type -> FeatureService.Empty
	0,  // 15: FeatureService.Feature.EditFeature:output_type -> FeatureService.Empty
	4,  // 16: FeatureService.Feature.Get:output_type -> FeatureService.HibrydFeatureList
	3,  // 17: FeatureService.Feature.GetFeaturesById:output_type -> FeatureService.HibrydFeature
	1,  // 18: FeatureService.Feature.GetFeaturesByName:output_type -> FeatureService.FeatureStruct
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_feature_proto_init() }
func file_feature_proto_init() {
	if File_feature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_feature_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureStruct); i {
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
		file_feature_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PriorityStruct); i {
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
		file_feature_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HibrydFeature); i {
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
		file_feature_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HibrydFeatureList); i {
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
		file_feature_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdStruct); i {
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
		file_feature_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameStruct); i {
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
			RawDescriptor: file_feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_feature_proto_goTypes,
		DependencyIndexes: file_feature_proto_depIdxs,
		MessageInfos:      file_feature_proto_msgTypes,
	}.Build()
	File_feature_proto = out.File
	file_feature_proto_rawDesc = nil
	file_feature_proto_goTypes = nil
	file_feature_proto_depIdxs = nil
}
