// пакет с самим сервером

package serverwithapi

import (
	"io"
	"net/http"

	"github.com/Shin0kari/TestAPIonGo/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// ServerAPI - структура самого сервера
type ServerAPI struct {
	// параметры сервера
	config *Config
	// параметры вывода в консоль
	logger *logrus.Logger
	// роутер
	router *mux.Router
	store  *store.Store
}

// New - заполняем структуру сервера
func New(config *Config) *ServerAPI {
	return &ServerAPI{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start - функция для запуска сервера
func (s *ServerAPI) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")

	// функция для запуска веб-приложения, 1 параметр - адрес, 2 - интерфейс(обработчик запроса)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// настройка(конфигурация логов)
func (s *ServerAPI) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	// задаём логу уровень
	s.logger.SetLevel(level)

	return nil
}

// конфигурация роутера(в данном случае добавления "листа" /hello)
func (s *ServerAPI) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *ServerAPI) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	return nil
}

// как я понял, это функция, которая обрабатывает запрос
func (s *ServerAPI) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
