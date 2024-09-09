package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func NewConnection() (*sql.DB, error) {
    godotenv.Load("/app/.env")
    host := os.Getenv("POSTGRE_HOST")
    user := os.Getenv("POSTGRE_USER")
    pwd := os.Getenv("POSTGRE_PWD")
    db := os.Getenv("POSTGRE_DB")
    conString := fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable", host, user, pwd, db)
    database, err := sql.Open("postgres", conString)
    if err != nil {
        log.Fatal("Err trying to connect to database")
    }
    err = database.Ping()
    return database, err
}


