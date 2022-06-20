package robot

import (
	"ROBOT/src/config"
	"ROBOT/src/log"
	protodef "ROBOT/src/network/protogo"
	"math/rand"
	"strconv"
	"time"
)

type Robot struct {
	PlayerId int64
	Token    string
	Guid     string
	Idfa     string
	Net      *RobotNetwork
	roleId   int32
	cards    []int32
	opera    int32
	Players  []*protodef.EntityPlayer
}

func NewRobot(id int64) *Robot {
	return &Robot{
		Guid: "gillar-" + strconv.FormatInt(id, 10),
		Idfa: "idfa" + strconv.FormatInt(id, 10),
	}
}

func (r *Robot) Init() {
	r.Net = NewNetwork(r)
	log.Info("机器人初始化成功")
	rand.Seed(time.Now().UnixNano())
	r.roleId = int32(rand.Intn(14))
	r.cards = make([]int32, 8)
	for i := 0; i < 8; i++ {
		r.cards[i] = int32(100 + i)
	}
	go r.Update()
}

func (r *Robot) CreateLeague() {
	log.Info("创建工会")
	req := &protodef.ReqCreateLeague{}
	req.LeagueName = "L" + strconv.FormatInt(r.PlayerId, 10)
	req.LeagueDesc = "desc-" + strconv.FormatInt(r.PlayerId, 10)
	r.Net.ReqCreateLeague(req)
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

func (r *Robot) OnLoginSuccess(rspLogin *protodef.RspLogin) {
	r.PlayerId = rspLogin.PlayerId
	r.Token = rspLogin.Token

	r.CreateLeague()
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

func (r *Robot) OnGameOver() {
	r.opera = 0
	r.MatchGame()
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
			if r.opera > 200 {
				r.KillAll()
				return
			}
		}
	}
}

func (r *Robot) Disconnect() {
	log.Info("重连server")
	r.Net.connect(config.GetConfig().Server)
}

func (r *Robot) Login() {
	log.Info("Login:%d", r.PlayerId)
	req := &protodef.ReqLogin{}
	req.Idfa = r.Idfa
	req.Idfv = r.Idfa
	req.Guid = r.Guid
	req.Token = r.Token
	req.Device = "linux"
	req.Platform = "linux-robot"
	req.AppVersion = "2.2.2"
	r.Net.ReqLogin(req)
}

func (r *Robot) HeatBeat() {
	req := &protodef.ReqHeart{}
	r.Net.ReqHeart(req)
}

func (r *Robot) MatchGame() {
	log.Info("MatchGame:%d", r.PlayerId)

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
	log.Info("Reconnect:%d", r.PlayerId)
	req := &protodef.ReqGameReconnect{}
	r.Net.ReqReconnect(req)
}

func (r *Robot) ReconnectFinish() {
	log.Info("ReconnectFinish:%d", r.PlayerId)
	req := &protodef.ReqReconnectFinish{}
	r.Net.ReqReconnectFinish(req)
}

func (r *Robot) GameReady() {
	log.Info("GameReady:%d", r.PlayerId)

	req := &protodef.ReqGameReady{}
	r.Net.ReqGameStart(req)
}

func (r *Robot) SelectCards() {
	//log.Info("SelectCards:%d", r.PlayerId)

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
	//log.Info("SelectRole:%d", r.PlayerId)

	req := &protodef.ReqSelectRole{}
	req.Frame = 0
	r.Net.ReqSelectRole(req)
}

func (r *Robot) FingerMove() {
	//log.Info("FingerMove:%d", r.PlayerId)

	req := &protodef.ReqFingerMove{}
	req.Dir = &protodef.Vector2{}
	req.Power = 0.85
	req.Frame = 0
	r.Net.ReqFingerMove(req)
}

func (r *Robot) ShootBulelt() {
	//log.Info("ShootBulelt:%d", r.PlayerId)

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
	//log.Info("ActionEnd:%d", r.PlayerId)

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

func (r *Robot) KillAll() {
	//log.Info("ActionEnd:%d", r.PlayerId)
	req := &protodef.ReqActionEnd{}
	req.Type = 1
	req.Entities = &protodef.EntityAll{
		Players: make([]*protodef.EntityPlayer, len(r.Players)),
	}
	for i, player := range r.Players {
		req.Entities.Players[i] = player
		for _, role := range player.Roles {
			role.Hp = 0
		}
	}
	r.Net.ReqActionEnd(req)
}
