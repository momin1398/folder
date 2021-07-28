package student

import (
	"fmt"

	"github.com/labstack/echo"
)

var e = echo.New()

func Start() {
	fmt.Println(StudentDetails)

}
