package student

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Student struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
	Id   int    `json:"Id"`
}

var StudentData []Student

func init() {
	StudentD := []Student{{"momin", 23, 44}, {"basha", 24, 56}}
	StudentData = append(StudentData, StudentD...)
}
func findEmployee(iD int) (Student, int) {
	n := 0
	var stud Student
	for i := 0; i < len(StudentData); i++ {
		if iD == StudentData[i].Id {
			stud = StudentData[i]
			n = i
		}
	}
	return stud, n
}
func getStudents(c echo.Context) error {
	return c.JSON(http.StatusOK, StudentData)

}
func getStudent(c echo.Context) error {
	sID, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		return err
	}
	stud, _ := findEmployee(sID)
	if stud == (Student{}) {
		return c.JSON(http.StatusNotFound, "id not available")

	}
	return c.JSON(http.StatusOK, stud)
}
func addStudent(c echo.Context) error {
	var reqBody, stud1 Student
	if err := c.Bind(&reqBody); err != nil {
		return err

	}
	stud1.Id = reqBody.Id
	stud1.Age = reqBody.Age
	stud1.Name = reqBody.Name
	StudentData = append(StudentData, stud1)
	return c.JSON(http.StatusOK, stud1)

}
