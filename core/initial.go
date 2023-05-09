package core

import (
	"log"
	"os"
)

type NanaYaml struct {
	Loading  string `yaml:"loading"`
	LogLevel string `yaml:"log_level"`
}

func CreateNanaYaml() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := os.Stat(dirname + "/.nana/nana.yaml"); os.IsNotExist(err) {
		_, err := os.Create(dirname + "/.nana/nana.yaml")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func CreateNanaFolder() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return
	}

	if _, err := os.Stat(dirname + "/.nana"); os.IsNotExist(err) {
		os.Mkdir(dirname+"/.nana", os.ModePerm)
	}
}
