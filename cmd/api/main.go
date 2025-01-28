package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marco-fabian/api-crud-go/internal/database"
	"github.com/marco-fabian/api-crud-go/internal/post"
)

func main() {
	// Estabelecendo uma conexão com o banco de dados PostgreSQL
	connectionString := "postgresql://posts:p0stgr3s@db:5432/posts"
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	repo := post.Repository{
		Conn: conn,
	}
	repo = repo

	// Definindo rota para o endpoint via GET
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	g.Run(":3000")
}
