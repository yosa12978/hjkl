package types

import "time"

type Account struct {
	Id       uint64    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"-"`
	Created  time.Time `json:"created"`
	Salt     string    `json:"-"`
	IsAdmin  bool      `json:"is_admin"`
}

type AccountSql struct {
	Id       uint64
	Username string
	Password string
	Created  time.Time
	Salt     string
	IsAdmin  bool
}

type AccountCreateDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccountUpdateDto struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
