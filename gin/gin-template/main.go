package gin_template

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./statics")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/about.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", nil)
	})
	r.GET("/contact.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", nil)
	})

	r.Run(":8080")
}
