package http_transport

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"zzucturum-app/pkg/http-transport/router"
	"zzucturum-app/pkg/utils"
)

var (
	errTokenMismatch = errors.New("token mismatched")
	errMissingToken  = errors.New("missing token")
)

type Server struct {
	SecureToken      string
	EnableCORS       bool
	Routes           []router.Route
}

type ctxKey struct{}

func NewServer(token string, routes []router.Route) Server {
	return Server{
		SecureToken:      token,
		Routes: 		  routes,
	}
}

func (s Server) checkToken(r *http.Request) error {
	token := r.URL.Query().Get("token")
	if token == "" {
		token = r.FormValue("token")
	}
	if token == "" {
		return errMissingToken
	}

	if token != s.SecureToken {
		return errTokenMismatch
	}
	return nil
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err:= s.checkToken(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		utils.WriteError(w, err)
		return
	}

	var allow []string
	for _, route := range s.Routes {
		matches := route.Regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.Method {
				allow = append(allow, route.Method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.Handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	utils.WriteError(w, fmt.Errorf("Route \"%s\" is not found", r.URL.Path))

}
