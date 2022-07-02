package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ // nolint:exhaustruct
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
	}))
	e.POST("/post", Post)
	e.Logger.Fatal(e.Start(":1323"))
}
func Post(c echo.Context) error {
	v := make(map[string]interface{})
	err := c.Bind(&v)
	if err != nil {
		fmt.Println("error: ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	fmt.Println("request: ", v)
	// return c.JSON(http.StatusBadRequest, err)
	return c.JSON(http.StatusOK, v)
}
