package testing

import "config/define"

var MysqlServerConfig define.MysqlServerConfig

var KafkaServerConfig define.KafakaServerConfig


func init()  {

	slaveConfig := define.MysqlServerConfig{}
	slaveConfig.Host = "192.168.64.160"
	slaveConfig.Port = 3306
	slaveConfig.User = "root"
	slaveConfig.PassWord = "123456"

	MysqlServerConfig = slaveConfig

	KafkaServerConfig = define.KafakaServerConfig{"192.168.10.145",39092}

}