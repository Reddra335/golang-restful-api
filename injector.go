//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	app "restful_api/APP"
	"restful_api/controller"
	"restful_api/middleware"
	"restful_api/repository"
	"restful_api/service"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(repository.NewCategoryRepository, wire.Bind(new(repository.CategoryRepository),
	new(*repository.CategoryRepositoryImpln)), service.NewCategoryService, wire.Bind(new(service.CategoryService),
	new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController),
		new(*controller.CategoryControllerImpln)),
)

func InitializedServer() *http.Server {

	wire.Build(
		app.NewDB,
		app.RedValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(
			new(http.Handler),
			new(*httprouter.Router),
		),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
