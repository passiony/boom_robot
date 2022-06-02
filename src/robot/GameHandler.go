package robot

import (
	"ROBOT/src/log"
	protodef "ROBOT/src/network/protogo"
	"github.com/golang/protobuf/proto"
)

type GameHander struct {
}

func (l *GameHander) Init(events map[int32]func(robot *Robot, msg *protodef.Packet)) {
	events[protodef.GAME_NTF_GAME_RECONNECT] = l.OnGameReconnect
	events[protodef.GAME_NTF_RECONNECT_FINISH] = l.OnReconnectFinish
	events[protodef.GAME_NTF_GAME_START] = l.OnGameStart
	events[protodef.GAME_NTF_GAME_OVER] = l.OnGameOver
}

func (l *GameHander) OnGameReconnect(robot *Robot, msg *protodef.Packet) {
	log.Info("重连 1")
	robot.ReconnectFinish()
}

func (l *GameHander) OnReconnectFinish(robot *Robot, msg *protodef.Packet) {
	log.Info("重连 2")
	serverPacket := &protodef.NtfReconnectFinish{}
	err := proto.Unmarshal(msg.Data, serverPacket)
	if err != nil {
		return
	}
	robot.OnEnterRoom(serverPacket.Entities.Players)
	robot.OnGameStart()
}

func (l *GameHander) OnGameStart(robot *Robot, msg *protodef.Packet) {
	log.Info("游戏开始")
	robot.OnGameStart()
}

func (l *GameHander) OnGameOver(robot *Robot, msg *protodef.Packet) {
	log.Info("游戏结束")
	robot.OnGameOver()
}
