package robot

import (
	"ROBOT/src/log"
	protodef "ROBOT/src/network/protogo"
	"github.com/golang/protobuf/proto"
)

type LoginHander struct {
}

func (l *LoginHander) Init(events map[int32]func(robot *Robot, msg *protodef.Packet)) {
	events[protodef.LOGIN_RSP_LOGIN] = l.OnLogin
}

func (l *LoginHander) OnLogin(robot *Robot, msg *protodef.Packet) {
	log.Info("登录成功")

	serverPacket := &protodef.RspLogin{}
	err := proto.Unmarshal(msg.Data, serverPacket)
	if err != nil {
		return
	}

	if serverPacket.ErrorCode != 0 {
		log.Info("登录错误：", serverPacket.ErrorCode)
	}

	robot.OnLoginSuccess(serverPacket)
	if serverPacket.GameState == 0 {
		robot.MatchGame()
	} else {
		robot.Reconnect()
	}
}
