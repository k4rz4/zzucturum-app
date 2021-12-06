package router

import (
	"net/http"
	"regexp"
)

type Route struct {
	Method  string
	Regex   *regexp.Regexp
	Handler http.HandlerFunc
}

func NewRoute(method, pattern string, handler http.HandlerFunc) Route {
	return Route{method, regexp.MustCompile("^" + pattern + "$"), handler}
}