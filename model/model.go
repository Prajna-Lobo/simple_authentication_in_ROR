package model

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type Token struct {
	Token string `json:"token"`
}
type Response struct {
	Data string `json:"data"`
}
type AccessToken struct {
	Token string      `json:"token"`
}

