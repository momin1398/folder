package student

import (
	"github.com/labstack/echo"
)

var E = echo.New()

func Call() {
	E.GET("/student", getStudents)
	E.GET("/student/:id", getStudent)
	E.POST("/student", addStudent)
	E.Logger.Print("listenning to server")
	E.Logger.Fatal(E.Start("localhost:8080"))

}
