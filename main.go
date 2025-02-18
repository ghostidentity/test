package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "content": "<h1>Hello Leapcell</h1>",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
