package main

import (
	"github.com/anandh86/instagram/api/handlers"
	"github.com/anandh86/instagram/repository"
	"github.com/anandh86/instagram/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize the repository and service
	repo := repository.NewInMemoryRepo()
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	r.POST("/api/posts", handler.CreatePost)
	r.GET("/api/posts/:id", handler.GetPostById)
	r.GET("/api/posts", handler.GetAllPosts)
	r.POST("/api/comments", handler.CommentOnPost)
	r.DELETE("/api/comments/:id", handler.DeleteComment)
	r.GET("/api/timeline", handler.ViewTimeline)

	r.Run(":8080")
}
