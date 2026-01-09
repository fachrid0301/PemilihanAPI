package main

import "net/http"

func dashboard1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome admin! Login successful.</h1>"))
}
