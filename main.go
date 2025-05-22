package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "Hello from vulnerable Gin v1.4.0")
    })
    r.Run(":8080")
}
