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

func razorTemplate() {
	var html string
	var w bytes.Buffer
	Index(user).WriteTo(&w)
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

func BenchmarkRazor(b *testing.B) {
	max := b.N
	for n := 0; n < max; n++ {
		razorTemplate()
	}
}
