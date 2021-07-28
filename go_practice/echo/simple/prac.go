package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type employee struct {
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Id     int    `json:"id"`
}

func main() {
	e := echo.New()
	var employeeDetails []employee
	var emp employee

	emp.Name = "xxx"
	emp.Id = 11
	emp.Salary = 46567
	employeeDetails = append(employeeDetails, emp)
	emp.Name = "yyy"
	emp.Id = 45
	emp.Salary = 87655
	employeeDetails = append(employeeDetails, emp)
	e.GET("/employee", func(c echo.Context) error {
		return c.JSON(http.StatusOK, employeeDetails)
	})
	e.GET("/employee/:id", func(c echo.Context) error {
		var empl employee
		eID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i := 0; i < len(employeeDetails); i++ {
			if eID == employeeDetails[i].Id {
				empl = employeeDetails[i]
			}

		}
		if empl == (employee{}) {
			return c.JSON(http.StatusNotFound, "given id not available")
		}

		return c.JSON(http.StatusOK, empl)

	})
	e.POST("/employee", func(c echo.Context) error {
		var reqBody, empl employee
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		empl.Id = reqBody.Id
		empl.Name = reqBody.Name
		empl.Salary = reqBody.Salary
		employeeDetails = append(employeeDetails, empl)
		return c.JSON(http.StatusOK, empl)

	})
	e.PUT("/employee/:id", func(c echo.Context) error {
		var empl, reqBody employee
		n := 0
		eID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i := 0; i < len(employeeDetails); i++ {
			if eID == employeeDetails[i].Id {
				empl = employeeDetails[i]
				n = i
			}

		}
		if empl == (employee{}) {
			return c.JSON(http.StatusNotFound, "given id not available")
		}
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		employeeDetails[n].Id = reqBody.Id
		employeeDetails[n].Name = reqBody.Name
		employeeDetails[n].Salary = reqBody.Salary
		return c.JSON(http.StatusOK, employeeDetails[n])
	})
	e.DELETE("/employee/:id", func(c echo.Context) error {
		var empl employee
		var n int
		eID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for i := 0; i < len(employeeDetails); i++ {
			if eID == employeeDetails[i].Id {
				empl = employeeDetails[i]
				n = i
			}

		}
		if empl == (employee{}) {
			return c.JSON(http.StatusNotFound, "given id not available")
		}
		delet := func(s []employee, index int) []employee {
			return append(s[:index], s[index+1:]...)
		}
		employeeDetails = delet(employeeDetails, n)
		return c.JSON(http.StatusOK, empl)

	})
	e.Logger.Print("listenning to server 8080")
	e.Logger.Fatal(e.Start("localhost:8080"))

}
