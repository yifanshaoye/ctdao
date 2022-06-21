package main

import (
	"fmt"
	"github.com/yifanshaoye/ctdao/gorm"
	"github.com/yifanshaoye/ctdao/zaplog"
)

type User struct {
	ID int `gorm:"primary_key"`
	Name string `gorm:"not_null"`
}

func main() {
	fmt.Println("Hello, World !!!")

	//rcli,_ := goredis.GetRedisInstance("", "", "")
	//res := rcli.Get("a")
	//fmt.Println( res)
	//
	//res = rcli.Get("b")
	//fmt.Println( res)

	//gorm.Done()
	dsn := "root:long123456@tcp(127.0.0.1:3306)/collections?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.GetMysqlInstance(dsn)
	fmt.Println("err: ", err)
	user := []User{}
	db.Find(&user)
	fmt.Println(user)
	userr := User{
		Name: "log",
	}
	db.Create(&userr)
	db.Find(&user)
	fmt.Println(user)

	//zaplog.SetLogFilePath("./database/access.log")
	zaplog.InitZaplog()
	zaplog.Info("test", "test log %+v", "zaplog")
	zaplog.Error("second", "send log: %+v", "hahah error")
}
