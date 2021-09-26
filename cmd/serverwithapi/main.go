package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"

	"github.com/Shin0kari/APITest.git/internal/app/serverwithapi"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/serverapi.toml", "path to config file")
}

func main() {
	flag.Parse()
	//
	config := serverwithapi.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	//
	s := serverwithapi.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
