package api

import "net/http"

// HeadersMiddleware устанавливает заголовки ответа сервера.
func (api *API) HeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Manual-Header", "I love you")
		next.ServeHTTP(w, r)
	})
}