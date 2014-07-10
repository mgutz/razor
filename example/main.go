package main

import (
	"fmt"
	"net/http"

	"github.com/mgutz/razor/example/models"
	"github.com/mgutz/razor/example/views"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "Foo"}
	views.Index(user).WriteTo(w)
}

func main() {
	http.HandleFunc("/", viewHandler)
	//http.Handle("/{{version}}/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	http.Handle("/{{version}}/", http.FileServer(http.Dir("public")))
	port := ":8080"
	fmt.Printf("Browse 127.0.0.1%s\n", port)
	http.ListenAndServe(":8080", nil)
}
