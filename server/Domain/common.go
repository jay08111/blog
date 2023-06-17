package domain

import (
	"strconv"

	"github.com/go-sql-driver/mysql"
)

// MySQLBuildQueryString builds a query string for MySQL.
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

	// MySQL is a bad, and by default reads date/datetime into a []byte
	// instead of a time.Time. Tell it to stop being a bad.
	config.ParseTime = true

	return config.FormatDSN()
}
