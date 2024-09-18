package apiserver

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {

	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	s.configRouter()
	return http.ListenAndServe(s.config.BindAddr, s.router)

}

func (s *APIServer) configRouter() {

	s.router.HandleFunc("/hello", s.handleHello())

}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
