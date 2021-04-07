package main

import "github.com/gogf/gf/frame/g"

type DemoComponent struct {
}

func (c *DemoComponent) Render(name string) string {
	content := `{{.name}} dafasdfasdfasdfasdfasdfasdf`
	result1, _ := g.View().ParseContent(content, g.Map{
		"name": name,
	})
	return result1
}

func NewDemo() *DemoComponent {
	c := &DemoComponent{}
	return c
}
