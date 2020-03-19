package model

type Categories struct {
	Model
	Title       string `orm:"category" json:"category"`
	Slug        string `orm:"slug" json:"slug"`
	Description string `orm:"description" json:"description"`
	BlogId      int    `orm:"blog_id" json:"blog_id"`
}
