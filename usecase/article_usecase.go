package usecase

import (
	"context"

	"github.com/krobus00/sg-rpl-echo/model"
)

type articleUsecase struct {
	articleRepository model.ArticleRepository
}

func NewArticleUsecase() model.ArticleUsecase {
	return new(articleUsecase)
}

func (u *articleUsecase) RegisterArticleRepository(repo model.ArticleRepository) {
	u.articleRepository = repo
}

func (u *articleUsecase) FindAll(ctx context.Context) ([]*model.Article, error) {
	articles, err := u.articleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (u *articleUsecase) Create(ctx context.Context, article *model.Article) error {
	err := u.articleRepository.Create(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (u *articleUsecase) FindByID(ctx context.Context, id int64) (*model.Article, error) {
	article, err := u.articleRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (u *articleUsecase) UpdateByID(ctx context.Context, article *model.Article) error {
	err := u.articleRepository.UpdateByID(ctx, article)
	if err != nil {
		return err
	}
	return nil
}

func (u *articleUsecase) DeleteByID(ctx context.Context, id int64) error {
	_, err := u.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = u.articleRepository.DeleteByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
