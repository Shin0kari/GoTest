package serverwithapi

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//ServerAPI ...
type ServerAPI struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

//New ...
func New(config *Config) *ServerAPI {
	return &ServerAPI{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *ServerAPI) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *ServerAPI) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *ServerAPI) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ServerAPI) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
