package repository

import (
	"image"
	"testing"
	"time"

	"github.com/anandh86/instagram/models"
	"github.com/stretchr/testify/assert"
)

func TestSaveImage(t *testing.T) {
	repo := NewInMemoryRepo()
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))

	imgID, err := repo.SaveImage(img)

	assert.NoError(t, err)
	assert.NotEmpty(t, imgID)

	savedImg, err := repo.GetImageByID(imgID)
	assert.NoError(t, err)
	assert.Equal(t, img, savedImg)
}

func TestGetImageByID_NotFound(t *testing.T) {
	repo := NewInMemoryRepo()

	_, err := repo.GetImageByID("nonexistent")

	assert.Error(t, err)
	assert.EqualError(t, err, "image not found")
}

func TestSavePostMeta(t *testing.T) {
	repo := NewInMemoryRepo()
	postMeta := models.PostMetaDTO{
		Caption: "Test Post",
		Creator: "user123",
		ImageId: "img123",
	}

	postID, err := repo.SavePostMeta(postMeta)

	assert.NoError(t, err)
	assert.NotEmpty(t, postID)

	savedPostMeta, err := repo.GetPostMetaByID(postID)
	assert.NoError(t, err)
	assert.Equal(t, postMeta.Caption, savedPostMeta.Caption)
	assert.Equal(t, postMeta.Creator, savedPostMeta.Creator)
	assert.Equal(t, postMeta.ImageId, savedPostMeta.ImageId)
}

func TestGetPostMetaByID_NotFound(t *testing.T) {
	repo := NewInMemoryRepo()

	_, err := repo.GetPostMetaByID("nonexistent")

	assert.Error(t, err)
	assert.EqualError(t, err, "post metadata not found")
}

func TestGetAllPostMetas(t *testing.T) {
	repo := NewInMemoryRepo()

	postMeta1 := models.PostMetaDTO{
		Caption: "Post 1",
		Creator: "user1",
		ImageId: "img1",
	}
	postMeta2 := models.PostMetaDTO{
		Caption: "Post 2",
		Creator: "user2",
		ImageId: "img2",
	}

	_, _ = repo.SavePostMeta(postMeta1)
	_, _ = repo.SavePostMeta(postMeta2)

	allPosts, err := repo.GetAllPostMetas()

	assert.NoError(t, err)
	assert.Len(t, allPosts, 2)
}

func TestSaveComment(t *testing.T) {
	repo := NewInMemoryRepo()
	postMeta := models.PostMetaDTO{
		Caption: "Test Post",
		Creator: "user123",
		ImageId: "img123",
	}

	postID, _ := repo.SavePostMeta(postMeta)

	commentReq := models.CommentRequestDTO{
		PostId:   postID,
		Comment:  "Nice post!",
		AuthorId: "user456",
	}

	commentID, err := repo.SaveComment(commentReq)

	assert.NoError(t, err)
	assert.NotEmpty(t, commentID)

	savedComment, err := repo.GetCommentByID(commentID)
	assert.NoError(t, err)
	assert.Equal(t, "Nice post!", savedComment.Content)
	assert.Equal(t, "user456", savedComment.Creator)
}

func TestGetCommentByID_NotFound(t *testing.T) {
	repo := NewInMemoryRepo()

	_, err := repo.GetCommentByID("nonexistent")

	assert.Error(t, err)
	assert.EqualError(t, err, "comment not found")
}

func TestDeleteCommentByID(t *testing.T) {
	repo := NewInMemoryRepo()
	postMeta := models.PostMetaDTO{
		Caption: "Test Post",
		Creator: "user123",
		ImageId: "img123",
	}

	postID, _ := repo.SavePostMeta(postMeta)

	commentReq := models.CommentRequestDTO{
		PostId:   postID,
		Comment:  "Nice post!",
		AuthorId: "user456",
	}

	commentID, _ := repo.SaveComment(commentReq)

	err := repo.DeleteCommentByID(commentID)
	assert.NoError(t, err)

	_, err = repo.GetCommentByID(commentID)
	assert.Error(t, err)
	assert.EqualError(t, err, "comment not found")
}

func TestGetPostLatestComments(t *testing.T) {
	repo := NewInMemoryRepo()
	postMeta := models.PostMetaDTO{
		Caption: "Test Post",
		Creator: "user123",
		ImageId: "img123",
	}

	postID, _ := repo.SavePostMeta(postMeta)

	commentReq1 := models.CommentRequestDTO{
		PostId:   postID,
		Comment:  "First comment",
		AuthorId: "user1",
	}

	commentReq2 := models.CommentRequestDTO{
		PostId:   postID,
		Comment:  "Second comment",
		AuthorId: "user2",
	}

	_, _ = repo.SaveComment(commentReq1)
	time.Sleep(1 * time.Second) // Ensure different timestamps
	_, _ = repo.SaveComment(commentReq2)

	latestComments, err := repo.GetPostLatestComments(postID, 1)

	assert.NoError(t, err)
	assert.Len(t, latestComments, 1)
	assert.Equal(t, "Second comment", latestComments[0].Content)
}
