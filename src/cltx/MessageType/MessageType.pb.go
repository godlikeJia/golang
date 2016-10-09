// Code generated by protoc-gen-go.
// source: MessageType.proto
// DO NOT EDIT!

/*
Package MessageType is a generated protocol buffer package.

It is generated from these files:
	MessageType.proto

It has these top-level messages:
	ErrorInfo
	IdInfo
	RichInfo
	TipsMsgId
*/
package MessageType

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type ErrorInfo struct {
	ErrCode          *int64 `protobuf:"varint,1,opt,name=err_code,def=2147483648" json:"err_code,omitempty"`
	ErrMsg           []byte `protobuf:"bytes,2,opt,name=err_msg" json:"err_msg,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ErrorInfo) Reset()         { *m = ErrorInfo{} }
func (m *ErrorInfo) String() string { return proto.CompactTextString(m) }
func (*ErrorInfo) ProtoMessage()    {}

const Default_ErrorInfo_ErrCode int64 = 2147483648

func (m *ErrorInfo) GetErrCode() int64 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return Default_ErrorInfo_ErrCode
}

func (m *ErrorInfo) GetErrMsg() []byte {
	if m != nil {
		return m.ErrMsg
	}
	return nil
}

type IdInfo struct {
	Content          []byte      `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Extra            []*RichInfo `protobuf:"bytes,2,rep,name=extra" json:"extra,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *IdInfo) Reset()         { *m = IdInfo{} }
func (m *IdInfo) String() string { return proto.CompactTextString(m) }
func (*IdInfo) ProtoMessage()    {}

func (m *IdInfo) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *IdInfo) GetExtra() []*RichInfo {
	if m != nil {
		return m.Extra
	}
	return nil
}

type RichInfo struct {
	Begin            *int32 `protobuf:"varint,1,req,name=begin" json:"begin,omitempty"`
	End              *int32 `protobuf:"varint,2,req,name=end" json:"end,omitempty"`
	Tag              *int32 `protobuf:"varint,3,opt,name=tag" json:"tag,omitempty"`
	TagId            *int32 `protobuf:"varint,4,opt,name=tag_id" json:"tag_id,omitempty"`
	TagExtra         []byte `protobuf:"bytes,5,opt,name=tag_extra" json:"tag_extra,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RichInfo) Reset()         { *m = RichInfo{} }
func (m *RichInfo) String() string { return proto.CompactTextString(m) }
func (*RichInfo) ProtoMessage()    {}

func (m *RichInfo) GetBegin() int32 {
	if m != nil && m.Begin != nil {
		return *m.Begin
	}
	return 0
}

func (m *RichInfo) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

func (m *RichInfo) GetTag() int32 {
	if m != nil && m.Tag != nil {
		return *m.Tag
	}
	return 0
}

func (m *RichInfo) GetTagId() int32 {
	if m != nil && m.TagId != nil {
		return *m.TagId
	}
	return 0
}

func (m *RichInfo) GetTagExtra() []byte {
	if m != nil {
		return m.TagExtra
	}
	return nil
}

type TipsMsgId struct {
	TipsType         *int32 `protobuf:"varint,1,req,name=tipsType" json:"tipsType,omitempty"`
	PlaceId          *int64 `protobuf:"varint,2,opt,name=placeId" json:"placeId,omitempty"`
	TipContent       []byte `protobuf:"bytes,3,opt,name=tipContent" json:"tipContent,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TipsMsgId) Reset()         { *m = TipsMsgId{} }
func (m *TipsMsgId) String() string { return proto.CompactTextString(m) }
func (*TipsMsgId) ProtoMessage()    {}

func (m *TipsMsgId) GetTipsType() int32 {
	if m != nil && m.TipsType != nil {
		return *m.TipsType
	}
	return 0
}

func (m *TipsMsgId) GetPlaceId() int64 {
	if m != nil && m.PlaceId != nil {
		return *m.PlaceId
	}
	return 0
}

func (m *TipsMsgId) GetTipContent() []byte {
	if m != nil {
		return m.TipContent
	}
	return nil
}

func init() {
}
