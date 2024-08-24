package types

import (
	"context"
	"strings"

	"github.com/yosa12978/hjkl/validation"
)

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
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *PostCreateDto) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
	if err := validation.ValidatePresence("content", p.Content); err != nil {
		problems["content"] = err.Error()
	}
	if err := validation.ValidatePresence("title", p.Title); err != nil {
		problems["title"] = err.Error()
	}
	return problems
}

type PostUpdateDto struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (p *PostUpdateDto) Validate(ctx context.Context) map[string]string {
	problems := make(map[string]string)
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
	if err := validation.ValidatePresence("content", p.Content); err != nil {
		problems["content"] = err.Error()
	}
	if err := validation.ValidatePresence("title", p.Title); err != nil {
		problems["title"] = err.Error()
	}
	return problems
}
