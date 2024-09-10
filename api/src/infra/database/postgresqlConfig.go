package database

import (
	"database/sql"
	"errors"
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

func Delete(table string, targetId int, db *sql.DB) error {
    query := fmt.Sprintf("DELETE FROM %s WHERE id IN (SELECT id FROM person WHERE id = %d LIMIT 1)", table, targetId)
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    res, err := tx.Exec(query)
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return err
        }
        return err
    }
    countRows, err := res.RowsAffected()
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return err
        }
        return err
    }
    if countRows > 1 {
        if err := tx.Rollback(); err != nil {
            return err
        }
        return errors.New("Delete count bigger than one")
    }
    if err := tx.Commit(); err != nil {
        return err
    }
    return nil
}

func Exists(table string, targetId int, db *sql.DB) (bool, error) {
    query := fmt.Sprintf("SELECT id FROM %s WHERE id = %d LIMIT 1", table, targetId)
    row := db.QueryRow(query)
    if row.Err() != nil {
        return false, row.Err()
    }
    var dummy int
    if err := row.Scan(&dummy); err != nil {
        if err.Error() == sql.ErrNoRows.Error() {
            return false, nil
        }
        return false, err
    }
    return true, nil
}
