package service

import (
	"image"

	"github.com/anandh86/instagram/models"
)

type IService interface {
	/*------------------------------------------------------------------------
	*                             Post
	------------------------------------------------------------------------*/
	// Create a new post by uploading an image
	CreatePost(post_img image.Image, post_info models.PostRequestDTO) (post_id string, err error)

	// Get the image specific to a post by its id
	GetPostById(post_id string) (post_img image.Image, post_info models.PostResponseDTO, err error)

	// Get all the posts
	GetAllPosts() (posts []models.PostMetaDTO, err error)

	/*------------------------------------------------------------------------
	*                             Comment
	------------------------------------------------------------------------*/

	// Comment on a specific post
	CommentOnPost(comment models.CommentRequestDTO) (comment_id string, err error)

	// Delete a comment; Only author's would be able to delete
	DeleteComment(comment_id, author_id string) (err error)

	// Get the comments corresponding to a specific post
	GetPostComments(post_id string) (comments []models.CommentDTO, err error)
}
