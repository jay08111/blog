package domain

import (
	"database/sql"
	"server/e"
	"server/env"
	"time"

	"github.com/sirupsen/logrus"
)

type PostDomain struct {
}

var Post *PostDomain

var MainDB *sql.DB

func Init() {
	var err error

	MainDB, err := sql.Open("mysql", MySQLBuildQueryString(
		env.DatabaseSetting.User,
		env.DatabaseSetting.Password,
		env.DatabaseSetting.DBName,
		env.DatabaseSetting.Host,
		env.DatabaseSetting.Port,
	))

	if err != nil {
		logrus.Error(e.DomainErrInit1, err)
	}

	MainDB.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	MainDB.SetMaxIdleConns(100)
}

func Close() {
	MainDB.Close()
}
