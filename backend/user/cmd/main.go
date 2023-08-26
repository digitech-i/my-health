package main

import (
	"fmt"
	"myhealth/backend/user/internal/handler"
	"myhealth/backend/user/internal/service"
	"net/http"
)

func main() {
	userService := service.NewUserService()
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", userHandler.GetAllUsers)
	http.HandleFunc("/users/add", userHandler.AddUser)

	port := 8080
	fmt.Printf("Server started on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
