package student

type Student struct {
	Name string `json:"Name"`
	Id   int    `json:"Id"`
	Age  int    `json:"ID"`
}

var StudentDetails []Student

func init() {
	stud := []Student{{Name: "xxx", Id: 23, Age: 20}, {"bbbb", 20, 19}}
	StudentDetails = append(StudentDetails, stud...)

}
func main() {
	Start()
}
