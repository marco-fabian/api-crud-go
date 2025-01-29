package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marco-fabian/api-crud-go/internal"
	"github.com/marco-fabian/api-crud-go/internal/database"
	"github.com/marco-fabian/api-crud-go/internal/post"
)

func main() {
	// Estabelecendo uma conexão com o banco de dados PostgreSQL
	connectionString := "postgresql://posts:p0stgr3s@db:5432/posts"
	//Puxando o pool de conexões 'conn'
	conn, err := database.NewConnection(connectionString)
	if err != nil {
		panic(err)
	}
	//Setando para definir qual query SQL aplicar baseado no contexto
	defer conn.Close()

	repo := post.Repository{
		Conn: conn,
	}
	//Para não dar error
	service := post.Service{
		Repository: repo,
	}

	// Alterando a requisição JSON para um objeto POST
	g := gin.Default()
	g.POST("/posts", func(ctx *gin.Context) {
		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//Chamando o método CREATE para inserir um novo post
		if err := service.Create(post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

	})
	g.Run(":3000")
}
