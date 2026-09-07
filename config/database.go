package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func Connect() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?parseTime=True"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err = db.Ping(); err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	log.Println("Berhasil terhubung ke database")
	return db
}