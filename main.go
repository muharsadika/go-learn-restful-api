package main

import (
	"muharsadika/golang-belajar-restful-api/app"
	"muharsadika/golang-belajar-restful-api/controller"
	"muharsadika/golang-belajar-restful-api/helper"
	"muharsadika/golang-belajar-restful-api/middleware"
	"muharsadika/golang-belajar-restful-api/repository"
	"muharsadika/golang-belajar-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
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

	// Start the server
	go func() {
		err := server.ListenAndServe()
		helper.PanicIfError(err)
	}()

	// Print a message once the server is running
	println("Server is running on port", server.Addr)

	// Block indefinitely to keep the server running
	select {}
}
