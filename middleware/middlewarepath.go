package middleware

import (
	"net/http"
)

func OnlyPath(path string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w,r)
			return 
		}
		next(w,r)
	}
}