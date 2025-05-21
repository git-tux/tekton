package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()
	router := app.Group("/")
	router.GET("/", Hello)
	app.Run(":8000")
}
func Hello(c *gin.Context) {
	c.Writer.Write([]byte("<h1>Hello Panos, Bodgan, Maria !</h1>"))
}
