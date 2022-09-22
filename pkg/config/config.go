package config

type ServerSection struct {
	Port uint
	Host string
	Mode string
}

type DBSection struct {
	UserName string
	Password string
	Schema   string
}

type Config struct {
	Server *ServerSection
	Mongo  *DBSection
}
