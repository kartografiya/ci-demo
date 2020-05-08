package main

import (
	"log"
	"net/http"

	env "github.com/caarlos0/env/v6"
	"github.com/kartografiya/ci-demo/functions"
)

type config struct {
	Port string `env:"PORT" envDefault:"3000"`
}

func main() {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", functions.CalcHash)
	log.Printf("Listening on port %s", cfg.Port)
	http.ListenAndServe(":"+cfg.Port, nil)
}
