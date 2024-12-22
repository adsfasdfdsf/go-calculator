package server

import (
	"encoding/json"
	"fmt"
	"go-calculator/internal/models"
	"io"
	"net/http"
)

type Repo interface {
	Calc(expr string) (float64, error)
}

type Server struct {
	port int
	repo Repo
}

func NewServer(port int, repo Repo) *Server {
	return &Server{port: port, repo: repo}
}

func (s *Server) Handler(r http.ResponseWriter, req *http.Request) {
	data, err := io.ReadAll(req.Body)
	if err != nil {
		var resp models.ErrorResponse
		resp.Error = "Expression is not valid"
		fmt.Println(err)
		data, _ = json.Marshal(resp)
		_, _ = r.Write(data)
		return
	}
	var expr models.Request
	err = json.Unmarshal(data, &expr)
	if err != nil {
		var resp models.ErrorResponse
		resp.Error = "Internal server error"
		fmt.Println(err, data, expr)
		data, _ = json.Marshal(resp)
		r.WriteHeader(http.StatusInternalServerError)
		_, _ = r.Write(data)
		return
	}
	answer, err := s.repo.Calc(expr.Expr)

	if err != nil {
		var resp models.ErrorResponse
		resp.Error = "Expression is not valid"
		fmt.Println(err)
		data, _ = json.Marshal(resp)
		r.WriteHeader(http.StatusInternalServerError)
		_, _ = r.Write(data)
		return
	}

	var resp models.Response
	resp.Result = fmt.Sprint(answer)
	data, err = json.Marshal(resp)

	r.WriteHeader(http.StatusOK)
	_, _ = r.Write(data)

	defer req.Body.Close()
}

func (s *Server) Start() error {
	http.Handle("/api/v1/calculate", http.HandlerFunc(s.Handler))

	return http.ListenAndServe(fmt.Sprintf(":%v", s.port), nil)
}
