package repository

import (
	"image"

	"github.com/anandh86/instagram/models"
)

type IRepository interface {
	// Save image to database
	SaveImage(image image.Image) (img_id string, err error)
	// Get Image by ID
	GetImageByID(img_id string) (img image.Image, err error)
	// Save Post's metadata to database
	SavePostMeta(postMeta models.PostMeta) (post_id string, err error)
	// Get Post Metadata by ID
	GetPostMetaByID(post_id string) (postMeta models.PostMeta, err error)
	// Get All Post Metadata
	GetAllPostMetas() ([]models.PostMeta, error)
	// Save a Comment on a Post
	SaveComment(comment models.CommentRequestDTO) (comment_id string, err error)
	// Retrieve a Comment on a Post
	GetCommentByID(comment_id string) (comment models.Comment, err error)
	// Delete a Comment on a Post
	DeleteCommentByID(comment_id string) error
	GetPostCommentsMap(post_id string) ([]string, error)
}
