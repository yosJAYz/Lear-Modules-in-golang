package user

import (
	"encoding/json"
	"net/http"
)

var dummyUsers = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only Get allowed", http.StatusMethodNotAllowed)
		return
	}
	json.NewEncoder(w).Encode(dummyUsers)
}
