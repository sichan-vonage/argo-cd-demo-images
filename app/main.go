package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	env := os.Getenv("ENV")
	if env == "" {
		env = "N/A"
	}
	log.Printf("RUNNING APP IN ENV: %q", env)

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello from env=%q", env)
	})

	r.Run(":8080")
}