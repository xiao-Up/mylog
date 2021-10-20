package mylog

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	myLog := GetLogger()
	defer DestoryLogger()
	myLog.logConfig = myLog.logConfig.Logger2FileEnable(true).
		Logger2FileLevel(InfoLevel).
		Logger2FileNameFormat("20060102.log")
	fmt.Println(myLog.logConfig.format)
}

func TestLevel(t *testing.T) {
	myLog := GetLogger()
	defer DestoryLogger()
	myLog.logConfig = myLog.logConfig.Logger2FileEnable(true).
		Logger2FileLevel(InfoLevel).
		Logger2FileNameFormat("20060102.log")

	myLog.Debug("我是输出日志的地方")
	myLog.Error("sdfdsfds")
}
