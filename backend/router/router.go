package router

import (
    "backend/handlers"
    "backend/handlers/todo"
    "backend/handlers/user"
    "backend/handlers/auth"
    "github.com/gorilla/mux"
)

func Router() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/api/resource", handlers.GetResource).Methods("GET")
    r.HandleFunc("/api/resource", handlers.CreateResource).Methods("POST")

    //todo routes
    r.HandleFunc("/api/todos", todoHandlers.GetAllTodos).Methods("GET")
    r.HandleFunc("/api/todo", todoHandlers.CreateTodoHandler).Methods("POST")
    r.HandleFunc("/api/todo/{userId}", todoHandlers.GetTodoByUserId).Methods("GET")
    r.HandleFunc("/api/todo/{todoId}", todoHandlers.DeleteTodoHandler).Methods("DELETE")
    r.HandleFunc("/api/todo/{todoId}", todoHandlers.UpdateTodoHandler).Methods("PUT")
    r.HandleFunc("/api/todo/mark-complete/{todoId}", todoHandlers.MarkTodoAsCompleteHandler).Methods("PUT")


    //user routes
    r.HandleFunc("/api/users", userHandlers.GetAllUsers).Methods("GET")

    
    // auth routes
    r.HandleFunc("/api/auth/login", authHandlers.Login).Methods("POST")
    r.HandleFunc("/api/auth/register", authHandlers.Register).Methods("POST")
    
    return r
}
