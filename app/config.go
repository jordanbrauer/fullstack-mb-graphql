package app

import (
	"log"

	"gopkg.in/ini.v1"
)

func Config(path string) *ini.File {
	config, err := ini.Load(path)

	if err != nil {
		log.Panic("cannot load config file at: ", path, err)
	}

	return config
}
