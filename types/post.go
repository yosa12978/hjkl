package types

type Post struct {
	Id       uint64 `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Created  string `json:"created"`
	Author   string `json:"author"`
	AuthorId uint64 `json:"author_id"`
	Comments int    `json:"comments_count"`
}

type PostSql struct {
	Id       uint64
	Title    string
	Content  string
	Created  string
	Author   string
	AuthorId uint64
	Comments int
}

type PostCreateDto struct {
	Title    string
	Content  string
	Created  string
	AuthorId string
}

type PostUpdateDto struct {
	Title   string
	Content string
}
