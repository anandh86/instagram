package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreatePost(c *gin.Context) {
	// Fetch the caption from the form data
	caption := c.PostForm("caption")
	if caption == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Caption is required"})
		return
	}

	// Fetch the image file from the form data
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}

	if fileHeader == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error processing image file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post_Id": "1234"})

}

func (h *Handler) GetPostById(c *gin.Context) {
	post_Id := c.Param("id")

	if post_Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post_id": post_Id})

}

func (h *Handler) CommentOnPost(c *gin.Context) {

	var requestBody struct {
		Comment string `json:"comment"`
		PostId  string `json:"post_id"`
		UserId  string `json:"user_id"`
	}

	// Bind the JSON request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Comment == "" || requestBody.PostId == "" || requestBody.UserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check input parameters again"})
		return
	}

	// Respond to the client
	c.JSON(http.StatusCreated, gin.H{"comment_id": "1234"})

}

func (h *Handler) DeleteComment(c *gin.Context) {
	comment_Id := c.Param("id")

	if comment_Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment deleted": comment_Id})

}

func (h *Handler) ViewTimeline(c *gin.Context) {
	var requestBody struct {
		UserId string `json:"user_id"`
	}

	// Bind the JSON request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.UserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check input parameters again"})
		return
	}

	// Respond to the client
	c.JSON(http.StatusOK, gin.H{"timeline for user": requestBody.UserId})

}
