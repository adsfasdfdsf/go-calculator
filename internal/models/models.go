package models

type Request struct {
	Expr string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
