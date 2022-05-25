package main

import (
	"fmt"
	"github.com/yifanshaoye/ctdao/zaplog"

	"github.com/yifanshaoye/ctdao/database"
)

func main() {
	fmt.Println("Hello, World !!!")
	database.Done()

	//zaplog.SetLogFilePath("./database/access.log")
	zaplog.InitZaplog()
	zaplog.Info("test", "test log %+v", "zaplog")
	zaplog.Error("second", "send log: %+v", "hahah")
}
