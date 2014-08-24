package main

import (
	. "github.com/mgutz/gosu"
	"github.com/mgutz/gosu/util"
)

func Tasks(p *Project) {
	p.Task("default", Pre{"views"})

	p.Task("views", Watch{"example/views/**/*.go.html"}, Pre{"build"}, func() {
		util.Run("razor views views", M{"Dir": "example"})
	})

	p.Task("example", Pre{"views"}, Watch{"example/**/*.go"}, func() {
		util.Run("go build -o example main.go", M{"Dir": "example"})
		util.Start("example", M{"Dir": "example"})
	})

	p.Task("build", func() {
		util.Run("go install", M{"Dir": "cmd/razor"})
	})
}

func main() {
	Run(Tasks)
}
