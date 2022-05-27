package robot

import (
	protodef "ROBOT/src/network/protogo"
	"log"
)

type GameHander struct {
}

func (l *GameHander) Init(events map[int32]func(robot *Robot, msg *protodef.Packet)) {
	events[protodef.GAME_NTF_GAME_RECONNECT] = l.OnGameReconnect
	events[protodef.GAME_NTF_RECONNECT_FINISH] = l.OnReconnectFinish
	events[protodef.GAME_NTF_GAME_START] = l.OnGameStart
}

func (l *GameHander) OnGameReconnect(robot *Robot, msg *protodef.Packet) {
	log.Println("重连 1")
	robot.ReconnectFinish()
}

func (l *GameHander) OnReconnectFinish(robot *Robot, msg *protodef.Packet) {
	log.Println("重连 2")
	robot.OnGameStart()
}

func (l *GameHander) OnGameStart(robot *Robot, msg *protodef.Packet) {
	log.Println("游戏开始")
	robot.OnGameStart()
}
