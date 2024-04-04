package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/helpers"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/repository"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/server"
	"github.com/segmentio/ksuid"
)

type InsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	ID          string `json:"id"`
	PostContent string `json:"post_content"`
}

type UpdatePostRequest struct {
	PostContent string `json:"post_content"`
}

type UpdatePostResponse struct {
	Message string `json:"message"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := helpers.GetJWTAuthorizationInfo(s, w, r)
		if err != nil {
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var postRequest = InsertPostRequest{}
			err := json.NewDecoder(r.Body).Decode(&postRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			post := models.Posts{
				ID:          id.String(),
				PostContent: postRequest.PostContent,
				UserID:      claims.UserId,
			}
			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(PostResponse{
				ID:          post.ID,
				PostContent: post.PostContent,
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetPostByIDHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetPostByID(r.Context(), params["id"])
		if err != nil {
			log.Println("error getting post by id", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		token, err := helpers.GetJWTAuthorizationInfo(s, w, r)
		if err != nil {
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var updatePostRequest = UpdatePostRequest{}
			err := json.NewDecoder(r.Body).Decode(&updatePostRequest)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			post := models.Posts{
				ID:          params["id"],
				PostContent: updatePostRequest.PostContent,
			}
			err = repository.UpdatePost(r.Context(), &post, claims.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(UpdatePostResponse{
				Message: "Post updated successfully",
			})
		}
	}
}

func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		token, err := helpers.GetJWTAuthorizationInfo(s, w, r)
		if err != nil {
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			err := repository.DeletePost(r.Context(), params["id"], claims.UserId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(UpdatePostResponse{
				Message: "Post deleted successfully",
			})
		}
	}
}
