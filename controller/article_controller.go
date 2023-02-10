package controller

import (
	"net/http"

	"github.com/krobus00/sg-rpl-echo/model"
	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	articleUscase model.ArticleUsecase
}

func NewArticleController() *ArticleController {
	return new(ArticleController)
}

func (c *ArticleController) RegisterArticleUsecase(usecase model.ArticleUsecase) {
	c.articleUscase = usecase
}

func (c *ArticleController) FindAll(eCtx echo.Context) (err error) {
	ctx := eCtx.Request().Context()

	res := new(model.Response)

	articles, err := c.articleUscase.FindAll(ctx)

	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	res = model.NewResponse().WithData(articles)

	return eCtx.JSON(http.StatusOK, res)
}

func (c *ArticleController) Create(eCtx echo.Context) (err error) {
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	req := new(model.CreateArticleRequest)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	article := req.ToArticle()

	err = c.articleUscase.Create(ctx, article)

	if err != nil {
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	res = model.NewResponse().WithData(article)

	return eCtx.JSON(http.StatusCreated, res)
}

func (c *ArticleController) FindByID(eCtx echo.Context) (err error) {
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	req := new(model.FindByIDRequest)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	article, err := c.articleUscase.FindByID(ctx, req.ID)

	switch err {
	case nil:
	case model.ErrArticleNotFound:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusNotFound, res)
	default:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	res = model.NewResponse().WithData(article)

	return eCtx.JSON(http.StatusOK, res)
}

func (c *ArticleController) UpdateByID(eCtx echo.Context) (err error) {
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	req := new(model.UpdateArticleRequest)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	article := req.ToArticle()

	err = c.articleUscase.UpdateByID(ctx, article)

	switch err {
	case nil:
	case model.ErrArticleNotFound:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusNotFound, res)
	default:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	res = model.NewResponse().WithData(article)

	return eCtx.JSON(http.StatusOK, res)
}

func (c *ArticleController) DeleteByID(eCtx echo.Context) (err error) {
	ctx := eCtx.Request().Context()

	res := new(model.Response)
	req := new(model.FindByIDRequest)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	err = c.articleUscase.DeleteByID(ctx, req.ID)

	switch err {
	case nil:
	case model.ErrArticleNotFound:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusNotFound, res)
	default:
		res = model.NewResponse().WithMessage(err.Error())
		return eCtx.JSON(http.StatusInternalServerError, res)
	}

	res = model.NewResponse().WithMessage("deleted")

	return eCtx.JSON(http.StatusOK, res)
}
