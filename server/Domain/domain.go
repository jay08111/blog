package domain

import (
	"database/sql"

	"github.com/jay08111/blog/e"
	"github.com/sirupsen/logrus"
)

var MainDB *sql.DB

func Init() {
	var err error

	MainDB, err := sql.Open("mysql", dsn)

	if err != nil {
		logrus.Error(e.DomainErrInit1, err)
	}
}
