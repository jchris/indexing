// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	common.proto
	index.proto
	partn_single.proto
	partn_tp.proto
	projector.proto

It has these top-level messages:
	Error
	Vbuckets
	Snapshot
	TsVb
	TsVbFull
	TsVbuuid
	FailoverLog
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Error message can be sent back as response or
// encapsulated in response packets.
type Error struct {
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// list of vbucket numbers
type Vbuckets struct {
	Vbnos            []uint32 `protobuf:"varint,1,rep,name=vbnos" json:"vbnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Vbuckets) Reset()         { *m = Vbuckets{} }
func (m *Vbuckets) String() string { return proto.CompactTextString(m) }
func (*Vbuckets) ProtoMessage()    {}

func (m *Vbuckets) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

// Start and end of DCP snapshot
type Snapshot struct {
	Start            *uint64 `protobuf:"varint,1,req,name=start" json:"start,omitempty"`
	End              *uint64 `protobuf:"varint,2,req,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}

func (m *Snapshot) GetStart() uint64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *Snapshot) GetEnd() uint64 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

// logical clock for a subset of vbuckets
type TsVb struct {
	Pool             *string  `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string  `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32 `protobuf:"varint,3,rep,name=vbnos" json:"vbnos,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,4,rep,name=seqnos" json:"seqnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TsVb) Reset()         { *m = TsVb{} }
func (m *TsVb) String() string { return proto.CompactTextString(m) }
func (*TsVb) ProtoMessage()    {}

func (m *TsVb) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *TsVb) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *TsVb) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

func (m *TsVb) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

// logical clock for full set of vbuckets, starting from 0 to MaxVbucket.
type TsVbFull struct {
	Pool             *string  `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string  `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,3,rep,name=seqnos" json:"seqnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TsVbFull) Reset()         { *m = TsVbFull{} }
func (m *TsVbFull) String() string { return proto.CompactTextString(m) }
func (*TsVbFull) ProtoMessage()    {}

func (m *TsVbFull) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *TsVbFull) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *TsVbFull) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

// logical clock for a subset of vbuckets along with branch-id and snapshot
// information.
type TsVbuuid struct {
	Pool             *string     `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string     `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32    `protobuf:"varint,3,rep,name=vbnos" json:"vbnos,omitempty"`
	Seqnos           []uint64    `protobuf:"varint,4,rep,name=seqnos" json:"seqnos,omitempty"`
	Vbuuids          []uint64    `protobuf:"varint,5,rep,name=vbuuids" json:"vbuuids,omitempty"`
	Snapshots        []*Snapshot `protobuf:"bytes,6,rep,name=snapshots" json:"snapshots,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *TsVbuuid) Reset()         { *m = TsVbuuid{} }
func (m *TsVbuuid) String() string { return proto.CompactTextString(m) }
func (*TsVbuuid) ProtoMessage()    {}

func (m *TsVbuuid) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *TsVbuuid) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *TsVbuuid) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

func (m *TsVbuuid) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

func (m *TsVbuuid) GetVbuuids() []uint64 {
	if m != nil {
		return m.Vbuuids
	}
	return nil
}

func (m *TsVbuuid) GetSnapshots() []*Snapshot {
	if m != nil {
		return m.Snapshots
	}
	return nil
}

// failover log for a vbucket.
type FailoverLog struct {
	Vbno             *uint32  `protobuf:"varint,1,req,name=vbno" json:"vbno,omitempty"`
	Vbuuids          []uint64 `protobuf:"varint,2,rep,name=vbuuids" json:"vbuuids,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,3,rep,name=seqnos" json:"seqnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FailoverLog) Reset()         { *m = FailoverLog{} }
func (m *FailoverLog) String() string { return proto.CompactTextString(m) }
func (*FailoverLog) ProtoMessage()    {}

func (m *FailoverLog) GetVbno() uint32 {
	if m != nil && m.Vbno != nil {
		return *m.Vbno
	}
	return 0
}

func (m *FailoverLog) GetVbuuids() []uint64 {
	if m != nil {
		return m.Vbuuids
	}
	return nil
}

func (m *FailoverLog) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

func init() {
}
