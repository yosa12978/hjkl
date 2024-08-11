package types

import (
	"github.com/yosa12978/hjkl/validation"
)

type Comment struct {
	Id       string
	Content  string
	Author   string
	AuthorId string
	PostId   string
	Created  string
	Replies  []Comment
}

type CommentCreateDto struct {
	Content  string
	AuthorId string
	ParentId string //(Maybe just use separate route to repoly and regular comment.)
	PostId   string
}

func (c CommentCreateDto) Validate() []string {
	errs := []error{
		validation.ValidatePresence("content", c.Content),
	}
	return validation.ExtractErrors(errs)
}

type CommentUpdateDto struct {
	Content  string
	AuthorId string
}

func (c CommentUpdateDto) Validate() []string {
	errs := []error{
		validation.ValidatePresence("content", c.Content),
	}
	return validation.ExtractErrors(errs)
}
