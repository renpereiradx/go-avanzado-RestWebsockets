package helpers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/server"
)

func GetJWTAuthorizationInfo(s server.Server, w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Config().JWTSecret), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	return token, err
}
