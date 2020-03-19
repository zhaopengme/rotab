package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	config := g.Config().GetMap("server")
	s := g.Server()
	s.SetConfigWithMap(config)
	s.Run()
}
