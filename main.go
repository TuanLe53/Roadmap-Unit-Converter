package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/converter/length", func(c *gin.Context) {
		c.HTML(http.StatusOK, "length.html", nil)
	})
	r.GET("/converter/weight", func(c *gin.Context) {
		c.HTML(http.StatusOK, "weight.html", nil)
	})
	r.GET("/converter/temperature", func(c *gin.Context) {
		c.HTML(http.StatusOK, "temperature.html", nil)
	})

	log.Fatal(r.Run())
}
