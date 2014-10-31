package main

import (
	. "github.com/mgutz/gosu"
)

func Tasks(p *Project) {
	p.Task("default", D{"example"})

	p.Task("views", D{"build"}, W{"example/views/**/*.go.html"}, func() {
		Run("razor example/views")
	})

	p.Task("example", D{"views"}, W{"example/**/*.go"}, Debounce(3000), func() {
		Start("main.go", In{"example"})
	})

	p.Task("bench", D{"build"}, func() {
		Run("razor benchfiles")
		Run("go test -bench . -benchmem", In{"benchfiles"})
	})

	p.Task("build?", func() {
		Run("go install", In{"cmd/razor"})
	})
}

func main() {
	Gosu(Tasks)
}
