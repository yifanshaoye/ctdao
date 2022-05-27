package gorm

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// dsn: user:password@tcp(ip:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local
func GetMysqlInstance(dsn string) (*gorm.DB, error) {
	if len(dsn) == 0 {
		return nil, errors.New("invalid MySQL dsn info !!!")
	}

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
}

func Done() {
	fmt.Println("db done")
}
