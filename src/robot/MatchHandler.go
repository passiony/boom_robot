package robot

import (
	protodef "ROBOT/src/network/protogo"
	"github.com/golang/protobuf/proto"
	"log"
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
		log.Println("匹配错误：", serverPacket.ErrorCode)
	}
}

func (l *MatchHander) OnRoomReady(robot *Robot, msg *protodef.Packet) {
	log.Println("匹配成功")
	robot.GameReady()
}
