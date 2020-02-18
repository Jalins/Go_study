package main

import (
	"Go_study/goweb-framework/gout"
	"net/http"
)

func main()  {
	r := gout.New()

	r.GET("/", func(c *gout.Context) {

		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		c.String(http.StatusOK, "hello %s, you're at %s\n")
		c.JSON(http.StatusOK, gout.H{
			"method" : c.Req.Method,
			"path" : c.Req.URL.Path,
		})

		c.Data(http.StatusOK, []byte("jalins"))
	})

	r.Run(":8080")
}