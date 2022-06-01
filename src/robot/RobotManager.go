package robot

import (
	"ROBOT/src/log"
	"sync"
	"time"
)

type RobotManager struct {
	robots      map[int64]*Robot
	mutex       sync.RWMutex
	maxRobotNum uint32
	freq        int64
	robotNum    uint32
	finish      bool
}

var robMgr *RobotManager

func GetRobotMgr() *RobotManager {
	if robMgr == nil {
		robMgr = &RobotManager{
			robots: make(map[int64]*Robot),
		}
	}
	return robMgr
}

func (m *RobotManager) Init(maxRobotNum uint32, freq int64) {
	m.maxRobotNum = maxRobotNum
	m.freq = freq
	go m.timeAction()
}

func (m *RobotManager) createRobot(id int64) {
	log.Info("添加机器人:%d", id)
	m.mutex.Lock()
	defer m.mutex.Unlock()
	rob := NewRobot(id)
	m.robots[rob.PlayerId] = rob
	rob.Init()
}

func (m *RobotManager) timeAction() {
	createTick := time.NewTicker(time.Millisecond * time.Duration(m.freq))
	defer createTick.Stop()
	for {
		select {
		case <-createTick.C:
			if m.robotNum < m.maxRobotNum {
				m.createRobot(int64(10000000 + m.robotNum))
				m.robotNum++
			}
		}
	}
}

func (m *RobotManager) IsFinish() bool {
	return m.finish
}
