package main

type User struct {
	Password string
	Role     int
}

var users = map[string]User{
	"haidar": {
		Password: "1738",
		Role:     1,
	},
	"exejar": {
		Password: "1234",
		Role:     2,
	},
	"finn": {
		Password: "1234",
		Role:     2,
	},
	"meme": {
		Password: "1010",
		Role:     1,
	},
}
