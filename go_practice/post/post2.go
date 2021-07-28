package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", serve)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func serve(w http.ResponseWriter, r *http.Request) {
	var res map[string]interface{}

	json.NewDecoder(r.Body).Decode(&res)

	fmt.Println(res["json"])
	fmt.Println(res)
	fmt.Fprintf(w, "received data")

}
