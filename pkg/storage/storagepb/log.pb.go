// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/storagepb/log.proto

package storagepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import time "time"
import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RangeLogEventType int32

const (
	// These are lower case to maintain compatibility with how they were
	// originally stored.
	// Split is the event type recorded when a range splits.
	RangeLogEventType_split RangeLogEventType = 0
	// Merge is the event type recorded when a range merges.
	RangeLogEventType_merge RangeLogEventType = 3
	// Add is the event type recorded when a range adds a new replica.
	RangeLogEventType_add RangeLogEventType = 1
	// Remove is the event type recorded when a range removed an existing replica.
	RangeLogEventType_remove RangeLogEventType = 2
)

var RangeLogEventType_name = map[int32]string{
	0: "split",
	3: "merge",
	1: "add",
	2: "remove",
}
var RangeLogEventType_value = map[string]int32{
	"split":  0,
	"merge":  3,
	"add":    1,
	"remove": 2,
}

func (x RangeLogEventType) String() string {
	return proto.EnumName(RangeLogEventType_name, int32(x))
}
func (RangeLogEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_log_ddb6a32f80f69138, []int{0}
}

type RangeLogEvent struct {
	Timestamp            time.Time                                            `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	RangeID              github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,2,opt,name=range_id,json=rangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"range_id,omitempty"`
	StoreID              github_com_cockroachdb_cockroach_pkg_roachpb.StoreID `protobuf:"varint,3,opt,name=store_id,json=storeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.StoreID" json:"store_id,omitempty"`
	EventType            RangeLogEventType                                    `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3,enum=cockroach.storage.RangeLogEventType" json:"event_type,omitempty"`
	OtherRangeID         github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,5,opt,name=other_range_id,json=otherRangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"other_range_id,omitempty"`
	Info                 *RangeLogEvent_Info                                  `protobuf:"bytes,6,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
}

