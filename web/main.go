package main

import (
	"fmt"
	"frame"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	//V1_3()
	V1_4()
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func V1_4() {
	r := frame.New()
	r.Use(frame.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	r.GET("/", func(c *frame.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.Run(":8080")
}

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

func V1_3() {
	r := frame.New()

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
