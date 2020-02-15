package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kapustaprusta/simple_rest_server/internal/app/store"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// NewAPIServer ...
func NewAPIServer(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	err := s.configureLogger()
	if err != nil {
		return err
	}

	s.configureRouter()
	err = s.configureStore()
	if err != nil {
		return err
	}

	s.logger.Info("starting api server at ", s.config.BindAddr)

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.NewStore(s.config.StoreConfig)
	err := st.Open()
	if err != nil {
		return err
	}

	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	}
}
