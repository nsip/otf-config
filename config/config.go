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

func (cfg *Config) validate() (err error) {

	record := lk.WarnOnErr

	// natsstreaming
	for i := 0; i < len(cfg.NatsStreamings); i++ {
		if err = (&cfg.NatsStreamings[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// nias3
	for i := 0; i < len(cfg.Nias3s); i++ {
		if err = (&cfg.Nias3s[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// benthos
	for i := 0; i < len(cfg.Benthoses); i++ {
		if err = (&cfg.Benthoses[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// reader
	for i := 0; i < len(cfg.Readers); i++ {
		if err = (&cfg.Readers[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// align
	for i := 0; i < len(cfg.Aligns); i++ {
		if err = (&cfg.Aligns[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// text classifier
	for i := 0; i < len(cfg.TxtClassifiers); i++ {
		if err = (&cfg.TxtClassifiers[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// level
	for i := 0; i < len(cfg.Levels); i++ {
		if err = (&cfg.Levels[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// weight
	for i := 0; i < len(cfg.Weights); i++ {
		if err = (&cfg.Weights[i]).Validate(); err != nil {
			record("%v", err)
			return err
		}
	}

	// hub
	for i := 0; i < len(cfg.Hubs); i++ {
		if err = (&cfg.Hubs[i]).Validate(); err != nil {
			record("%v", err)
			return err
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

		err = cfg.validate()
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
	if err = cfg.validate(); err == nil {
		io.MustWriteFile(cfg.using+".json", bytes)
	}
	return err
}

func (cfg *Config) SaveToml() (err error) {
	var buf bytes.Buffer
	err = toml.NewEncoder(&buf).Encode(*cfg)
	lk.FailOnErr("%v", err)
	if err = cfg.validate(); err == nil {
		io.MustWriteFile(cfg.using, buf.Bytes())
	}
	return err
}

func (cfg *Config) Dispense() {

	// reader
	for _, reader := range cfg.Readers {
		reader.Dispense()
	}

	// align
	for _, align := range cfg.Aligns {
		align.Dispense()
	}

	// text classifier
	for _, textclassifier := range cfg.TxtClassifiers {
		textclassifier.Dispense()
	}

	// level
	for _, level := range cfg.Levels {
		level.Dispense()
	}

	// weight
	for _, weight := range cfg.Weights {
		weight.Dispense()
	}

	// hub
	for _, hub := range cfg.Hubs {
		hub.Dispense()
	}
}
