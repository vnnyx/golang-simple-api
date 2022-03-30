package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"simple-api/app"
	"simple-api/controller"
	"simple-api/exception"
	"simple-api/helper"
	"simple-api/middleware"
	"simple-api/repository"
	"simple-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	router := httprouter.New()
	router.GET("/api/customer", customerController.FindAll)
	router.GET("/api/customer/:customerId", customerController.FindById)
	router.POST("/api/customer", customerController.Create)
	router.PUT("/api/customer/:customerId", customerController.Update)
	router.DELETE("/api/customer/:customerId", customerController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
