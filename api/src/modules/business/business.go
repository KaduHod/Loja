package business

import (
	"api-loja/src/infra/database"
	"api-loja/src/utils"
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

type Person struct {
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
func CreatePerson(person Person, db *sql.DB) (Person, error) {
    var owner Person
    tx, err := db.Begin()
    if err != nil {
        return owner, err
    }
    _, err = mail.ParseAddress(person.Email)
    if err != nil {
        return owner, errors.New("Email invalid")
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
type Business struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Cnpj string `json:"cnpj"`
}
func GetBusinessByCnpj(cnpj string, db *sql.DB) (Business, error) {
    var business Business
    row := db.QueryRow("SELECT id, name, cnpj FROM businesses WHERE cnpj = '$1' LIMIT 1", cnpj)
    if row.Err() != nil {
        return business, row.Err()
    }
    if err := row.Scan(&business.Id, &business.Name, &business.Cnpj); err != nil {
        return business, err
    }
    return business, nil
}
func CreateBusiness(business Business, db *sql.DB) (Business, error) {
    business.Cnpj = utils.OnlyNumbers(business.Cnpj)
    if len(business.Cnpj) < 14 {
        return business, errors.New("Cnpj invalid")
    }
    tx, err := db.Begin()
    if err != nil {
        return business, err
    }
    _, err = tx.Exec("INSERT INTO businesses (name, cnpj) VALUES ('$1', '$2')",  business.Name, business.Cnpj)
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return business, err
        }
        return business, err
    }
    if err := tx.Commit(); err != nil {
        return business, err
    }
    business, err = GetBusinessByCnpj(business.Cnpj, db)
    if err != nil {
        return business, err
    }
    return business, nil
}
func GetBusinessBy[V any](filterColumn string, filterValue V, db *sql.DB) (Business, error) {
    var business Business
    config := database.GetByConfig[V]{
        FilterColumn: filterColumn,
        FilterValue: filterValue,
        Table: "businesses",
        ReturnColumns:  []string{"id", "name", "cnpj"},
    }
    row, err := database.GetBy[V](config, db)
    if err != nil {
        return business, err
    }
    if err := row.Scan(&business.Id, &business.Name, &business.Cnpj); err != nil {
        return business, err
    }
    return business, nil
}
func GetPersonBy[V any](filterColumn string, filterValue V, db *sql.DB) (Person, error) {
    var person Person
    config := database.GetByConfig[V]{
        FilterColumn: filterColumn,
        FilterValue: filterValue,
        Table: "person",
        ReturnColumns:  []string{"id", "name", "email"},
    }
    row, err := database.GetBy[V](config, db)
    if err != nil {
        return person, err
    }
    if err := row.Scan(&person.Id, &person.Name, &person.Email); err != nil {
        return person, err
    }
    return person, nil
}
func RelateBusinessToPersons(businessId int, ids []int, db *sql.DB) error {
    var personsDb []Person
    var values []string
    business, err := GetBusinessBy[int]("id" ,businessId, db)
    if err != nil {
        return err
    }
    var argsAux [][]int
    var args []interface{}
    for _, id := range ids {
        person, err := GetPersonBy[int]("id", id, db)
        if err != nil {
            return err
        }
        personsDb = append(personsDb, person)
        argsAux = append(argsAux, []int{person.Id, business.Id})
        args = append(args, person.Id)
        args = append(args, business.Id)
    }
    for i := range argsAux {
        first := i + 1
        second := first + 1
        values = append(values, fmt.Sprintf("($%d, $%d)", first, second))
    }
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    query := fmt.Sprintf("INSERT INTO user_businesses (person_id, business_id) VALUES %s", strings.Join(values, ","))
    _, err = tx.Exec(query, args...)
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return err
        }
        return err
    }
    if err := tx.Commit(); err != nil {
        return err
    }
    return nil
}
