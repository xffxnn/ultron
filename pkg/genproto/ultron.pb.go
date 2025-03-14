// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: ultron.proto

package genproto

import (
	statistics "github.com/wosai/ultron/v2/pkg/statistics"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventType int32

const (
	EventType_UNKNOWN            EventType = 0
	EventType_PING               EventType = 1
	EventType_CONNECTED          EventType = 2 //已连接
	EventType_DISCONNECT         EventType = 3 // 要求slave断开连接
	EventType_PLAN_STARTED       EventType = 4 // 测试计划开始
	EventType_PLAN_FINISHED      EventType = 5 // 测试计划结束
	EventType_PLAN_INTERRUPTED   EventType = 6 // 测试计划中断执行
	EventType_NEXT_STAGE_STARTED EventType = 7 // 开始执行计划中的下一阶段
	EventType_STATS_AGGREGATE    EventType = 8 // 上报统计对象
	EventType_STATUS_REPORT      EventType = 9 // 上报运行状态
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "UNKNOWN",
		1: "PING",
		2: "CONNECTED",
		3: "DISCONNECT",
		4: "PLAN_STARTED",
		5: "PLAN_FINISHED",
		6: "PLAN_INTERRUPTED",
		7: "NEXT_STAGE_STARTED",
		8: "STATS_AGGREGATE",
		9: "STATUS_REPORT",
	}
	EventType_value = map[string]int32{
		"UNKNOWN":            0,
		"PING":               1,
		"CONNECTED":          2,
		"DISCONNECT":         3,
		"PLAN_STARTED":       4,
		"PLAN_FINISHED":      5,
		"PLAN_INTERRUPTED":   6,
		"NEXT_STAGE_STARTED": 7,
		"STATS_AGGREGATE":    8,
		"STATUS_REPORT":      9,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_ultron_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_ultron_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{0}
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlaveId string            `protobuf:"bytes,1,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	Extras  map[string]string `protobuf:"bytes,2,rep,name=extras,proto3" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribeRequest) GetSlaveId() string {
	if x != nil {
		return x.SlaveId
	}
	return ""
}

func (x *SubscribeRequest) GetExtras() map[string]string {
	if x != nil {
		return x.Extras
	}
	return nil
}

type TimerDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Timer []byte `protobuf:"bytes,2,opt,name=timer,proto3" json:"timer,omitempty"`
}

func (x *TimerDTO) Reset() {
	*x = TimerDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimerDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimerDTO) ProtoMessage() {}

func (x *TimerDTO) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimerDTO.ProtoReflect.Descriptor instead.
func (*TimerDTO) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{1}
}

func (x *TimerDTO) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *TimerDTO) GetTimer() []byte {
	if x != nil {
		return x.Timer
	}
	return nil
}

type AttackStrategyDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type           string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	AttackStrategy []byte `protobuf:"bytes,2,opt,name=attack_strategy,json=attackStrategy,proto3" json:"attack_strategy,omitempty"`
}

func (x *AttackStrategyDTO) Reset() {
	*x = AttackStrategyDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackStrategyDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackStrategyDTO) ProtoMessage() {}

func (x *AttackStrategyDTO) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackStrategyDTO.ProtoReflect.Descriptor instead.
func (*AttackStrategyDTO) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{2}
}

