package store
import (
	"api-loja/src/infra/database"
	"database/sql"
	"fmt"
	"os"
	"testing"
)

var db *sql.DB
func TestMain(m *testing.M) {
    var err error
    db, err = database.NewConnection()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer db.Close()
    exitCode := m.Run()
    os.Exit(exitCode)
}
func TestCreateStore(t *testing.T) {
    store, err := CreateStore(Store{
        Name: "TEST",
        BusinessId: 1,
    }, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
    fmt.Println(store)
    if store.Id == 0 {
        t.Log("Store id equals 0")
        t.Fail()
    }
}
func TestCreateStoreWithInvalidBusiness(t *testing.T) {
    _, err := CreateStore(Store{
        Name: "TEST",
        BusinessId: 9999999999999999,
    }, db)
    if err == nil {
        fmt.Println(err)
        t.Fail()
    }
}
func TestGetStoreBy(t *testing.T) {
    _, err := GetStoreBy[int]("id", 1, db)
    if err != nil {
        fmt.Println(err)
        t.Fail()
    }
}
