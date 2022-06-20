package robot

import (
	"ROBOT/src/config"
	"ROBOT/src/log"
	protodef "ROBOT/src/network/protogo"
	"errors"
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	"github.com/golang/protobuf/proto"
	"time"
)

type RobotNetwork struct {
	robot   *Robot
	session cellnet.Session
	events  map[int32]func(robot *Robot, msg *protodef.Packet)
	login   *LoginHander
	match   *MatchHander
	game    *GameHander
}

func NewNetwork(r *Robot) *RobotNetwork {
	n := &RobotNetwork{}
	n.robot = r
	n.events = map[int32]func(robot *Robot, msg *protodef.Packet){}

	n.login = &LoginHander{}
	n.match = &MatchHander{}
	n.game = &GameHander{}
	n.login.Init(n.events)
	n.match.Init(n.events)
	n.game.Init(n.events)

	n.connect(config.GetConfig().Server)
	return n
}

func (n *RobotNetwork) connect(netCfg *config.NetConfig) {
	addr := fmt.Sprintf("%s:%d", netCfg.Ip, netCfg.Port)
	queue := cellnet.NewEventQueue()
	p := peer.NewGenericPeer("tcp.Connector", "client", addr, queue)
	p.(cellnet.TCPConnector).SetReconnectDuration(time.Second)
	proc.BindProcessorHandler(p, "tcp.myltv", func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			n.onConnect(ev.Session())
		case *cellnet.SessionClosed:
			n.onDisconnect(ev.Session())
		case *protodef.Packet:
			n.onServerPacket(msg)
		}
	})
	p.Start()
	queue.StartLoop()
}

func (n *RobotNetwork) onConnect(session cellnet.Session) {
	n.session = session
	log.Info("onConnect:%d", n.robot.PlayerId)
	n.robot.Login()
}
func (n *RobotNetwork) onDisconnect(session cellnet.Session) {
	log.Info("onDisconnect:%d", n.robot.PlayerId)
	n.robot.Disconnect()
}

func (n *RobotNetwork) onServerPacket(msg *protodef.Packet) {
	function := n.events[msg.CmdId]
	if function != nil {
		function(n.robot, msg)
	}
}

func GetProtoData(msg interface{}) ([]byte, error) {
	if m, ok := msg.(proto.Message); ok {
		return proto.Marshal(m)
	}
	return nil, errors.New("获取proto失败")
}

func (n *RobotNetwork) sendMsg(cmdId int32, msg interface{}) {
	bytes, err := GetProtoData(msg)
	if err != nil {
		log.Info("%d sendMsg %s", err.Error())
		return
	}

	packet := &protodef.Packet{
		CmdId: cmdId,
		Data:  bytes,
	}
	if n.session != nil {
		n.session.Send(packet)
	}
}

func (n *RobotNetwork) ReqLogin(req *protodef.ReqLogin) {
	n.sendMsg(protodef.LOGIN_REQ_LOGIN, req)
}

func (n *RobotNetwork) ReqCreateLeague(req *protodef.ReqCreateLeague) {
	n.sendMsg(protodef.LEAGUE_REQ_CREATE_LEAGUE, req)
}

func (n *RobotNetwork) ReqHeart(req *protodef.ReqHeart) {
	n.sendMsg(protodef.LOGIN_REQ_HEART, req)
}

func (n *RobotNetwork) ReqMatchGame(req *protodef.ReqMatchGame) {
	n.sendMsg(protodef.MATCH_REQ_MATCH_GAME, req)
}

func (n *RobotNetwork) ReqReconnect(req *protodef.ReqGameReconnect) {
	n.sendMsg(protodef.GAME_REQ_GAME_RECONNECT, req)
}

func (n *RobotNetwork) ReqReconnectFinish(req *protodef.ReqReconnectFinish) {
	n.sendMsg(protodef.GAME_REQ_RECONNECT_FINISH, req)
}

func (n *RobotNetwork) ReqGameStart(req *protodef.ReqGameReady) {
	n.sendMsg(protodef.GAME_REQ_GAME_READY, req)
}

func (n *RobotNetwork) ReqSelectCards(req *protodef.ReqSelectCards) {
	n.sendMsg(protodef.GAME_REQ_SELECT_CARDS, req)
}

func (n *RobotNetwork) ReqSelectRole(req *protodef.ReqSelectRole) {
	n.sendMsg(protodef.GAME_REQ_SELECT_ROLE, req)
}

func (n *RobotNetwork) ReqFingerMove(req *protodef.ReqFingerMove) {
	n.sendMsg(protodef.GAME_REQ_FINGER_MOVE, req)
}

func (n *RobotNetwork) ReqShootBullet(req *protodef.ReqShootBullet) {
	n.sendMsg(protodef.GAME_REQ_SHOOT_BULLET, req)
}

func (n *RobotNetwork) ReqActionEnd(req *protodef.ReqActionEnd) {
	n.sendMsg(protodef.GAME_REQ_ACTION_END, req)
}
