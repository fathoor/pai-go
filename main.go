package main

import (
	"log"
	"net/http"
)

type wrappedWriter struct {
	http.ResponseWriter
	Status int
}

func (w *wrappedWriter) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
	w.Status = code
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writer := &wrappedWriter{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		next.ServeHTTP(writer, r)
		log.Printf("[%v] %v %v\n", writer.Status, r.Method, r.URL.Path)
	})
}

func home(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	if _, err := w.Write([]byte("are you lost? try visiting /flag you might find something useful")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func flag(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("flag: pai{p12aK7ikUm_w1r3sH4rk}")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func main() {
	router := http.NewServeMux()
	{
		router.HandleFunc("GET /", home)
		router.HandleFunc("GET /flag", flag)
	}

	server := http.Server{
		Addr:    "0.0.0.0:1337",
		Handler: logger(router),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
