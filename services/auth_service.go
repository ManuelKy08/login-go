package services

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"login-go/config"
	"login-go/models"
)

type AuthService struct {
	DB *sql.DB
}

func NewAuthService() *AuthService {
	return &AuthService{DB: config.Connect()}
}

// HashPassword menghash password menggunakan bcrypt
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCostFactor)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword membandingkan password dengan hash
func (s *AuthService) VerifyPassword(plainPassword, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))
	return err == nil
}

// Login memeriksa kredensial
func (s *AuthService) Login(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email, role, created_at FROM users WHERE email = ?"
	row := s.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("akun tidak ditemukan")
		}
		return nil, errors.New("error internal server")
	}
	return &user, nil
}

// Register membuat user baru
func (s *AuthService) Register(name, email, password string) error {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	row := s.DB.QueryRow(query, email)
	if err := row.Scan(&count); err != nil {
		return errors.New("error validasi email")
	}
	if count > 0 {
		return errors.New("email sudah terdaftar")
	}

	hashedPW, err := s.HashPassword(password)
	if err != nil {
		return errors.New("error membuat password hash")
	}

	insertQuery := "INSERT INTO users (name, email, password_hash, role, created_at) VALUES (?, ?, ?, 'member', NOW())"
	_, err = s.DB.Exec(insertQuery, name, email, hashedPW)
	if err != nil {
		return errors.New("error menyimpan ke database")
	}
	return nil
}

// GetUserByID
func (s *AuthService) GetUserByID(id int) (*models.User, error) {
	query := "SELECT id, name, email, role, created_at FROM users WHERE id = ?"
	row := s.DB.QueryRow(query, id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}
	return &user, nil
}

// UpdateLastLogin
func (s *AuthService) UpdateLastLogin(id int) error {
	query := "UPDATE users SET last_login = NOW() WHERE id = ?"
	_, err := s.DB.Exec(query, id)
	return err
}

const bcryptCostFactor = 12