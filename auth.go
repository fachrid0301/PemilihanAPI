package main

import (
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, exists := users[username]
	if !exists || user.Password != password {
		w.Write([]byte("Login failed"))
		return
	}

	if user.Role == 1 {
		http.Redirect(w, r, "/dashboard1", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/dashboard2", http.StatusSeeOther)
	}
}
