// Code generated by protoc-gen-go.
// source: tableacl.proto
// DO NOT EDIT!

/*
Package tableacl is a generated protocol buffer package.

It is generated from these files:
	tableacl.proto

It has these top-level messages:
	TableGroupSpec
	Config
*/
package tableacl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TableGroupSpec defines ACLs for a group of tables.
type TableGroupSpec struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// either tables or a table name prefixes (if it ends in a %)
	TableNamesOrPrefixes []string `protobuf:"bytes,2,rep,name=table_names_or_prefixes,json=tableNamesOrPrefixes" json:"table_names_or_prefixes,omitempty"`
	Readers              []string `protobuf:"bytes,3,rep,name=readers" json:"readers,omitempty"`
	Writers              []string `protobuf:"bytes,4,rep,name=writers" json:"writers,omitempty"`
	Admins               []string `protobuf:"bytes,5,rep,name=admins" json:"admins,omitempty"`
}

func (m *TableGroupSpec) Reset()                    { *m = TableGroupSpec{} }
func (m *TableGroupSpec) String() string            { return proto.CompactTextString(m) }
func (*TableGroupSpec) ProtoMessage()               {}
func (*TableGroupSpec) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TableGroupSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableGroupSpec) GetTableNamesOrPrefixes() []string {
	if m != nil {
		return m.TableNamesOrPrefixes
	}
	return nil
}

func (m *TableGroupSpec) GetReaders() []string {
	if m != nil {
		return m.Readers
	}
	return nil
}

func (m *TableGroupSpec) GetWriters() []string {
	if m != nil {
		return m.Writers
	}
	return nil
}

func (m *TableGroupSpec) GetAdmins() []string {
	if m != nil {
		return m.Admins
	}
	return nil
}

type Config struct {
	TableGroups []*TableGroupSpec `protobuf:"bytes,1,rep,name=table_groups,json=tableGroups" json:"table_groups,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetTableGroups() []*TableGroupSpec {
	if m != nil {
		return m.TableGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*TableGroupSpec)(nil), "tableacl.TableGroupSpec")
	proto.RegisterType((*Config)(nil), "tableacl.Config")
}

func init() { proto.RegisterFile("tableacl.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x49, 0x4c, 0xca,
	0x49, 0x4d, 0x4c, 0xce, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x96,
	0x33, 0x72, 0xf1, 0x85, 0x80, 0x38, 0xee, 0x45, 0xf9, 0xa5, 0x05, 0xc1, 0x05, 0xa9, 0xc9, 0x42,
	0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6,
	0x90, 0x29, 0x97, 0x38, 0x58, 0x4b, 0x3c, 0x88, 0x57, 0x1c, 0x9f, 0x5f, 0x14, 0x5f, 0x50, 0x94,
	0x9a, 0x96, 0x59, 0x91, 0x5a, 0x2c, 0xc1, 0xa4, 0xc0, 0xac, 0xc1, 0x19, 0x24, 0x02, 0x96, 0xf6,
	0x03, 0xc9, 0xfa, 0x17, 0x05, 0x40, 0xe5, 0x84, 0x24, 0xb8, 0xd8, 0x8b, 0x52, 0x13, 0x53, 0x52,
	0x8b, 0x8a, 0x25, 0x98, 0xc1, 0xca, 0x60, 0x5c, 0x90, 0x4c, 0x79, 0x51, 0x66, 0x09, 0x48, 0x86,
	0x05, 0x22, 0x03, 0xe5, 0x0a, 0x89, 0x71, 0xb1, 0x25, 0xa6, 0xe4, 0x66, 0xe6, 0x15, 0x4b, 0xb0,
	0x82, 0x25, 0xa0, 0x3c, 0x25, 0x57, 0x2e, 0x36, 0xe7, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x21, 0x6b,
	0x2e, 0x1e, 0x88, 0x63, 0xd2, 0x41, 0x6e, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92,
	0xd0, 0x83, 0x7b, 0x12, 0xd5, 0x43, 0x41, 0xdc, 0x25, 0x70, 0x7e, 0x71, 0x12, 0x1b, 0x38, 0x04,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x64, 0x57, 0xee, 0xfe, 0x13, 0x01, 0x00, 0x00,
}
