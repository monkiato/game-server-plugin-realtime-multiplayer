// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package multiplayer

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type CreateMatchInfo struct {
	_tab flatbuffers.Table
}

func GetRootAsCreateMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *CreateMatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CreateMatchInfo{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsCreateMatchInfo(buf []byte, offset flatbuffers.UOffsetT) *CreateMatchInfo {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &CreateMatchInfo{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *CreateMatchInfo) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *CreateMatchInfo) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *CreateMatchInfo) PlayerId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func CreateMatchInfoStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func CreateMatchInfoAddPlayerId(builder *flatbuffers.Builder, playerId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(playerId), 0)
}
func CreateMatchInfoEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
