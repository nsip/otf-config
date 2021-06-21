package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	fd "github.com/digisan/gotk/filedir"
	"github.com/digisan/gotk/io"
	lk "github.com/digisan/logkit"
)

func init() {
	lk.Log2F(true, "./otf-config.log")
}

// Config :
type Config struct {
	Name string

	NatsStreaming struct {
		CfgName string
		Path    string
	}

	Nias3 struct {
		CfgName string
		Path    string
	}

	Benthos struct {
		CfgName string
		Path    string
	}

	Reader struct {
		CfgName       string
		Path          string
		Name          string `json:"name"`
		ID            string `json:"id"`
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

	Align struct {
		CfgName   string
		Path      string
		Name      string `json:"name"`
		ID        string `json:"id"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		NiasHost  string `json:"niasHost"`
		NiasPort  int    `json:"niasPort"`
		NiasToken string `json:"niasToken"`
		TcHost    string `json:"tcHost"`
		TcPort    int    `json:"tcPort"`
	}

	TxtClassifier struct {
		CfgName string
		Path    string
	}

	Level struct {
		CfgName   string
		Path      string
		Name      string `json:"name"`
		ID        string `json:"id"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		NiasHost  string `json:"niasHost"`
		NiasPort  int    `json:"niasPort"`
		NiasToken string `json:"niasToken"`
	}

	Weight struct {
		CfgName     string
		Path        string
		FailWhenErr bool `json:"failWhenErr"`
		Service     struct {
			Name string `json:"name"`
			ID   string `json:"id"`
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

	Hub struct {
		CfgName string
		Path    string
	}
}

func (cfg *Config) dirProc(check bool) {

	var err error

	// natsstreaming
	cfg.NatsStreaming.Path, err = fd.AbsPath(cfg.NatsStreaming.Path, check)
	lk.FailOnErr("%v", err)

	// nias3
	cfg.Nias3.Path, err = fd.AbsPath(cfg.Nias3.Path, check)
	lk.FailOnErr("%v", err)

	// benthos
	cfg.Benthos.Path, err = fd.AbsPath(cfg.Benthos.Path, check)
	lk.FailOnErr("%v", err)

	// reader
	cfg.Reader.Path, err = fd.AbsPath(cfg.Reader.Path, check)
	lk.FailOnErr("%v", err)

	// align
	cfg.Align.Path, err = fd.AbsPath(cfg.Align.Path, check)
	lk.FailOnErr("%v", err)

	// text classifier
	cfg.TxtClassifier.Path, err = fd.AbsPath(cfg.TxtClassifier.Path, check)
	lk.FailOnErr("%v", err)

	// level
	cfg.Level.Path, err = fd.AbsPath(cfg.Level.Path, check)
	lk.FailOnErr("%v", err)

	// weight
	cfg.Weight.Path, err = fd.AbsPath(cfg.Weight.Path, check)
	lk.FailOnErr("%v", err)

	// hub
	cfg.Hub.Path, err = fd.AbsPath(cfg.Hub.Path, check)
	lk.FailOnErr("%v", err)
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
	cfg.Weight.Service.API = withSlash(cfg.Weight.Service.API)

	// hub
}

func (cfg *Config) envProc() {

	// natsstreaming

	// nias3

	// benthos

	// reader

	// align

	// text classifier

	// level

	// weight
	os.Setenv("OTF_WEIGHT_FAILWHENERR", strconv.FormatBool(cfg.Weight.FailWhenErr))

	// hub
}

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
		cfg.envProc()
		return cfg
	}
	lk.FailOnErr("%v", fmt.Errorf("otf root config file is missing or error"))
	return nil
}

func (cfg *Config) Dispense() {

	// reader
	cr, err := json.Marshal(cfg.Reader)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.Reader.Path)+"/"+cfg.Reader.CfgName+".json", cr)

	// align
	ca, err := json.Marshal(cfg.Align)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.Align.Path)+"/"+cfg.Align.CfgName+".json", ca)

	// text classifier
	ct, err := json.Marshal(cfg.TxtClassifier)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.TxtClassifier.Path)+"/"+cfg.TxtClassifier.CfgName+".json", ct)

	// level
	cl, err := json.Marshal(cfg.Level)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.Level.Path)+"/"+cfg.Level.CfgName+".json", cl)

	// weight
	cw, err := json.Marshal(cfg.Weight)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.Weight.Path)+"/"+cfg.Weight.CfgName+".json", cw)

	// hub
	ch, err := json.Marshal(cfg.Hub)
	lk.FailOnErr("%v", err)
	io.MustWriteFile(filepath.Dir(cfg.Hub.Path)+"/"+cfg.Hub.CfgName+".json", ch)
}
