package main

import (
	"github.com/mgutz/gosu"
	"github.com/mgutz/gosu/util"
)

func Project(p *gosu.Project) {
	p.Task("default", []string{"views"})

	p.Task("views", gosu.Files{"views/**/*.go.html"}, func() {
		util.Exec("razor views views")
	})
}

func main() {
	gosu.Run(Project)
}
