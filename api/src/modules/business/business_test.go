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
var businessIdsToDelete []int
var businessPersonIdsToDelete []int
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
func TestGetBusiness(t *testing.T) {
    _, err := GetBusinessByCnpj("12345678912345", db)
    if err != nil {
        fmt.Println(err)
        t.Log("Failed to get business")
        t.Fail()
    }
}
func TestCreateBusiness(t *testing.T) {
    business, err := CreateBusiness(Business{
        Name: "TESTE_CREATE_BUSINESS",
        Cnpj: "12345678912346",
    }, db)
    if err != nil {
        fmt.Println(err)
        t.Log("Err trying to create business")
        t.Fail()
    }
    businessIdsToDelete = append(businessIdsToDelete, business.Id)
}
func TestGetBusinessBy(t *testing.T) {
    _, err := GetBusinessBy[int]("id", 1, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}
func TestShouldNotGetBusinessBy(t *testing.T) {
    _, err := GetBusinessBy[string]("cnpj", "00011122233344", db)
    if err == nil {
        fmt.Println(err)
        t.Fail()
    }
    if err != sql.ErrNoRows {
        fmt.Println(err)
        t.Fail()
    }
}
func TestGetPersonBy(t *testing.T) {
    _, err := GetPersonBy[string]("email", "root@mail.com", db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}
func TestShouldNotGetPersonBy(t *testing.T) {
    _, err := GetPersonBy[string]("email", "root_@mail.com", db)
    if err == nil {
        fmt.Println(err)
        t.Fail()
    }
    if err != sql.ErrNoRows {
        fmt.Println(err)
        t.Fail()
    }
}
func TestRelatePersonToBusiness(t *testing.T) {
    err := RelateBusinesToPersons(1, []int{6}, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    row := db.QueryRow("SELECT id FROM user_businesses WHERE person_id = 6 AND business_id = 1 LIMIT 1")
    if row.Err() != nil {
        fmt.Println(row.Err().Error())
        t.Fail()
    }
    var id int
    if err := row.Scan(&id); err != nil {
        fmt.Println(err)
        t.Fail()
    }
    businessPersonIdsToDelete = append(businessPersonIdsToDelete, id)
}
func TestRelatePersonToBusinessShouldNotSucced(t *testing.T) {
    err := RelateBusinesToPersons(1, []int{0}, db)
    if err == nil {
        t.Fail()
    }
    fmt.Println(err)
}
func postTests() {
    var has_error bool
    var errors []string
    for _, email := range emailsToDelete {
        person, err := GetPersonByEmail(email, db)
        if err != nil {
            fmt.Println(err)
            errors = append(errors, err.Error())
            fmt.Println("Err trying to query email", email)
            has_error = true
        }
        if err := database.Delete("person", person.Id, db); err != nil {
            errors = append(errors, err.Error())
            fmt.Println(err)
            fmt.Println("Err trying to DELETE email", email)
            has_error = true
        }
    }
    for _, id := range businessIdsToDelete {
        if err := database.Delete("businesses", id, db); err != nil {
            errors = append(errors, err.Error())
            fmt.Println(err)
            fmt.Println("Err trying to DELETE business", id)
            has_error = true
        }
    }
    for _, id := range businessPersonIdsToDelete {
        if err := database.Delete("user_businesses", id, db); err != nil {
            errors = append(errors, err.Error())
            fmt.Println(err)
            fmt.Println("Err trying to DELETE bususer_businesses", id)
            has_error = true
        }
    }
    if has_error {
        fmt.Println(errors)
        os.Exit(1)
    }
}
