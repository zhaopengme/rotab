package main

func HelloInit() {
	Register("hello", New)
}

type HelloComponent struct {
}

func (c *HelloComponent) Render(name string) string {
	return "hello " + name
}

func New() *HelloComponent {
	c := &HelloComponent{}
	return c
}
