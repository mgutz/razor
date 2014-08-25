package main

import (
	. "github.com/mgutz/gosu"
)

func Tasks(p *Project) {
	p.Task("default", D{"views"})

	p.Task("views", W{"example/views/**/*.go.html"}, func() {
		Run("razor views views", M{"Dir": "example"})
	})

	p.Task("example", D{"views"}, W{"example/**/*.go"}, Debounce(3000), func() {
		Start("main.go", M{"Dir": "example"})
	})
}

func main() {
	Gosu(Tasks)
}
