package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
)

var numberOfReqs = 0

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body)
		in, err := strconv.Atoi(string(body))
		if err == nil {
			numberOfReqs += in
		}
		return c.String(http.StatusOK, strconv.Itoa(numberOfReqs))
	})
	e.GET("/count", func(c echo.Context) error {
		return c.String(http.StatusOK, strconv.Itoa(numberOfReqs))
	})

	_ = e.Start(":80")

}
