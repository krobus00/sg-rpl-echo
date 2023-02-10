package model

import (
	"context"
	"errors"
)

var (
	ErrArticleNotFound = errors.New("article not found")
)

type Article struct {
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CreateArticleRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type FindByIDRequest struct {
	ID int64 `param:"id" query:"id" form:"id" json:"id"`
}

type UpdateArticleRequest struct {
	ID     int64  `param:"id" query:"id" form:"id" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (m *UpdateArticleRequest) ToArticle() *Article {
	return &Article{
		ID:     m.ID,
		Title:  m.Title,
		Author: m.Author,
	}
}

func (m *CreateArticleRequest) ToArticle() *Article {
	return &Article{
		ID:     0,
		Title:  m.Title,
		Author: m.Author,
	}
}

type ArticleRepository interface {
	FindAll(ctx context.Context) ([]*Article, error)
	Create(ctx context.Context, article *Article) error
	FindByID(ctx context.Context, id int64) (*Article, error)
	UpdateByID(ctx context.Context, article *Article) error
	DeleteByID(ctx context.Context, id int64) error
}

type ArticleUsecase interface {
	FindAll(ctx context.Context) ([]*Article, error)
	Create(ctx context.Context, article *Article) error
	FindByID(ctx context.Context, id int64) (*Article, error)
	UpdateByID(ctx context.Context, article *Article) error
	DeleteByID(ctx context.Context, id int64) error

	RegisterArticleRepository(repo ArticleRepository)
}
