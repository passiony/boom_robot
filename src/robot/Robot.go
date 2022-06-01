package robot

import (
	protodef "ROBOT/src/network/protogo"
	"log"
	"math/rand"
	"time"
)

type Robot struct {
	PlayerId uint64
	Net      *RobotNetwork
	roleId   uint32
	cards    []uint32
	opera    int32
	Players  []*protodef.EntityPlayer
}

func NewRobot(id uint64) *Robot {
	return &Robot{
		PlayerId: id,
	}
}

func (r *Robot) Init() {
	r.Net = NewNetwork(r)
	log.Println("机器人初始化成功")
	rand.Seed(time.Now().UnixNano())
	r.roleId = uint32(rand.Intn(14))
	r.cards = make([]uint32, 8)
	for i := 0; i < 8; i++ {
		r.cards[i] = uint32(100 + i)
	}
	go r.Update()
}

func (r *Robot) Update() {
	createTick := time.NewTicker(2 * time.Second)
	defer createTick.Stop()
	for {
		select {
		case <-createTick.C:
			r.HeatBeat()
		}
	}
}

func (r *Robot) OnEnterRoom(players []*protodef.EntityPlayer) {
	r.Players = make([]*protodef.EntityPlayer, len(players))
	for i, player := range players {
		r.Players[i] = player
	}
}

func (r *Robot) OnGameStart() {
	r.opera = 0
	go r.operation()
}

func (r *Robot) operation() {
	createTick := time.NewTicker(100 * time.Millisecond)
	defer createTick.Stop()
	for {
		select {
		case <-createTick.C:
			r.opera++
			op := r.opera % 6
			if op == 1 {
				r.SelectCards()
			} else if op == 2 {
				r.SelectRole()
			} else if op == 3 {
				r.FingerMove()
			} else if op == 4 {
				r.ShootBulelt()
			} else if op == 5 {
				r.ActionEnd()
			}
		}
	}
}

func (r *Robot) Login() {
	log.Println("Login", r.PlayerId)
	req := &protodef.ReqLogin{}
	req.Idfa = string(r.PlayerId)
	req.Idfv = string(r.PlayerId)
	req.Guid = string(r.PlayerId)
	req.Device = "linux"
	req.Platform = "robot"
	req.AppVersion = "2.2.2"
	r.Net.ReqLogin(req)
}

func (r *Robot) HeatBeat() {
	req := &protodef.ReqHeart{}
	r.Net.ReqHeart(req)
}

func (r *Robot) MatchGame() {
	log.Println("MatchGame", r.PlayerId)

	req := &protodef.ReqMatchGame{}
	req.SceneLevel = 1
	req.GameType = 1
	req.Player = &protodef.EntityPlayer{
		Id:       int64(r.PlayerId),
		Lv:       2,
		Mp:       0,
		MaxMp:    0,
		CupNum:   200,
		Nickname: string(r.PlayerId),
		RoleType: 0,
		Rank:     100,
		Flag:     "cn",
		HasSkill: true,
		Roles:    make([]*protodef.EntityRole, 1),
		Cards:    make([]*protodef.EntityCard, 8),
	}
	for i, v := range r.cards {
		req.Player.Cards[i] = &protodef.EntityCard{
			Index: int32(i),
			Id:    int32(v),
			Lv:    1,
			Flag:  0,
		}
	}
	role := &protodef.EntityRole{
		Id:      1,
		CfgId:   int32(r.roleId),
		SkinIdx: 0,
		Hp:      200,
		MaxHp:   200,
		Name:    string(r.roleId),
	}
	req.Player.Roles[0] = role
	r.Net.ReqMatchGame(req)
}

func (r *Robot) Reconnect() {
	log.Println("Reconnect", r.PlayerId)
	req := &protodef.ReqGameReconnect{}
	r.Net.ReqReconnect(req)
}

func (r *Robot) ReconnectFinish() {
	log.Println("ReconnectFinish", r.PlayerId)
	req := &protodef.ReqReconnectFinish{}
	r.Net.ReqReconnectFinish(req)
}

func (r *Robot) GameReady() {
	log.Println("GameReady", r.PlayerId)

	req := &protodef.ReqGameReady{}
	r.Net.ReqGameStart(req)
}

func (r *Robot) SelectCards() {
	//log.Println("SelectCards", r.PlayerId)

	req := &protodef.ReqSelectCards{}
	req.Cards = make([]*protodef.EntityCard, 1)
	req.Cards[0] = &protodef.EntityCard{
		Id:    101,
		Index: 1,
		Lv:    1,
		Flag:  0,
	}
	r.Net.ReqSelectCards(req)
}

func (r *Robot) SelectRole() {
	//log.Println("SelectRole", r.PlayerId)

	req := &protodef.ReqSelectRole{}
	req.Frame = 0
	r.Net.ReqSelectRole(req)
}

func (r *Robot) FingerMove() {
	//log.Println("FingerMove", r.PlayerId)

	req := &protodef.ReqFingerMove{}
	req.Dir = &protodef.Vector2{}
	req.Power = 0.85
	req.Frame = 0
	r.Net.ReqFingerMove(req)
}

func (r *Robot) ShootBulelt() {
	log.Println("ShootBulelt", r.PlayerId)

	req := &protodef.ReqShootBullet{}
	req.Dir = &protodef.Vector2{}
	req.Power = 0.85
	req.Frame = 0
	req.InstanceId = 1
	req.OriginCard = &protodef.EntityCard{Id: 101}
	req.ActualCard = &protodef.EntityCard{Id: 101}
	req.Entities = &protodef.EntityAll{
		Players: make([]*protodef.EntityPlayer, len(r.Players)),
	}
	for i, player := range r.Players {
		req.Entities.Players[i] = player
	}
	r.Net.ReqShootBullet(req)
}

func (r *Robot) ActionEnd() {
	//log.Println("ActionEnd", r.PlayerId)

	req := &protodef.ReqActionEnd{}
	req.Type = 1
	req.Entities = &protodef.EntityAll{
		Players: make([]*protodef.EntityPlayer, len(r.Players)),
	}
	for i, player := range r.Players {
		req.Entities.Players[i] = player
	}
	r.Net.ReqActionEnd(req)
}
