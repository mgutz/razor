package main

import (
	. "github.com/mgutz/gosu"
	"github.com/mgutz/gosu/util"
)

func Tasks(p *Project) {
	p.Task("default", Pre{"views"})

	p.Task("views", Watch{"example/views/**/*.go.html"}, Pre{"build"}, func() {
		util.Exec("razor views views", M{"Dir": "example"})
	})

	p.Task("example", func() {
		util.Exec("go run main.go", M{"Dir": "example"})
	})

	p.Task("build", func() {
		util.Exec("go install", M{"Dir": "cmd/razor"})
	})
}

func main() {
	Run(Tasks)
}
