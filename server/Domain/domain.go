package domain

import (
	"database/sql"
	"fmt"
	"server/e"
	"server/env"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type PostDomain struct {
}

var Post *PostDomain

var MainDB *sql.DB

func Init() {
	var err error

	dbHost := env.DatabaseSetting.Host
	dbPort := env.DatabaseSetting.Port
	dbUser := env.DatabaseSetting.User
	dbName := env.DatabaseSetting.DBName
	dbPassword := env.DatabaseSetting.Password

	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	MainDB, err := sql.Open("mysql", dbConnection)

	if err != nil {
		logrus.Error(e.DomainErrInit1, err)
	}

	MainDB.SetConnMaxLifetime(time.Duration(2 * time.Minute))
	MainDB.SetMaxIdleConns(100)
}

func Close() {
	MainDB.Close()
}
