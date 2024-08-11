package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/antoniohauren/raffle/config"
	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	host := config.DB_Host
	user := config.DB_User
	pass := config.DB_Pass
	name := config.DB_Name
	port := 5432

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("connected to %s", name)

	return db
}
