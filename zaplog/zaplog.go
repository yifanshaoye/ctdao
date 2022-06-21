package zaplog

import (
	"errors"
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/yifanshaoye/ctdao/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"time"
)

type coreLogger struct {
	logPath string
	logErrorPath string
	zaplog *zap.Logger

}

const (
	accessLog = "/access.log"
	errorLog = "/error.log"
)

var logger *coreLogger

func init() {
	dirPath := utils.GetLogDir()
	path, _ := filepath.Abs(dirPath + accessLog)
	logger = &coreLogger{logPath: path}
	path, _ = filepath.Abs(dirPath + errorLog)
	logger.logErrorPath = path
}

func (lg *coreLogger) initLogger() {
	enconf := zap.NewProductionEncoderConfig()
	enconf.EncodeTime = func(tm time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(tm.Format("2006-01-02 15:04:05.000"))
	}
	encoder := zapcore.NewConsoleEncoder(enconf)

	//fdir := filepath.Dir(lg.logPath)
	//os.MkdirAll(fdir, 0766)
	//file, err := os.OpenFile(lg.logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	lg.logPath, _= filepath.Abs(accessLogPath)
	//	fdir := filepath.Dir(lg.logPath)
	//	os.MkdirAll(fdir, 0766)
	//	file, _ = os.OpenFile(lg.logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//}
	//wrteSyncer := zapcore.AddSync(file)

	accessSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lg.logPath,
		MaxSize:    1000,
		MaxAge:     7,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	})

	infol := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level <= zapcore.InfoLevel
	})
	coreInfo := zapcore.NewCore(encoder, accessSyncer, infol)


	//file, err = os.OpenFile(lg.logErrorPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	lg.logErrorPath, _= filepath.Abs(errorLogPath)
	//	fdir := filepath.Dir(lg.logErrorPath)
	//	os.MkdirAll(fdir, 0766)
	//	file, _ = os.OpenFile(lg.logErrorPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	//}
	//writeSyncer := zapcore.AddSync(file)

	errorSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lg.logErrorPath,
		MaxSize:    1000,
		MaxAge:     7,
		MaxBackups: 3,
		LocalTime:  true,
		Compress:   false,
	})

	errorl := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level > zapcore.InfoLevel
	})

	coreError := zapcore.NewCore(encoder, errorSyncer, errorl)

	lg.zaplog = zap.New(zapcore.NewTee(coreInfo, coreError), zap.AddCallerSkip(1), zap.AddCaller())
}

// 初始化zaplog, 在写日志之前必须初始化
func InitZaplog() *coreLogger {
	if logger.zaplog == nil {
		logger.initLogger()
	}
	return logger
}

func SetLogFilePath(path string) error {
	apath, err := filepath.Abs(path)
	if err != nil {
		return errors.New("invalie path !!!")
	}
	logger.logPath = apath
	return nil
}

func SetErrorLogFilePath(path string) error {
	apath, err := filepath.Abs(path)
	if err != nil {
		return errors.New("invalie path !!!")
	}
	logger.logErrorPath = apath
	return nil
}

func Info(tag, format string, args ...interface{})  {
	msg := fmt.Sprintf(format, args...)
	logger.zaplog.Info(msg,
		zap.String("tag", tag),)
}

func Error(tag, format string, args ...interface{})  {
	msg := fmt.Sprintf(format, args...)
	logger.zaplog.Error(msg,
		zap.String("tag", tag),)
}