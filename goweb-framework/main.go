package main

import (
	"Go_study/goweb-framework/gout"
	"net/http"
)

func main()  {
	r := gout.New()

	r.GET("/", func(c *gout.Context) {

		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")

	})

	r.GET("/:name", func(c *gout.Context) {

		c.Data(http.StatusOK, []byte("jalins"))
	})

	r.Run(":8080")
}