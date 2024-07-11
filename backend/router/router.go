package router

import (
	"backend/handlers"
	authHandlers "backend/handlers/auth"
	todoHandlers "backend/handlers/todo"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/resource", handlers.GetResource).Methods("GET")
	r.HandleFunc("/api/resource", handlers.CreateResource).Methods("POST")

	r.HandleFunc("/api/todo", middleware.VerifyToken(todoHandlers.CreateTodoHandler)).Methods("POST")
	r.HandleFunc("/api/user/todos", middleware.VerifyToken(todoHandlers.GetTodoByUserId)).Methods("GET")
	r.HandleFunc("/api/todo/{todoId}", middleware.VerifyToken(todoHandlers.DeleteTodoHandler)).Methods("DELETE")
	r.HandleFunc("/api/todo/{todoId}", middleware.VerifyToken(todoHandlers.UpdateTodoHandler)).Methods("PUT")
	r.HandleFunc("/api/todo/{todoId}/complete", middleware.VerifyToken(todoHandlers.MarkTodoAsCompleteHandler)).Methods("PUT")

	// auth routes
	r.HandleFunc("/api/auth/login", authHandlers.Login).Methods("POST")
	r.HandleFunc("/api/auth/register", authHandlers.Register).Methods("POST")

	return r
}
