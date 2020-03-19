package dao

import (
	"gf-app/app/model"
	"gf-app/gfboot/gfdao"
)

type userinfoDao struct {
	gfdao.Base
}

func UserinfoDao() *userinfoDao {
	d := new(userinfoDao)
	d.TableName = "rotab_userinfo"
	return d
}

func (this *userinfoDao) FindByUserId(userId int64) (*model.Userinfo, error) {
	userinfo := &model.Userinfo{}
	e := this.T().Where("user_id=?", userId).Struct(userinfo)
	return userinfo, e
}
