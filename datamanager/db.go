package datamanager

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB 

func init() {
    var err error 
    db, err = sql.Open("postgres",
        "postgres://dist:dist@localhost/dist?sslmode=disable")

    if err != nil {
        panic(err.Error())
    }
}