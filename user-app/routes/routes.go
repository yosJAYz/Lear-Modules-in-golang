package routes

import (
	"net/http"
	"user-app/auth"
	"user-app/user"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", auth.LoginHandler)
	mux.HandleFunc("/users", user.GetUserHandler)
	return mux
}
