package repository

import (
	"image"

	"github.com/anandh86/instagram/models"
)

type IRepository interface {
	/*------------------------------------------------------------------------
	*                             Image
	------------------------------------------------------------------------*/

	// Save image to database
	SaveImage(image image.Image) (imgId string, err error)

	// Get Image by ID
	GetImageByID(imgId string) (img image.Image, err error)

	/*------------------------------------------------------------------------
	*                             Post
	------------------------------------------------------------------------*/

	// Save Post's metadata to database
	SavePostMeta(postMeta models.PostMetaDTO) (post_id string, err error)

	// Get Post Metadata by ID
	GetPostMetaByID(post_id string) (postMeta models.PostMetaDTO, err error)

	// Get All Post Metadata
	GetAllPostMetas() ([]models.PostMetaDTO, error)

	/*------------------------------------------------------------------------
	*                             Comment
	------------------------------------------------------------------------*/

	// Save a Comment on a Post
	SaveComment(comment models.CommentRequestDTO) (comment_id string, err error)

	// Retrieve a Comment on a Post
	GetCommentByID(comment_id string) (comment models.CommentDTO, err error)

	// Read all latest posts for a particular comment
	GetPostLatestComments(post_id string, numberOfComments int) ([]models.CommentDTO, error)

	// Delete a Comment on a Post
	DeleteCommentByID(comment_id string) error
}
