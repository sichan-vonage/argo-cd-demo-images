package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const serviceName = "vonage-scheduler"

const version = "1.0.2"

func main() {
	r := gin.Default()

	env := os.Getenv("ENV")
	if env == "" {
		env = "N/A"
	}
	log.Printf("RUNNING %q IN ENV: %q", serviceName, env)

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello from serviceName=%q env=%q", serviceName, env)
	})

	r.Run(":8080")
}
