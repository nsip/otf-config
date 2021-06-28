package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	fd "github.com/digisan/gotk/filedir"
	"github.com/digisan/gotk/io"
	lk "github.com/digisan/logkit"
)

var (
	OtfCfg = GetConfig("../config.toml", "./config.toml")
)

func init() {
	lk.Log2F(true, "../otf-config.log")
}

// Config :
type Config struct {
	Name       string
	Port       int
	API        string
	PageFolder string

	NatsStreamings []struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	}

	Nias3s []struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	}

	Benthoses []struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	}

	Readers []struct {
		CfgName       string `json:"name"`
		Path          string `json:"path"`
		Name          string `json:"svrname"`
		ID            string `json:"svrid"`
		Provider      string `json:"provider"`
		InputFmt      string `json:"inputFormat"`
		AlignMethod   string `json:"alignMethod"`
		LevelMethod   string `json:"levelMethod"`
		GenCapability string `json:"capability"`
		NatsHost      string `json:"natsHost"`
		NatsPort      int    `json:"natsPort"`
		NatsCluster   string `json:"natsCluster"`
		Topic         string `json:"topic"`
		Folder        string `json:"folder"`
		FileSuffix    string `json:"suffix"`
		Interval      string `json:"interval"`
		Recursive     bool   `json:"recursive"`
		DotFiles      bool   `json:"dotfiles"`
		Ignore        string `json:"ignore"`
		ConcurrFiles  int    `json:"concurrFiles"`
	}

	Aligns []struct {
		CfgName   string `json:"name"`
		Path      string `json:"path"`
		Name      string `json:"svrname"`
		ID        string `json:"svrid"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		NiasHost  string `json:"niasHost"`
		NiasPort  int    `json:"niasPort"`
		NiasToken string `json:"niasToken"`
		TcHost    string `json:"tcHost"`
		TcPort    int    `json:"tcPort"`
	}

	TxtClassifiers []struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	}

	Levels []struct {
		CfgName   string `json:"name"`
		Path      string `json:"path"`
		Name      string `json:"svrname"`
		ID        string `json:"svrid"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		NiasHost  string `json:"niasHost"`
		NiasPort  int    `json:"niasPort"`
		NiasToken string `json:"niasToken"`
	}

	Weights []struct {
		CfgName     string `json:"name"`
		Path        string `json:"path"`
		FailWhenErr bool   `json:"failWhenErr"`
		Service     struct {
			Name string `json:"svrname"`
			ID   string `json:"svrid"`
			Port int    `json:"port"`
			API  string `json:"api"`
		}
		Weighting struct {
			StudentIDPath string `json:"studentIDPath"`
			DomainPath    string `json:"domainPath"`
			TimePath      string `json:"timePath"`
			ScorePath     string `json:"scorePath"`
		}
	}

	Hubs []struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	}
}

func (cfg *Config) dirProc(check bool) {

	var err error

	// natsstreaming
	for i := 0; i < len(cfg.NatsStreamings); i++ {
		sub := &cfg.NatsStreamings[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// nias3
	for i := 0; i < len(cfg.Nias3s); i++ {
		sub := &cfg.Nias3s[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// benthos
	for i := 0; i < len(cfg.Benthoses); i++ {
		sub := &cfg.Benthoses[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// reader
	for i := 0; i < len(cfg.Readers); i++ {
		sub := &cfg.Readers[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// align
	for i := 0; i < len(cfg.Aligns); i++ {
		sub := &cfg.Aligns[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// text classifier
	for i := 0; i < len(cfg.TxtClassifiers); i++ {
		sub := &cfg.TxtClassifiers[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// level
	for i := 0; i < len(cfg.Levels); i++ {
		sub := &cfg.Levels[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// weight
	for i := 0; i < len(cfg.Weights); i++ {
		sub := &cfg.Weights[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}

	// hub
	for i := 0; i < len(cfg.Hubs); i++ {
		sub := &cfg.Hubs[i]
		sub.Path, err = fd.AbsPath(sub.Path, check)
		lk.FailOnErr("%v", err)
	}
}

func (cfg *Config) apiPathProc() {

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

// func (cfg *Config) envProc() {

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
		cfg.dirProc(true)
		cfg.apiPathProc()
		// cfg.envProc()
		return cfg
	}
	lk.FailOnErr("%v", fmt.Errorf("otf root config file is missing or error"))
	return nil
}

func (cfg *Config) Dispense() {

	dir := filepath.Dir

	// reader
	for _, reader := range cfg.Readers {
		cr, err := json.Marshal(reader)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(reader.Path)+"/"+reader.CfgName+".json", cr)
	}

	// align
	for _, align := range cfg.Aligns {
		ca, err := json.Marshal(align)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(align.Path)+"/"+align.CfgName+".json", ca)
	}

	// text classifier
	for _, textclassifier := range cfg.TxtClassifiers {
		ct, err := json.Marshal(textclassifier)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(textclassifier.Path)+"/"+textclassifier.CfgName+".json", ct)
	}

	// level
	for _, level := range cfg.Levels {
		cl, err := json.Marshal(level)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(level.Path)+"/"+level.CfgName+".json", cl)
	}

	// weight
	for _, weight := range cfg.Weights {
		cw, err := json.Marshal(weight)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(weight.Path)+"/"+weight.CfgName+".json", cw)
	}

	// hub
	for _, hub := range cfg.Hubs {
		ch, err := json.Marshal(hub)
		lk.FailOnErr("%v", err)
		io.MustWriteFile(dir(hub.Path)+"/"+hub.CfgName+".json", ch)
	}
}
