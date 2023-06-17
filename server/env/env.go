package env

import (
	"path/filepath"

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

type Server struct {
	Port string
}

var DatabaseSetting Database

var ServerSetting Server

func Setup(filePath string) error {
	rootDir, _ := filepath.Abs("./")

	var configToml *toml.Tree
	var err error

	if len(filePath) > 0 {
		configToml, err = toml.LoadFile(rootDir + filePath)
	} else {
		configToml, err = toml.LoadFile(rootDir + "\\config\\dev-env.toml")
	}

	// configToml, err := toml.LoadFile("db.toml")
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
