package types

type Category struct {
	Id            string
	Subcategories []Category
	Posts         []Post
	Name          string
}
