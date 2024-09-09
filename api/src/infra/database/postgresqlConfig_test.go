package database

import (
	"fmt"
	"testing"
)
type dummy struct {
    Id int
    Status string
    Description string
}
func TestConnection(t *testing.T) {
    db, err := NewConnection()
    if err != nil {
        fmt.Println(err)
        t.Logf("Err tryying to connect to databse")
    }
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
