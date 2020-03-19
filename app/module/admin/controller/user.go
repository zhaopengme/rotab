package controller

import (
	"gf-app/app/dao"
	"gf-app/library/auth"
	"gf-app/library/response"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)

type UserCtr struct{}

func init() {
	g.Server().BindObject("/admin/user", new(UserCtr))
}

func (o *UserCtr) Info(r *ghttp.Request) {
	userinfo, e := dao.UserinfoDao().FindByUserId(auth.GetUserId(r))
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", userinfo)
}

func (o *UserCtr) UpdateInfo(r *ghttp.Request) {
	data := r.GetPostMapStrStr(g.Map{"nickname": "", "email": "", "description": "",})
	validationRules := map[string]string{
		"nickname":    "required",
		"email":       "required",
		"description": "required",
	}
	if e := gvalid.CheckMap(data, validationRules); e != nil {
		response.JsonFailGError(r, e, nil)
	}
	_, e := dao.UserinfoDao().Update(data, g.Map{"user_id": auth.GetUserId(r)})
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", nil)
}

func (o *UserCtr) UpdatePassword(r *ghttp.Request) {
	data := r.GetMapStrStr()
	validationRules := map[string]string{
		"originPassword": "required",
		"password":       "required",
		"password2":      "required",
	}
	if e := gvalid.CheckMap(data, validationRules); e != nil {
		response.JsonFailGError(r, e, nil)
	}
	user, e := dao.UserDao().FindByUsername(auth.GetUsername(r))
	response.JsonFailError(r, e, nil)
	pwd, _ := gmd5.EncryptString(data["originPassword"] + user.Salt)
	if pwd != user.Password {
		response.JsonFail(r, "password not match!", nil)
	}

	newPwd, _ := gmd5.EncryptString(data["password"] + user.Salt)
	_, e = dao.UserDao().Update(g.Map{"password": newPwd}, g.Map{"id": user.ID})
	response.JsonFailError(r, e, nil)
	response.JsonOK(r, "", nil)
}
