package dao

import (
	"gf-app/gfboot/gfdao"
	"github.com/gogf/gf/frame/g"
)

type optionDao struct {
	gfdao.Base
}

func OptionDao() *optionDao {
	d := new(optionDao)
	d.TableName = "rotab_options"
	return d
}

func (this *optionDao) FindByCategory(category string) (*map[string]interface{}, error) {
	result, e := this.T().Where("category=?", category).All()
	option := g.Map{category: category}
	if result != nil {
		for _, item := range result.List() {
			option[item["name"].(string)] = item["value"]
		}
	}
	return &option, e
}

func (this *optionDao) SaveOrUpdate(option *map[string]interface{}) error {
	count, e := this.T().Where(option).Count()
	if count == 0 {
		_, e = this.T().Data(option).Insert()
	} else {
		_, e = this.T().Where(option).Update()
	}
	return e
}
