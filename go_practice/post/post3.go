package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type student struct {
	name    string
	age     int
	section string
}

var p1 student

//var studen student

func addStudent(a interface{}) {
	fmt.Println("adding")
	p1, _ = a.(student)
	fmt.Println(p1.name)
	studen := student{name: p1.name, age: p1.age, section: p1.section}

	fmt.Printf("The details are %v\n", studen)
	fmt.Println(p1.age)

}

func main() {
	http.HandleFunc("/", serv)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
func serv(w http.ResponseWriter, r *http.Request) {
	var res map[string]interface{}

	json.NewDecoder(r.Body).Decode(&res)

	//	fmt.Println(res["json"])
	fmt.Println(res)
	fmt.Fprintf(w, "received data")
	addStudent(res)

}
