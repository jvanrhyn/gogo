package storage

import "littleapi/types"

type MemoryStorage struct {
	users []*types.User
}

func NewMemoryStorage() *MemoryStorage {

	return &MemoryStorage{
		users: []*types.User{
			{ID: 1, Name: "One"},
			{ID: 2, Name: "Two"},
			{ID: 3, Name: "Three"},
		},
	}
}

func (s *MemoryStorage) Get(id int) *types.User {
	for _, user := range s.users {
		if user.ID == id {
			return user
		}
	}

	return nil
}

func (s *MemoryStorage) Add(user types.User) *types.User {

	id := s.findHighestId() + 1

	user.ID = id

	s.users = append(s.users, &user)
	return &user
}

func (s *MemoryStorage) Delete(id int) error {

	user := s.Get(id)
	s.removeIt(user)

	return nil
}

func (s *MemoryStorage) findHighestId() int {
	var maxId int

	for _, user := range s.users {
		if user.ID > maxId {
			maxId = user.ID
		}
	}
	return maxId
}

func (s *MemoryStorage) removeIt(user *types.User) {
	for idx, v := range s.users {
		if v == user {
			s.users = append(s.users[0:idx], s.users[idx+1:]...)
		}
	}
}
