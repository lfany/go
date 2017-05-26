package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"net/http"
)

func main() {

	e := echo.New()
	e.Get("/", helloworld)
	//e.POST("/users", saveUser)

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	//e.Get("/user/:uid", getuser)

	e.Run(standard.New(":1234"))
}

func helloworld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world!")
}

/*
func saveUser(c echo.Context) error {

}*/

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
}
