package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port    string
	DB_Host string
	DB_User string
	DB_Pass string
	DB_Name string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("missing .env")
	}

	if Port = os.Getenv("PORT"); len(Port) == 0 {
		log.Fatal("missing [PORT] env variable")
	}

	if DB_Host = os.Getenv("POSTGRES_HOST"); len(DB_Host) == 0 {
		log.Fatal("missing [POSTGRES_HOST] env variable")
	}

	if DB_User = os.Getenv("POSTGRES_USER"); len(DB_User) == 0 {
		log.Fatal("missing [POSTGRES_USER] env variable")
	}

	if DB_Pass = os.Getenv("POSTGRES_PASSWORD"); len(DB_Pass) == 0 {
		log.Fatal("missing [POSTGRES_PASSWORD] env variable")
	}

	if DB_Name = os.Getenv("POSTGRES_DB"); len(DB_Name) == 0 {
		log.Fatal("missing [POSTGRES_DB] env variable")
	}

}
