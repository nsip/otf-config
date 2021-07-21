package config

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	lk "github.com/digisan/logkit"
)

func TestConfig(t *testing.T) {
	cfg := &Config{}
	_, err := toml.DecodeFile("../config.toml", cfg)
	lk.FailOnErr("%v", err)
	fmt.Println("-------------------------------")
	spew.Dump(cfg)
}

func TestGetConfig(t *testing.T) {
	cfg := GetConfig("../config.toml")
	spew.Dump(cfg)
	cfg.Dispense("Hub")
}

func TestSave(t *testing.T) {
	cfg := GetConfig("../config.toml")

	cfg.NatsStreamings = append(cfg.NatsStreamings, NatsStreaming{
		Name: "test_name",
		Path: "test_path",
	})

	cfg.SaveAsJson()
	cfg.SaveToml()
}
