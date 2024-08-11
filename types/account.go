package types

import (
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

func (a AccountCreateDto) Validate() []string {
	errs := []error{
		validation.ValidateLength(3, 20)("username", a.Username),
		validation.ValidateLength(3, 20)("password", a.Password),
	}
	return validation.ExtractErrors(errs)
}

type AccountUpdateDto struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func (a AccountUpdateDto) Validate() []string {
	errs := []error{
		validation.ValidateLength(3, 20)("old password", a.OldPassword),
		validation.ValidateLength(3, 20)("new password", a.NewPassword),
	}
	return validation.ExtractErrors(errs)
}
