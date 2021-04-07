package main

import "github.com/gogf/gf/frame/g"


type HelloComponent struct {
}

func (c *HelloComponent) Render(name string) string {
	content := `{{.name}} says nihao {{gowire "demo"}}`
	result1, _ := g.View().ParseContent(content, g.Map{
		"name": name,
	})
	return result1
}

func New() *HelloComponent {
	c := &HelloComponent{}
	return c
}
