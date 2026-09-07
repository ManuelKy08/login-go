package routes

import (
	"login-go/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Halaman utama (cek login)
	mux.HandleFunc("/", handlers.IndexHandler)

	// Halaman auth
	mux.HandleFunc("/login", handlers.LoginHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)

	// Logout
	mux.HandleFunc("/logout", handlers.LogoutHandler)

	// Dashboard (auth check inside handler)
	mux.HandleFunc("/dashboard", handlers.DashboardHandler)

	// Static files
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))
}
