package types

import "database/sql"

type Comment struct {
	Id       uint64
	Content  string
	Author   string
	AuthorId uint64
	PostId   uint64
	Replies  []Comment
}

type CommentSql struct {
	Id       uint64
	Content  string
	AuthorId sql.NullInt64
	ParentId sql.NullInt64
	PostId   uint64
}

type CommentCreateDto struct {
	Content  string
	AuthorId string
	// ParentId sql.NullInt64 (Maybe just use separate route to repoly and regular comment.)
	// PostId   uint64
}

type CommentUpdateDto struct {
	Content  string
	AuthorId string
}
