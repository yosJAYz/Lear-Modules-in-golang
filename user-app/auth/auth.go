package auth

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
	}

	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Username == "admin" && req.Password == "1234" {
		w.Write([]byte("Login Berhasil"))
	} else {
		http.Error(w, "Login Gagal", http.StatusUnauthorized)
	}
}
