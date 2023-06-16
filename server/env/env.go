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
	Port     int
}

var DatabaseSetting Database

func Setup() error {
	var configToml *toml.Tree

	dbToml := configToml.Get("db").(*toml.Tree)

	if err := dbToml.Unmarshal(&DatabaseSetting); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
