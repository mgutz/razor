package benchfiles

import (
	"bytes"
	//"fmt"
	"html/template"
	"testing"
)

var indexTmpl = template.Must(template.New("index").ParseFiles("layout.tmpl", "index.tmpl"))

func goTemplate() bytes.Buffer {
	var w bytes.Buffer
	indexTmpl.ExecuteTemplate(&w, "layout", user)
	return w
}

func razorRenderByFunc() bytes.Buffer {
	var w bytes.Buffer
	Index(user).WriteTo(&w)
	return w
}

func razorRenderByName() bytes.Buffer {
	var w bytes.Buffer
	Render("index", user).WriteTo(&w)
	return w
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
