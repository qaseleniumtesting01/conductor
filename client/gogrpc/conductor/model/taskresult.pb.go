// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/taskresult.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TaskResult_Status int32

const (
	TaskResult_IN_PROGRESS TaskResult_Status = 0
	TaskResult_FAILED      TaskResult_Status = 1
	TaskResult_COMPLETED   TaskResult_Status = 2
	TaskResult_SCHEDULED   TaskResult_Status = 3
)

var TaskResult_Status_name = map[int32]string{
	0: "IN_PROGRESS",
	1: "FAILED",
	2: "COMPLETED",
	3: "SCHEDULED",
}
var TaskResult_Status_value = map[string]int32{
	"IN_PROGRESS": 0,
	"FAILED":      1,
	"COMPLETED":   2,
	"SCHEDULED":   3,
}

func (x TaskResult_Status) String() string {
	return proto.EnumName(TaskResult_Status_name, int32(x))
}
func (TaskResult_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_taskresult_544d5e3612411ce1, []int{0, 0}
}

type TaskResult struct {
	WorkflowInstanceId    string                    `protobuf:"bytes,1,opt,name=workflow_instance_id,json=workflowInstanceId" json:"workflow_instance_id,omitempty"`
	TaskId                string                    `protobuf:"bytes,2,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	ReasonForIncompletion string                    `protobuf:"bytes,3,opt,name=reason_for_incompletion,json=reasonForIncompletion" json:"reason_for_incompletion,omitempty"`
	CallbackAfterSeconds  int64                     `protobuf:"varint,4,opt,name=callback_after_seconds,json=callbackAfterSeconds" json:"callback_after_seconds,omitempty"`
	WorkerId              string                    `protobuf:"bytes,5,opt,name=worker_id,json=workerId" json:"worker_id,omitempty"`
	Status                TaskResult_Status         `protobuf:"varint,6,opt,name=status,enum=com.netflix.conductor.proto.TaskResult_Status" json:"status,omitempty"`
	OutputData            map[string]*_struct.Value `protobuf:"bytes,7,rep,name=output_data,json=outputData" json:"output_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral  struct{}                  `json:"-"`
	XXX_unrecognized      []byte                    `json:"-"`
	XXX_sizecache         int32                     `json:"-"`
}

func (m *TaskResult) Reset()         { *m = TaskResult{} }
func (m *TaskResult) String() string { return proto.CompactTextString(m) }
func (*TaskResult) ProtoMessage()    {}
func (*TaskResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_taskresult_544d5e3612411ce1, []int{0}
}
func (m *TaskResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResult.Unmarshal(m, b)
}
func (m *TaskResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResult.Marshal(b, m, deterministic)
}
func (dst *TaskResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResult.Merge(dst, src)
}
func (m *TaskResult) XXX_Size() int {
	return xxx_messageInfo_TaskResult.Size(m)
}
func (m *TaskResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResult.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResult proto.InternalMessageInfo

func (m *TaskResult) GetWorkflowInstanceId() string {
	if m != nil {
		return m.WorkflowInstanceId
	}
	return ""
}

func (m *TaskResult) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskResult) GetReasonForIncompletion() string {
	if m != nil {
		return m.ReasonForIncompletion
	}
	return ""
}

func (m *TaskResult) GetCallbackAfterSeconds() int64 {
	if m != nil {
		return m.CallbackAfterSeconds
	}
	return 0
}

func (m *TaskResult) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func (m *TaskResult) GetStatus() TaskResult_Status {
	if m != nil {
		return m.Status
	}
	return TaskResult_IN_PROGRESS
}

func (m *TaskResult) GetOutputData() map[string]*_struct.Value {
	if m != nil {
		return m.OutputData
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskResult)(nil), "com.netflix.conductor.proto.TaskResult")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "com.netflix.conductor.proto.TaskResult.OutputDataEntry")
	proto.RegisterEnum("com.netflix.conductor.proto.TaskResult_Status", TaskResult_Status_name, TaskResult_Status_value)
}

func init() { proto.RegisterFile("model/taskresult.proto", fileDescriptor_taskresult_544d5e3612411ce1) }

