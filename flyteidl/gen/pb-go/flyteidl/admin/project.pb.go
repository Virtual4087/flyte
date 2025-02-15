// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/project.proto

package admin

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

// The state of the project is used to control its visibility in the UI and validity.
type Project_ProjectState int32

const (
	// By default, all projects are considered active.
	Project_ACTIVE Project_ProjectState = 0
	// Archived projects are no longer visible in the UI and no longer valid.
	Project_ARCHIVED Project_ProjectState = 1
	// System generated projects that aren't explicitly created or managed by a user.
	Project_SYSTEM_GENERATED Project_ProjectState = 2
)

var Project_ProjectState_name = map[int32]string{
	0: "ACTIVE",
	1: "ARCHIVED",
	2: "SYSTEM_GENERATED",
}

var Project_ProjectState_value = map[string]int32{
	"ACTIVE":           0,
	"ARCHIVED":         1,
	"SYSTEM_GENERATED": 2,
}

func (x Project_ProjectState) String() string {
	return proto.EnumName(Project_ProjectState_name, int32(x))
}

func (Project_ProjectState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{1, 0}
}

// Namespace within a project commonly used to differentiate between different service instances.
// e.g. "production", "development", etc.
type Domain struct {
	// Globally unique domain name.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Domain) Reset()         { *m = Domain{} }
func (m *Domain) String() string { return proto.CompactTextString(m) }
func (*Domain) ProtoMessage()    {}
func (*Domain) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{0}
}

func (m *Domain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Domain.Unmarshal(m, b)
}
func (m *Domain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Domain.Marshal(b, m, deterministic)
}
func (m *Domain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Domain.Merge(m, src)
}
func (m *Domain) XXX_Size() int {
	return xxx_messageInfo_Domain.Size(m)
}
func (m *Domain) XXX_DiscardUnknown() {
	xxx_messageInfo_Domain.DiscardUnknown(m)
}

var xxx_messageInfo_Domain proto.InternalMessageInfo

func (m *Domain) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Domain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Top-level namespace used to classify different entities like workflows and executions.
type Project struct {
	// Globally unique project name.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name.
	Name        string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Domains     []*Domain `protobuf:"bytes,3,rep,name=domains,proto3" json:"domains,omitempty"`
	Description string    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Leverage Labels from flyteidl.admin.common.proto to
	// tag projects with ownership information.
	Labels               *Labels              `protobuf:"bytes,5,opt,name=labels,proto3" json:"labels,omitempty"`
	State                Project_ProjectState `protobuf:"varint,6,opt,name=state,proto3,enum=flyteidl.admin.Project_ProjectState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{1}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetDomains() []*Domain {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetLabels() *Labels {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Project) GetState() Project_ProjectState {
	if m != nil {
		return m.State
	}
	return Project_ACTIVE
}

// Represents a list of projects.
// See :ref:`ref_flyteidl.admin.Project` for more details
type Projects struct {
	Projects []*Project `protobuf:"bytes,1,rep,name=projects,proto3" json:"projects,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Projects) Reset()         { *m = Projects{} }
func (m *Projects) String() string { return proto.CompactTextString(m) }
func (*Projects) ProtoMessage()    {}
func (*Projects) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{2}
}

func (m *Projects) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Projects.Unmarshal(m, b)
}
func (m *Projects) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Projects.Marshal(b, m, deterministic)
}
func (m *Projects) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Projects.Merge(m, src)
}
func (m *Projects) XXX_Size() int {
	return xxx_messageInfo_Projects.Size(m)
}
func (m *Projects) XXX_DiscardUnknown() {
	xxx_messageInfo_Projects.DiscardUnknown(m)
}

var xxx_messageInfo_Projects proto.InternalMessageInfo

func (m *Projects) GetProjects() []*Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *Projects) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Request to retrieve a list of projects matching specified filters.
// See :ref:`ref_flyteidl.admin.Project` for more details
type ProjectListRequest struct {
	// Indicates the number of projects to be returned.
	// +required
	Limit uint32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, this server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,3,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,4,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectListRequest) Reset()         { *m = ProjectListRequest{} }
func (m *ProjectListRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectListRequest) ProtoMessage()    {}
func (*ProjectListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{3}
}

