package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/marco-fabian/api-crud-go/internal"
	"github.com/marco-fabian/api-crud-go/internal/database"
	"github.com/marco-fabian/api-crud-go/internal/post"
)

func main() {
	connectionString := os.Getenv("DATABASE_URL")
	if connectionString == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	conn, err := database.NewConnection(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close()

	repo := post.Repository{Conn: conn}
	service := post.Service{Repository: repo}

	g := gin.Default()

	g.POST("/posts", func(ctx *gin.Context) {
		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.Create(post); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "Post criado com sucesso", "post": post})
	})

	g.DELETE("/posts/:id", func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		id, err := uuid.Parse(idStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
			return
		}

		err = service.Delete(id)
		if err == post.ErrPostNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	})

	g.GET("/posts", func(ctx *gin.Context) {
		posts, err := service.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, posts)
	})

	g.PATCH("/posts/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		idUUID, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = service.Update(idUUID, post)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Post atualizado com sucesso", "post": post})
	})

	g.Run(":3000")
}
