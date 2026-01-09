package main

import "net/http"

func dashboard2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome! Login successful.</h1>"))
}
