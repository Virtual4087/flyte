// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/mpi.proto

package plugins

import (
	fmt "fmt"
	math "math"

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

// MPI operator proposal https://github.com/kubeflow/community/blob/master/proposals/mpi-operator-proposal.md
// Custom proto for plugin that enables distributed training using https://github.com/kubeflow/mpi-operator
type DistributedMPITrainingTask struct {
	// number of worker spawned in the cluster for this job
	NumWorkers int32 `protobuf:"varint,1,opt,name=num_workers,json=numWorkers,proto3" json:"num_workers,omitempty"`
	// number of launcher replicas spawned in the cluster for this job
	// The launcher pod invokes mpirun and communicates with worker pods through MPI.
	NumLauncherReplicas int32 `protobuf:"varint,2,opt,name=num_launcher_replicas,json=numLauncherReplicas,proto3" json:"num_launcher_replicas,omitempty"`
	// number of slots per worker used in hostfile.
	// The available slots (GPUs) in each pod.
	Slots                int32    `protobuf:"varint,3,opt,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedMPITrainingTask) Reset()         { *m = DistributedMPITrainingTask{} }
func (m *DistributedMPITrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedMPITrainingTask) ProtoMessage()    {}
func (*DistributedMPITrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_13cf3fae00e5b069, []int{0}
}

func (m *DistributedMPITrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedMPITrainingTask.Unmarshal(m, b)
}
func (m *DistributedMPITrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedMPITrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedMPITrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedMPITrainingTask.Merge(m, src)
}
func (m *DistributedMPITrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedMPITrainingTask.Size(m)
}
func (m *DistributedMPITrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedMPITrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedMPITrainingTask proto.InternalMessageInfo

func (m *DistributedMPITrainingTask) GetNumWorkers() int32 {
	if m != nil {
		return m.NumWorkers
	}
	return 0
}

func (m *DistributedMPITrainingTask) GetNumLauncherReplicas() int32 {
	if m != nil {
		return m.NumLauncherReplicas
	}
	return 0
}

func (m *DistributedMPITrainingTask) GetSlots() int32 {
	if m != nil {
		return m.Slots
	}
	return 0
}

func init() {
	proto.RegisterType((*DistributedMPITrainingTask)(nil), "flyteidl.plugins.DistributedMPITrainingTask")
}

func init() { proto.RegisterFile("flyteidl/plugins/mpi.proto", fileDescriptor_13cf3fae00e5b069) }

var fileDescriptor_13cf3fae00e5b069 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x31, 0x4b, 0xc6, 0x30,
	0x10, 0x40, 0xa9, 0xf2, 0x39, 0xc4, 0x45, 0xa2, 0x42, 0xe9, 0xa2, 0x38, 0xb9, 0xd8, 0x80, 0xce,
	0x22, 0x88, 0x8b, 0xa0, 0x20, 0xa5, 0x20, 0xb8, 0x94, 0xa4, 0x8d, 0xe9, 0xd1, 0xe4, 0x12, 0x2e,
	0x09, 0xe2, 0x2f, 0xf0, 0x6f, 0x8b, 0x69, 0x75, 0x70, 0xbb, 0xbb, 0xf7, 0x86, 0x7b, 0xac, 0x79,
	0xb7, 0x9f, 0x49, 0xc3, 0x64, 0x45, 0xb0, 0xd9, 0x00, 0x46, 0xe1, 0x02, 0xb4, 0x81, 0x7c, 0xf2,
	0xfc, 0xe8, 0x97, 0xb5, 0x1b, 0xbb, 0xf8, 0xaa, 0x58, 0xf3, 0x00, 0x31, 0x11, 0xa8, 0x9c, 0xf4,
	0xf4, 0xfc, 0xf2, 0xd8, 0x93, 0x04, 0x04, 0x34, 0xbd, 0x8c, 0x0b, 0x3f, 0x63, 0x87, 0x98, 0xdd,
	0xf0, 0xe1, 0x69, 0xd1, 0x14, 0xeb, 0xea, 0xbc, 0xba, 0xdc, 0x75, 0x0c, 0xb3, 0x7b, 0x5d, 0x2f,
	0xfc, 0x9a, 0x9d, 0xfe, 0x08, 0x56, 0x66, 0x1c, 0x67, 0x4d, 0x03, 0xe9, 0x60, 0x61, 0x94, 0xb1,
	0xde, 0x2b, 0xea, 0x31, 0x66, 0xf7, 0xb4, 0xb1, 0x6e, 0x43, 0xfc, 0x84, 0xed, 0xa2, 0xf5, 0x29,
	0xd6, 0xfb, 0xc5, 0x59, 0x97, 0xfb, 0xbb, 0xb7, 0x5b, 0x03, 0x69, 0xce, 0xaa, 0x1d, 0xbd, 0x13,
	0xe5, 0x51, 0x4f, 0x66, 0x1d, 0xc4, 0x5f, 0x93, 0xd1, 0x28, 0x82, 0xba, 0x32, 0x5e, 0xfc, 0xcf,
	0x54, 0x07, 0xa5, 0xf1, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x65, 0x77, 0xd7, 0x89, 0x01, 0x01,
	0x00, 0x00,
}
