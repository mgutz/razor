package main

import (
	. "github.com/mgutz/gosu"
)

func Tasks(p *Project) {
	p.Task("default", D{"views"})

	p.Task("views", W{"example/views/**/*.go.html"}, func() {
		Run("razor views", M{"Dir": "example"})
	})

	p.Task("example", D{"views", "build"}, W{"example/**/*.go"}, Debounce(3000), func() {
		Start("main.go", M{"Dir": "example"})
	})

	p.Task("bench", D{"build"}, func() {
		Run("razor benchfiles")
		Run("go test -bench=.", M{"Dir": "benchfiles"})
	})

	p.Task("build", func() {
		Run("go install")
		Run("go install", M{"Dir": "cmd/razor"})
	})
}

func main() {
	Gosu(Tasks)
}
