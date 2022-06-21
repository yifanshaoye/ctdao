package gorm

import (
	"errors"
	"fmt"
	"github.com/yifanshaoye/ctdao/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"path/filepath"
	"time"
)

const mysqlLog = "/mysql.log"

var logPath string

func init() {
	dirPath := utils.GetLogDir()
	path, _ := filepath.Abs(dirPath + mysqlLog)
	logPath = path
}

// dsn: user:password@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func GetMysqlInstance(dsn string) (*gorm.DB, error) {
	if len(dsn) == 0 {
		return nil, errors.New("invalid MySQL dsn info !!!")
	}

	//sqlog := MysqlLog{}
	writer, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		writer = os.Stdout
	}
	//sqlog.File = writer
	fmt.Println(writer, err)
	newLogger := logger.New(
		log.New(writer, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             1 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		},
		)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger: newLogger,
	})
}


//func (lw MysqlLog) Printf(format string, v ...interface{}) {
//	Logger.SetPrefix("")
//	setLogFile()
//	Logger.Printf(format, v...)
//}

func Done() {
	fmt.Println("db done")
}
