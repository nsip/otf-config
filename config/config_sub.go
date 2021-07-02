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

func isEmpty(field string) bool {
	return strings.Trim(field, " \t") == ""
}

func fileMissing(field string) bool {
	return !filedir.FileExists(field)
}

func dirMissing(field string) bool {
	return !filedir.DirExists(field)
}

// for element
type IEle interface {
	GetName() string
	Validate() error
	Dispense() error
}

// for Group
type IGrp interface {
	GetAllNames() []string
	GetElem(name string) IEle
	AddElem(elem IEle)
}

///////////////////////////////////////////////

type NatsStreaming struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *NatsStreaming) GetName() string {
	return cfg.Name
}

func (cfg *NatsStreaming) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *NatsStreaming) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type NatsStreamingGrp []NatsStreaming

func (grp *NatsStreamingGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *NatsStreamingGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *NatsStreamingGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*NatsStreaming)))
}

////////////////////////////////

type Nias3 struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Nias3) GetName() string {
	return cfg.Name
}

func (cfg *Nias3) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Nias3) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type Nias3Grp []Nias3

func (grp *Nias3Grp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *Nias3Grp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *Nias3Grp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Nias3)))
}

////////////////////////////////

type Benthos struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Benthos) GetName() string {
	return cfg.Name
}

func (cfg *Benthos) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Benthos) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type BenthosGrp []Benthos

func (grp *BenthosGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *BenthosGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *BenthosGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Benthos)))
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

func (cfg *Reader) GetName() string {
	return cfg.Name
}

func (cfg *Reader) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Reader) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type ReaderGrp []Reader

func (grp *ReaderGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *ReaderGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *ReaderGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Reader)))
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

func (cfg *Align) GetName() string {
	return cfg.Name
}

func (cfg *Align) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Align) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type AlignGrp []Align

func (grp *AlignGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *AlignGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *AlignGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Align)))
}

////////////////////////////////

type TxtClassifier struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *TxtClassifier) GetName() string {
	return cfg.Name
}

func (cfg *TxtClassifier) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *TxtClassifier) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type TxtClassifierGrp []TxtClassifier

func (grp *TxtClassifierGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *TxtClassifierGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *TxtClassifierGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*TxtClassifier)))
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

func (cfg *Level) GetName() string {
	return cfg.Name
}

func (cfg *Level) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Level) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type LevelGrp []Level

func (grp *LevelGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *LevelGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *LevelGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Level)))
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

func (cfg *Weight) GetName() string {
	return cfg.Name
}

func (cfg *Weight) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Weight) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type WeightGrp []Weight

func (grp *WeightGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *WeightGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *WeightGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Weight)))
}

////////////////////////////////

type Hub struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (cfg *Hub) GetName() string {
	return cfg.Name
}

func (cfg *Hub) Validate() error {
	switch {
	case isEmpty(cfg.Name):
		return fmt.Errorf("not valid %s config, empty name", objType(cfg))
	case fileMissing(cfg.Path):
		return fmt.Errorf("not valid %s config, executable cannot be found via path", objType(cfg))
	default:
		return nil
	}
}

func (cfg *Hub) Dispense() error {
	cf, err := json.Marshal(cfg)
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type HubGrp []Hub

func (grp *HubGrp) GetAllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *HubGrp) GetElem(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *HubGrp) AddElem(elem IEle) {
	*grp = append(*grp, *(elem.(*Hub)))
}
