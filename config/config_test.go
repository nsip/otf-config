package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"github.com/digisan/gotk/io"
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
	cfg.Dispense()
}

func TestCreateJSON(t *testing.T) {
	cfg := GetConfig("../config.toml")
	bytes, err := json.Marshal(cfg)
	lk.FailOnErr("%v", err)
	io.MustWriteFile("config.json", bytes)
}
