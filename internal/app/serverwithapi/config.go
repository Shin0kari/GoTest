// внутренние параметры сервера
// задаётся через toml
package serverwithapi

import "github.com/Shin0kari/TestAPIonGo/internal/app/store"

// Config - формирует параметры, чтобы их легче было использовать и не писать
// сотню раз.
type Config struct {
	// место, где расположен аддрес нашего сервера
	BindAddr string `toml:"bind_addr"`
	// задаём значение, при котором будет решаться, выводить ли
	// в консоль инфо или нет(подразделяется на уровни логирования)
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig - заполняет параметры нашего сервера
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		// уровни искать в интернете
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
