package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Students struct {
	Name   string `json:"name"`
	Id     int    `json:"id"`
	Rollno int    `json:"rollno"`
}

var stud []Students

func getstud(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting json as content type
	json.NewEncoder(w).Encode(stud)
}
func createstud(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s Students
	_ = json.NewDecoder(r.Body).Decode(&s)
	//s.Id = rand.Int(1000000)
	stud = append(stud, s)
	fmt.Println(s)

}

func main() {
	stud = append(stud, Students{Name: "xxxx", Id: 34, Rollno: 2345})
	stud = append(stud, Students{Name: "yyyy", Id: 67, Rollno: 6785})
	r := mux.NewRouter()
	r.HandleFunc("/", getstud).Methods("GET")
	r.HandleFunc("/", createstud).Methods("POST")
	fmt.Println("starting server")
	http.ListenAndServe(":8000", r) //starting func
}
