package handlers

import (
	"image"
	"image/png"
	_ "image/png"
	"mime/multipart"
	"net/http"

	"github.com/anandh86/instagram/models"
	"github.com/anandh86/instagram/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *services.Service
}

// TODO : Fix the interface handling
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		service: service,
	}
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

	post_img, format, imgErr := h.processImage(fileHeader)

	if imgErr != nil || post_img == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error processing image file"})
		return
	}

	if format != "png" && format != "jpeg" && format != "jpg" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid format"})
	}

	postRequestDTO := models.PostRequestDTO{
		Caption:  caption,
		AuthorId: "1234",
	}

	post_id, post_err := h.service.CreatePost(post_img, postRequestDTO)

	if post_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating post"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post_Id": post_id})

}

func (h *Handler) GetPostById(c *gin.Context) {
	post_Id := c.Param("id")

	if post_Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get post"})
		return
	}

	post_img, _, err := h.service.GetPostById(post_Id)

	if err != nil || post_img == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting post"})
		return
	}

	c.Header("Content-Type", "image/png")
	encodeErr := png.Encode(c.Writer, post_img)

	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to encode image"})
		return
	}
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

	commentRequestDTO := models.CommentRequestDTO{
		Comment:  requestBody.Comment,
		PostId:   requestBody.PostId,
		AuthorId: requestBody.UserId,
	}

	comment_id, err := h.service.CommentOnPost(commentRequestDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
		return
	}

	// Respond to the client
	c.JSON(http.StatusCreated, gin.H{"comment_id": comment_id})

}

func (h *Handler) DeleteComment(c *gin.Context) {
	comment_Id := c.Param("id")

	if comment_Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var requestBody struct {
		PostId   string `json:"post_id"`
		AuthorId string `json:"author_id"`
	}

	// Bind the JSON request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.DeleteComment(comment_Id, requestBody.AuthorId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting comment"})
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

func (h *Handler) processImage(fileHeader *multipart.FileHeader) (image.Image, string, error) {
	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	// Decode the image
	return image.Decode(file)
}
