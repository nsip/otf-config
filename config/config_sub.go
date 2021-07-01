package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/digisan/gotk/filedir"
	"github.com/digisan/gotk/io"
	"github.com/digisan/logkit"
)

var (
	record = logkit.WarnOnErr
	dir    = filepath.Dir
)

func objType(obj interface{}) string {
	return reflect.ValueOf(obj).Elem().Type().Name()
}

func isNotEmpty(field string) bool {
	return strings.Trim(field, " \t") != ""
}

func fileExists(field string) bool {
	return filedir.FileExists(field)
}

func dirExists(field string) bool {
	return filedir.DirExists(field)
}

///////////////////////////////////////////////

type NatsStreaming struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *NatsStreaming) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *NatsStreaming) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Nias3 struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Nias3) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Nias3) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Benthos struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Benthos) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Benthos) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Reader struct {
	Name          string `json:"name"`
	Path          string `json:"path"`
	SvrName       string `json:"svrname"`
	SvrID         string `json:"svrid"`
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

func (cfg *Reader) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Reader) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Align struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	SvrName   string `json:"svrname"`
	SvrID     string `json:"svrid"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	NiasHost  string `json:"niasHost"`
	NiasPort  int    `json:"niasPort"`
	NiasToken string `json:"niasToken"`
	TcHost    string `json:"tcHost"`
	TcPort    int    `json:"tcPort"`
}

func (cfg *Align) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Align) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type TxtClassifier struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *TxtClassifier) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *TxtClassifier) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Level struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	SvrName   string `json:"svrname"`
	SvrID     string `json:"svrid"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	NiasHost  string `json:"niasHost"`
	NiasPort  int    `json:"niasPort"`
	NiasToken string `json:"niasToken"`
}

func (cfg *Level) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Level) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Weight struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	FailWhenErr bool   `json:"failWhenErr"`
	Service     struct {
		SvrName string `json:"svrname"`
		SvrID   string `json:"svrid"`
		Port    int    `json:"port"`
		API     string `json:"api"`
	}
	Weighting struct {
		StudentIDPath string `json:"studentIDPath"`
		DomainPath    string `json:"domainPath"`
		TimePath      string `json:"timePath"`
		ScorePath     string `json:"scorePath"`
	}
}

func (cfg *Weight) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Weight) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}

////////////////////////////////

type Hub struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Hub) Validate() error {
	if isNotEmpty(cfg.Name) && fileExists(cfg.Path) {
		return nil
	}
	return fmt.Errorf("not valid %s config", objType(cfg))
}

func (cfg *Hub) Dispense() {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
}
