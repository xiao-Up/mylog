# mylog
一个简单的日志中间件
# Simple example

func TestLevel(t *testing.T) {
	myLog := GetLogger()
	defer DestoryLogger()
	myLog.logConfig = myLog.logConfig.Logger2FileEnable(true).
		Logger2FileLevel(InfoLevel).
		Logger2FileNameFormat("20060102.log")

	myLog.Debug("我是输出日志的地方")
	myLog.Error("sdfdsfds")
}

Output:

[2021-10-20 17:11:56][D]:我是输出日志的地方

[2021-10-20 17:11:56][E]:sdfdsfds
