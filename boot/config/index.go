package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func init()  {
	glog.Debug("init config")
	g.Cfg().SetFileName("config.yaml")
}
