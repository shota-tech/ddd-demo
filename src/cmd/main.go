package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/shota-tech/layered-architecture-demo/src/infra"
	"github.com/shota-tech/layered-architecture-demo/src/interfaces/handler"
	"github.com/shota-tech/layered-architecture-demo/src/usecase"
)

func main() {
	dsn := "host=db user=sample_user password=sample_password dbname=sample_db sslmode=disable"
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

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
