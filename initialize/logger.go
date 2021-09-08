package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func InitLogger() {
	hook := lumberjack.Logger{
		Filename:   "./log/user.log", //日志文件路径
		MaxSize:    128,              //每个文件日志保存的最大尺寸 单位:M
		MaxBackups: 30,               //日志最多保存备份数量
		MaxAge:     7,                //日志文件保存天数
		Compress:   true,             //是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	//设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           //编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), //打印到控制台和文件
		atomicLevel, //日志级别
	)

	//开启开发者模式,堆栈跟踪
	caller := zap.AddCaller()
	//开启文件和行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development)

	zap.ReplaceGlobals(logger)
}
func init() {
	InitLogger()
}
