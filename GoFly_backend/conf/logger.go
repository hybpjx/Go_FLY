package conf

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.InfoLevel

	if viper.GetBool("model.development") {
		logMode = zapcore.DebugLevel
	}
	// 第一个参数是输出的格式 第二个参数 输出的位置

	//zapcore.NewMultiWriteSyncer 输出到多个终端 比如 文件 console中
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriterSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	return zap.New(core).Sugar()
}

// def 输出日志的格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	{
		// LevelKey值变为 level
		encoderConfig.LevelKey = "level"
		// MessageKey值变为 msg
		encoderConfig.MessageKey = "msg"
		// TimeKey值 变成time
		encoderConfig.TimeKey = "time"
		// 把输出的info 变成INFO 只需要丢对象 不许执行
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		// 对时间进行格式化处理
		encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Local().Format("2006-01-02 15:04:05"))
		}
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

// def 日志要输出到什么地方
func getWriterSyncer() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "log" + stSeparator + time.Now().Format("2006-01-02") + ".log"
	fmt.Println(stLogFilePath)

	// 日志分割
	hook := lumberjack.Logger{
		Filename:   stLogFilePath,                  // 日志文件路径，默认 os.TempDir()
		MaxSize:    viper.GetInt("log.MaxSize"),    // 每个日志文件保存500M，默认 100M
		MaxBackups: viper.GetInt("log.MaxBackups"), // 保留3个备份，默认不限
		MaxAge:     viper.GetInt("log.MaxAge"),     // 保留28天，默认不限
		Compress:   viper.GetBool("log.Compress"),  // 是否压缩，默认不压缩
	}

	return zapcore.AddSync(&hook)
}
