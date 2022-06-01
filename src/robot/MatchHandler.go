package robot

import (
	"ROBOT/src/log"
	protodef "ROBOT/src/network/protogo"
	"github.com/golang/protobuf/proto"
)

type MatchHander struct {
}

func (l *MatchHander) Init(events map[int32]func(robot *Robot, msg *protodef.Packet)) {
	events[protodef.MATCH_NTF_ROOM_READY] = l.OnRoomReady
	events[protodef.MATCH_RSP_MATCH_GAME] = l.OnRspMatch
}

func (l *MatchHander) OnRspMatch(robot *Robot, msg *protodef.Packet) {
	serverPacket := &protodef.RspMatchGame{}
	err := proto.Unmarshal(msg.Data, serverPacket)
	if err != nil {
		return
	}
	if serverPacket.ErrorCode != 0 {
		log.Info("匹配错误：", serverPacket.ErrorCode)
	}
}

func (l *MatchHander) OnRoomReady(robot *Robot, msg *protodef.Packet) {
	log.Info("匹配成功")
	serverPacket := &protodef.NtfRoomReady{}
	err := proto.Unmarshal(msg.Data, serverPacket)
	if err != nil {
		return
	}

	robot.OnEnterRoom(serverPacket.Players)
	robot.GameReady()
}
