package services

import (
	"errors"
	"image"

	"github.com/anandh86/instagram/models"
	"github.com/anandh86/instagram/repository"
)

type PostService interface {
	// TODO : Remove all response DTOs? It should only be part of handler?
	CreatePost(post_img image.Image, post_info models.PostRequestDTO) (post_id string, err error)

	GetPostById(id string) (post_img image.Image, post_info models.PostResponseDTO, err error)

	CommentOnPost(comment models.CommentRequestDTO) (comment_id string, err error)

	DeleteComment(comment_id, author_id string) (err error)
}

type Service struct {
	repo repository.IRepository
}

func NewService(repository repository.IRepository) *Service {
	return &Service{
		repo: repository,
	}
}

func (s *Service) CreatePost(post_img image.Image, post_info models.PostRequestDTO) (post_id string, err error) {
	// Implement the logic to create a new post
	img_id, img_err := s.repo.SaveImage(post_img)

	if img_err != nil {
		return "", errors.New("error saving image")
	}

	post_meta := models.PostMeta{
		Caption:  post_info.Caption,
		ImageId:  img_id,
		AuthorId: post_info.AuthorId,
	}

	return s.repo.SavePostMeta(post_meta)
}

func (s *Service) GetPostById(post_id string) (post_img image.Image, post_info models.PostResponseDTO, err error) {
	// Implement the logic to retrieve a post by ID
	post_meta, post_err := s.repo.GetPostMetaByID(post_id)

	if post_err != nil {
		return nil, models.PostResponseDTO{}, errors.New("error retrieving post")
	}

	post_img, err = s.repo.GetImageByID(post_meta.ImageId)

	if err != nil {
		return nil, models.PostResponseDTO{}, errors.New("error retrieving image")
	}

	post_info = models.PostResponseDTO{
		Id:       post_meta.Id,
		Caption:  post_meta.Caption,
		AuthorId: post_meta.AuthorId,
	}

	return post_img, post_info, nil
}

func (s *Service) GetAllPosts() (posts []models.PostMeta, err error) {
	// Implement the logic to retrieve all posts
	postMetaDatas, err := s.repo.GetAllPostMetas()
	var RetPostMetaDatas []models.PostMeta

	if err != nil {
		return nil, errors.New("error retrieving posts")
	}

	for _, post_meta := range postMetaDatas {

		comments, _ := s.GetPostComments(post_meta.Id)

		var commentTexts []string
		for _, c := range comments {
			commentTexts = append(commentTexts, c.Comment)
		}

		post_meta.Comments = commentTexts

		RetPostMetaDatas = append(RetPostMetaDatas, post_meta)
	}

	return RetPostMetaDatas, nil
}

func (s *Service) CommentOnPost(comment models.CommentRequestDTO) (comment_id string, err error) {
	// Implement the logic to create a new post

	// Check for validity of post id
	_, post_err := s.repo.GetPostMetaByID(comment.PostId)

	if post_err != nil {
		return "", errors.New("error retrieving post")
	}

	return s.repo.SaveComment(comment)
}

func (s *Service) DeleteComment(comment_id, author_id string) (err error) {
	// Implement the logic to retrieve a post by ID
	comment, err := s.repo.GetCommentByID(comment_id)

	if err != nil {
		return errors.New("error retrieving comment")
	}

	if comment.AuthorId != author_id {
		return errors.New("unauthorized")
	}

	return s.repo.DeleteCommentByID(comment_id)
}

func (s *Service) GetPostComments(post_id string) (comments []models.Comment, err error) {
	// Implement the logic to retrieve a post by ID
	post_comments, err := s.repo.GetPostCommentsMap(post_id)

	if err != nil {
		return nil, errors.New("error retrieving comments")
	}

	for _, comment_id := range post_comments {
		comment, _ := s.repo.GetCommentByID(comment_id)

		comments = append(comments, comment)
	}

	return comments, nil
}
