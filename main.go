package main

import (
  "github.com/anandh86/instagram/repository"
  "github.com/anandh86/instagram/service"
  "github.com/anandh86/instagram/api/handlers"
  "github.com/gin-gonic/gin"
)

func main()  {

  r := gin.Default()

  repo := repository.NewInMemoryRepo()
  serv := service.NewService(repo)
  handler := handlers.NewHandler(serv)

  // User stories and their corresponding APIs

  // As a user, I should be able to create posts with images (1 post - 1 image)
  // As a user, I should be able to set a text caption when I create a post
  r.POST("/api/posts", handler.CreatePost)

	r.GET("/api/posts/:id", handler.GetPostById)

  // As a user, I should be able to get the list of all posts along with the
  // last 2 comments on each post
	r.GET("/api/posts", handler.GetAllPosts)

  // As a user, I should be able to comment on a post
	r.POST("/api/posts/:postId/comments", handler.CommentOnPost)
  // As a user, I should be able to delete a comment (created by me) from a post
	r.DELETE("/api/comments/:id", handler.DeleteComment)

  r.Run(":8080")
}
