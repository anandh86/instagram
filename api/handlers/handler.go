package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreatePost(c *gin.Context) {

}

func (h *Handler) GetPost(c *gin.Context) {

}

func (h *Handler) CommentOnPost(c *gin.Context) {

}

func (h *Handler) DeleteComment(c *gin.Context) {

}

func (h *Handler) ViewTimeline(c *gin.Context) {

}