func (m *ProjectListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectListRequest.Unmarshal(m, b)
}
func (m *ProjectListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectListRequest.Marshal(b, m, deterministic)
}
func (m *ProjectListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectListRequest.Merge(m, src)
}
func (m *ProjectListRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectListRequest.Size(m)
}
func (m *ProjectListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectListRequest proto.InternalMessageInfo

func (m *ProjectListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ProjectListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ProjectListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *ProjectListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

// Adds a new user-project within the Flyte deployment.
// See :ref:`ref_flyteidl.admin.Project` for more details
type ProjectRegisterRequest struct {
	// +required
	Project              *Project `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRegisterRequest) Reset()         { *m = ProjectRegisterRequest{} }
func (m *ProjectRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectRegisterRequest) ProtoMessage()    {}
func (*ProjectRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{4}
}

func (m *ProjectRegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRegisterRequest.Unmarshal(m, b)
}
func (m *ProjectRegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRegisterRequest.Marshal(b, m, deterministic)
}
func (m *ProjectRegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRegisterRequest.Merge(m, src)
}
func (m *ProjectRegisterRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectRegisterRequest.Size(m)
}
func (m *ProjectRegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRegisterRequest proto.InternalMessageInfo

func (m *ProjectRegisterRequest) GetProject() *Project {
	if m != nil {
		return m.Project
	}
	return nil
}

// Purposefully empty, may be updated in the future.
type ProjectRegisterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectRegisterResponse) Reset()         { *m = ProjectRegisterResponse{} }
func (m *ProjectRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectRegisterResponse) ProtoMessage()    {}
func (*ProjectRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{5}
}

func (m *ProjectRegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectRegisterResponse.Unmarshal(m, b)
}
func (m *ProjectRegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectRegisterResponse.Marshal(b, m, deterministic)
}
func (m *ProjectRegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectRegisterResponse.Merge(m, src)
}
func (m *ProjectRegisterResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectRegisterResponse.Size(m)
}
func (m *ProjectRegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectRegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectRegisterResponse proto.InternalMessageInfo

// Purposefully empty, may be updated in the future.
type ProjectUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectUpdateResponse) Reset()         { *m = ProjectUpdateResponse{} }
func (m *ProjectUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectUpdateResponse) ProtoMessage()    {}
func (*ProjectUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2db065ce03bf106d, []int{6}
}

func (m *ProjectUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectUpdateResponse.Unmarshal(m, b)
}
func (m *ProjectUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectUpdateResponse.Marshal(b, m, deterministic)
}
func (m *ProjectUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectUpdateResponse.Merge(m, src)
}
func (m *ProjectUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectUpdateResponse.Size(m)
}
func (m *ProjectUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectUpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("flyteidl.admin.Project_ProjectState", Project_ProjectState_name, Project_ProjectState_value)
	proto.RegisterType((*Domain)(nil), "flyteidl.admin.Domain")
	proto.RegisterType((*Project)(nil), "flyteidl.admin.Project")
	proto.RegisterType((*Projects)(nil), "flyteidl.admin.Projects")
	proto.RegisterType((*ProjectListRequest)(nil), "flyteidl.admin.ProjectListRequest")
	proto.RegisterType((*ProjectRegisterRequest)(nil), "flyteidl.admin.ProjectRegisterRequest")
	proto.RegisterType((*ProjectRegisterResponse)(nil), "flyteidl.admin.ProjectRegisterResponse")
	proto.RegisterType((*ProjectUpdateResponse)(nil), "flyteidl.admin.ProjectUpdateResponse")
}

func init() { proto.RegisterFile("flyteidl/admin/project.proto", fileDescriptor_2db065ce03bf106d) }

var fileDescriptor_2db065ce03bf106d = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xb1, 0xdb, 0xd8, 0x61, 0x52, 0xa2, 0x68, 0x14, 0x1a, 0xf3, 0x71, 0xb0, 0x2c, 0x0e,
	0x39, 0x50, 0x1b, 0xd2, 0x1b, 0x08, 0xa4, 0xb4, 0xb1, 0xa0, 0xa2, 0x20, 0xb4, 0x49, 0x2b, 0xc1,
	0xa5, 0xf2, 0xc7, 0xd6, 0x2c, 0xd8, 0x5e, 0xe3, 0xdd, 0x1e, 0xf2, 0x08, 0x3c, 0x0c, 0xef, 0x88,
	0xbc, 0x5e, 0x57, 0x6d, 0x08, 0xa8, 0x27, 0xef, 0xcc, 0xfc, 0x76, 0xe6, 0x3f, 0x33, 0x5e, 0x78,
	0x7a, 0x99, 0xaf, 0x25, 0x65, 0x69, 0x1e, 0x44, 0x69, 0xc1, 0xca, 0xa0, 0xaa, 0xf9, 0x77, 0x9a,
	0x48, 0xbf, 0xaa, 0xb9, 0xe4, 0x38, 0xec, 0xa2, 0xbe, 0x8a, 0x3e, 0x7e, 0xb2, 0x41, 0x27, 0xbc,
	0x28, 0x78, 0xd9, 0xc2, 0xde, 0x73, 0xb0, 0x16, 0xbc, 0x88, 0x58, 0x89, 0x43, 0x30, 0x59, 0xea,
	0x18, 0xae, 0x31, 0xbd, 0x4f, 0x4c, 0x96, 0x22, 0xc2, 0x6e, 0x19, 0x15, 0xd4, 0x31, 0x95, 0x47,
	0x9d, 0xbd, 0xdf, 0x26, 0xd8, 0x9f, 0xdb, 0x62, 0x77, 0xe1, 0xf1, 0x05, 0xd8, 0xa9, 0xca, 0x2e,
	0x9c, 0x1d, 0x77, 0x67, 0x3a, 0x98, 0xed, 0xfb, 0xb7, 0xc5, 0xf9, 0x6d, 0x71, 0xd2, 0x61, 0xe8,
	0xc2, 0x20, 0xa5, 0x22, 0xa9, 0x59, 0x25, 0x19, 0x2f, 0x9d, 0x5d, 0x95, 0xec, 0xa6, 0x0b, 0x7d,
	0xb0, 0xf2, 0x28, 0xa6, 0xb9, 0x70, 0x7a, 0xae, 0xb1, 0x2d, 0xe5, 0xa9, 0x8a, 0x12, 0x4d, 0xe1,
	0x2b, 0xe8, 0x09, 0x19, 0x49, 0xea, 0x58, 0xae, 0x31, 0x1d, 0xce, 0x9e, 0x6d, 0xe2, 0xba, 0x9f,
	0xee, 0xbb, 0x6c, 0x58, 0xd2, 0x5e, 0xf1, 0xde, 0xc2, 0xde, 0x4d, 0x37, 0x02, 0x58, 0xf3, 0xe3,
	0xd5, 0xc9, 0x79, 0x38, 0xba, 0x87, 0x7b, 0xd0, 0x9f, 0x93, 0xe3, 0xf7, 0x27, 0xe7, 0xe1, 0x62,
	0x64, 0xe0, 0x18, 0x46, 0xcb, 0x2f, 0xcb, 0x55, 0xf8, 0xf1, 0xe2, 0x5d, 0xf8, 0x29, 0x24, 0xf3,
	0x55, 0xb8, 0x18, 0x99, 0xde, 0x19, 0xf4, 0xf5, 0x7d, 0x81, 0x87, 0xd0, 0xd7, 0x7b, 0x12, 0x8e,
	0xa1, 0x86, 0x31, 0xf9, 0x87, 0x14, 0x72, 0x0d, 0xe2, 0x18, 0x7a, 0x92, 0xff, 0xa0, 0xa5, 0x9e,
	0x6a, 0x6b, 0x78, 0xbf, 0x0c, 0x40, 0xcd, 0x9e, 0x32, 0x21, 0x09, 0xfd, 0x79, 0x45, 0x85, 0x6c,
	0xe0, 0x9c, 0x15, 0x4c, 0xaa, 0xa5, 0x3c, 0x20, 0xad, 0xb1, 0x3d, 0x05, 0x3a, 0x60, 0x5f, 0xb2,
	0x5c, 0xd2, 0xba, 0xd9, 0x4c, 0xe3, 0xef, 0x4c, 0x3c, 0x00, 0x5b, 0xf0, 0x5a, 0x5e, 0xc4, 0x6b,
	0x35, 0xfd, 0xc1, 0x6c, 0xbc, 0x29, 0x73, 0xc9, 0x6b, 0x49, 0xac, 0x06, 0x3a, 0x5a, 0x7b, 0x1f,
	0x60, 0xbf, 0x93, 0x4d, 0x33, 0x26, 0x24, 0xad, 0x3b, 0x39, 0x2f, 0xc1, 0xd6, 0x7d, 0x28, 0x41,
	0xff, 0xe9, 0xb7, 0xe3, 0xbc, 0x47, 0x30, 0xf9, 0x2b, 0x99, 0xa8, 0x78, 0x29, 0xa8, 0x37, 0x81,
	0x87, 0x3a, 0x74, 0x56, 0xa5, 0xcd, 0x8a, 0x74, 0xe0, 0xe8, 0xcd, 0xd7, 0xd7, 0x19, 0x93, 0xdf,
	0xae, 0x62, 0x3f, 0xe1, 0x45, 0xa0, 0x2a, 0xf0, 0x3a, 0x6b, 0x0f, 0xc1, 0xf5, 0xaf, 0x9f, 0xd1,
	0x32, 0xa8, 0xe2, 0x83, 0x8c, 0x07, 0xb7, 0x5f, 0x43, 0x6c, 0xa9, 0x77, 0x70, 0xf8, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x33, 0x79, 0x69, 0x68, 0x54, 0x03, 0x00, 0x00,
}
