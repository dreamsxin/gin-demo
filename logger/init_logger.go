package logger

import (
	"log"
	"os"
	"time"

	"gin-demo/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func InitLogger(v *viper.Viper) *zap.Logger {
	var err error
	debug := v.GetBool("debug")
	if debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		log.Fatalf("logger init failed, err: %v", err)
	}
	log.Println("logger init success", debug)
	return logger
}

func Sugar() *zap.SugaredLogger {
	return logger.Sugar()
}

func Debug(v ...any) {
	logger.Sugar().Debug(v...)
}

func Debugln(v ...any) {
	logger.Sugar().Debugln(v...)
}

// 日志文件切割
func Rotatelog(filename string) (zapcore.WriteSyncer, error) {
	//保存日志30天，每1分钟分割一次日志
	hook, err := rotatelogs.New(
		filename,
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(24*time.Hour*time.Duration(config.Log().MaxAge)),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if config.Log().Console {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), err
	}
	return zapcore.AddSync(hook), err
}
