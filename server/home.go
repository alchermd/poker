package server

import (
	"fmt"
	"net/http"
)

// Home servers the home page content.
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "hello, world!")
	if err != nil {
		panic(err)
	}
}
