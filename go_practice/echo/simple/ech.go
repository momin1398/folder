package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	products := []map[int]string{{1: "phone"}, {2: "tv"}, {3: "laptop"}}
	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, products)

	})
	e.PUT("/products/:id", func(c echo.Context) error {
		var prod map[int]string
		pID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return err
		}
		for _, p := range products {
			for k := range p {
				if pID == k {
					prod = p
				}
			}
		}
		if prod != nil {
			type body struct {
				Name string `json:"name"`
			}
			var reqBody body
			if err := c.Bind(&reqBody); err != nil {
				return err
			}
			prod[pID] = reqBody.Name
		}
		return c.JSON(http.StatusOK, prod)

	})
	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"name"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		prod := map[int]string{
			len(products) + 1: reqBody.Name,
		}
		products = append(products, prod)
		return c.JSON(http.StatusOK, prod)

	})
	e.GET("/products/:id", func(c echo.Context) error {
		var prod map[int]string
		for _, p := range products {
			for k := range p {
				pid, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pid == k {
					prod = p
				}
			}
		}
		if prod == nil {
			return c.JSON(http.StatusNotFound, "product not available")
		}
		return c.JSON(http.StatusOK, prod)

	})
	e.DELETE("/products/:id", func(c echo.Context) error {
		var prod map[int]string
		var index int
		for i, p := range products {
			for k := range p {
				pid, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pid == k {
					prod = p
					index = i
				}
			}
		}
		if prod == nil {
			return c.JSON(http.StatusNotFound, "product not available")
		}
		com := func(s []map[int]string, index int) []map[int]string {
			return append(s[:index], s[index+1:]...)

		}
		products = com(products, index)
		return c.JSON(http.StatusOK, prod)

	})
	e.Logger.Print("listenning to server 8080")
	e.Logger.Fatal(e.Start("localhost:8080"))
}
