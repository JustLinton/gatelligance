package utils

type RegisterResponse struct {
	IsSuccess string
	ErrorMsg  string
}

type LoginResponse struct {
	Token     string
	IsSuccess string
	ErrorMsg  string
}

type WorkSubmitResponse struct {
	TaskList  string
	IsSuccess string
	ErrorMsg  string
}
