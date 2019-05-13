package utils

import (
	"encoding/json"
	"fmt"
	log "github.com/smilemakc/utils/logging"
	"net/http"
	"os"
)

var (
	authorizationString = os.Getenv("BOT_AUTHORIZATION_STRING")
	useragentString     = os.Getenv("BOT_USER_AGENT")
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MarshalResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8\n\n")
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func WriteErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"success": false, "message": "%s"}`, message)))
}

func SetSystemHeaders(req *http.Request) http.Request {
	req.Header.Set("Content-Type", "application/json")
	if useragentString != "" {
		req.Header.Set("User-Agent", useragentString)
	}
	if authorizationString != "" {
		req.Header.Set("Authorization", authorizationString)
	}
	return *req
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug.Printf("request %s from  %s", r.URL.String(), r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
func WriteHeadersMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		next.ServeHTTP(w, r)
	}
}

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) < 1 {
		return h
	}
	wrapped := h
	for i := len(m) - 1; i >= 0; i-- {
		wrapped = m[i](wrapped)
	}
	return wrapped
}

func CommonMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return MultipleMiddleware(h, LoggingMiddleware, WriteHeadersMiddleware)
}
