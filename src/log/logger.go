package log

import (
	"ROBOT/src/config"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"runtime/debug"
)

var log *logs.BeeLogger

func InitLog(logConfig *config.LogConfig) {
	logConf := make(map[string]interface{})
	logConf["filename"] = logConfig.Log_path
	// logConf["level"] = logConfig.Log_level
	// logConf["maxlines"] = logConfig.Maxlines
	logConf["perm"] = "666"
	logConf["rotateperm"] = "666"
	logConf["daily"] = true
	logConf["MaxDays"] = 30

	confStr, err := json.Marshal(logConf)
	if err != nil {
		fmt.Println("marshal failed,err:", err)
		return
	}
	log = logs.NewLogger(10000)
	log.Async()
	log.SetLogger(logs.AdapterFile, string(confStr))
	if logConfig.Log_Console == 1 {
		log.SetLogger(logs.AdapterConsole)
	}
	log.SetLogFuncCallDepth(3)
	log.EnableFuncCallDepth(true)
}

func Info(format string, a ...interface{}) {
	log.Info(format, a...)
}

func Error(format string, a ...interface{}) {
	log.Error(format, a...)
	log.Error("%v", string(debug.Stack()))
}
