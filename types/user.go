package types

import "time"

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func (u *User) Validate() bool { return true }
