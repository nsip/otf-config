package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
	fd "github.com/digisan/gotk/filedir"
	"github.com/digisan/gotk/io"
	lk "github.com/digisan/logkit"
)

func init() {
	lk.Log2F(true, "../otf-config.log")
}

// Config :
type Config struct {
	using          string
	Name           string
	Port           int
	API            string
	PageFolder     string
	NatsStreamings NatsStreamingGrp
	Nias3s         Nias3Grp
	Benthoses      BenthosGrp
	Readers        ReaderGrp
	Aligns         AlignGrp
	TxtClassifiers TxtClassifierGrp
	Levels         LevelGrp
	Weights        WeightGrp
	Hubs           HubGrp
}

func (cfg *Config) Validate() (err error) {

	// natsstreaming
	if err = cfg.NatsStreamings.Validate(); err != nil {
		return
	}

	// nias3
	if err = cfg.Nias3s.Validate(); err != nil {
		return
	}

	// benthos
	if err = cfg.Benthoses.Validate(); err != nil {
		return
	}

	// reader
	if err = cfg.Readers.Validate(); err != nil {
		return
	}

	// align
	if err = cfg.Aligns.Validate(); err != nil {
		return
	}

	// text classifier
	if err = cfg.TxtClassifiers.Validate(); err != nil {
		return
	}

	// level
	if err = cfg.Levels.Validate(); err != nil {
		return
	}

	// weight
	if err = cfg.Weights.Validate(); err != nil {
		return
	}

	// hub
	if err = cfg.Hubs.Validate(); err != nil {
		return
	}

	return
}

func (cfg *Config) procAPI() {

	withSlash := func(str string) string {
		return "/" + strings.Trim(str, "/")
	}

	// natsstreaming

	// nias3

	// benthos

	// reader

	// align

	// text classifier

	// level

	// weight
	for i := 0; i < len(cfg.Weights); i++ {
		sub := &cfg.Weights[i]
		sub.Service.API = withSlash(sub.Service.API)
	}

	// hub
}

// func (cfg *Config) procEnv() {

// 	// natsstreaming

// 	// nias3

// 	// benthos

// 	// reader

// 	// align

// 	// text classifier

// 	// level

// 	// weight
// 	// os.Setenv("OTF_WEIGHT_FAILWHENERR", strconv.FormatBool(cfg.Weight.FailWhenErr))

// 	// hub
// }

// GetConfig :
func GetConfig(configs ...string) *Config {
	for _, config := range configs {
		cfg := &Config{}
		_, err := toml.DecodeFile(config, cfg)
		if err != nil {
			continue
		}

		err = cfg.Validate()
		lk.FailOnErr("%v", err)

		cfg.procAPI()
		// cfg.procEnv()

		absconfig, _ := fd.AbsPath(config, false)
		cfg.using = absconfig
		return cfg
	}
	lk.FailOnErr("%v", fmt.Errorf("otf root config file is missing or error"))
	return nil
}

func (cfg *Config) SaveAsJson() (err error) {
	bytes, err := json.Marshal(cfg)
	lk.FailOnErr("%v", err)
	if err = cfg.Validate(); err == nil {
		io.MustWriteFile(cfg.using+".json", bytes)
	}
	return err
}

func (cfg *Config) SaveToml() (err error) {
	var buf bytes.Buffer
	err = toml.NewEncoder(&buf).Encode(*cfg)
	lk.FailOnErr("%v", err)
	if err = cfg.Validate(); err == nil {
		io.MustWriteFile(cfg.using, buf.Bytes())
	}
	return err
}

func (cfg *Config) Dispense() {
	cfg.NatsStreamings.Dispense()
	cfg.Nias3s.Dispense()
	cfg.Benthoses.Dispense()
	cfg.Readers.Dispense()
	cfg.Aligns.Dispense()
	cfg.TxtClassifiers.Dispense()
	cfg.Levels.Dispense()
	cfg.Weights.Dispense()
	cfg.Hubs.Dispense()
}
