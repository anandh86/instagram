package service

import (
	"errors"
	"image"

	"github.com/anandh86/instagram/models"
	"github.com/anandh86/instagram/repository"
)

type Service struct {
	repo repository.IRepository
}

func NewService(repository repository.IRepository) *Service {

	// compile-time check to ensure we implement the interface
	var _ IService = (*Service)(nil)

	return &Service{
		repo: repository,
	}
}

/*------------------------------------------------------------------------
*                             Post
------------------------------------------------------------------------*/

func (s *Service) CreatePost(post_img image.Image, post_info models.PostRequestDTO) (post_id string, err error) {
	// Implement the logic to create a new post
	img_id, img_err := s.repo.SaveImage(post_img)

	if img_err != nil {
		return "", errors.New("error saving image")
	}

	post_meta := models.PostMetaDTO{
		Caption: post_info.Caption,
		ImageId: img_id,
		Creator: post_info.AuthorId,
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
		AuthorId: post_meta.Creator,
	}

	return post_img, post_info, nil
}

func (s *Service) GetAllPosts() (posts []models.PostMetaDTO, err error) {
	// Implement the logic to retrieve all posts
	postMetaDatas, err := s.repo.GetAllPostMetas()
	var RetPostMetaDatas []models.PostMetaDTO

	if err != nil {
		return nil, errors.New("error retrieving posts")
	}

	for _, post_meta := range postMetaDatas {

		comments, _ := s.GetPostComments(post_meta.Id)

		var respComments []models.CommentResponseDTO

		for _, c := range comments {
			resComment := models.CommentResponseDTO{
				Id:        c.Id,
				Comment:   c.Content,
				AuthorId:  c.Creator,
				CreatedAt: c.CreatedAt,
			}

			respComments = append(respComments, resComment)
		}

		post_meta.Comments = respComments

		RetPostMetaDatas = append(RetPostMetaDatas, post_meta)
	}

	return RetPostMetaDatas, nil
}

/*
------------------------------------------------------------------------
*                             Comment
------------------------------------------------------------------------
*/
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

	if comment.Creator != author_id {
		return errors.New("unauthorized")
	}

	return s.repo.DeleteCommentByID(comment_id)
}

func (s *Service) GetPostComments(post_id string) (comments []models.CommentDTO, err error) {
	// TODO : Use config to set number of comments to retrieve; also the max file size
	return s.repo.GetPostLatestComments(post_id, 2)
}
