package main

type Component interface {
	Render(name string) string
}
