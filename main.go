package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/4k1k0/go-rest-sql/config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var db *sql.DB

func main() {
	conf := config.New()

	// SQLServer
	// Conecction string
	connSring := fmt.Sprintf("sqlserver://%s:%s@%s:%d", conf.SQL.Username, conf.SQL.Password, conf.SQL.Host, conf.SQL.Port)

	fmt.Println(connSring)
	db, err := sql.Open("sqlserver", connSring)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	fmt.Println("Conectado con exito")
	fmt.Println(db)
	defer db.Close()
}
