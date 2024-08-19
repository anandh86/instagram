package models

import "time"

type PostMetaDTO struct {
	Id        string               `json:"id"`
	Caption   string               `json:"caption"`
	CreatedAt time.Time            `json:"created_at"`
	ImageId   string               `json:"image_id"`
	Creator   string               `json:"creator_id"`
	Comments  []CommentResponseDTO `json:"comments"`
}

type PostRequestDTO struct {
	Id        string    `json:"id"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	ImageId   string    `json:"image_id"`
	AuthorId  string    `json:"creator_id"`
	Comments  []string  `json:"comments"`
}

type PostResponseDTO struct {
	Id       string               `json:"id"`
	Caption  string               `json:"caption"`
	AuthorId string               `json:"creator_id"`
	ImageId  string               `json:"image_id"`
	Comments []CommentResponseDTO `json:"comments"`
}

type CommentDTO struct {
	Id        string    `json:"id"`
	PostId    string    `json:"post_id"`
	Content   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	Creator   string    `json:"creator_id"`
}

type CommentRequestDTO struct {
	PostId   string `json:"post_id"`
	Comment  string `json:"comment"`
	AuthorId string `json:"creator_id"`
}

type CommentResponseDTO struct {
	Id        string    `json:"id"`
	Comment   string    `json:"comment"`
	AuthorId  string    `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
}
