package routes

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func newProxy(target string) http.HandlerFunc {
	parsedURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("invalid proxy target %q: %v", target, err)
	}

	proxy := httputil.NewSingleHostReverseProxy(parsedURL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error: %v", err)
		http.Error(w, "bad gateway", http.StatusBadGateway)
	}
	return proxy.ServeHTTP
}

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	auth := "http://auth-service:8081"
	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/register", newProxy(auth))
		r.Post("/login", newProxy(auth))
		r.Post("/refresh", newProxy(auth))
		r.Post("/logout", newProxy(auth))
	})

	reviewService := "http://review:8082"
	router.Route("/api/review", func(r chi.Router) {
		r.Post("/create", newProxy(reviewService))
		r.Post("/update", newProxy(reviewService))
		r.Post("/delete/{id}", newProxy(reviewService))
		r.Post("/get", newProxy(reviewService))
	})

	mentorService := "http://mentor:8083"
	router.Route("/api/mentors", func(r chi.Router) {
		r.Get("/get", newProxy(mentorService))
	})

	return router
}
