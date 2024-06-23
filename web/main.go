package main

import (
	"fmt"
	"frame"
	"net/http"
)

func main() {
	r := frame.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "URL Path: %s\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	_ = r.Run(":8080")
}
