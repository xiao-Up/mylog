package mylog

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type LoggerLevel int

var loggerTmp = LoggerImp{logConfig: &LoggerConfigImp{}}
var logFile *os.File = nil
var lastDay string = ""
var logMutex sync.RWMutex

const (
	UnsetLevel = iota - 2
	DebugLevel
	InfoLevel
	WarningLevel
	ErrorLevel
)

const (
	LogTimeFormat = "2006-01-02 15:04:05"
	LogPath       = "log/"
)

type Logger interface {
	Info(content string)
	Debug(content string)
	Error(content string)
	Warning(content string)
}

type LoggerImp struct {
	logConfig *LoggerConfigImp
}

func (lm *LoggerImp) Info(content string) {
	lm.printLog(InfoLevel, content)
}

func (lm *LoggerImp) Debug(content string) {
	lm.printLog(DebugLevel, content)
}

func (lm *LoggerImp) Error(content string) {
	lm.printLog(ErrorLevel, content)
}

func (lm *LoggerImp) Warning(content string) {
	lm.printLog(WarningLevel, content)
}

func (lm *LoggerImp) printToFile(level LoggerLevel, content string) {
	err := lm.initLogPath()
	if err != nil {
		panic("initLogPaht defeated")
	}
	if lm.logConfig.enable {
		if lm.logConfig.level >= level {
			err := writeLog(logFile, content)
			if err != nil {
				panic("Write to logFile failed: ")
			}
		}
	}
}

func (lm *LoggerImp) initLogPath() error {
	var err error
	if !IsExist(LogPath) {
		if err = CreateDir(LogPath); err != nil {
			return err
		}
	}

	if lastDay == "" {
		lastDay = time.Now().Format(lm.logConfig.format)
	}

	nowDay := time.Now().Format(lm.logConfig.format)

	//不等于0说明日期不对应 需要更新file句柄
	if strings.Compare(lastDay, nowDay) != 0 || logFile == nil {
		logFile, err = os.OpenFile(LogPath+nowDay, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}

	return err
}

func (lm *LoggerImp) printLog(level LoggerLevel, content string) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	str := ""
	//TODO
	switch level {
	case DebugLevel:
		str = fmt.Sprintf("[%s][%s]:%s", time.Now().Format(LogTimeFormat), "D", content+"\r\n")
		fmt.Println(str)
	case InfoLevel:
		str = fmt.Sprintf("[%s][%s]:%s", time.Now().Format(LogTimeFormat), "I", content+"\r\n")
		fmt.Println(str)
	case WarningLevel:
		str = fmt.Sprintf("[%s][%s]:%s", time.Now().Format(LogTimeFormat), "W", content+"\r\n")
		fmt.Println(str)
	case ErrorLevel:
		str = fmt.Sprintf("[%s][%s]:%s", time.Now().Format(LogTimeFormat), "E", content+"\r\n")
		fmt.Println(str)
	}

	lm.printToFile(level, str)
}

func GetLogger() LoggerImp {
	return loggerTmp
}

func DestoryLogger() {
	logFile.Close()
	logFile = nil
}

func writeLog(f *os.File, msg string) error {
	logMutex.Lock()
	defer logMutex.Unlock()
	_, err := fmt.Fprintf(f, msg)
	return err
}
