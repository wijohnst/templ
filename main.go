package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/wijohnst/templ-poc/views"
)

func main() {
	http.Handle("/", templ.Handler(views.Hello("Foo")))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
