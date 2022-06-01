package main

import (
	"ROBOT/src/config"
	"ROBOT/src/log"
	_ "ROBOT/src/network/mytcp"
	"ROBOT/src/robot"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	confPath := flag.String("conf", "", "Server configuration file path")
	flag.Parse()
	config.LoadConfig(*confPath)
	log.InitLog(config.GetConfig().Log)

	log.Info("[启动] 开始机器人测试")
	robot.GetRobotMgr().Init(2000, 1000)
	wait()
}

func wait() {
	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	sig := <-c

	timeout := time.NewTimer(10 * time.Second)
	wait := make(chan struct{})
	go func() {
		wait <- struct{}{}
	}()
	select {
	case <-timeout.C:
		panic(fmt.Sprintf("close timeout (signal: %v)", sig))
	case <-wait:
		log.Info("closing down (signal: %v)", sig)
	}
}
