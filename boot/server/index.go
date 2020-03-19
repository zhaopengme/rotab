package server

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.Debug("init server")
	config := g.Config().GetMap("server")
	s := g.Server()
	s.SetConfigWithMap(config)
	if gfile.Exists("./rtadmin") {
		s.AddStaticPath("/rtadmin/", "./rtadmin")
	}
}
