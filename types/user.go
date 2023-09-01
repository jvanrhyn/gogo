package types

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (u *User) Validate() bool { return true }
