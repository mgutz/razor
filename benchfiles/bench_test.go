package benchfiles

import (
	"bytes"
	//"fmt"
	"html/template"
	"testing"
)

var indexTmpl = template.Must(template.New("index").ParseFiles("layout.tmpl", "index.tmpl"))

func goTemplate() {
	var html string
	var w bytes.Buffer
	indexTmpl.ExecuteTemplate(&w, "layout", user)
	html = w.String()
	//fmt.Println("go", html)
	if html == "" {
	}
}

func razorRenderByFunc() {
	var html string
	var w bytes.Buffer
	Index(user).WriteTo(&w)
	html = w.String()
	//fmt.Println("razor", html)
	if html == "" {
	}
}

func razorRenderByName() {
	var html string
	var w bytes.Buffer
	Render("index", user).WriteTo(&w)
	html = w.String()
	//fmt.Println("razor", html)
	if html == "" {
	}
}

func BenchmarkGoTemplate(b *testing.B) {
	max := b.N
	for n := 0; n < max; n++ {
		goTemplate()
	}
}
func BenchmarkRazorByName(b *testing.B) {
	max := b.N
	for n := 0; n < max; n++ {
		razorRenderByName()
	}
}

func BenchmarkRazorByFunc(b *testing.B) {
	max := b.N
	for n := 0; n < max; n++ {
		razorRenderByFunc()
	}
}
