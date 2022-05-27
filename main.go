package main

import (
	"fmt"
	"github.com/yifanshaoye/ctdao/gorm"
)

type User struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"not_null"`
}

func main() {
	fmt.Println("Hello, World !!!")
	gorm.Done()
	dsn := "root:long123456@tcp(127.0.0.1:3306)/collections?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.GetMysqlInstance(dsn)
	fmt.Println("err: ", err)
	user := []User{{Name: "teng"}, {Name: "chen"}}
	db.Create(user)

	//zaplog.SetLogFilePath("./database/access.log")
	//zaplog.InitZaplog()
	//zaplog.Info("test", "test log %+v", "zaplog")
	//zaplog.Error("second", "send log: %+v", "hahah")
}
