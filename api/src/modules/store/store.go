package store

import (
	"api-loja/src/infra/database"
	"api-loja/src/modules/business"
	"database/sql"
	"fmt"
)

type Store struct {
    Id int `json:"id"`
    Name string `json:"name"`
    BusinessId int `json:"business_id"`
}
func GetStoreBy[V any](filterColumn string, filterValue V, db *sql.DB) (Store, error) {
    var store Store
    config := database.GetByConfig[V]{
        FilterColumn: filterColumn,
        FilterValue: filterValue,
        Table: "store",
        ReturnColumns:  []string{"id", "name", "business_id"},
    }
    row, err := database.GetBy[V](config, db)
    if err != nil {
        return store, err
    }
    if err := row.Scan(&store.Id, &store.Name, &store.BusinessId); err != nil {
        return store, err
    }
    return store, nil
}
func CreateStore(store Store, db *sql.DB) (Store, error) {
    _, err := business.GetBusinessBy[int]("id", store.BusinessId, db)
    if err != nil {
        return store, err
    }
    tx, err := db.Begin()
    if err != nil {
        return store, err
    }
    row := tx.QueryRow("INSERT INTO store (name, business_id) VALUES ($1, $2) RETURNING id;", store.Name, store.BusinessId)
    err = row.Err()
    if err != nil {
        if err := tx.Rollback(); err != nil {
            return store, err
        }
        return store, err
    }
    if err := row.Scan(&store.Id); err != nil {
        return store, err
    }
    if err:= tx.Commit(); err != nil {
        return store, err
    }
    return store, nil
}
