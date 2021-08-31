// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.2
// source: ai/traceable/agent/config/v1/config.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Parent configuration object for Opa & Blocking Config
type AgentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opa            *Opa            `protobuf:"bytes,1,opt,name=opa,proto3" json:"opa,omitempty"`
	BlockingConfig *BlockingConfig `protobuf:"bytes,2,opt,name=blocking_config,json=blockingConfig,proto3" json:"blocking_config,omitempty"`
}

func (x *AgentConfig) Reset() {
	*x = AgentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentConfig) ProtoMessage() {}

func (x *AgentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentConfig.ProtoReflect.Descriptor instead.
func (*AgentConfig) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *AgentConfig) GetOpa() *Opa {
	if x != nil {
		return x.Opa
	}
	return nil
}

func (x *AgentConfig) GetBlockingConfig() *BlockingConfig {
	if x != nil {
		return x.BlockingConfig
	}
	return nil
}

// Opa covers the options related to the mechanics for getting Open Policy Agent configuration file.
// The client should use secure and token option from reporting settings.
type Opa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// when `true` Open Policy Agent evaluation is enabled to block request
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// endpoint represents the endpoint for polling OPA config file e.g. http://opa.traceableai:8181/
	Endpoint *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// poll period in seconds to query OPA service
	PollPeriodSeconds *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=poll_period_seconds,json=pollPeriodSeconds,proto3" json:"poll_period_seconds,omitempty"`
}

func (x *Opa) Reset() {
	*x = Opa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Opa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Opa) ProtoMessage() {}

func (x *Opa) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Opa.ProtoReflect.Descriptor instead.
func (*Opa) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *Opa) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Opa) GetEndpoint() *wrapperspb.StringValue {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *Opa) GetPollPeriodSeconds() *wrapperspb.Int32Value {
	if x != nil {
		return x.PollPeriodSeconds
	}
	return nil
}

// Blocking config
type BlockingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled        *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	DebugLog       *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=debug_log,json=debugLog,proto3" json:"debug_log,omitempty"`
	Modsecurity    *ModsecurityConfig    `protobuf:"bytes,3,opt,name=modsecurity,proto3" json:"modsecurity,omitempty"`
	EvaluateBody   *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=evaluate_body,json=evaluateBody,proto3" json:"evaluate_body,omitempty"`
	RegionBlocking *RegionBlockingConfig `protobuf:"bytes,5,opt,name=region_blocking,json=regionBlocking,proto3" json:"region_blocking,omitempty"`
	RemoteConfig   *RemoteConfig         `protobuf:"bytes,6,opt,name=remote_config,json=remoteConfig,proto3" json:"remote_config,omitempty"`
}

func (x *BlockingConfig) Reset() {
	*x = BlockingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockingConfig) ProtoMessage() {}

func (x *BlockingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockingConfig.ProtoReflect.Descriptor instead.
func (*BlockingConfig) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *BlockingConfig) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *BlockingConfig) GetDebugLog() *wrapperspb.BoolValue {
	if x != nil {
		return x.DebugLog
	}
	return nil
}

func (x *BlockingConfig) GetModsecurity() *ModsecurityConfig {
	if x != nil {
		return x.Modsecurity
	}
	return nil
}

func (x *BlockingConfig) GetEvaluateBody() *wrapperspb.BoolValue {
	if x != nil {
		return x.EvaluateBody
	}
	return nil
}

func (x *BlockingConfig) GetRegionBlocking() *RegionBlockingConfig {
	if x != nil {
		return x.RegionBlocking
	}
	return nil
}

func (x *BlockingConfig) GetRemoteConfig() *RemoteConfig {
	if x != nil {
		return x.RemoteConfig
	}
	return nil
}

type ModsecurityConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *ModsecurityConfig) Reset() {
	*x = ModsecurityConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModsecurityConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModsecurityConfig) ProtoMessage() {}

func (x *ModsecurityConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModsecurityConfig.ProtoReflect.Descriptor instead.
func (*ModsecurityConfig) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *ModsecurityConfig) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

type RegionBlockingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *RegionBlockingConfig) Reset() {
	*x = RegionBlockingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionBlockingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionBlockingConfig) ProtoMessage() {}

func (x *RegionBlockingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionBlockingConfig.ProtoReflect.Descriptor instead.
func (*RegionBlockingConfig) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *RegionBlockingConfig) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

// RemoteConfig defines the remote endpoint where the config is fetched from
type RemoteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// enabled denotes if config needs to be fetched from remote or not
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// endpoint denotes the agentmanager endpoint to connect to for config. eg: localhost:5441
	Endpoint *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// poll period in seconds to query for config updates
	PollPeriodSeconds *wrapperspb.Int32Value `protobuf:"bytes,3,opt,name=poll_period_seconds,json=pollPeriodSeconds,proto3" json:"poll_period_seconds,omitempty"`
}

func (x *RemoteConfig) Reset() {
	*x = RemoteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig) ProtoMessage() {}

func (x *RemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ai_traceable_agent_config_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig.ProtoReflect.Descriptor instead.
func (*RemoteConfig) Descriptor() ([]byte, []int) {
	return file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *RemoteConfig) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *RemoteConfig) GetEndpoint() *wrapperspb.StringValue {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *RemoteConfig) GetPollPeriodSeconds() *wrapperspb.Int32Value {
	if x != nil {
		return x.PollPeriodSeconds
	}
	return nil
}

var File_ai_traceable_agent_config_v1_config_proto protoreflect.FileDescriptor

var file_ai_traceable_agent_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x29, 0x61, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x61, 0x69, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x0b, 0x41, 0x67,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x0a, 0x03, 0x6f, 0x70, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x61, 0x52, 0x03, 0x6f, 0x70, 0x61, 0x12, 0x55,
	0x0a, 0x0f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xc2, 0x01, 0x0a, 0x03, 0x4f, 0x70, 0x61, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4b, 0x0a,
	0x13, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x70, 0x6f, 0x6c, 0x6c, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x22, 0xc1, 0x03, 0x0a, 0x0e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x09, 0x64, 0x65, 0x62, 0x75, 0x67, 0x5f, 0x6c, 0x6f, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x64, 0x65, 0x62, 0x75, 0x67, 0x4c, 0x6f, 0x67, 0x12, 0x51, 0x0a, 0x0b,
	0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x0b, 0x6d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x3f, 0x0a, 0x0d, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0c, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x5b, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x61, 0x69, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x4f, 0x0a,
	0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x49,
	0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x4c, 0x0a, 0x14, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0xcb, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x38,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x4b, 0x0a, 0x13, 0x70, 0x6f, 0x6c, 0x6c,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x11, 0x70, 0x6f, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x72, 0x61, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x61, 0x69, 0x2f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ai_traceable_agent_config_v1_config_proto_rawDescOnce sync.Once
	file_ai_traceable_agent_config_v1_config_proto_rawDescData = file_ai_traceable_agent_config_v1_config_proto_rawDesc
)

func file_ai_traceable_agent_config_v1_config_proto_rawDescGZIP() []byte {
	file_ai_traceable_agent_config_v1_config_proto_rawDescOnce.Do(func() {
		file_ai_traceable_agent_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_ai_traceable_agent_config_v1_config_proto_rawDescData)
	})
	return file_ai_traceable_agent_config_v1_config_proto_rawDescData
}

