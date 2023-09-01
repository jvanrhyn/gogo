package storage

import "littleapi/types"

type MemoryStorage struct {
	users []types.User
}

func NewMemoryStorage() *MemoryStorage {

	return &MemoryStorage{
		users: []types.User{
			{ID: 1, Name: "One"},
			{ID: 2, Name: "Two"},
			{ID: 3, Name: "Three"},
		},
	}
}

func (s *MemoryStorage) Get(id int) *types.User {
	for _, user := range s.users {
		if user.ID == id {
			return &user
		}
	}

	return nil
}

func (s *MemoryStorage) Add(user types.User) *types.User {

	id := len(s.users) + 1

	user.ID = id

	s.users = append(s.users, user)
	return &user
}
