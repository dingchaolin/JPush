package config

import (
	"os"
)

var env string

func InitEnv()  {
	sysEnv := os.Getenv("GOENV")
	if len(sysEnv) > 0 && (sysEnv == "production" || sysEnv == "testing" || sysEnv == "development"){
		env =  sysEnv
	}else{
		env = "testing"
	}
}

func GetEnv() string{
	return env
}