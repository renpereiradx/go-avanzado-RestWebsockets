package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/models"
	"github.com/renpereiradx/go-avanzado-RestWebsocket/server"
)

// Creo un slice de Strings con los paths
var (
	NO_AUTH_NEEDED = []string{
		"login",
		"signup",
	}
)

// Revisamos si algun path debe ser validado
func shouldCheckToken(route string) bool {
	for _, p := range NO_AUTH_NEEDED {
		if strings.Contains(route, p) {
			return false
		}
	}
	return true
}

// Aca el Middleware checkeara si se debe realizar la comprobación del token de autorización para una solicitud HTTP entrante.
func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Aca hace la condicion
			if !shouldCheckToken(r.URL.Path) {
				// Si no lo tiene que checkear, retorna el handler
				next.ServeHTTP(w, r)
				return
			}
			// Del encabezado (Header) de la solicitud HTTP se OBTIENE el "Authorization" donde esta el Token y se le saca los espacios
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			// Analiza (Parsea) el tokenString con los claims que creamos en Models.
			// Tambien realiza una funcion anonima donde retornara el JWTSecret del "s" (Server)
			_, err := jwt.ParseWithClaims(tokenString, models.AppClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			// SI tiene error, el token esta mal y tira Unauthorized
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// Retorna el handler
			next.ServeHTTP(w, r)
		})
	}
}
