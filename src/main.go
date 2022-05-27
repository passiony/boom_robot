package main

import (
	"ROBOT/src/config"
	_ "ROBOT/src/network/mytcp"
	"ROBOT/src/robot"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("[启动] 开始机器人测试")
	confPath := flag.String("conf", "", "Server configuration file path")
	flag.Parse()
	config.LoadConfig(*confPath)

	robot.GetRobotMgr().Init(1000, 1000)

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
		log.Println("closing down (signal: %v)", sig)
	}
}
