package handlers

import (
	"image"
	"image/png"
	"mime/multipart"
	"net/http"

	"github.com/anandh86/instagram/models"
	"github.com/anandh86/instagram/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.IService
}

func NewHandler(serv service.IService) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h *Handler) CreatePost(c *gin.Context) {
	// Fetch the caption from the form data
	caption := c.PostForm("caption")

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

	// Check file size (100MB = 100 * 1024 * 1024 bytes)
	const MaxFileSize = 100 * 1024 * 1024 // 100MB
	if fileHeader.Size > MaxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds limit (100MB)"})
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

func (h *Handler) GetAllPosts(c *gin.Context) {

	postsMetaDatas, err := h.service.GetAllPosts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get posts"})
		return
	}

	var responses []models.PostResponseDTO

	for _, postMeta := range postsMetaDatas {

		postResponse := models.PostResponseDTO{
			Id:       postMeta.Id,
			Caption:  postMeta.Caption,
			AuthorId: postMeta.Creator,
			ImageId:  postMeta.ImageId,
			Comments: postMeta.Comments,
		}

		responses = append(responses, postResponse)
	}

	c.JSON(http.StatusOK, gin.H{"posts": responses})

}

func (h *Handler) CommentOnPost(c *gin.Context) {

	// Fetch the postId from the URL
	postId := c.Param("postId")
	// TODO: user id can be fetched from JWT token
	var requestBody struct {
		Comment string `json:"comment"`
		UserId  string `json:"user_id"`
	}

	// Bind the JSON request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if requestBody.Comment == "" || postId == "" || requestBody.UserId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Check input parameters again"})
		return
	}

	commentRequestDTO := models.CommentRequestDTO{
		Comment:  requestBody.Comment,
		PostId:   postId,
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

	// TODO : This can be fetched from JWT token (when implemented)
	var requestBody struct {
		AuthorId string `json:"author_id"`
	}

	// Bind the JSON request
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.DeleteComment(comment_Id, requestBody.AuthorId)

	if err != nil {
		if err.Error() == "unauthorized" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		} else if err.Error() == "error retrieving comment" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Comment not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment deleted": comment_Id})

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
