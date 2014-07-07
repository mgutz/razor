package main

import (
	"os/exec"

	"github.com/mgutz/gosu"
)

func Project(p *gosu.Project) {
	p.Task("default", []string{"views", "styles"})

	p.Task("views", gosu.Files{"views/**/*.go.html"}, func() {
		exec.Command("razor", "views", "views").Run()
	})

	p.Task("styles", gosu.Files{"public/**/*.less"}, func() {
		exec.Command("lessc", "public/{{version}}/css/style.less", "public/{{version}}/css/style.css").Run()
	})
}

func main() {
	gosu.Run(Project)
}
