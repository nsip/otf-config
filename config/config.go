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

func (cfg *Config) getSubMap() map[string]IGrp {
	return map[string]IGrp{
		"NatsStreaming": &cfg.NatsStreamings,
		"Nias3":         &cfg.Nias3s,
		"Benthos":       &cfg.Benthoses,
		"Reader":        &cfg.Readers,
		"Align":         &cfg.Aligns,
		"TxtClassifier": &cfg.TxtClassifiers,
		"Level":         &cfg.Levels,
		"Weight":        &cfg.Weights,
		"Hub":           &cfg.Hubs,
	}
}

func (cfg *Config) Validate() (err error) {
	m := cfg.getSubMap()
	for _, sub := range m {
		if err = sub.Validate(); err != nil {
			return
		}
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

func (cfg *Config) Dispense(proj string) {
	m := cfg.getSubMap()
	m[proj].Dispense()
}

func (cfg *Config) Withdraw(proj string) {
	m := cfg.getSubMap()
	m[proj].Withdraw()
}
