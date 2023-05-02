package main

import (
	"github.com/df-mc/structure"
	"github.com/schphe/libabel/library"
	"github.com/schphe/libabel/library/session"
	"log"
)

func main() {
	str, err := structure.ReadFile("./assets/library.mcstructure")
	if err != nil {
		log.Fatalf("couldn't read structure file: %v", err)
	}
	gen := library.NewGenerator(str)

	cfg, err := library.DefaultConfig(gen)
	if err != nil {
		log.Fatalf("couldn't create config: %v", err)
	}

	pro, err := session.NewProvider("./players")
	if err != nil {
		log.Fatalf("couldn't create session data provider: %v", err)
	}

	lib := library.New(cfg, pro)
	lib.Start()
}
