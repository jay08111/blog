package main

import (
	"flag"
	"server/env"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	flag.Parse()

	env.Setup()

	var echoRouter Router

	echoRouter.Init()

	// // dsn := "hoyeoun:028137jy@tcp(localhost:3306)/blog"

	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal("Failed to connect to the database:", err)
	// }

	// err = db.Ping()

	// if err != nil {
	// 	log.Fatal("Failed to pint to the database:", err)
	// }

	// fmt.Println("Connected to the database!")

	// defer db.Close()

	// e := echo.New()

}
