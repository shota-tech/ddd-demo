package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shota-tech/layered-architecture-demo/src/infra"
	"github.com/shota-tech/layered-architecture-demo/src/interfaces/handler"
	"github.com/shota-tech/layered-architecture-demo/src/usecase"
)

func main() {
	userRepository := infra.NewUserRepository()
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/user", userHandler.GetUserList).Methods("GET")
	router.HandleFunc("/user", userHandler.AddUser).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
