package storage

import "littleapi/types"

type Storage interface {
	Get(int) *types.User
	Add(types.User) *types.User
	Delete(int) int
}
