package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Shin0kari/TestAPIonGo/internal/app/serverwithapi"
)

var (
	// создаём переменную, в которую сунем параметры сервера из
	// toml
	configPath string
)

// в функции init, код запускается перед любой частью кода
func init() {
	// суём в configPath штуки из toml
	flag.StringVar(&configPath, "config-path", "configs/serverapi.toml", "path to config file")
}

func main() {
	// для того, чтобы все флаги разобрались(анализировались)
	flag.Parse()
	// как я понял, соединяем config с присваиванием новых
	// параметров сервера
	config := serverwithapi.NewConfig()
	// из configPath суём в config
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	// создаём апи сервер c параметрами из config
	s := serverwithapi.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
