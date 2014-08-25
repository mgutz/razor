package main

import (
	. "github.com/mgutz/gosu"
)

func Tasks(p *Project) {
	p.Task("default", D{"views"})

	p.Task("views", W{"example/views/**/*.go.html"}, func() {
		Run("razor views views", M{"Dir": "example"})
	})

	p.Task("example", D{"views", "cmd"}, W{"example/**/*.go"}, Debounce(3000), func() {
		Start("main.go", M{"Dir": "example"})
	})

	p.Task("bench", D{"cmd"}, func() {
		Run("razor benchfiles benchfiles")
		Run("go test -bench=.", M{"Dir": "benchfiles"})
	})

	p.Task("cmd", "build razor command", func() {
		Run("go install", M{"Dir": "cmd/razor"})
	})
}

func main() {
	Gosu(Tasks)
}
