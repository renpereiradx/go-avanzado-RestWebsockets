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

func BindRoutes(s server.Server, r *mux.Router) {
	middleware.CheckAuthMiddleware(s)

	r.HandleFunc("/", handler.HomeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/signup", handler.SignUpHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/login", handler.LoginHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/me", handler.MeHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts", handler.InsertPostHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/posts/{id}", handler.GetPostByIDHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/posts/{id}", handler.UpdatePostHandler(s)).Methods(http.MethodPut)
	r.HandleFunc("/posts/{id}", handler.DeletePostHandler(s)).Methods(http.MethodDelete)
}
