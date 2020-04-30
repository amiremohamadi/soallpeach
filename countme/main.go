package main

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"strconv"
	"sync/atomic"
)

var numberOfReqs = new(int64)

func main() {
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body)
		in, _ := strconv.Atoi(string(body))
		atomic.AddInt64(numberOfReqs, int64(in))
		return c.NoContent(200)
	})
	e.GET("/count", func(c echo.Context) error {
		return c.String(200, strconv.Itoa(int(atomic.LoadInt64(numberOfReqs))))
	})

	_ = e.Start(":80")

}
