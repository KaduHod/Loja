package store

import (
	"api-loja/src/infra/database"
	"api-loja/src/modules/business"
	"database/sql"
)

type Store struct {
    Id int `json:"id"`
    Name string `json:"name"`
    BusinessId int `json:"business_id"`
}
type Product struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
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
func CreateStoreProduct(store Store, product Product, db *sql.DB) (Product, error) {
    store, err := GetStoreBy[int]("id" ,store.Id, db)
    if err != nil {
        return product, err
    }
    tx, err := db.Begin()
    if err != nil {
        return product, err
    }
    row := tx.QueryRow("INSERT INTO product (name, description) VALUES ($1, $2) RETURNING id", product.Name, product.Description)
    if row.Err() != nil {
        if err := tx.Rollback(); err != nil {
            return product, err
        }
        return product, err
    }
    if err := row.Scan(&product.Id); err != nil {
        if err := tx.Rollback(); err != nil {
            return product, err
        }
        return product, err
    }
    row = tx.QueryRow("INSERT INTO store_product (store_id, product_id) VALUES ($1, $2)", store.Id, product.Id)
    if row.Err() != nil {
        if err := tx.Rollback(); err != nil {
            return product, err
        }
        return product, err
    }
    if err := tx.Commit(); err != nil {
        return product, err
    }
    return product, nil
}
