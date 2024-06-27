package middlewares

import "net/http"

func securityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sessionValid(w, r) {
			if r.URL.Path == "/login" {
				next.ServeHTTP(w, r)
				return
			}
		}

		if hasBeenAuthenticated(w, r) {
			next.ServeHTTP(w, r)
			return
		}

		//otherwise it will need to be redirected to /login
		storeAuthenticated(w, r, false)
		http.Redirect(w, r, "/login", 307)
	})
}
