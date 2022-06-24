package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", welcome)

	http.ListenAndServe(":3000", nil)
}
