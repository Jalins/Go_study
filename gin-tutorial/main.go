package main

import (
	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	router.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()

		go func() {
			time.Sleep(5 * time.Second)

			log.Panicln("async",cCp.Request.URL.Path)
		}()
	})

	router.GET("/sync", func(c *gin.Context) {

			time.Sleep(5 * time.Second)

			log.Panicln("sync",c.Request.URL.Path)

	})

	manners.ListenAndServe(":8080", router)
}
