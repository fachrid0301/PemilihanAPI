package main

import (
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) {
	html := `
	<!DOCTYPE html>
	<html>
	<body>
		<h2>Login</h2>
		<form method="POST" action="/login">
			Username: <input type="text" name="username"><br><br>
			Password: <input type="password" name="password"><br><br>
			<button type="submit">Login</button>
		</form>
	</body>
	</html>
	`
	w.Write([]byte(html))
}
