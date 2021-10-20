# mylog
一个简单的日志组件,
1. 提供 INFO/DEBUG/ERROR/WARN 类型的日志打印。
2. 日志打印控制台并且输出到日志文件, 文件句柄需手动控制关闭。
3. 考虑到并发情况，加了读写锁。
4. 日志函数出错时不影响当前线程的正常工作。
# Simple example
```
func TestLevel(t *testing.T) {
	myLog := GetLogger()
	//关闭文件句柄
	defer DestoryLogger()
	myLog.logConfig = myLog.logConfig.Logger2FileEnable(true).
		Logger2FileLevel(InfoLevel).
		Logger2FileNameFormat("20060102.log")

	myLog.Debug("我是输出日志的地方")
	myLog.Error("sdfdsfds")
}
```
## Output:

```
[2021-10-20 17:11:56][D]:我是输出日志的地方

[2021-10-20 17:11:56][E]:sdfdsfds
```
