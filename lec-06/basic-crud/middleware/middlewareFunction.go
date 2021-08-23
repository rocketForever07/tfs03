package middleware

import (
	"fmt"
	"net/http"
)

func ContentTypeChecking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")

		if reqContentType != "application/json"{
			fmt.Fprintf(w, "request only alow content type application/json")
			http.Redirect(w, r, "http://www.google.com", 301)
		}else{
			w.Write([]byte("alow content type application/json"))
		}

		next.ServeHTTP(w,r)
	})
}

func LoginChecking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		if r.Header.Get("Authorization") == ""{
			fmt.Fprintf(w, "you need to login\n")
			// http.Redirect(w,r,"http://localhost:10000/peoples",http.StatusOK)
		}
	})
}