func (x *AttackStrategyDTO) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AttackStrategyDTO) GetAttackStrategy() []byte {
	if x != nil {
		return x.AttackStrategy
	}
	return nil
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=wosai.ultron.EventType" json:"type,omitempty"`
	// Types that are assignable to Data:
	//	*SubscribeResponse_PlanName
	//	*SubscribeResponse_AttackStrategy
	//	*SubscribeResponse_BatchId
	Data  isSubscribeResponse_Data `protobuf_oneof:"data"`
	Timer *TimerDTO                `protobuf:"bytes,5,opt,name=timer,proto3" json:"timer,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeResponse) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_UNKNOWN
}

func (m *SubscribeResponse) GetData() isSubscribeResponse_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SubscribeResponse) GetPlanName() string {
	if x, ok := x.GetData().(*SubscribeResponse_PlanName); ok {
		return x.PlanName
	}
	return ""
}

func (x *SubscribeResponse) GetAttackStrategy() *AttackStrategyDTO {
	if x, ok := x.GetData().(*SubscribeResponse_AttackStrategy); ok {
		return x.AttackStrategy
	}
	return nil
}

func (x *SubscribeResponse) GetBatchId() uint32 {
	if x, ok := x.GetData().(*SubscribeResponse_BatchId); ok {
		return x.BatchId
	}
	return 0
}

func (x *SubscribeResponse) GetTimer() *TimerDTO {
	if x != nil {
		return x.Timer
	}
	return nil
}

type isSubscribeResponse_Data interface {
	isSubscribeResponse_Data()
}

type SubscribeResponse_PlanName struct {
	PlanName string `protobuf:"bytes,2,opt,name=plan_name,json=planName,proto3,oneof"`
}

type SubscribeResponse_AttackStrategy struct {
	AttackStrategy *AttackStrategyDTO `protobuf:"bytes,3,opt,name=attack_strategy,json=attackStrategy,proto3,oneof"`
}

type SubscribeResponse_BatchId struct {
	BatchId uint32 `protobuf:"varint,4,opt,name=batch_id,json=batchId,proto3,oneof"`
}

func (*SubscribeResponse_PlanName) isSubscribeResponse_Data() {}

func (*SubscribeResponse_AttackStrategy) isSubscribeResponse_Data() {}

func (*SubscribeResponse_BatchId) isSubscribeResponse_Data() {}

type SubmitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlaveId string                           `protobuf:"bytes,1,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	BatchId uint32                           `protobuf:"varint,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	Stats   *statistics.StatisticianGroupDTO `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *SubmitRequest) Reset() {
	*x = SubmitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitRequest) ProtoMessage() {}

func (x *SubmitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitRequest.ProtoReflect.Descriptor instead.
func (*SubmitRequest) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{4}
}

func (x *SubmitRequest) GetSlaveId() string {
	if x != nil {
		return x.SlaveId
	}
	return ""
}

func (x *SubmitRequest) GetBatchId() uint32 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *SubmitRequest) GetStats() *statistics.StatisticianGroupDTO {
	if x != nil {
		return x.Stats
	}
	return nil
}

type SendStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlaveId         string `protobuf:"bytes,1,opt,name=slave_id,json=slaveId,proto3" json:"slave_id,omitempty"`
	ConcurrentUsers int32  `protobuf:"varint,2,opt,name=concurrent_users,json=concurrentUsers,proto3" json:"concurrent_users,omitempty"`
}

func (x *SendStatusRequest) Reset() {
	*x = SendStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ultron_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStatusRequest) ProtoMessage() {}

func (x *SendStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ultron_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStatusRequest.ProtoReflect.Descriptor instead.
func (*SendStatusRequest) Descriptor() ([]byte, []int) {
	return file_ultron_proto_rawDescGZIP(), []int{5}
}

func (x *SendStatusRequest) GetSlaveId() string {
	if x != nil {
		return x.SlaveId
	}
	return ""
}

func (x *SendStatusRequest) GetConcurrentUsers() int32 {
	if x != nil {
		return x.ConcurrentUsers
	}
	return 0
}

var File_ultron_proto protoreflect.FileDescriptor

var file_ultron_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x77, 0x6f, 0x73, 0x61, 0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x1a, 0x10, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a, 0x10,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x06, 0x65,
	0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x6f,
	0x73, 0x61, 0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x34, 0x0a, 0x08, 0x54, 0x69,
	0x6d, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x6d, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72,
	0x22, 0x50, 0x0a, 0x11, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x44, 0x54, 0x4f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x22, 0xfe, 0x01, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x77, 0x6f, 0x73, 0x61, 0x69, 0x2e, 0x75,
	0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x77, 0x6f, 0x73, 0x61, 0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74,
	0x61, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x54, 0x4f, 0x48, 0x00,
	0x52, 0x0e, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x1b, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x00, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77,
	0x6f, 0x73, 0x61, 0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x72, 0x44, 0x54, 0x4f, 0x52, 0x05, 0x74, 0x69, 0x6d, 0x65, 0x72, 0x42, 0x06, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x7f, 0x0a, 0x0d, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x77, 0x6f, 0x73, 0x61,
	0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x69, 0x61, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x54, 0x4f, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6c, 0x61,
	0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6c, 0x61,
	0x76, 0x65, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x2a,
	0xbc, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43,
	0x54, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x4c, 0x41, 0x4e, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x41, 0x4e,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x52, 0x55, 0x50, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x16,
	0x0a, 0x12, 0x4e, 0x45, 0x58, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x52, 0x54, 0x45, 0x44, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x53, 0x5f,
	0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x45, 0x10, 0x08, 0x12, 0x11, 0x0a, 0x0d, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x09, 0x32, 0xe7,
	0x01, 0x0a, 0x09, 0x55, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x41, 0x50, 0x49, 0x12, 0x50, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x1e, 0x2e, 0x77, 0x6f, 0x73, 0x61,
	0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x77, 0x6f, 0x73, 0x61,
	0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3f,
	0x0a, 0x06, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x2e, 0x77, 0x6f, 0x73, 0x61, 0x69,
	0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x2e,
	0x77, 0x6f, 0x73, 0x61, 0x69, 0x2e, 0x75, 0x6c, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x6f, 0x73, 0x61, 0x69, 0x2f, 0x75, 0x6c, 0x74,
	0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ultron_proto_rawDescOnce sync.Once
	file_ultron_proto_rawDescData = file_ultron_proto_rawDesc
)

func file_ultron_proto_rawDescGZIP() []byte {
	file_ultron_proto_rawDescOnce.Do(func() {
		file_ultron_proto_rawDescData = protoimpl.X.CompressGZIP(file_ultron_proto_rawDescData)
	})
	return file_ultron_proto_rawDescData
}

var file_ultron_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ultron_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ultron_proto_goTypes = []interface{}{
	(EventType)(0),                          // 0: wosai.ultron.EventType
	(*SubscribeRequest)(nil),                // 1: wosai.ultron.SubscribeRequest
	(*TimerDTO)(nil),                        // 2: wosai.ultron.TimerDTO
	(*AttackStrategyDTO)(nil),               // 3: wosai.ultron.AttackStrategyDTO
	(*SubscribeResponse)(nil),               // 4: wosai.ultron.SubscribeResponse
	(*SubmitRequest)(nil),                   // 5: wosai.ultron.SubmitRequest
	(*SendStatusRequest)(nil),               // 6: wosai.ultron.SendStatusRequest
	nil,                                     // 7: wosai.ultron.SubscribeRequest.ExtrasEntry
	(*statistics.StatisticianGroupDTO)(nil), // 8: wosai.ultron.StatisticianGroupDTO
	(*emptypb.Empty)(nil),                   // 9: google.protobuf.Empty
}
var file_ultron_proto_depIdxs = []int32{
	7, // 0: wosai.ultron.SubscribeRequest.extras:type_name -> wosai.ultron.SubscribeRequest.ExtrasEntry
	0, // 1: wosai.ultron.SubscribeResponse.type:type_name -> wosai.ultron.EventType
	3, // 2: wosai.ultron.SubscribeResponse.attack_strategy:type_name -> wosai.ultron.AttackStrategyDTO
	2, // 3: wosai.ultron.SubscribeResponse.timer:type_name -> wosai.ultron.TimerDTO
	8, // 4: wosai.ultron.SubmitRequest.stats:type_name -> wosai.ultron.StatisticianGroupDTO
	1, // 5: wosai.ultron.UltronAPI.Subscribe:input_type -> wosai.ultron.SubscribeRequest
	5, // 6: wosai.ultron.UltronAPI.Submit:input_type -> wosai.ultron.SubmitRequest
	6, // 7: wosai.ultron.UltronAPI.SendStatus:input_type -> wosai.ultron.SendStatusRequest
	4, // 8: wosai.ultron.UltronAPI.Subscribe:output_type -> wosai.ultron.SubscribeResponse
	9, // 9: wosai.ultron.UltronAPI.Submit:output_type -> google.protobuf.Empty
	9, // 10: wosai.ultron.UltronAPI.SendStatus:output_type -> google.protobuf.Empty
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ultron_proto_init() }
func file_ultron_proto_init() {
	if File_ultron_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ultron_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_ultron_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimerDTO); i {
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
		file_ultron_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackStrategyDTO); i {
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
		file_ultron_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_ultron_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitRequest); i {
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
		file_ultron_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStatusRequest); i {
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
	file_ultron_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SubscribeResponse_PlanName)(nil),
		(*SubscribeResponse_AttackStrategy)(nil),
		(*SubscribeResponse_BatchId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ultron_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ultron_proto_goTypes,
		DependencyIndexes: file_ultron_proto_depIdxs,
		EnumInfos:         file_ultron_proto_enumTypes,
		MessageInfos:      file_ultron_proto_msgTypes,
	}.Build()
	File_ultron_proto = out.File
	file_ultron_proto_rawDesc = nil
	file_ultron_proto_goTypes = nil
	file_ultron_proto_depIdxs = nil
}
