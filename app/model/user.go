package model

// User model.
type User struct {
	Model
	Username string `orm:"username" json:"username"`
	Password string `orm:"password" json:"password"`
	Salt     string `orm:"salt" json:"salt"`
}