var file_ai_traceable_agent_config_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ai_traceable_agent_config_v1_config_proto_goTypes = []interface{}{
	(*AgentConfig)(nil),            // 0: ai.traceable.agent.config.v1.AgentConfig
	(*Opa)(nil),                    // 1: ai.traceable.agent.config.v1.Opa
	(*BlockingConfig)(nil),         // 2: ai.traceable.agent.config.v1.BlockingConfig
	(*ModsecurityConfig)(nil),      // 3: ai.traceable.agent.config.v1.ModsecurityConfig
	(*RegionBlockingConfig)(nil),   // 4: ai.traceable.agent.config.v1.RegionBlockingConfig
	(*RemoteConfig)(nil),           // 5: ai.traceable.agent.config.v1.RemoteConfig
	(*wrapperspb.BoolValue)(nil),   // 6: google.protobuf.BoolValue
	(*wrapperspb.StringValue)(nil), // 7: google.protobuf.StringValue
	(*wrapperspb.Int32Value)(nil),  // 8: google.protobuf.Int32Value
}
var file_ai_traceable_agent_config_v1_config_proto_depIdxs = []int32{
	1,  // 0: ai.traceable.agent.config.v1.AgentConfig.opa:type_name -> ai.traceable.agent.config.v1.Opa
	2,  // 1: ai.traceable.agent.config.v1.AgentConfig.blocking_config:type_name -> ai.traceable.agent.config.v1.BlockingConfig
	6,  // 2: ai.traceable.agent.config.v1.Opa.enabled:type_name -> google.protobuf.BoolValue
	7,  // 3: ai.traceable.agent.config.v1.Opa.endpoint:type_name -> google.protobuf.StringValue
	8,  // 4: ai.traceable.agent.config.v1.Opa.poll_period_seconds:type_name -> google.protobuf.Int32Value
	6,  // 5: ai.traceable.agent.config.v1.BlockingConfig.enabled:type_name -> google.protobuf.BoolValue
	6,  // 6: ai.traceable.agent.config.v1.BlockingConfig.debug_log:type_name -> google.protobuf.BoolValue
	3,  // 7: ai.traceable.agent.config.v1.BlockingConfig.modsecurity:type_name -> ai.traceable.agent.config.v1.ModsecurityConfig
	6,  // 8: ai.traceable.agent.config.v1.BlockingConfig.evaluate_body:type_name -> google.protobuf.BoolValue
	4,  // 9: ai.traceable.agent.config.v1.BlockingConfig.region_blocking:type_name -> ai.traceable.agent.config.v1.RegionBlockingConfig
	5,  // 10: ai.traceable.agent.config.v1.BlockingConfig.remote_config:type_name -> ai.traceable.agent.config.v1.RemoteConfig
	6,  // 11: ai.traceable.agent.config.v1.ModsecurityConfig.enabled:type_name -> google.protobuf.BoolValue
	6,  // 12: ai.traceable.agent.config.v1.RegionBlockingConfig.enabled:type_name -> google.protobuf.BoolValue
	6,  // 13: ai.traceable.agent.config.v1.RemoteConfig.enabled:type_name -> google.protobuf.BoolValue
	7,  // 14: ai.traceable.agent.config.v1.RemoteConfig.endpoint:type_name -> google.protobuf.StringValue
	8,  // 15: ai.traceable.agent.config.v1.RemoteConfig.poll_period_seconds:type_name -> google.protobuf.Int32Value
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_ai_traceable_agent_config_v1_config_proto_init() }
func file_ai_traceable_agent_config_v1_config_proto_init() {
	if File_ai_traceable_agent_config_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentConfig); i {
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
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Opa); i {
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
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockingConfig); i {
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
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModsecurityConfig); i {
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
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionBlockingConfig); i {
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
		file_ai_traceable_agent_config_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig); i {
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
			RawDescriptor: file_ai_traceable_agent_config_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ai_traceable_agent_config_v1_config_proto_goTypes,
		DependencyIndexes: file_ai_traceable_agent_config_v1_config_proto_depIdxs,
		MessageInfos:      file_ai_traceable_agent_config_v1_config_proto_msgTypes,
	}.Build()
	File_ai_traceable_agent_config_v1_config_proto = out.File
	file_ai_traceable_agent_config_v1_config_proto_rawDesc = nil
	file_ai_traceable_agent_config_v1_config_proto_goTypes = nil
	file_ai_traceable_agent_config_v1_config_proto_depIdxs = nil
}
