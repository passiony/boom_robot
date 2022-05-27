package mytcp

import (
	protodef "ROBOT/src/network/protogo"
	"encoding/binary"
	"errors"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/util"
	"io"
	"net"
)

const (
	bodySize  = 4 // 包体大小字段
	msgIDSize = 2
)

var (
	ErrMaxPacket  = errors.New("packet over size")
	ErrMinPacket  = errors.New("packet short size")
	ErrShortMsgID = errors.New("short msgid")
)

type MyTCPMessageTransmitter struct {
}

type mysocketOpt interface {
	MaxPacketSize() int
	ApplySocketReadTimeout(conn net.Conn, callback func())
	ApplySocketWriteTimeout(conn net.Conn, callback func())
}

func (MyTCPMessageTransmitter) OnRecvMessage(ses cellnet.Session) (msg interface{}, err error) {
	reader, ok := ses.Raw().(io.Reader)

	// 转换错误，或者连接已经关闭时退出
	if !ok || reader == nil {
		return nil, nil
	}
	opt := ses.Peer().(mysocketOpt)
	if conn, ok := reader.(net.Conn); ok {
		// 有读超时时，设置超时
		opt.ApplySocketReadTimeout(conn, func() {

			msg, err = MyRecvLTVPacket(reader, opt.MaxPacketSize())

		})
	}
	return
}

func MyRecvLTVPacket(reader io.Reader, maxPacketSize int) (msg interface{}, err error) {

	// Size为uint16，占2字节
	var sizeBuffer = make([]byte, bodySize)

	// 持续读取Size直到读到为止
	_, err = io.ReadFull(reader, sizeBuffer)

	// 发生错误时返回
	if err != nil {
		return
	}

	if len(sizeBuffer) < bodySize {
		return nil, ErrMinPacket
	}

	// 用小端格式读取Size
	size := binary.LittleEndian.Uint32(sizeBuffer)

	if maxPacketSize > 0 && size >= uint32(maxPacketSize) {
		return nil, ErrMaxPacket
	}

	// 分配包体大小
	body := make([]byte, size)

	// 读取包体数据
	_, err = io.ReadFull(reader, body)

	// 发生错误时返回
	if err != nil {
		return
	}
	if len(body) < msgIDSize {
		return nil, ErrShortMsgID
	}

	cmdID := binary.LittleEndian.Uint16(body)

	packet := &protodef.Packet{
		CmdId: int32(cmdID),
		Data:  body[msgIDSize:],
	}
	return packet, nil
}

func (MyTCPMessageTransmitter) OnSendMessage(ses cellnet.Session, msg interface{}) (err error) {

	writer, ok := ses.Raw().(io.Writer)

	// 转换错误，或者连接已经关闭时退出
	if !ok || writer == nil {
		return nil
	}

	opt := ses.Peer().(mysocketOpt)

	// 有写超时时，设置超时
	opt.ApplySocketWriteTimeout(writer.(net.Conn), func() {

		err = MySendLTVPacket(writer, ses.(cellnet.ContextSet), msg)

	})

	return
}

// 发送Length-Type-Value格式的封包流程
func MySendLTVPacket(writer io.Writer, ctx cellnet.ContextSet, data interface{}) error {

	msgData, ok := data.(*protodef.Packet)
	if !ok {
		return nil
	}

	pkt := make([]byte, bodySize+msgIDSize+len(msgData.Data))

	// Length
	binary.LittleEndian.PutUint32(pkt, uint32(msgIDSize+len(msgData.Data)))

	binary.LittleEndian.PutUint16(pkt[bodySize:], uint16(msgData.CmdId))

	// Value
	copy(pkt[bodySize+msgIDSize:], msgData.Data)

	// 将数据写入Socket
	err := util.WriteFull(writer, pkt)

	return err
}
