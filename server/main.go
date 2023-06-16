package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"server/env"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	flag.Parse()

	env.Setup()

	// db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Failed to pint to the database:", err)
	}

	fmt.Println("Connected to the database!")

	defer db.Close()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM posts")
		if err != nil {
			log.Println("Failed to execute query:", err)
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		defer rows.Close()

		ret := "hi"

		return c.JSON(http.StatusOK, ret)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
