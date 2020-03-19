package model

// User model.
type Userinfo struct {
	Model
	UserId      int64  `orm:"user_id" json:"userId"`
	NickName    string `orm:"nickname" json:"nickname"`
	Email       string `orm:"email" json:"email"`
	Description string `orm:"description" json:"description"`
}
