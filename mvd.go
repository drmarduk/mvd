package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "HAI.")
}
func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe("localhost:8000", nil)
}
