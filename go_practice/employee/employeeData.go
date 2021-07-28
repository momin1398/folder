package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type employee struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Salary int    `json:"salary"`
	Id     int    `json:"id"`
}

var EmployeeDetails []employee
var e = echo.New()

func init() {
	emp := []employee{{Name: "xxxx", Age: 23, Salary: 357899, Id: 45}, {Name: "yyyy", Age: 24, Salary: 567899, Id: 34}}
	EmployeeDetails = append(EmployeeDetails, emp...)

}

func Call() {

	e.GET("/employee", getEmployees)
	e.GET("/employee/:id", getEmployee)
	e.POST("/employee", addEmployee)
	e.PUT("/employee/:id", updateEmployee)
	e.DELETE("/employee/:id", deleteEmployee)

}
func getEmployees(c echo.Context) error {
	return c.JSON(http.StatusOK, EmployeeDetails)
}
func findId(id int) (employee, int) {
	n := 0
	var m employee
	for i := 0; i < len(EmployeeDetails); i++ {
		if id == EmployeeDetails[i].Id {
			m = EmployeeDetails[i]
			n = i

		}
	}
	return m, n

}
func getEmployee(c echo.Context) error {

	eID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	empl, _ := findId(eID)

	if empl == (employee{}) {
		return c.JSON(http.StatusNotFound, "id not available")

	}
	return c.JSON(http.StatusOK, empl)
}
func addEmployee(c echo.Context) error {
	var reqBody, empl employee
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	empl.Id = reqBody.Id
	empl.Age = reqBody.Age
	empl.Salary = reqBody.Salary
	empl.Name = reqBody.Name
	EmployeeDetails = append(EmployeeDetails, empl)
	return c.JSON(http.StatusOK, empl)
}

func updateEmployee(c echo.Context) error {
	var reqBody employee
	var n int
	eID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	empl, n := findId(eID)
	if empl == (employee{}) {
		return c.JSON(http.StatusNotFound, "id not available")
	}
	if err := c.Bind(&reqBody); err != nil {
		return err
	}
	EmployeeDetails[n].Age = reqBody.Age
	EmployeeDetails[n].Id = reqBody.Id
	EmployeeDetails[n].Name = reqBody.Name
	EmployeeDetails[n].Salary = reqBody.Salary
	return c.JSON(http.StatusOK, EmployeeDetails[n])

}
func deleteEmployee(c echo.Context) error {
	eID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	empl, n := findId(eID)
	if empl == (employee{}) {
		return c.JSON(http.StatusNotFound, "id not availble")
	}

	del := func(s []employee, index int) []employee {
		return append(s[:index], s[index+1:]...)
	}
	EmployeeDetails = del(EmployeeDetails, n)
	return c.JSON(http.StatusOK, empl)
}
func main() {

	Call()
	e.Logger.Print("listenning to server")
	e.Logger.Fatal(e.Start("localhost:8080"))
}
