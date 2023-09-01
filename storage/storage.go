package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"littleapi/types"
)

type Storage interface {
	All() *[]types.User
	Get(int) *types.User
	Add(types.User) *types.User
	Delete(int) int
}

var db *sql.DB

func init() {
	var err error

	conn := "postgres://postgres@localhost?sslmode=disable"
	db, err = sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	// this will be printed in the terminal, confirming the connection to the database
	fmt.Println("The database is connected")
}
