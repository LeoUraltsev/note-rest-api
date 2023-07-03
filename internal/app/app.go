package app

import (
	"NoteRestApi/config"
	"log"
)

func Run() {
	log.Println("Service start...")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: postgresql

	// TODO: init router: chi

	// TODO: run server
}
