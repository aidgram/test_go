package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// LoggingMiddleware логирует запросы и коды ответов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Оборачиваем ResponseWriter, чтобы перехватить статус-код
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)

		log.Printf("%s %s %d %s", r.Method, r.URL.Path, lrw.statusCode, duration)
	})
}

// Обёртка ResponseWriter для перехвата statusCode
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Хэндлер
func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"test": "tesalanoskusovvtigran",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", testHandler)

	loggedMux := LoggingMiddleware(mux)

	log.Println("test30")
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
