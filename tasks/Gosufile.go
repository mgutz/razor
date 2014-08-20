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

	p.Task("build", func() {
		util.Exec("go install", M{"Dir": "cmd/razor"})
	})

	p.Task("test", func() {
		util.Exec("go test", M{"Dir": "razor"})
	})
}

func main() {
	Run(Tasks)
}
