package middleware

import (
	"net/http"
	"strings"

	"login-go/models"
)

// SessionData merepresentasikan data session
type SessionData struct {
	UserID *int
}

// AuthMiddleware memverifikasi session yang valid
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if r.URL.Path != "/login" && r.URL.Path != "/register" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
			next.ServeHTTP(w, r)
			return
		}

		values := strings.Split(cookie.Value, "|")
		if len(values) != 2 {
			http.Error(w, "Sesi tidak valid", http.StatusUnauthorized)
			return
		}

		var userID int
		_, err := fmt.Sscanf(values[0], "%d", &userID)
		if err != nil {
			http.Error(w, "Sesi tidak valid", http.StatusUnauthorized)
			return
		}

		// Sementara: ambil user dengan service (akan dioptimalkan di handler)
		// Untuk sekarang, hanya set context jika userID > 0
		if userID > 0 {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", userID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		if r.URL.Path != "/login" && r.URL.Path != "/register" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
		next.ServeHTTP(w, r)
	})
}

// SecurityHeaders menambahkan header keamanan
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w, r)
	})
}