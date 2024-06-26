package main

import (
	"frame"
	"log"
	"net/http"
	"time"
)

func onlyForV1_3() frame.HandlerFunc {
	return func(c *frame.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Request.RequestURI, time.Since(t))
	}
}

func main() {
	r := frame.New()
	r.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, `<h1>Hello SleepWeb</h1>`)
	})

	r.GET("/assets/*filepath", func(c *frame.Context) {
		c.JSON(http.StatusOK, frame.J{
			"file_path": c.Param("filepath"),
		})
	})

	v1_3 := r.Group("/v1.3")
	v1_3.Use(onlyForV1_3())
	{
		v1_3.GET("/hello/:name", func(c *frame.Context) {
			c.JSON(http.StatusOK, frame.J{
				"name": c.Param("name"),
				"path": c.Path,
			})
		})
	}

	_ = r.Run(":8080")
}
