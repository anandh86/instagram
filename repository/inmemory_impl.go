package repository

import (
	"errors"
	"image"
	"sort"
	"time"

	"github.com/anandh86/instagram/models"
	"github.com/google/uuid"
)

// InMemoryRepo is an in-memory implementation of IRepository
type InMemoryRepo struct {
	images          map[string]image.Image
	posts           map[string]models.PostMetaDTO
	comments        map[string]models.CommentDTO
	postCommentsMap map[string][]string
}

// NewInMemoryRepo creates a new instance of InMemoryRepo
func NewInMemoryRepo() *InMemoryRepo {

  // compile-time check to ensure we implement the interface
  var _ IRepository = (*InMemoryRepo)(nil)

	return &InMemoryRepo{
		images:          make(map[string]image.Image),
		posts:           make(map[string]models.PostMetaDTO),
		comments:        make(map[string]models.CommentDTO),
		postCommentsMap: make(map[string][]string),
	}
}

/*------------------------------------------------------------------------
*                             Image
------------------------------------------------------------------------*/
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

/*------------------------------------------------------------------------
*                             Post
------------------------------------------------------------------------*/
// SavePostMeta saves a post's metadata to the in-memory database
func (repo *InMemoryRepo) SavePostMeta(postMeta models.PostMetaDTO) (string, error) {
	postID := uuid.New().String()
	postMeta.Id = postID
	repo.posts[postID] = postMeta
	return postID, nil
}

// GetPostMetaByID retrieves a post's metadata by its ID
func (repo *InMemoryRepo) GetPostMetaByID(postID string) (models.PostMetaDTO, error) {
	postMeta, exists := repo.posts[postID]
	if !exists {
		return models.PostMetaDTO{}, errors.New("post metadata not found")
	}
	return postMeta, nil
}

func (repo *InMemoryRepo) GetAllPostMetas() ([]models.PostMetaDTO, error) {
	posts := make([]models.PostMetaDTO, 0, len(repo.posts))
	for _, post := range repo.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

/*------------------------------------------------------------------------
*                             Comment
------------------------------------------------------------------------*/
// SaveComment saves a comment to the in-memory database
func (repo *InMemoryRepo) SaveComment(reqComment models.CommentRequestDTO) (string, error) {
	comment := models.CommentDTO{
		Id:        uuid.New().String(),
		Content:   reqComment.Comment,
		PostId:    reqComment.PostId,
		Creator:   reqComment.AuthorId,
		CreatedAt: time.Now(),
	}

	repo.appendPostCommentsMap(reqComment.PostId, comment.Id)

	repo.comments[comment.Id] = comment
	return comment.Id, nil
}

// GetCommentByID retrieves a comment by its ID
func (repo *InMemoryRepo) GetCommentByID(commentID string) (models.CommentDTO, error) {
	comment, exists := repo.comments[commentID]
	if !exists {
		return models.CommentDTO{}, errors.New("comment not found")
	}
	return comment, nil
}

// DeleteCommentByID deletes a comment by its ID
func (repo *InMemoryRepo) DeleteCommentByID(commentID string) error {
	db_comment, exists := repo.comments[commentID]

	if !exists {
		return errors.New("comment not found")
	}

	delete(repo.comments, commentID)
	repo.removePostCommentsMap(db_comment.PostId, commentID)

	return nil
}

func (repo *InMemoryRepo) GetPostLatestComments(post_id string, numberOfComments int) ([]models.CommentDTO, error) {
	postComments, exists := repo.postCommentsMap[post_id]
	if !exists {
		return nil, errors.New("post not found")
	}

	var comments []models.CommentDTO

	for _, commentId := range postComments {
		comment, err := repo.GetCommentByID(commentId)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	// Sort comments by CreatedAt in descending order
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.After(comments[j].CreatedAt)
	})

	// Return the latest 'numberOfComments' comments
	if len(comments) > numberOfComments {
		return comments[:numberOfComments], nil
	}

	return comments, nil
}

/*------------------------------------------------------------------------
*                             Private functions
------------------------------------------------------------------------*/

func (repo *InMemoryRepo) appendPostCommentsMap(post_id string, comment_id string) error {
	postComments, exists := repo.postCommentsMap[post_id]
	if !exists {
		repo.postCommentsMap[post_id] = []string{comment_id}
		return nil
	}
	repo.postCommentsMap[post_id] = append(postComments, comment_id)
	return nil
}

func (repo *InMemoryRepo) removePostCommentsMap(post_id string, comment_id string) error {
	postComments, exists := repo.postCommentsMap[post_id]

	if !exists {
		return errors.New("post not found")
	}

	for i, comment := range postComments {
		if comment == comment_id {
			repo.postCommentsMap[post_id] = append(postComments[:i], postComments[i+1:]...)
			return nil
		}
	}

	return errors.New("comment not found")
}
