package model

type User struct {
	ID int64 `db:"id"`
}

func NewUser(id int64) *User {
	return &User{
		ID: id,
	}
}
