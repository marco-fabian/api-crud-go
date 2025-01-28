package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marco-fabian/api-crud-go/internal/database"
)

func main() {
	connectionString := "postgresql://posts:p0stgr3s@db:5432/posts"
	_, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":3000")
}
