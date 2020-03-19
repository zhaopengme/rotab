package model

type Option struct {
	Model
	Category string `orm:"category" json:"category"`
	Name     string `orm:"name" json:"name"`
	Value    string `orm:"value" json:"value"`
	BlogId   int    `orm:"blog_id" json:"blog_id"`
}
