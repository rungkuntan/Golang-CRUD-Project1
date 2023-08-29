package main

import (
	"fmt"
	"go_project/config"
	"go_project/controller"
	"go_project/errorhelper"
	"go_project/repository"
	"go_project/router"
	"go_project/service"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	
	db := config.DatabaseConnection() //run db
	
	bookRepository := repository.NewBookRepository(db) //repo

	bookService := service.NewBookServiceImpl(bookRepository) //service 

	bookController := controller.NewBookController(bookService) //controller 
	

	routes := router.NewRouter(bookController)	

	
	server := http.Server{Addr: "localhost:8000",Handler:routes}
	err := server.ListenAndServe()
	errorhelper.PanicIfErr(err)

}





