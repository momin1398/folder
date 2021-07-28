package main

import (
	"fmt"
	"net/http"
)

func simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi Guys!! %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", simple)
	http.ListenAndServe(":8080", nil)

}
