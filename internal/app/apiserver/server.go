package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

// ServeHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUserCreate()).Methods("POST")
}

func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
