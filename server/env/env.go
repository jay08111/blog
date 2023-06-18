package env

import (
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

type Database struct {
	User     string
	Password string
	Host     string
	DBName   string
	Port     string
}

type Server struct {
	Port string
}

var DatabaseSetting Database

var ServerSetting Server

func Setup() error {

	configToml, err := toml.LoadFile("config/dev-env.toml")

	if err != nil {
		logrus.Error(err)
		return err
	}

	dbToml := configToml.Get("db").(*toml.Tree)
	if err := dbToml.Unmarshal(&DatabaseSetting); err != nil {
		logrus.Error(err)
		return err
	}

	serverToml := configToml.Get("server").(*toml.Tree)
	if err := serverToml.Unmarshal(&ServerSetting); err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
