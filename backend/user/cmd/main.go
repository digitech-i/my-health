package main

import (
	"database/sql"
	"fmt"
	"log"
	"myhealth/backend/user/internal/handler"
	"myhealth/backend/user/internal/repository"
	"myhealth/backend/user/internal/service"
	"net/http"
)

func main() {
	// Database connection setup
	db, err := sql.Open("mysql", "username:password@tcp(host:port)/database")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Repository
	userRepository := repository.NewUserRepository(db)

	// Service
	userService := service.NewUserService(userRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/createUser", userHandler.CreateUser)
	http.HandleFunc("/getUsers", userHandler.GetUsers)

	port := 8080
	fmt.Printf("Server listening on port %d...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
