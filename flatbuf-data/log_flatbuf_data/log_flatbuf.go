// Code generated by the FlatBuffers compiler. DO NOT EDIT.
package log_flatbuf_data

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type LogMessage struct {
	_tab flatbuffers.Table
}

func GetRootAsLogMessage(buf []byte, offset flatbuffers.UOffsetT) *LogMessage {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &LogMessage{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *LogMessage) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *LogMessage) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *LogMessage) ContainerId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LogMessage) ContainerName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *LogMessage) Payload() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func LogMessageStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func LogMessageAddContainerId(builder *flatbuffers.Builder, ContainerId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ContainerId), 0)
}
func LogMessageAddContainerName(builder *flatbuffers.Builder, ContainerName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(ContainerName), 0)
}
func LogMessageAddPayload(builder *flatbuffers.Builder, Payload flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(Payload), 0)
}
func LogMessageEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
