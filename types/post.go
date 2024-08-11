package types

import "github.com/yosa12978/hjkl/validation"

type Post struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Created  string `json:"created"`
	Author   string `json:"author"`
	AuthorId string `json:"author_id"`
	Comments int    `json:"comments_count"`
}

type PostCreateDto struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorId string `json:"author_id"`
}

func (p PostCreateDto) Validate() []string {
	errs := []error{
		validation.ValidatePresence("content", p.Content),
		validation.ValidatePresence("title", p.Title),
		validation.ValidatePresence("author", p.AuthorId),
	}
	return validation.ExtractErrors(errs)
}

type PostUpdateDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p PostUpdateDto) Validate() []string {
	errs := []error{
		validation.ValidatePresence("content", p.Content),
		validation.ValidatePresence("title", p.Title),
	}
	return validation.ExtractErrors(errs)
}
