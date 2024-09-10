package database

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
)
type dummy struct {
    Id int
    Status string
    Description string
}
var db *sql.DB
func TestMain(m *testing.M) {
    var err error
    db, err = NewConnection()
    defer db.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(db.Stats().OpenConnections)
    // Executar os testes
    exitCode := m.Run()

    // Fazer "cleanup" (opcional)
    fmt.Println("Finalizando os testes")
    // Sair com o código de saída dos testes
    os.Exit(exitCode)
}

func TestConnection(t *testing.T) {
    rows, err := db.Query("SELECT id, status, status_description FROM purchase_status;")
    if err != nil {
        fmt.Println(err)
        t.Logf("Err while querying")
    }
    defer rows.Close()
    var result []dummy
    for rows.Next() {
        var row dummy
        if err := rows.Scan(&row.Id, &row.Status, &row.Description); err != nil {
            fmt.Println(err)
            t.Logf("Err scanning row")
        }
        result = append(result, row)
    }
}
func TestExists(t *testing.T) {
    exists, err := Exists("purchase_status", 1, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    if !exists {
        t.Log("Do not find register that should exists")
        t.Fail()
    }
}
func TestRegisterThatShouldNotExists(t *testing.T) {
    exists, err := Exists("purchase_status", 9999, db)
    if err != nil {
        fmt.Println(err)
        t.Log("Error while trying to query unexisting row")
        t.Fail()
    }
    if exists {
        t.Log("Find register that should not exists")
        t.Fail()
    }
}
func TestDelete(t *testing.T) {
    query := "INSERT INTO person (name, email) VALUES ('TESTE_DELETE', 'TESTE_DELETE@mail.com')"
    _, err := db.Exec(query)
    if err != nil {
        t.Log("Err while trying to insert value to delete")
        t.Fail()
    }
    row := db.QueryRow("SELECT id FROM person WHERE email = 'TESTE_DELETE@mail.com'")
    if row.Err() != nil {
        t.Log("Err while trying to fetch value to delete")
        t.Fail()
    }
    var id int
    if err := row.Scan(&id); err != nil {
        t.Log("Err while trying to scan value to delete")
        t.Fail()
    }
    if err := Delete("person", id, db); err != nil {
        t.Log("Err while trying to delete value")
        t.Fail()
    }
}
