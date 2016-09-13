// Code generated by protoc-gen-gogo.
// source: lsd-config.proto
// DO NOT EDIT!

/*
Package lsd is a generated protocol buffer package.

It is generated from these files:
	lsd-config.proto

It has these top-level messages:
	LsdConfig
*/
package lsd

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LsdConfig struct {
	ServerConfig *LsdConfigServerConfigT `protobuf:"bytes,1,opt,name=server_config" json:"server_config,omitempty"`
	ClientConfig *LsdConfigClientConfigT `protobuf:"bytes,2,opt,name=client_config" json:"client_config,omitempty"`
}

func (m *LsdConfig) Reset()         { *m = LsdConfig{} }
func (m *LsdConfig) String() string { return proto.CompactTextString(m) }
func (*LsdConfig) ProtoMessage()    {}

func (m *LsdConfig) GetServerConfig() *LsdConfigServerConfigT {
	if m != nil {
		return m.ServerConfig
	}
	return nil
}

func (m *LsdConfig) GetClientConfig() *LsdConfigClientConfigT {
	if m != nil {
		return m.ClientConfig
	}
	return nil
}

type LsdConfigServerConfigT struct {
	TargetDir           *string                                  `protobuf:"bytes,1,req,name=target_dir" json:"target_dir,omitempty"`
	MaxFileSize         *uint64                                  `protobuf:"varint,2,opt,name=max_file_size,def=5000000" json:"max_file_size,omitempty"`
	FileRotateInterval  *uint64                                  `protobuf:"varint,3,opt,name=file_rotate_interval,def=5" json:"file_rotate_interval,omitempty"`
	PerCategorySettings []*LsdConfigServerConfigTServerSettingsT `protobuf:"bytes,4,rep,name=per_category_settings" json:"per_category_settings,omitempty"`
	ChunkOutput         *bool                                    `protobuf:"varint,5,opt,name=chunk_output,def=1" json:"chunk_output,omitempty"`
}

func (m *LsdConfigServerConfigT) Reset()         { *m = LsdConfigServerConfigT{} }
func (m *LsdConfigServerConfigT) String() string { return proto.CompactTextString(m) }
func (*LsdConfigServerConfigT) ProtoMessage()    {}

const Default_LsdConfigServerConfigT_MaxFileSize uint64 = 5000000
const Default_LsdConfigServerConfigT_FileRotateInterval uint64 = 5
const Default_LsdConfigServerConfigT_ChunkOutput bool = true

func (m *LsdConfigServerConfigT) GetTargetDir() string {
	if m != nil && m.TargetDir != nil {
		return *m.TargetDir
	}
	return ""
}

func (m *LsdConfigServerConfigT) GetMaxFileSize() uint64 {
	if m != nil && m.MaxFileSize != nil {
		return *m.MaxFileSize
	}
	return Default_LsdConfigServerConfigT_MaxFileSize
}

func (m *LsdConfigServerConfigT) GetFileRotateInterval() uint64 {
	if m != nil && m.FileRotateInterval != nil {
		return *m.FileRotateInterval
	}
	return Default_LsdConfigServerConfigT_FileRotateInterval
}

func (m *LsdConfigServerConfigT) GetPerCategorySettings() []*LsdConfigServerConfigTServerSettingsT {
	if m != nil {
		return m.PerCategorySettings
	}
	return nil
}

func (m *LsdConfigServerConfigT) GetChunkOutput() bool {
	if m != nil && m.ChunkOutput != nil {
		return *m.ChunkOutput
	}
	return Default_LsdConfigServerConfigT_ChunkOutput
}

