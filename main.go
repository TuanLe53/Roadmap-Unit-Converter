package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/css", "./css")

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

	r.POST("/convert/length", func(c *gin.Context) {
		source_value := c.PostForm("source_value")
		source_unit := c.PostForm("source_unit")
		target_unit := c.PostForm("target_unit")

		value, err := strconv.ParseFloat(source_value, 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		target_value, err := convertLength(value, source_unit, target_unit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		formatted_target_value := formatReadableValue(target_value)
		c.HTML(http.StatusOK, "result.html", gin.H{
			"conversion_type": "length",
			"source_unit":     source_unit,
			"target_unit":     target_unit,
			"source_value":    source_value,
			"target_value":    formatted_target_value,
		})
	})

	r.POST("/convert/weight", func(c *gin.Context) {
		source_value := c.PostForm("source_value")
		source_unit := c.PostForm("source_unit")
		target_unit := c.PostForm("target_unit")

		value, err := strconv.ParseFloat(source_value, 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		target_value, err := convertWeight(value, source_unit, target_unit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		formatted_target_value := formatReadableValue(target_value)
		c.HTML(http.StatusOK, "result.html", gin.H{
			"conversion_type": "weight",
			"source_unit":     source_unit,
			"target_unit":     target_unit,
			"source_value":    source_value,
			"target_value":    formatted_target_value,
		})
	})

	r.POST("/convert/temperature", func(c *gin.Context) {
		source_value := c.PostForm("source_value")
		source_unit := c.PostForm("source_unit")
		target_unit := c.PostForm("target_unit")

		value, err := strconv.ParseFloat(source_value, 64)
		if err != nil {
			fmt.Println("Error converting string to float64:", err)
			return
		}

		target_value, err := convertTemperature(value, source_unit, target_unit)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		formatted_target_value := formatReadableValue(target_value)
		c.HTML(http.StatusOK, "result.html", gin.H{
			"conversion_type": "temperature",
			"source_unit":     source_unit,
			"target_unit":     target_unit,
			"source_value":    source_value,
			"target_value":    formatted_target_value,
		})
	})

	log.Fatal(r.Run())
}
