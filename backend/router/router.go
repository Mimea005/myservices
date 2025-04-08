package router

import (
	"log"
	"net/http"
	"strings"
)

type Router struct {
	*http.ServeMux
	original *http.ServeMux
}

func (router *Router) Handle(pattern string, handler http.Handler)  {
	router.original.Handle(pattern, handler)
}

func (router *Router) HandleFunc(pattern string, handler http.HandlerFunc) {
	router.original.HandleFunc(pattern, handler)
}

func NewRouter() *Router {
	mux := http.NewServeMux()
	return &Router{mux, mux}
}

// Strips the prefix from the request path before invoking the handler
func (r *Router) OnSubPath(path string, handler http.Handler) {
	pathSlice := strings.Split(path, "/")
	pathSlice = pathSlice[:len(pathSlice)-1]
	prefix := strings.Join(pathSlice, "/")
	log.Printf("Prefix: %s", prefix)
	r.Handle(path, http.StripPrefix(prefix, handler))
}


// Register a middleware for all handlers of this router
func (r *Router) UseMiddleware(middleware func(http.Handler) http.Handler) {
	nr := http.NewServeMux()
	nr.Handle("/", middleware(r.ServeMux))
	r.ServeMux = nr
}


func WithMiddleware(handler http.HandlerFunc, middleware ...func(http.Handler) http.Handler) http.Handler {
	var next http.Handler = http.HandlerFunc(handler)
	for k:=len(middleware)-1;k>=0;k++ {
		next=middleware[k](next)
	}

	return next
}
