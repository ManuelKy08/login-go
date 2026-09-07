package handlers

import (
	"net/http"
	"strings"
	"text/template"
	"time"
	"strconv"
	"encoding/base64"
	"login-go/services"
)

func saveSession(w http.ResponseWriter, userID int) {
	token := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(userID) + "|" + time.Now().Format("2006-01-02")))
	cookie := &http.Cookie{
		Name:     "session_token",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 7,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
}

func getUserIDFromContext(r *http.Request) (int, bool) {
	v := r.Context().Value("userID")
	if v == nil {
		return 0, false
	}
	id, ok := v.(int)
	return id, ok
}

func renderTemplate(w http.ResponseWriter, tmplName string, data interface{}) {
	tmplPath := "templates/" + tmplName + ".html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template not found: "+tmplPath, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

// LoginHandler menangani permintaan login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "login", nil)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		email := strings.TrimSpace(r.FormValue("email"))
		password := r.FormValue("password")

		if email == "" || password == "" {
			renderTemplate(w, "login", map[string]string{
				"Error": "Email dan password wajib diisi",
			})
			return
		}

		service := services.NewAuthService()
		user, err := service.Login(email)
		if err != nil {
			renderTemplate(w, "login", map[string]string{
				"Error": "Email atau password yang Anda masukkan salah",
			})
			return
		}

		if !service.VerifyPassword(password, user.PasswordHash) {
			renderTemplate(w, "login", map[string]string{
				"Error": "Email atau password yang Anda masukkan salah",
			})
			return
		}

		saveSession(w, user.ID)

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

// RegisterHandler menangani registrasi pengguna baru
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	service := services.NewAuthService()
	if r.Method == http.MethodGet {
		renderTemplate(w, "register", nil)
		return
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		email := strings.TrimSpace(r.FormValue("email"))
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		if name == "" || email == "" || password == "" {
			renderTemplate(w, "register", map[string]string{
				"Error": "Semua field wajib diisi",
			})
			return
		}

		if password != confirmPassword {
			renderTemplate(w, "register", map[string]string{
				"Error": "Password dan konfirmasi password tidak cocok",
			})
			return
		}

		if len(password) < 8 {
			renderTemplate(w, "register", map[string]string{
				"Error": "Password minimal 8 karakter",
			})
			return
		}

		err := service.Register(name, email, password)
		if err != nil {
			if strings.Contains(err.Error(), "sudah terdaftar") {
				renderTemplate(w, "register", map[string]string{
					"Error": "Email sudah terdaftar",
				})
			} else {
				renderTemplate(w, "register", map[string]string{
					"Error": "Terjadi error: " + err.Error(),
				})
			}
			return
		}

		renderTemplate(w, "login", map[string]string{
			"Success": "Registrasi berhasil! Silakan login.",
		})
	}
}

// DashboardHandler halaman dashboard
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserIDFromContext(r)
	if userID == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	service := services.NewAuthService()
	user, err := service.GetUserByID(userID)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	data := map[string]interface{}{
		"FullName": user.Name,
		"Email":    user.Email,
		"Role":     user.Role,
	}
	renderTemplate(w, "dashboard", data)
}

// LogoutHandler menghapus session
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		HttpOnly: true,
		Path:   "/",
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// IndexHandler redirect berdasarkan status login
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := getUserIDFromContext(r)
	if userID != 0 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
