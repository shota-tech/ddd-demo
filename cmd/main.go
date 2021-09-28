package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/shota-tech/ddd-demo/infra"
	"github.com/shota-tech/ddd-demo/interfaces/handler"
	"github.com/shota-tech/ddd-demo/usecase"
)

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		"db", "sample_user", "sample_password", "sample_db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := infra.NewUserRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
	router.HandleFunc("/users", userHandler.GetUserList).Methods("GET")
	router.HandleFunc("/users", userHandler.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", userHandler.EditUser).Methods("PUT")
	router.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
