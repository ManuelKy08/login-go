package main

import (
	"log"
	"net/http"

	"login-go/config"
	"login-go/routes"
)

func main() {
	// Ambil konfigurasi dari environment
	port := config.GetEnv("APP_PORT", "8080")

	// Setup router
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Printf("🚀 Server berjalan di http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
