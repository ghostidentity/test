package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.Writer.WriteHeader(200)
        c.Writer.Write([]byte("<html><body><h1>Hello Leapcell</h1></body></html>"))
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
