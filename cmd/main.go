package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/WatchJani/configuration/cmd/internal/cli"
)

var (
	ErrNotSupported = errors.New("pq: Unsupported command")
)

func ReadConf(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type Config struct {
	Mod string `json:"mod"`
}

type Mod interface {
	DoSomething()
}

type Alive struct{}

func (a *Alive) DoSomething() {
	fmt.Println("alive")
}

type Active struct{}

func (a *Active) DoSomething() {
	fmt.Println("active")
}

func ChoseMode(mod string) (Mod, error) {
	switch mod {
	case "active":
		return &Active{}, nil
	case "alive":
		return &Alive{}, nil
	default:
		return nil, ErrNotSupported
	}
}

func main() {
	data, err := ReadConf(cli.SetupFlags())
	if err != nil {
		os.Exit(1)
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Println(err)
	}

	mod, err := ChoseMode(config.Mod)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	mod.DoSomething()
}
