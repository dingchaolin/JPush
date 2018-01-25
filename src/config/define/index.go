package define


type MysqlServerConfig struct {
	Host string
	Port int
	User string
	PassWord string
	Protocol string
}

type KafakaServerConfig struct {
	Host string
	Port int
}

type RedisServerConfig struct {
	Host string
	Port int
	Auth string
}

