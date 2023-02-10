package repository

import (
	"context"

	"github.com/krobus00/sg-rpl-echo/model"
)

type articleRepository struct {
	articles      []*model.Article
	autoIncrement int64
}

func NewArticleRepository() model.ArticleRepository {
	return &articleRepository{
		articles:      make([]*model.Article, 0),
		autoIncrement: 1,
	}
}

func (r *articleRepository) FindAll(ctx context.Context) ([]*model.Article, error) {
	return r.articles, nil
}

func (r *articleRepository) Create(ctx context.Context, article *model.Article) error {
	if article.ID == 0 {
		article.ID = r.autoIncrement
		r.autoIncrement++
	}
	r.articles = append(r.articles, article)

	return nil
}

func (r *articleRepository) FindByID(ctx context.Context, id int64) (*model.Article, error) {
	for _, article := range r.articles {
		if article.ID == id {
			return article, nil
		}
	}
	return nil, model.ErrArticleNotFound
}

func (r *articleRepository) UpdateByID(ctx context.Context, article *model.Article) error {
	for i, m := range r.articles {
		if m.ID == article.ID {
			r.articles[i] = article
			return nil
		}
	}
	return model.ErrArticleNotFound
}

func (r *articleRepository) DeleteByID(ctx context.Context, id int64) error {
	tmp := make([]*model.Article, 0)
	for _, article := range r.articles {
		if article.ID != id {
			tmp = append(tmp, article)
		}
	}
	r.articles = tmp
	return nil
}
