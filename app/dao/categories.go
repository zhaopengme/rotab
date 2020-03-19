package dao

import (
	"gf-app/gfboot/gfdao"
)

type navigationsDao struct {
	gfdao.Base
}

func NavigationsDao() *navigationsDao {
	d := new(navigationsDao)
	d.TableName = "rotab_navigations"
	return d
}
