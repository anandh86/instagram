package models

import "time"

type PostMeta struct {
	Id        string    `json:"id"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	ImageId   string    `json:"image_id"`
	AuthorId  string    `json:"author_id"`
}

type PostRequestDTO struct {
	Id        string    `json:"id"`
	Caption   string    `json:"caption"`
	CreatedAt time.Time `json:"created_at"`
	ImageId   string    `json:"image_id"`
	AuthorId  string    `json:"author_id"`
	Comments  []string  `json:"comments"`
}

type PostResponseDTO struct {
	Id       string   `json:"id"`
	Caption  string   `json:"caption"`
	AuthorId string   `json:"author_id"`
	Comments []string `json:"comments"`
}

type Comment struct {
	Id        string    `json:"id"`
	PostId    string    `json:"post_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId  string    `json:"author_id"`
}
