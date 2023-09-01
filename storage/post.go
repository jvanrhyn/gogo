package storage

import (
	"database/sql"
	"littleapi/types"
	"log"
)

type DbStorage struct {
	users []*types.User
}

func NewDbStorage() *DbStorage {
	return &DbStorage{}
}

func (s *DbStorage) All() *[]types.User {
	rowsRs, _ := db.Query("SELECT * FROM person")

	defer func(rowsRs *sql.Rows) {
		err := rowsRs.Close()
		if err != nil {
		}
	}(rowsRs)

	users := make([]types.User, 0)

	for rowsRs.Next() {
		user := types.User{}
		e := rowsRs.Scan(&user.ID, &user.Name, &user.Created)
		if e != nil {
			log.Panic(e.Error())
		}

		users = append(users, user)

	}

	return &users
}

func (s *DbStorage) Get(id int) *types.User {
	return &types.User{ID: id, Name: "Host"}
}

func (s *DbStorage) Delete(id int) int {
	return id
}

func (s DbStorage) Add(user types.User) *types.User {
	return &types.User{ID: 1, Name: user.Name}
}
