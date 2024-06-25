package main

import (
	"frame"
	"net/http"
)

func main() {
	r := frame.New()
	r.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, `<h1>Hello SleepWeb</h1>`)
	})

	r.GET("/hello", func(c *frame.Context) {
		c.String(http.StatusOK, "Welcome to %s, %s", c.Path, c.Query("name"))
	})

	r.GET("/hello/:name", func(c *frame.Context) {
		c.JSON(http.StatusOK, frame.J{
			"name": c.Param("name"),
			"path": c.Path,
		})
	})

	r.GET("/assets/*filepath", func(c *frame.Context) {
		c.JSON(http.StatusOK, frame.J{
			"file_path": c.Param("filepath"),
		})
	})

	_ = r.Run(":8080")
}
