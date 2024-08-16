package repository

import (
	"errors"
	"image"

	"github.com/anandh86/instagram/models"
	"github.com/google/uuid"
)

// InMemoryRepo is an in-memory implementation of IRepository
type InMemoryRepo struct {
	images   map[string]image.Image
	posts    map[string]models.PostMeta
	comments map[string]models.Comment
}

// NewInMemoryRepo creates a new instance of InMemoryRepo
func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{
		images:   make(map[string]image.Image),
		posts:    make(map[string]models.PostMeta),
		comments: make(map[string]models.Comment),
	}
}

// SaveImage saves an image to the in-memory database
func (repo *InMemoryRepo) SaveImage(img image.Image) (string, error) {
	imgID := uuid.New().String()
	repo.images[imgID] = img
	return imgID, nil
}

// GetImageByID retrieves an image by its ID
func (repo *InMemoryRepo) GetImageByID(imgID string) (image.Image, error) {
	img, exists := repo.images[imgID]
	if !exists {
		return nil, errors.New("image not found")
	}
	return img, nil
}

// SavePostMeta saves a post's metadata to the in-memory database
func (repo *InMemoryRepo) SavePostMeta(postMeta models.PostMeta) (string, error) {
	postID := uuid.New().String()
	postMeta.Id = postID
	repo.posts[postID] = postMeta
	return postID, nil
}

// GetPostMetaByID retrieves a post's metadata by its ID
func (repo *InMemoryRepo) GetPostMetaByID(postID string) (models.PostMeta, error) {
	postMeta, exists := repo.posts[postID]
	if !exists {
		return models.PostMeta{}, errors.New("post metadata not found")
	}
	return postMeta, nil
}

// SaveComment saves a comment to the in-memory database
func (repo *InMemoryRepo) SaveComment(comment models.Comment) (string, error) {
	commentID := uuid.New().String()
	comment.Id = commentID
	repo.comments[commentID] = comment
	return commentID, nil
}

// GetCommentByID retrieves a comment by its ID
func (repo *InMemoryRepo) GetCommentByID(commentID string) (models.Comment, error) {
	comment, exists := repo.comments[commentID]
	if !exists {
		return models.Comment{}, errors.New("comment not found")
	}
	return comment, nil
}

// DeleteCommentByID deletes a comment by its ID
func (repo *InMemoryRepo) DeleteCommentByID(commentID string) error {
	if _, exists := repo.comments[commentID]; !exists {
		return errors.New("comment not found")
	}
	delete(repo.comments, commentID)
	return nil
}
