package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"learn-restful-api/app"
	"learn-restful-api/controller"
	"learn-restful-api/helper"
	"learn-restful-api/middleware"
	"learn-restful-api/repository"
	"learn-restful-api/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server Starting . . .")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
