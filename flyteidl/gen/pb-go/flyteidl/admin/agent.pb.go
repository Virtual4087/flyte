// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/agent.proto

package admin

import (
	fmt "fmt"
	math "math"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The state of the execution is used to control its visibility in the UI/CLI.
type State int32

const (
	State_RETRYABLE_FAILURE State = 0
	State_PERMANENT_FAILURE State = 1
	State_PENDING           State = 2
	State_RUNNING           State = 3
	State_SUCCEEDED         State = 4
)

var State_name = map[int32]string{
	0: "RETRYABLE_FAILURE",
	1: "PERMANENT_FAILURE",
	2: "PENDING",
	3: "RUNNING",
	4: "SUCCEEDED",
}

var State_value = map[string]int32{
	"RETRYABLE_FAILURE": 0,
	"PERMANENT_FAILURE": 1,
	"PENDING":           2,
	"RUNNING":           3,
	"SUCCEEDED":         4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

// Represents a subset of runtime task execution metadata that are relevant to external plugins.
type TaskExecutionMetadata struct {
	// ID of the task execution
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=task_execution_id,json=taskExecutionId,proto3" json:"task_execution_id,omitempty"`
	// k8s namespace where the task is executed in
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Labels attached to the task execution
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations attached to the task execution
	Annotations map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// k8s service account associated with the task execution
	K8SServiceAccount string `protobuf:"bytes,5,opt,name=k8s_service_account,json=k8sServiceAccount,proto3" json:"k8s_service_account,omitempty"`
	// Environment variables attached to the task execution
	EnvironmentVariables map[string]string `protobuf:"bytes,6,rep,name=environment_variables,json=environmentVariables,proto3" json:"environment_variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskExecutionMetadata) Reset()         { *m = TaskExecutionMetadata{} }
func (m *TaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionMetadata) ProtoMessage()    {}
func (*TaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{0}
}

func (m *TaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionMetadata.Unmarshal(m, b)
}
func (m *TaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *TaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionMetadata.Merge(m, src)
}
func (m *TaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionMetadata.Size(m)
}
func (m *TaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionMetadata proto.InternalMessageInfo

func (m *TaskExecutionMetadata) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.TaskExecutionId
	}
	return nil
}

func (m *TaskExecutionMetadata) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *TaskExecutionMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TaskExecutionMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TaskExecutionMetadata) GetK8SServiceAccount() string {
	if m != nil {
		return m.K8SServiceAccount
	}
	return ""
}

func (m *TaskExecutionMetadata) GetEnvironmentVariables() map[string]string {
	if m != nil {
		return m.EnvironmentVariables
	}
	return nil
}

// Represents a request structure to create task.
type CreateTaskRequest struct {
	// The inputs required to start the execution. All required inputs must be
	// included in this map. If not required and not provided, defaults apply.
	// +optional
	Inputs *core.LiteralMap `protobuf:"bytes,1,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Template of the task that encapsulates all the metadata of the task.
	Template *core.TaskTemplate `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	// Prefix for where task output data will be written. (e.g. s3://my-bucket/randomstring)
	OutputPrefix string `protobuf:"bytes,3,opt,name=output_prefix,json=outputPrefix,proto3" json:"output_prefix,omitempty"`
	// subset of runtime task execution metadata.
	TaskExecutionMetadata *TaskExecutionMetadata `protobuf:"bytes,4,opt,name=task_execution_metadata,json=taskExecutionMetadata,proto3" json:"task_execution_metadata,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{1}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *CreateTaskRequest) GetTemplate() *core.TaskTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *CreateTaskRequest) GetOutputPrefix() string {
	if m != nil {
		return m.OutputPrefix
	}
	return ""
}

func (m *CreateTaskRequest) GetTaskExecutionMetadata() *TaskExecutionMetadata {
	if m != nil {
		return m.TaskExecutionMetadata
	}
	return nil
}

// Represents a create response structure.
type CreateTaskResponse struct {
	// Metadata is created by the agent. It could be a string (jobId) or a dict (more complex metadata).
	ResourceMeta         []byte   `protobuf:"bytes,1,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{2}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// A message used to fetch a job resource from flyte agent server.
type GetTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{3}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *GetTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to get an individual task resource.
type GetTaskResponse struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetTaskResponse) Reset()         { *m = GetTaskResponse{} }
func (m *GetTaskResponse) String() string { return proto.CompactTextString(m) }
func (*GetTaskResponse) ProtoMessage()    {}
func (*GetTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{4}
}

func (m *GetTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskResponse.Unmarshal(m, b)
}
func (m *GetTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskResponse.Marshal(b, m, deterministic)
}
func (m *GetTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskResponse.Merge(m, src)
}
func (m *GetTaskResponse) XXX_Size() int {
	return xxx_messageInfo_GetTaskResponse.Size(m)
}
func (m *GetTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskResponse proto.InternalMessageInfo

func (m *GetTaskResponse) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

type Resource struct {
	// The state of the execution is used to control its visibility in the UI/CLI.
	State State `protobuf:"varint,1,opt,name=state,proto3,enum=flyteidl.admin.State" json:"state,omitempty"`
	// The outputs of the execution. It's typically used by sql task. Agent service will create a
	// Structured dataset pointing to the query result table.
	// +optional
	Outputs              *core.LiteralMap `protobuf:"bytes,2,opt,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{5}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RETRYABLE_FAILURE
}

func (m *Resource) GetOutputs() *core.LiteralMap {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A message used to delete a task.
type DeleteTaskRequest struct {
	// A predefined yet extensible Task type identifier.
	TaskType string `protobuf:"bytes,1,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Metadata about the resource to be pass to the agent.
	ResourceMeta         []byte   `protobuf:"bytes,2,opt,name=resource_meta,json=resourceMeta,proto3" json:"resource_meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskRequest) Reset()         { *m = DeleteTaskRequest{} }
func (m *DeleteTaskRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskRequest) ProtoMessage()    {}
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{6}
}

func (m *DeleteTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskRequest.Unmarshal(m, b)
}
func (m *DeleteTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskRequest.Marshal(b, m, deterministic)
}
func (m *DeleteTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskRequest.Merge(m, src)
}
func (m *DeleteTaskRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskRequest.Size(m)
}
func (m *DeleteTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskRequest proto.InternalMessageInfo

func (m *DeleteTaskRequest) GetTaskType() string {
	if m != nil {
		return m.TaskType
	}
	return ""
}

func (m *DeleteTaskRequest) GetResourceMeta() []byte {
	if m != nil {
		return m.ResourceMeta
	}
	return nil
}

// Response to delete a task.
type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c434e52bb0028071, []int{7}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flyteidl.admin.State", State_name, State_value)
	proto.RegisterType((*TaskExecutionMetadata)(nil), "flyteidl.admin.TaskExecutionMetadata")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.EnvironmentVariablesEntry")
	proto.RegisterMapType((map[string]string)(nil), "flyteidl.admin.TaskExecutionMetadata.LabelsEntry")
	proto.RegisterType((*CreateTaskRequest)(nil), "flyteidl.admin.CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "flyteidl.admin.CreateTaskResponse")
	proto.RegisterType((*GetTaskRequest)(nil), "flyteidl.admin.GetTaskRequest")
	proto.RegisterType((*GetTaskResponse)(nil), "flyteidl.admin.GetTaskResponse")
	proto.RegisterType((*Resource)(nil), "flyteidl.admin.Resource")
	proto.RegisterType((*DeleteTaskRequest)(nil), "flyteidl.admin.DeleteTaskRequest")
	proto.RegisterType((*DeleteTaskResponse)(nil), "flyteidl.admin.DeleteTaskResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/agent.proto", fileDescriptor_c434e52bb0028071) }

var fileDescriptor_c434e52bb0028071 = []byte{
	// 726 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xed, 0x6e, 0xe2, 0x46,
	0x14, 0x2d, 0x10, 0x08, 0x5c, 0xf2, 0x01, 0xd3, 0xa0, 0x3a, 0x24, 0xad, 0x22, 0xaa, 0x56, 0x51,
	0xab, 0x1a, 0x85, 0x54, 0x6d, 0xd2, 0xaa, 0xad, 0x48, 0x70, 0x11, 0x12, 0x41, 0xd1, 0x04, 0xaa,
	0xb6, 0xd2, 0x2e, 0x1a, 0xcc, 0x85, 0xb5, 0x30, 0x63, 0xaf, 0x67, 0x8c, 0xc2, 0xef, 0x7d, 0x89,
	0x7d, 0xdc, 0x95, 0xc7, 0x1f, 0x01, 0xc4, 0xae, 0x12, 0xed, 0x3f, 0xcf, 0x3d, 0xe7, 0x9e, 0x39,
	0x73, 0xee, 0xd8, 0x86, 0xea, 0xc4, 0x5e, 0x4a, 0xb4, 0xc6, 0x76, 0x9d, 0x8d, 0xe7, 0x16, 0xaf,
	0xb3, 0x29, 0x72, 0xa9, 0xbb, 0x9e, 0x23, 0x1d, 0x72, 0x10, 0x63, 0xba, 0xc2, 0xaa, 0xa7, 0x09,
	0xd7, 0x74, 0x3c, 0xac, 0xdb, 0x96, 0x44, 0x8f, 0xd9, 0x22, 0x64, 0x57, 0x8f, 0xd7, 0x51, 0xc9,
	0xc4, 0x2c, 0x86, 0xbe, 0x5e, 0x87, 0x2c, 0x2e, 0xd1, 0x9b, 0x30, 0x13, 0x23, 0xf8, 0x9b, 0x0d,
	0x78, 0x8c, 0x5c, 0x5a, 0x13, 0x0b, 0xbd, 0x10, 0xaf, 0xbd, 0xcf, 0x42, 0xa5, 0xcf, 0xc4, 0xcc,
	0x78, 0x44, 0xd3, 0x97, 0x96, 0xc3, 0xef, 0x50, 0xb2, 0x31, 0x93, 0x8c, 0x50, 0x28, 0x07, 0xfb,
	0x0c, 0x31, 0x46, 0x86, 0xd6, 0x58, 0x4b, 0x9d, 0xa5, 0xce, 0x8b, 0x8d, 0xef, 0xf5, 0xc4, 0x7d,
	0xa0, 0xaa, 0xaf, 0x09, 0x74, 0x92, 0x2d, 0xe8, 0xa1, 0x5c, 0x07, 0xc8, 0x29, 0x14, 0x38, 0x9b,
	0xa3, 0x70, 0x99, 0x89, 0x5a, 0xfa, 0x2c, 0x75, 0x5e, 0xa0, 0x4f, 0x05, 0xd2, 0x81, 0x9c, 0xcd,
	0x46, 0x68, 0x0b, 0x2d, 0x73, 0x96, 0x39, 0x2f, 0x36, 0x2e, 0xf4, 0xf5, 0x90, 0xf4, 0xad, 0x46,
	0xf5, 0xae, 0xea, 0x31, 0xb8, 0xf4, 0x96, 0x34, 0x12, 0x20, 0xff, 0x42, 0x91, 0x71, 0xee, 0x48,
	0x16, 0x30, 0x85, 0xb6, 0xa3, 0xf4, 0x7e, 0x79, 0x9e, 0x5e, 0xf3, 0xa9, 0x31, 0x14, 0x5d, 0x95,
	0x22, 0x3a, 0x7c, 0x39, 0xbb, 0x12, 0x43, 0x81, 0xde, 0xc2, 0x32, 0x71, 0xc8, 0x4c, 0xd3, 0xf1,
	0xb9, 0xd4, 0xb2, 0xea, 0x30, 0xe5, 0xd9, 0x95, 0x78, 0x08, 0x91, 0x66, 0x08, 0x10, 0x09, 0x15,
	0xe4, 0x0b, 0xcb, 0x73, 0xf8, 0x1c, 0xb9, 0x1c, 0x2e, 0x98, 0x67, 0xb1, 0x91, 0x8d, 0x42, 0xcb,
	0x29, 0x4f, 0x7f, 0x3d, 0xcf, 0x93, 0xf1, 0x24, 0xf1, 0x4f, 0xac, 0x10, 0x9a, 0x3b, 0xc2, 0x2d,
	0x50, 0xf5, 0x1a, 0x8a, 0x2b, 0xb1, 0x90, 0x12, 0x64, 0x66, 0xb8, 0x54, 0xd3, 0x2b, 0xd0, 0xe0,
	0x91, 0x1c, 0x41, 0x76, 0xc1, 0x6c, 0x3f, 0x9e, 0x42, 0xb8, 0xf8, 0x2d, 0x7d, 0x95, 0xaa, 0xfe,
	0x09, 0xa5, 0xcd, 0x04, 0x5e, 0xd4, 0xdf, 0x86, 0xe3, 0x8f, 0xba, 0x7d, 0x89, 0x50, 0xed, 0x5d,
	0x1a, 0xca, 0xb7, 0x1e, 0x32, 0x89, 0x41, 0x26, 0x14, 0xdf, 0xfa, 0x28, 0x24, 0xb9, 0x80, 0x9c,
	0xc5, 0x5d, 0x5f, 0x8a, 0xe8, 0x2e, 0x1e, 0x6f, 0xdc, 0xc5, 0x6e, 0xf8, 0xe6, 0xdc, 0x31, 0x97,
	0x46, 0x44, 0xf2, 0x2b, 0xe4, 0x25, 0xce, 0x5d, 0x9b, 0xc9, 0x70, 0x97, 0x62, 0xe3, 0x64, 0xcb,
	0x05, 0xee, 0x47, 0x14, 0x9a, 0x90, 0xc9, 0xb7, 0xb0, 0xef, 0xf8, 0xd2, 0xf5, 0xe5, 0xd0, 0xf5,
	0x70, 0x62, 0x3d, 0x6a, 0x19, 0xe5, 0x71, 0x2f, 0x2c, 0xde, 0xab, 0x1a, 0x79, 0x05, 0x5f, 0x6d,
	0xbc, 0x27, 0xf3, 0x68, 0x6a, 0xda, 0x8e, 0xda, 0xec, 0xbb, 0x67, 0x8d, 0x98, 0x56, 0xe4, 0xb6,
	0x72, 0xed, 0x1a, 0xc8, 0x6a, 0x08, 0xc2, 0x75, 0xb8, 0x50, 0xce, 0x3c, 0x14, 0x8e, 0xef, 0x99,
	0xa8, 0xb6, 0x53, 0x61, 0xec, 0xd1, 0xbd, 0xb8, 0x18, 0xb4, 0xd7, 0x28, 0x1c, 0xb4, 0x51, 0xae,
	0x86, 0x77, 0x02, 0x05, 0xe5, 0x55, 0x2e, 0x5d, 0x8c, 0x86, 0x90, 0x0f, 0x0a, 0xfd, 0xa5, 0xbb,
	0x45, 0x33, 0xbd, 0x45, 0xb3, 0x0d, 0x87, 0x89, 0x66, 0xe4, 0xe5, 0x67, 0xc8, 0xc7, 0x94, 0x68,
	0x26, 0xda, 0xe6, 0x89, 0x69, 0x84, 0xd3, 0x84, 0x59, 0xb3, 0x21, 0x1f, 0x57, 0xc9, 0x8f, 0x90,
	0x15, 0x32, 0x98, 0x4e, 0xd0, 0x7e, 0xd0, 0xa8, 0x6c, 0xb6, 0x3f, 0x04, 0x20, 0x0d, 0x39, 0xe4,
	0x12, 0x76, 0xc3, 0xfc, 0x45, 0x34, 0xcc, 0x4f, 0xdc, 0x80, 0x98, 0x59, 0x1b, 0x40, 0xb9, 0x85,
	0x36, 0xae, 0x5f, 0xa5, 0xcf, 0x4f, 0xe3, 0x08, 0xc8, 0xaa, 0x6c, 0x18, 0xc8, 0x0f, 0xaf, 0x21,
	0xab, 0x1c, 0x93, 0x0a, 0x94, 0xa9, 0xd1, 0xa7, 0xff, 0x35, 0x6f, 0xba, 0xc6, 0xf0, 0xef, 0x66,
	0xa7, 0x3b, 0xa0, 0x46, 0xe9, 0x8b, 0xa0, 0x7c, 0x6f, 0xd0, 0xbb, 0x66, 0xcf, 0xe8, 0xf5, 0x93,
	0x72, 0x8a, 0x14, 0x61, 0xf7, 0xde, 0xe8, 0xb5, 0x3a, 0xbd, 0x76, 0x29, 0x1d, 0x2c, 0xe8, 0xa0,
	0xd7, 0x0b, 0x16, 0x19, 0xb2, 0x0f, 0x85, 0x87, 0xc1, 0xed, 0xad, 0x61, 0xb4, 0x8c, 0x56, 0x69,
	0xe7, 0xe6, 0x8f, 0xff, 0x7f, 0x9f, 0x5a, 0xf2, 0x8d, 0x3f, 0xd2, 0x4d, 0x67, 0x5e, 0x57, 0x87,
	0x77, 0xbc, 0x69, 0xf8, 0x50, 0x4f, 0xbe, 0xf7, 0x53, 0xe4, 0x75, 0x77, 0xf4, 0xd3, 0xd4, 0xa9,
	0xaf, 0xff, 0x86, 0x46, 0x39, 0xf5, 0xe5, 0xbf, 0xfc, 0x10, 0x00, 0x00, 0xff, 0xff, 0x14, 0xef,
	0x1e, 0x8b, 0x9f, 0x06, 0x00, 0x00,
}
