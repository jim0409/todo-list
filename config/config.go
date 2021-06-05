package config

import "github.com/go-ini/ini"

var Conf *Config

type Config struct {
	BaseConf
	DbConf
	LogConf
}

type BaseConf struct {
	HttpPort string `ini:"HttpPort"`
	Env      string `ini:"Env"`
}

type DbConf struct {
	DbName        string `ini:"DbName"`
	DbHost        string `ini:"DbHost"`
	DbPort        string `ini:"DbPort"`
	DbUser        string `ini:"DbUser"`
	DbPassword    string `ini:"DbPassword"`
	DbLogMode     int    `ini:"DbLogMode"`
	DbMaxConnect  int    `ini:"DbMaxConnect"`
	DbIdleConnect int    `ini:"DbIdleConnect"`
}

type LogConf struct {
	LogPath  string `ini:"LogPath"`
	LogLevel string `ini:"LogLevel"`
}

func InitConfig(confPath *string) (*Config, error) {
	Conf = new(Config)
	if err := ini.MapTo(Conf, *confPath); err != nil {
		return nil, err
	}
	return Conf, nil
}
