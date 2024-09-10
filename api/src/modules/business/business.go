package business

import (
	"database/sql"
	"errors"
	"fmt"
)

type Person struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}
type BusinessOwner struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

func GetPersonByEmail(email string, db *sql.DB) (Person, error) {
    var person Person
    query := fmt.Sprintf("SELECT id, name, email FROM person WHERE email = '%s' LIMIT 1", email)
    row := db.QueryRow(query)
    if row.Err() != nil {
        return person, row.Err()
    }
    if err := row.Scan(&person.Id, &person.Name, &person.Email); err != nil {
        return person, err
    }
    return person, nil
}

func CreateBusinessOwner(person BusinessOwner, db *sql.DB) (BusinessOwner, error) {
    var owner BusinessOwner
    tx, err := db.Begin()
    if err != nil {
        return owner, err
    }
    insertQuery := fmt.Sprintf("INSERT INTO person (name, email) VALUES ('%s', '%s')", person.Name, person.Email)
    _, err = tx.Exec(insertQuery)
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return owner, err
        }
        return owner, err
    }
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return owner, err
        }
        return owner, err
    }
    if err := tx.Commit(); err != nil {
        return owner, err
    }
    personCreated, err := GetPersonByEmail(person.Email, db)
    if err != nil {
        return owner, errors.New("Err while fetching created business owner")
    }
    owner.Id = personCreated.Id
    owner.Name = personCreated.Name
    owner.Email = personCreated.Email
    return owner, nil
}
