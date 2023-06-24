package domain

import (
	"database/sql"
	"server/e"
	"server/env"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type PostDomain struct {
}

var Post *PostDomain

var MainDB *sql.DB

func MySQLBuildQueryString(user, pass, dbname, host string, port int) string {
	config := mysql.NewConfig()

	config.User = user
	if len(pass) != 0 {
		config.Passwd = pass
	}
	config.DBName = dbname
	config.Net = "tcp"
	config.Addr = host
	if port == 0 {
		port = 3306
	}
	config.Addr += ":" + strconv.Itoa(port)

	config.ParseTime = true

	return config.FormatDSN()
}

func Init() {
	var err error

	MainDB, err = sql.Open("mysql", MySQLBuildQueryString(env.DatabaseSetting.User,
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
