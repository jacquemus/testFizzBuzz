package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jacquemus/Test_REEZOCAR/fizzer"
)

func main() {
	router := router()
	router.Run(":8080") //nolint
}

func router() *gin.Engine {
	r := gin.Default()
	r.GET("/fizzbuzz", GetAPIfizzbuzz)

	return r
}

func GetAPIfizzbuzz(c *gin.Context) {
	limitString := c.DefaultQuery("limit", "200")
	fizzModuloString := c.DefaultQuery("modulo1", "7")
	buzzModuloString := c.DefaultQuery("modulo2", "9")
	fizzWord := c.DefaultQuery("word1", "fizz")
	buzzWord := c.DefaultQuery("word2", "buzz")

	limitNumber, fizzModulo, buzzModulo, err := fizzer.Transformer(limitString, fizzModuloString, buzzModuloString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fizzerClient, err := fizzer.NewClient(limitNumber, fizzModulo, buzzModulo, fizzWord, buzzWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success!", "data": fizzerClient.Fizzbuzz()})
}
