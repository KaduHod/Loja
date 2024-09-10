package business

import (
	"api-loja/src/infra/database"
	"database/sql"
	"fmt"
	"os"
	"testing"
)
var db *sql.DB
var emailsToDelete []string
func TestMain(m *testing.M) {
    var err error
    db, err = database.NewConnection()
    defer db.Close()
    if err != nil {
        os.Exit(1)
    }
    fmt.Println(db.Stats().OpenConnections)
    exitCode := m.Run()
    postTests()
    os.Exit(exitCode)
}
func TestGetPersonByEmail(t *testing.T) {
    _, err := GetPersonByEmail("root@mail.com", db)
    if err != nil {
        t.Log(err)
        t.Fail()
    }
}
func TestCreateBusinessOwner(t *testing.T) {
    owner, err := CreateBusinessOwner(BusinessOwner{
        Name: "TEST__ user ",
        Email: "teste@mail.com",
    }, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    if owner.Id == 0 {
        t.Log("Err trying to create business owner")
        t.Fail()
    }
    emailsToDelete = append(emailsToDelete, "teste@mail.com")
}
func postTests() {
    for _, email := range emailsToDelete {
        person, err := GetPersonByEmail(email, db)
        if err != nil {
            fmt.Println(err)
            fmt.Println("Err trying to query email", email)
            os.Exit(1)
        }
        if err := database.Delete("person", person.Id, db); err != nil {
            fmt.Println(err)
            fmt.Println("Err trying to DELETE email", email)
            os.Exit(1)
        }
    }
}
