package config

import (
	"config/define"
	"config/env/production"
	"config/env/testing"
	"config/env/development"
	"fmt"
)

var MysqlServerConfig define.MysqlServerConfig
var KafkaServerConfig define.KafakaServerConfig
var Topic = "PushApiQueue_"
func init() {
	InitEnv()
	fmt.Println("evn===",env)
	switch env {
	case "production":
		MysqlServerConfig = production.MysqlServerConfig
		KafkaServerConfig = production.KafkaServerConfig
	case "testing":
		MysqlServerConfig = testing.MysqlServerConfig
		KafkaServerConfig = testing.KafkaServerConfig
	default:
		MysqlServerConfig = development.MysqlServerConfig
		KafkaServerConfig = development.KafkaServerConfig
	}

}
