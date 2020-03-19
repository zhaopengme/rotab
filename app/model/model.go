package model

import "time"

type Model struct {
	ID        int64     `orm:"id" json:"id"`
	CreatedAt time.Time `orm:"created_at" json:"createdAt"`
	UpdatedAt time.Time `orm:"updated_at" json:"updatedAt"`
	DeletedAt time.Time `orm:"deleted_at" json:"deletedAt"`
}
