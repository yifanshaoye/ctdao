package zaplog

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

type coreLogger struct {
	logPath string
	zaplog *zap.Logger

}

var logger *coreLogger

func init() {
	path, _ := filepath.Abs("./logs/logger.log")
	logger = &coreLogger{logPath: path}
}

func (lg *coreLogger) initLogger() {
	enconf := zap.NewProductionEncoderConfig()
	enconf.EncodeTime = func(tm time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(tm.Format("2006-01-02 15:04:05.000"))
	}
	encoder := zapcore.NewJSONEncoder(enconf)

	fdir := filepath.Dir(lg.logPath)
	os.MkdirAll(fdir, 0766)
	file, err := os.OpenFile(lg.logPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		lg.logPath, _= filepath.Abs("./logs/logger.log")
		fdir := filepath.Dir(lg.logPath)
		os.MkdirAll(fdir, 0766)
		file, _ = os.OpenFile(lg.logPath, os.O_WRONLY|os.O_CREATE, 0666)
	}
	wrteSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(encoder, wrteSyncer, zapcore.InfoLevel)


	lg.zaplog = zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())
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