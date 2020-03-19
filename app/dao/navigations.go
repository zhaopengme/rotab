package dao

import (
	"gf-app/gfboot/gfdao"
)

type categoriesDao struct {
	gfdao.Base
}

func CategoriesDao() *categoriesDao {
	d := new(categoriesDao)
	d.TableName = "rotab_categories"
	return d
}
