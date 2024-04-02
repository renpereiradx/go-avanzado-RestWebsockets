package middleware

import (
	"net/http"
	"strings"

	"github.com/renpereiradx/go-avanzado-RestWebsocket/helpers"
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
			_, err := helpers.GetJWTAuthorizationInfo(s, w, r)
			if err != nil {
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
