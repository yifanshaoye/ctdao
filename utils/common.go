package utils

import (
	"os"
	"os/user"
	"path/filepath"
)

var dirPath string

func getHomeDir() string {
	path, err := os.UserHomeDir()
	if err == nil {
		return path
	}
	user, err := user.Current()
	if err == nil {
		return user.HomeDir
	}
	path, _ = filepath.Abs("~")
	return path
}

func makeLogDir()  {
	home := getHomeDir()
	dirPath = home + "/collctionlogs"
	os.MkdirAll(dirPath, 0766)
}

func GetLogDir() string {
	makeLogDir()
	return dirPath
}