func (m *RangeLogEvent) Reset()         { *m = RangeLogEvent{} }
func (m *RangeLogEvent) String() string { return proto.CompactTextString(m) }
func (*RangeLogEvent) ProtoMessage()    {}
func (*RangeLogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_ddb6a32f80f69138, []int{0}
}
func (m *RangeLogEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RangeLogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *RangeLogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeLogEvent.Merge(dst, src)
}
func (m *RangeLogEvent) XXX_Size() int {
	return m.Size()
}
func (m *RangeLogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeLogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RangeLogEvent proto.InternalMessageInfo

type RangeLogEvent_Info struct {
	UpdatedDesc          *roachpb.RangeDescriptor   `protobuf:"bytes,1,opt,name=updated_desc,json=updatedDesc,proto3" json:"UpdatedDesc,omitempty"`
	NewDesc              *roachpb.RangeDescriptor   `protobuf:"bytes,2,opt,name=new_desc,json=newDesc,proto3" json:"NewDesc,omitempty"`
	RemovedDesc          *roachpb.RangeDescriptor   `protobuf:"bytes,7,opt,name=removed_desc,json=removedDesc,proto3" json:"RemovedDesc,omitempty"`
	AddedReplica         *roachpb.ReplicaDescriptor `protobuf:"bytes,3,opt,name=added_replica,json=addedReplica,proto3" json:"AddReplica,omitempty"`
	RemovedReplica       *roachpb.ReplicaDescriptor `protobuf:"bytes,4,opt,name=removed_replica,json=removedReplica,proto3" json:"RemovedReplica,omitempty"`
	Reason               RangeLogEventReason        `protobuf:"bytes,5,opt,name=reason,proto3,casttype=RangeLogEventReason" json:"Reason,omitempty"`
	Details              string                     `protobuf:"bytes,6,opt,name=details,proto3" json:"Details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
}

func (m *RangeLogEvent_Info) Reset()         { *m = RangeLogEvent_Info{} }
func (m *RangeLogEvent_Info) String() string { return proto.CompactTextString(m) }
func (*RangeLogEvent_Info) ProtoMessage()    {}
func (*RangeLogEvent_Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_ddb6a32f80f69138, []int{0, 0}
}
func (m *RangeLogEvent_Info) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RangeLogEvent_Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *RangeLogEvent_Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RangeLogEvent_Info.Merge(dst, src)
}
func (m *RangeLogEvent_Info) XXX_Size() int {
	return m.Size()
}
func (m *RangeLogEvent_Info) XXX_DiscardUnknown() {
	xxx_messageInfo_RangeLogEvent_Info.DiscardUnknown(m)
}

var xxx_messageInfo_RangeLogEvent_Info proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RangeLogEvent)(nil), "cockroach.storage.RangeLogEvent")
	proto.RegisterType((*RangeLogEvent_Info)(nil), "cockroach.storage.RangeLogEvent.Info")
	proto.RegisterEnum("cockroach.storage.RangeLogEventType", RangeLogEventType_name, RangeLogEventType_value)
}
func (m *RangeLogEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeLogEvent) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLog(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.RangeID != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RangeID))
	}
	if m.StoreID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.StoreID))
	}
	if m.EventType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.EventType))
	}
	if m.OtherRangeID != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.OtherRangeID))
	}
	if m.Info != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.Info.Size()))
		n2, err := m.Info.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *RangeLogEvent_Info) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeLogEvent_Info) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UpdatedDesc != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.UpdatedDesc.Size()))
		n3, err := m.UpdatedDesc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.NewDesc != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.NewDesc.Size()))
		n4, err := m.NewDesc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.AddedReplica != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.AddedReplica.Size()))
		n5, err := m.AddedReplica.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.RemovedReplica != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RemovedReplica.Size()))
		n6, err := m.RemovedReplica.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if len(m.Details) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintLog(dAtA, i, uint64(len(m.Details)))
		i += copy(dAtA[i:], m.Details)
	}
	if m.RemovedDesc != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintLog(dAtA, i, uint64(m.RemovedDesc.Size()))
		n7, err := m.RemovedDesc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func encodeVarintLog(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RangeLogEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovLog(uint64(l))
	if m.RangeID != 0 {
		n += 1 + sovLog(uint64(m.RangeID))
	}
	if m.StoreID != 0 {
		n += 1 + sovLog(uint64(m.StoreID))
	}
	if m.EventType != 0 {
		n += 1 + sovLog(uint64(m.EventType))
	}
	if m.OtherRangeID != 0 {
		n += 1 + sovLog(uint64(m.OtherRangeID))
	}
	if m.Info != nil {
		l = m.Info.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func (m *RangeLogEvent_Info) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UpdatedDesc != nil {
		l = m.UpdatedDesc.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.NewDesc != nil {
		l = m.NewDesc.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.AddedReplica != nil {
		l = m.AddedReplica.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	if m.RemovedReplica != nil {
		l = m.RemovedReplica.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovLog(uint64(l))
	}
	if m.RemovedDesc != nil {
		l = m.RemovedDesc.Size()
		n += 1 + l + sovLog(uint64(l))
	}
	return n
}

func sovLog(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLog(x uint64) (n int) {
	return sovLog(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RangeLogEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RangeLogEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeLogEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= (github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EventType", wireType)
			}
			m.EventType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EventType |= (RangeLogEventType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OtherRangeID", wireType)
			}
			m.OtherRangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OtherRangeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Info", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Info == nil {
				m.Info = &RangeLogEvent_Info{}
			}
			if err := m.Info.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RangeLogEvent_Info) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLog
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Info: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Info: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedDesc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatedDesc == nil {
				m.UpdatedDesc = &roachpb.RangeDescriptor{}
			}
			if err := m.UpdatedDesc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewDesc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NewDesc == nil {
				m.NewDesc = &roachpb.RangeDescriptor{}
			}
			if err := m.NewDesc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AddedReplica == nil {
				m.AddedReplica = &roachpb.ReplicaDescriptor{}
			}
			if err := m.AddedReplica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemovedReplica == nil {
				m.RemovedReplica = &roachpb.ReplicaDescriptor{}
			}
			if err := m.RemovedReplica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = RangeLogEventReason(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemovedDesc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLog
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLog
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RemovedDesc == nil {
				m.RemovedDesc = &roachpb.RangeDescriptor{}
			}
			if err := m.RemovedDesc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLog(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLog
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLog(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLog
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLog
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthLog
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLog
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipLog(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthLog = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLog   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/storagepb/log.proto", fileDescriptor_log_ddb6a32f80f69138) }

var fileDescriptor_log_ddb6a32f80f69138 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x9b, 0xb5, 0x5b, 0x5a, 0xaf, 0x1b, 0x9d, 0xd9, 0x50, 0x29, 0x28, 0xa9, 0x26, 0x10,
	0x15, 0xa0, 0x44, 0x1a, 0x5c, 0xb8, 0x20, 0xd1, 0x0d, 0x89, 0x4a, 0x08, 0xa4, 0x30, 0x2e, 0x1c,
	0x28, 0x6e, 0xfc, 0x2e, 0x8b, 0xd6, 0xc4, 0x91, 0xe3, 0x6e, 0xda, 0xb7, 0xd8, 0xc7, 0xda, 0x71,
	0x47, 0x4e, 0x01, 0xba, 0x0b, 0xea, 0x81, 0x0f, 0xb0, 0x13, 0xb2, 0xe3, 0xfe, 0x5b, 0x0f, 0x30,
	0x38, 0xf5, 0xcd, 0xe3, 0xf7, 0xfd, 0x39, 0x7e, 0x1e, 0x37, 0xe8, 0x5e, 0x2a, 0x18, 0x27, 0x01,
	0xb8, 0xfa, 0x37, 0xe9, 0xb9, 0x7d, 0x16, 0x38, 0x09, 0x67, 0x82, 0xe1, 0x0d, 0x9f, 0xf9, 0x47,
	0x9c, 0x11, 0xff, 0xd0, 0xd1, 0xcb, 0x8d, 0x3b, 0xea, 0x31, 0xe9, 0xb9, 0x11, 0x08, 0x42, 0x89,
	0x20, 0x79, 0x6b, 0x63, 0x33, 0x60, 0x01, 0x53, 0xa5, 0x2b, 0x2b, 0xad, 0xda, 0x01, 0x63, 0x41,
	0x1f, 0x5c, 0xf5, 0xd4, 0x1b, 0x1c, 0xb8, 0x22, 0x8c, 0x20, 0x15, 0x24, 0x4a, 0xf2, 0x86, 0xed,
	0x5f, 0x65, 0xb4, 0xe6, 0x91, 0x38, 0x80, 0xb7, 0x2c, 0x78, 0x7d, 0x0c, 0xb1, 0xc0, 0x6d, 0x54,
	0x99, 0x34, 0xd5, 0x8d, 0xa6, 0xd1, 0x5a, 0xdd, 0x69, 0x38, 0x39, 0xc6, 0x19, 0x63, 0x9c, 0xfd,
	0x71, 0x47, 0xbb, 0x7c, 0x9e, 0xd9, 0x85, 0xb3, 0x6f, 0xb6, 0xe1, 0x4d, 0xc7, 0xf0, 0x67, 0x54,
	0xe6, 0x12, 0xda, 0x0d, 0x69, 0x7d, 0xa9, 0x69, 0xb4, 0x8a, 0xed, 0xdd, 0x61, 0x66, 0x9b, 0x6a,
	0xa3, 0xce, 0xde, 0x55, 0x66, 0x3f, 0x0f, 0x42, 0x71, 0x38, 0xe8, 0x39, 0x3e, 0x8b, 0xdc, 0xc9,
	0x19, 0x69, 0x6f, 0x5a, 0xbb, 0xc9, 0x51, 0xe0, 0xea, 0xa3, 0x3a, 0x7a, 0xce, 0x33, 0x15, 0xb4,
	0x43, 0x25, 0x5f, 0xfa, 0xa1, 0xf8, 0xc5, 0xa6, 0xd1, 0x5a, 0xce, 0xf9, 0x1f, 0xa4, 0xf6, 0x0f,
	0x7c, 0x3d, 0xe7, 0x99, 0x0a, 0xda, 0xa1, 0x78, 0x17, 0x21, 0x90, 0x66, 0x74, 0xc5, 0x69, 0x02,
	0xf5, 0x52, 0xd3, 0x68, 0xad, 0xef, 0x3c, 0x70, 0x16, 0xc2, 0x70, 0xe6, 0x9c, 0xdb, 0x3f, 0x4d,
	0xc0, 0xab, 0xc0, 0xb8, 0xc4, 0x31, 0x5a, 0x67, 0xe2, 0x10, 0x78, 0x77, 0x62, 0xc5, 0xb2, 0xb2,
	0xe2, 0xcd, 0x30, 0xb3, 0xab, 0xef, 0xe5, 0xca, 0xff, 0xfa, 0x51, 0x65, 0x53, 0x0a, 0xc5, 0x2f,
	0x50, 0x29, 0x8c, 0x0f, 0x58, 0x7d, 0x45, 0x65, 0xf6, 0xf0, 0x4f, 0xaf, 0xeb, 0x74, 0xe2, 0x03,
	0xe6, 0xa9, 0x91, 0xc6, 0xcf, 0x12, 0x2a, 0xc9, 0x47, 0xfc, 0x05, 0x55, 0x07, 0x09, 0x25, 0x02,
	0x68, 0x97, 0x42, 0xea, 0xeb, 0xfc, 0xb7, 0x67, 0x58, 0x73, 0xef, 0xb0, 0x07, 0xa9, 0xcf, 0xc3,
	0x44, 0x30, 0xde, 0xbe, 0x3b, 0xca, 0xec, 0xad, 0x8f, 0xf9, 0xac, 0x94, 0x9f, 0xb2, 0x28, 0x14,
	0x10, 0x25, 0xe2, 0xd4, 0x5b, 0x1d, 0x4c, 0x65, 0xbc, 0x8f, 0xca, 0x31, 0x9c, 0xe4, 0xf4, 0xa5,
	0xbf, 0xa6, 0x6f, 0x8d, 0x32, 0x7b, 0xe3, 0x1d, 0x9c, 0x5c, 0x23, 0x9b, 0x71, 0x2e, 0x61, 0x1f,
	0xad, 0x11, 0x4a, 0x81, 0x76, 0x39, 0x24, 0xfd, 0xd0, 0x27, 0xea, 0x56, 0xac, 0xce, 0x65, 0x36,
	0x41, 0xe7, 0x1d, 0x33, 0xf0, 0xfa, 0x28, 0xb3, 0x37, 0x5f, 0x51, 0xaa, 0x57, 0x66, 0xf8, 0x55,
	0x05, 0xd5, 0x3a, 0x3e, 0x42, 0xb7, 0x38, 0x44, 0xec, 0x78, 0x66, 0x9b, 0xd2, 0x0d, 0xb6, 0xb9,
	0x3f, 0xca, 0xec, 0xba, 0x97, 0x03, 0x16, 0xb7, 0x5a, 0xe7, 0x73, 0x2b, 0x78, 0x17, 0xad, 0x70,
	0x20, 0x29, 0x8b, 0xd5, 0xad, 0xa9, 0xb4, 0x9f, 0x8c, 0x32, 0xbb, 0xe6, 0x29, 0x65, 0x3a, 0x75,
	0x95, 0xd9, 0xb7, 0xe7, 0x42, 0xcd, 0x1b, 0x3c, 0x3d, 0x8a, 0x5d, 0x64, 0x52, 0x10, 0x24, 0xec,
	0xa7, 0xea, 0x56, 0x54, 0x72, 0x1f, 0xf7, 0x72, 0x69, 0xd6, 0x47, 0xdd, 0x25, 0xf3, 0x1f, 0x1f,
	0x51, 0x25, 0x64, 0xde, 0x2c, 0x7f, 0x7d, 0xba, 0xeb, 0xf9, 0xf3, 0xa9, 0xfc, 0xf8, 0x25, 0xda,
	0x58, 0xf8, 0xd7, 0xe0, 0x0a, 0x5a, 0x4e, 0x93, 0x7e, 0x28, 0x6a, 0x05, 0x59, 0x46, 0xc0, 0x03,
	0xa8, 0x15, 0xb1, 0x89, 0x8a, 0x84, 0xd2, 0x9a, 0x81, 0x91, 0xf4, 0x42, 0x22, 0x6a, 0x4b, 0xed,
	0x47, 0xe7, 0x3f, 0xac, 0xc2, 0xf9, 0xd0, 0x32, 0x2e, 0x86, 0x96, 0xf1, 0x75, 0x68, 0x19, 0xdf,
	0x87, 0x96, 0x71, 0x76, 0x69, 0x15, 0x2e, 0x2e, 0xad, 0xc2, 0xa7, 0xca, 0xe4, 0x2b, 0xda, 0x5b,
	0x51, 0x1f, 0xab, 0x67, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x6f, 0x5c, 0x14, 0x61, 0x05,
	0x00, 0x00,
}
