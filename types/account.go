package types

import (
	"context"
	"strings"

	"github.com/yosa12978/hjkl/validation"
)

type Account struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
	Created  string `json:"created"`
	Salt     string `json:"-"`
	IsAdmin  bool   `json:"is_admin"`
}

type AccountCreateDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AccountCreateDto) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	a.Username = strings.TrimSpace(a.Username)
	a.Password = strings.TrimSpace(a.Password)
	if err := validation.ValidateLength(3, 20)("username", a.Username); err != nil {
		problems["username"] = err.Error()
	}
	if err := validation.ValidateLength(3, 20)("password", a.Password); err != nil {
		problems["password"] = err.Error()
	}
	return problems
}

type AccountUpdateDto struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func (a *AccountUpdateDto) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	a.NewPassword = strings.TrimSpace(a.NewPassword)
	a.OldPassword = strings.TrimSpace(a.OldPassword)
	if err := validation.ValidateLength(3, 20)("old password", a.OldPassword); err != nil {
		problems["oldPassword"] = err.Error()
	}
	if err := validation.ValidateLength(3, 20)("new password", a.NewPassword); err != nil {
		problems["newPassword"] = err.Error()
	}
	return problems
}
