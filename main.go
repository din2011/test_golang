package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	aa := 23;
	fmt.Fprint(w, "<h1>Welcome to my awesome  site!"+"test",aa,"</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
