package mylog

type LoggerConfig interface {
	Logger2FileEnable(enable bool) *LoggerConfig

	Logger2FileLevel(level LoggerLevel) *LoggerConfig

	Logger2FileNameFormat(format string) *LoggerConfig

	Logger2FilePath(path string) *LoggerConfig
}

type LoggerConfigImp struct {
	level    LoggerLevel
	enable   bool
	format   string
	filePath string
}

func (lc *LoggerConfigImp) Logger2FileEnable(enable bool) *LoggerConfigImp {
	return &LoggerConfigImp{level: lc.level,
		enable:   enable,
		format:   lc.format,
		filePath: lc.filePath}
}

func (lc *LoggerConfigImp) Logger2FileLevel(level LoggerLevel) *LoggerConfigImp {
	return &LoggerConfigImp{level: level,
		enable:   lc.enable,
		format:   lc.format,
		filePath: lc.filePath}
}

func (lc *LoggerConfigImp) Logger2FileNameFormat(format string) *LoggerConfigImp {
	return &LoggerConfigImp{level: lc.level,
		enable:   lc.enable,
		format:   format,
		filePath: lc.filePath}
}

func (lc *LoggerConfigImp) Logger2FilePath(path string) *LoggerConfigImp {
	return &LoggerConfigImp{level: lc.level,
		enable:   lc.enable,
		format:   lc.format,
		filePath: path}
}
