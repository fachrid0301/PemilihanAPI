package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", loginPage)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/dashboard1", dashboard1)
	http.HandleFunc("/dashboard2", dashboard2)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
