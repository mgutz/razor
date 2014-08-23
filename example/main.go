package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mgutz/razor"
	"github.com/mgutz/razor/example/models"
	"github.com/mgutz/razor/example/views/admin"
	"github.com/mgutz/razor/example/views/front"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "Admin"}
	admin.Index(user).WriteTo(w)
}

func frontHandler(w http.ResponseWriter, r *http.Request) {
	user := &models.User{Name: "You"}
	front.Index(user).WriteTo(w)
}

func main() {
	razor.SetAppState(razor.M{
		"version": "1.0.0",
		"now":     time.Now().UnixNano(),
		"pid":     os.Getpid(),
		"ppid":    os.Getppid(),
	})

	http.HandleFunc("/", frontHandler)
	http.HandleFunc("/admin", adminHandler)
	//http.Handle("/{{version}}/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	http.Handle("/{{version}}/", http.FileServer(http.Dir("public")))
	port := ":8080"
	fmt.Printf("Browse 127.0.0.1%s\n", port)
	http.ListenAndServe(":8080", nil)
}