type LsdConfigServerConfigTServerSettingsT struct {
	Categories         []string `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
	MaxFileSize        *uint64  `protobuf:"varint,2,opt,name=max_file_size" json:"max_file_size,omitempty"`
	FileRotateInterval *uint64  `protobuf:"varint,3,opt,name=file_rotate_interval" json:"file_rotate_interval,omitempty"`
	ChunkOutput        *bool    `protobuf:"varint,4,opt,name=chunk_output,def=1" json:"chunk_output,omitempty"`
}

func (m *LsdConfigServerConfigTServerSettingsT) Reset()         { *m = LsdConfigServerConfigTServerSettingsT{} }
func (m *LsdConfigServerConfigTServerSettingsT) String() string { return proto.CompactTextString(m) }
func (*LsdConfigServerConfigTServerSettingsT) ProtoMessage()    {}

const Default_LsdConfigServerConfigTServerSettingsT_ChunkOutput bool = true

func (m *LsdConfigServerConfigTServerSettingsT) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *LsdConfigServerConfigTServerSettingsT) GetMaxFileSize() uint64 {
	if m != nil && m.MaxFileSize != nil {
		return *m.MaxFileSize
	}
	return 0
}

func (m *LsdConfigServerConfigTServerSettingsT) GetFileRotateInterval() uint64 {
	if m != nil && m.FileRotateInterval != nil {
		return *m.FileRotateInterval
	}
	return 0
}

func (m *LsdConfigServerConfigTServerSettingsT) GetChunkOutput() bool {
	if m != nil && m.ChunkOutput != nil {
		return *m.ChunkOutput
	}
	return Default_LsdConfigServerConfigTServerSettingsT_ChunkOutput
}

type LsdConfigClientConfigT struct {
	SourceDir          *string                                 `protobuf:"bytes,1,req,name=source_dir" json:"source_dir,omitempty"`
	Routing            []*LsdConfigClientConfigTRoutingConfigT `protobuf:"bytes,2,rep,name=routing" json:"routing,omitempty"`
	MaxFileSize        *uint64                                 `protobuf:"varint,3,opt,name=max_file_size,def=1000000" json:"max_file_size,omitempty"`
	OffsetsDb          *string                                 `protobuf:"bytes,4,req,name=offsets_db" json:"offsets_db,omitempty"`
	UsageCheckInterval *uint64                                 `protobuf:"varint,5,opt,name=usage_check_interval,def=60" json:"usage_check_interval,omitempty"`
	AlwaysFlock        *bool                                   `protobuf:"varint,6,opt,name=always_flock,def=0" json:"always_flock,omitempty"`
}

func (m *LsdConfigClientConfigT) Reset()         { *m = LsdConfigClientConfigT{} }
func (m *LsdConfigClientConfigT) String() string { return proto.CompactTextString(m) }
func (*LsdConfigClientConfigT) ProtoMessage()    {}

const Default_LsdConfigClientConfigT_MaxFileSize uint64 = 1000000
const Default_LsdConfigClientConfigT_UsageCheckInterval uint64 = 60
const Default_LsdConfigClientConfigT_AlwaysFlock bool = false

func (m *LsdConfigClientConfigT) GetSourceDir() string {
	if m != nil && m.SourceDir != nil {
		return *m.SourceDir
	}
	return ""
}

func (m *LsdConfigClientConfigT) GetRouting() []*LsdConfigClientConfigTRoutingConfigT {
	if m != nil {
		return m.Routing
	}
	return nil
}

func (m *LsdConfigClientConfigT) GetMaxFileSize() uint64 {
	if m != nil && m.MaxFileSize != nil {
		return *m.MaxFileSize
	}
	return Default_LsdConfigClientConfigT_MaxFileSize
}

func (m *LsdConfigClientConfigT) GetOffsetsDb() string {
	if m != nil && m.OffsetsDb != nil {
		return *m.OffsetsDb
	}
	return ""
}

func (m *LsdConfigClientConfigT) GetUsageCheckInterval() uint64 {
	if m != nil && m.UsageCheckInterval != nil {
		return *m.UsageCheckInterval
	}
	return Default_LsdConfigClientConfigT_UsageCheckInterval
}

func (m *LsdConfigClientConfigT) GetAlwaysFlock() bool {
	if m != nil && m.AlwaysFlock != nil {
		return *m.AlwaysFlock
	}
	return Default_LsdConfigClientConfigT_AlwaysFlock
}

type LsdConfigClientConfigTReceiverT struct {
	Addr   *string `protobuf:"bytes,1,req,name=addr" json:"addr,omitempty"`
	Weight *uint64 `protobuf:"varint,2,req,name=weight" json:"weight,omitempty"`
}

func (m *LsdConfigClientConfigTReceiverT) Reset()         { *m = LsdConfigClientConfigTReceiverT{} }
func (m *LsdConfigClientConfigTReceiverT) String() string { return proto.CompactTextString(m) }
func (*LsdConfigClientConfigTReceiverT) ProtoMessage()    {}

func (m *LsdConfigClientConfigTReceiverT) GetAddr() string {
	if m != nil && m.Addr != nil {
		return *m.Addr
	}
	return ""
}

func (m *LsdConfigClientConfigTReceiverT) GetWeight() uint64 {
	if m != nil && m.Weight != nil {
		return *m.Weight
	}
	return 0
}

type LsdConfigClientConfigTRoutingConfigT struct {
	Categories      []string                           `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
	Receivers       []*LsdConfigClientConfigTReceiverT `protobuf:"bytes,2,rep,name=receivers" json:"receivers,omitempty"`
	PrefixSharding  *bool                              `protobuf:"varint,3,opt,name=prefix_sharding,def=0" json:"prefix_sharding,omitempty"`
	CutPrefix       *bool                              `protobuf:"varint,4,opt,name=cut_prefix,def=0" json:"cut_prefix,omitempty"`
	PrefixDelimiter *string                            `protobuf:"bytes,5,opt,name=prefix_delimiter,def=:" json:"prefix_delimiter,omitempty"`
}

func (m *LsdConfigClientConfigTRoutingConfigT) Reset()         { *m = LsdConfigClientConfigTRoutingConfigT{} }
func (m *LsdConfigClientConfigTRoutingConfigT) String() string { return proto.CompactTextString(m) }
func (*LsdConfigClientConfigTRoutingConfigT) ProtoMessage()    {}

const Default_LsdConfigClientConfigTRoutingConfigT_PrefixSharding bool = false
const Default_LsdConfigClientConfigTRoutingConfigT_CutPrefix bool = false
const Default_LsdConfigClientConfigTRoutingConfigT_PrefixDelimiter string = ":"

func (m *LsdConfigClientConfigTRoutingConfigT) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *LsdConfigClientConfigTRoutingConfigT) GetReceivers() []*LsdConfigClientConfigTReceiverT {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *LsdConfigClientConfigTRoutingConfigT) GetPrefixSharding() bool {
	if m != nil && m.PrefixSharding != nil {
		return *m.PrefixSharding
	}
	return Default_LsdConfigClientConfigTRoutingConfigT_PrefixSharding
}

func (m *LsdConfigClientConfigTRoutingConfigT) GetCutPrefix() bool {
	if m != nil && m.CutPrefix != nil {
		return *m.CutPrefix
	}
	return Default_LsdConfigClientConfigTRoutingConfigT_CutPrefix
}

func (m *LsdConfigClientConfigTRoutingConfigT) GetPrefixDelimiter() string {
	if m != nil && m.PrefixDelimiter != nil {
		return *m.PrefixDelimiter
	}
	return Default_LsdConfigClientConfigTRoutingConfigT_PrefixDelimiter
}
