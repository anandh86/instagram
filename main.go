package main

import (
	"github.com/anandh86/instagram/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	handler := handlers.NewHandler()

	r.POST("/api/posts", handler.CreatePost)
	r.GET("/api/posts/:id", handler.GetPostById)
	r.POST("/api/comments", handler.CommentOnPost)
	r.DELETE("/api/comments/:id", handler.DeleteComment)
	r.GET("/api/timeline", handler.ViewTimeline)

	r.Run(":8080")
}
