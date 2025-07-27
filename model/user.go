package model

type Response struct {
	Username string `json:"username"`
	Msg      string `json:"msg"`
	Token    string `json:"token"`
	Status   int    `json:"status"`
}

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

type User struct {
	ID       int
	Username string `json:"username"`
	Password string `json:"password"`
}
