package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/handler"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/middleware"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/server"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	server, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal("error creating server", err)
	}

	server.Start(BindRoutes)
}

func BindRoutes(server server.Server, r *mux.Router) {
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middleware.CheckAuthMiddleware(server))

	r.HandleFunc("/", handler.HomeHandler(server)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handler.SignUpHandler(server)).Methods(http.MethodPost)
	r.HandleFunc("/login", handler.LoginHandler(server)).Methods(http.MethodPost)
	api.HandleFunc("/me", handler.MeHandler(server)).Methods(http.MethodGet)
	api.HandleFunc("/posts", handler.InsertPostHandler(server)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id}", handler.GetPostByIDHandler(server)).Methods(http.MethodGet)
	api.HandleFunc("/posts/{id}", handler.UpdatePostHandler(server)).Methods(http.MethodPut)
	api.HandleFunc("/posts/{id}", handler.DeletePostHandler(server)).Methods(http.MethodDelete)
	r.HandleFunc("/posts", handler.ListPostsHandler(server)).Methods(http.MethodGet)
	r.HandleFunc("/ws", server.Hub().HandleWebsocket)
}