var fileDescriptor_taskresult_544d5e3612411ce1 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0xba, 0x65, 0xf4, 0x15, 0x58, 0x65, 0x8d, 0x2e, 0xda, 0x38, 0x54, 0x3b, 0xf5,
	0x80, 0x1c, 0x54, 0x10, 0xa0, 0xdd, 0xb6, 0x36, 0x85, 0xa0, 0x41, 0xab, 0x64, 0x43, 0x88, 0x4b,
	0xe4, 0x38, 0x4e, 0x88, 0xe2, 0xc6, 0x95, 0xed, 0x30, 0xf6, 0xdf, 0xf0, 0xa7, 0x22, 0x3b, 0xe9,
	0x36, 0x71, 0x40, 0xdc, 0xf2, 0xde, 0xef, 0xbd, 0xa7, 0xef, 0xfb, 0x1c, 0x18, 0xad, 0x45, 0xc6,
	0xb8, 0xaf, 0x89, 0xaa, 0x24, 0x53, 0x0d, 0xd7, 0x78, 0x23, 0x85, 0x16, 0xe8, 0x84, 0x8a, 0x35,
	0xae, 0x99, 0xce, 0x79, 0xf9, 0x0b, 0x53, 0x51, 0x67, 0x0d, 0xd5, 0x42, 0xb6, 0xf0, 0xf8, 0x45,
	0x21, 0x44, 0xc1, 0x99, 0x6f, 0xab, 0xb4, 0xc9, 0x7d, 0xa5, 0x65, 0x43, 0xbb, 0xd5, 0xd3, 0xdf,
	0xbb, 0x00, 0x57, 0x44, 0x55, 0x91, 0xbd, 0x87, 0x5e, 0xc1, 0xe1, 0x8d, 0x90, 0x55, 0xce, 0xc5,
	0x4d, 0x52, 0xd6, 0x4a, 0x93, 0x9a, 0xb2, 0xa4, 0xcc, 0x3c, 0x67, 0xec, 0x4c, 0xfa, 0x11, 0xda,
	0xb2, 0xb0, 0x43, 0x61, 0x86, 0x8e, 0x60, 0xdf, 0xe8, 0x31, 0x43, 0x3b, 0x76, 0xc8, 0x35, 0x65,
	0x98, 0xa1, 0xb7, 0x70, 0x24, 0x19, 0x51, 0xa2, 0x4e, 0x72, 0x21, 0x93, 0xb2, 0xa6, 0x62, 0xbd,
	0xe1, 0x4c, 0x97, 0xa2, 0xf6, 0x7a, 0x76, 0xf0, 0x79, 0x8b, 0x17, 0x42, 0x86, 0x0f, 0x20, 0x7a,
	0x03, 0x23, 0x4a, 0x38, 0x4f, 0x09, 0xad, 0x12, 0x92, 0x6b, 0x26, 0x13, 0xc5, 0x8c, 0x27, 0xe5,
	0xed, 0x8e, 0x9d, 0x49, 0x2f, 0x3a, 0xdc, 0xd2, 0x73, 0x03, 0xe3, 0x96, 0xa1, 0x13, 0xe8, 0x1b,
	0x71, 0x4c, 0x1a, 0x21, 0x7b, 0xf6, 0xfe, 0xe3, 0xb6, 0x11, 0x66, 0x68, 0x01, 0xae, 0xd2, 0x44,
	0x37, 0xca, 0x73, 0xc7, 0xce, 0xe4, 0xd9, 0x14, 0xe3, 0x7f, 0x04, 0x86, 0xef, 0xe3, 0xc0, 0xb1,
	0xdd, 0x8a, 0xba, 0x6d, 0xf4, 0x0d, 0x06, 0xa2, 0xd1, 0x9b, 0x46, 0x27, 0x19, 0xd1, 0xc4, 0xdb,
	0x1f, 0xf7, 0x26, 0x83, 0xe9, 0xbb, 0xff, 0x3d, 0xb6, 0xb4, 0xab, 0x73, 0xa2, 0x49, 0x50, 0x6b,
	0x79, 0x1b, 0x81, 0xb8, 0x6b, 0x1c, 0x5f, 0xc3, 0xc1, 0x5f, 0x18, 0x0d, 0xa1, 0x57, 0xb1, 0xdb,
	0x2e, 0x79, 0xf3, 0x89, 0x5e, 0xc2, 0xde, 0x4f, 0xc2, 0x1b, 0x66, 0x83, 0x1e, 0x4c, 0x47, 0xb8,
	0x7d, 0x59, 0xbc, 0x7d, 0x59, 0xfc, 0xd5, 0xd0, 0xa8, 0x1d, 0x3a, 0xdb, 0x79, 0xef, 0x9c, 0xce,
	0xc0, 0x6d, 0x2d, 0xa0, 0x03, 0x18, 0x84, 0x5f, 0x92, 0x55, 0xb4, 0xfc, 0x10, 0x05, 0x71, 0x3c,
	0x7c, 0x84, 0x00, 0xdc, 0xc5, 0x79, 0x78, 0x19, 0xcc, 0x87, 0x0e, 0x7a, 0x0a, 0xfd, 0xd9, 0xf2,
	0xf3, 0xea, 0x32, 0xb8, 0x0a, 0xe6, 0xc3, 0x1d, 0x53, 0xc6, 0xb3, 0x8f, 0xc1, 0xfc, 0xda, 0xd0,
	0xde, 0xc5, 0xa7, 0x8b, 0x27, 0xf7, 0x2e, 0x56, 0xe9, 0xf7, 0xb3, 0xa2, 0xd4, 0x3f, 0x9a, 0xd4,
	0xd8, 0xf6, 0x3b, 0xdb, 0xfe, 0x9d, 0x6d, 0x9f, 0xf2, 0x92, 0xd5, 0xda, 0x2f, 0x44, 0x21, 0x37,
	0xf4, 0x41, 0xdf, 0xfe, 0xba, 0xa9, 0x6b, 0xb5, 0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0xbe,
	0xcf, 0xb8, 0x26, 0xca, 0x02, 0x00, 0x00,
}
