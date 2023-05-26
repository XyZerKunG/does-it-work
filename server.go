package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	cachedHostname, err := os.Hostname()

	if err != nil {
		cachedHostname = "@xyzerkung/rollout"
	}

	r.Use(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Yes Mom, I’m alive!",
			"from":    cachedHostname,
			"method":  c.Request.Method,
			"path":    c.Request.RequestURI,
		})
	})
	r.Run()
}
