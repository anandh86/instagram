package service

import (
	"errors"
	"image"
	"testing"

	"github.com/anandh86/instagram/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository implementing IRepository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) SaveImage(img image.Image) (string, error) {
	args := m.Called(img)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) SavePostMeta(meta models.PostMetaDTO) (string, error) {
	args := m.Called(meta)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) GetPostMetaByID(postID string) (models.PostMetaDTO, error) {
	args := m.Called(postID)
	return args.Get(0).(models.PostMetaDTO), args.Error(1)
}

func (m *MockRepository) GetImageByID(imgID string) (image.Image, error) {
	args := m.Called(imgID)
	return args.Get(0).(image.Image), args.Error(1)
}

func (m *MockRepository) GetAllPostMetas() ([]models.PostMetaDTO, error) {
	args := m.Called()
	return args.Get(0).([]models.PostMetaDTO), args.Error(1)
}

func (m *MockRepository) SaveComment(comment models.CommentRequestDTO) (string, error) {
	args := m.Called(comment)
	return args.String(0), args.Error(1)
}

func (m *MockRepository) GetCommentByID(commentID string) (models.CommentDTO, error) {
	args := m.Called(commentID)
	return args.Get(0).(models.CommentDTO), args.Error(1)
}

func (m *MockRepository) DeleteCommentByID(commentID string) error {
	args := m.Called(commentID)
	return args.Error(0)
}

func (m *MockRepository) GetPostLatestComments(postID string, limit int) ([]models.CommentDTO, error) {
	args := m.Called(postID, limit)
	return args.Get(0).([]models.CommentDTO), args.Error(1)
}

func TestCreatePost_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	testImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	postInfo := models.PostRequestDTO{
		Caption:  "Test Caption",
		AuthorId: "user123",
	}

	mockRepo.On("SaveImage", testImg).Return("img123", nil)
	mockRepo.On("SavePostMeta", mock.Anything).Return("post123", nil)

	postID, err := svc.CreatePost(testImg, postInfo)

	assert.NoError(t, err)
	assert.Equal(t, "post123", postID)
	mockRepo.AssertExpectations(t)
}

func TestCreatePost_SaveImageError(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	testImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	postInfo := models.PostRequestDTO{
		Caption:  "Test Caption",
		AuthorId: "user123",
	}

	mockRepo.On("SaveImage", testImg).Return("", errors.New("error saving image"))

	postID, err := svc.CreatePost(testImg, postInfo)

	assert.Error(t, err)
	assert.Equal(t, "", postID)
	mockRepo.AssertExpectations(t)
}

func TestGetPostById_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	postID := "post123"
	postMeta := models.PostMetaDTO{
		Id:      postID,
		Caption: "Test Caption",
		ImageId: "img123",
		Creator: "user123",
	}
	testImg := image.NewRGBA(image.Rect(0, 0, 100, 100))

	mockRepo.On("GetPostMetaByID", postID).Return(postMeta, nil)
	mockRepo.On("GetImageByID", "img123").Return(testImg, nil)

	img, info, err := svc.GetPostById(postID)

	assert.NoError(t, err)
	assert.Equal(t, testImg, img)
	assert.Equal(t, "Test Caption", info.Caption)
	assert.Equal(t, "user123", info.AuthorId)
	mockRepo.AssertExpectations(t)
}

func TestGetPostById_GetPostMetaError(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	postID := "post123"

	mockRepo.On("GetPostMetaByID", postID).Return(models.PostMetaDTO{}, errors.New("error retrieving post"))

	img, info, err := svc.GetPostById(postID)

	assert.Error(t, err)
	assert.Nil(t, img)
	assert.Equal(t, models.PostResponseDTO{}, info)
	mockRepo.AssertExpectations(t)
}

func TestGetAllPosts_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	postsMeta := []models.PostMetaDTO{
		{
			Id:      "post1",
			Caption: "Post 1 Caption",
			ImageId: "img1",
			Creator: "user1",
		},
		{
			Id:      "post2",
			Caption: "Post 2 Caption",
			ImageId: "img2",
			Creator: "user2",
		},
	}

	mockRepo.On("GetAllPostMetas").Return(postsMeta, nil)
	mockRepo.On("GetPostLatestComments", "post1", 2).Return([]models.CommentDTO{}, nil)
	mockRepo.On("GetPostLatestComments", "post2", 2).Return([]models.CommentDTO{}, nil)

	posts, err := svc.GetAllPosts()

	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	assert.Equal(t, "Post 1 Caption", posts[0].Caption)
	assert.Equal(t, "Post 2 Caption", posts[1].Caption)
	mockRepo.AssertExpectations(t)
}

func TestCommentOnPost_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	comment := models.CommentRequestDTO{
		PostId:   "post123",
		Comment:  "Nice post!",
		AuthorId: "user456",
	}

	mockRepo.On("GetPostMetaByID", "post123").Return(models.PostMetaDTO{}, nil)
	mockRepo.On("SaveComment", comment).Return("comment123", nil)

	commentID, err := svc.CommentOnPost(comment)

	assert.NoError(t, err)
	assert.Equal(t, "comment123", commentID)
	mockRepo.AssertExpectations(t)
}

func TestDeleteComment_Success(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	comment := models.CommentDTO{
		Id:      "comment123",
		Content: "Nice post!",
		Creator: "user456",
	}

	mockRepo.On("GetCommentByID", "comment123").Return(comment, nil)
	mockRepo.On("DeleteCommentByID", "comment123").Return(nil)

	err := svc.DeleteComment("comment123", "user456")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteComment_Unauthorized(t *testing.T) {
	mockRepo := new(MockRepository)
	svc := NewService(mockRepo)

	comment := models.CommentDTO{
		Id:      "comment123",
		Content: "Nice post!",
		Creator: "user789",
	}

	mockRepo.On("GetCommentByID", "comment123").Return(comment, nil)

	err := svc.DeleteComment("comment123", "user456")

	assert.Error(t, err)
	assert.EqualError(t, err, "unauthorized")
	mockRepo.AssertExpectations(t)
}
