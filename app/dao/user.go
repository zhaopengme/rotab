package dao

import (
	"gf-app/app/model"
	"gf-app/gfboot/gfdao"
)

type userDao struct {
	gfdao.Base
}

func UserDao() *userDao {
	d := new(userDao)
	d.TableName = "rotab_users"
	return d
}

func (this *userDao) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	e := this.T().Where("username=?", username).Struct(user)
	return user, e
}
