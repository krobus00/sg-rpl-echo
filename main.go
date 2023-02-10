package main

import (
	"github.com/krobus00/sg-rpl-echo/controller"
	"github.com/krobus00/sg-rpl-echo/repository"
	"github.com/krobus00/sg-rpl-echo/usecase"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	articleRepository := repository.NewArticleRepository()

	articleUsecase := usecase.NewArticleUsecase()
	articleUsecase.RegisterArticleRepository(articleRepository)

	articleController := controller.NewArticleController()
	articleController.RegisterArticleUsecase(articleUsecase)

	// routing
	e.GET("/articles", articleController.FindAll)
	e.POST("/articles", articleController.Create)
	e.GET("/articles/:id", articleController.FindByID)
	e.PUT("/articles/:id", articleController.UpdateByID)
	e.DELETE("/articles/:id", articleController.DeleteByID)

	// run server
	e.Logger.Fatal(e.Start(":3000"))
}
