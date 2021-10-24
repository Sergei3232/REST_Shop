package main

import (
	"fmt"
	"github.com/Sergei3232/REST_Shop/app"
	"github.com/Sergei3232/REST_Shop/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication) // добавляем middleware проверки JWT-токена

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/product/new", controllers.CreateProduct).Methods("POST")

	port := "8000"
	fmt.Println(port)
	err := http.ListenAndServe(":"+port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}

}
