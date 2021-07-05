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
	record = logkit.FailOnErr
	dir    = filepath.Dir
)

func MarshalRemove(v interface{}, fields ...string) (bytes []byte, err error) {
	if bytes, err = json.Marshal(v); err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	json.Unmarshal(bytes, &m)
	for _, f := range fields {
		delete(m, f)
	}
	if bytes, err = json.Marshal(m); err != nil {
		return nil, err
	}
	return bytes, nil
}

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
	AllNames() []string
	Get(name string) IEle
	Add(newElem IEle)
	Update(name string, newElem IEle)
	Delete(name string)
	Validate() error
	Dispense() error
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type NatsStreamingGrp []NatsStreaming

func (grp *NatsStreamingGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *NatsStreamingGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *NatsStreamingGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*NatsStreaming)))
}

func (grp *NatsStreamingGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*NatsStreaming))
			return
		}
	}
}

func (grp *NatsStreamingGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *NatsStreamingGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *NatsStreamingGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type Nias3Grp []Nias3

func (grp *Nias3Grp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *Nias3Grp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *Nias3Grp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Nias3)))
}

func (grp *Nias3Grp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Nias3))
			return
		}
	}
}

func (grp *Nias3Grp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *Nias3Grp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *Nias3Grp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type BenthosGrp []Benthos

func (grp *BenthosGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *BenthosGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *BenthosGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Benthos)))
}

func (grp *BenthosGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Benthos))
			return
		}
	}
}

func (grp *BenthosGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *BenthosGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *BenthosGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type ReaderGrp []Reader

func (grp *ReaderGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *ReaderGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *ReaderGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Reader)))
}

func (grp *ReaderGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Reader))
			return
		}
	}
}

func (grp *ReaderGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *ReaderGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *ReaderGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type AlignGrp []Align

func (grp *AlignGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *AlignGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *AlignGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Align)))
}

func (grp *AlignGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Align))
			return
		}
	}
}

func (grp *AlignGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *AlignGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *AlignGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
}

////////////////////////////////

type TxtClassifier struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Port int    `json:"port"`
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type TxtClassifierGrp []TxtClassifier

func (grp *TxtClassifierGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *TxtClassifierGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *TxtClassifierGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*TxtClassifier)))
}

func (grp *TxtClassifierGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*TxtClassifier))
			return
		}
	}
}

func (grp *TxtClassifierGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *TxtClassifierGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *TxtClassifierGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type LevelGrp []Level

func (grp *LevelGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *LevelGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *LevelGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Level)))
}

func (grp *LevelGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Level))
			return
		}
	}
}

func (grp *LevelGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *LevelGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *LevelGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type WeightGrp []Weight

func (grp *WeightGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *WeightGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *WeightGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Weight)))
}

func (grp *WeightGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Weight))
			return
		}
	}
}

func (grp *WeightGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *WeightGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *WeightGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
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
	cf, err := MarshalRemove(cfg, "name", "path")
	record("%v", err)
	io.MustWriteFile(dir(cfg.Path)+"/"+cfg.Name+".json", cf)
	return err
}

type HubGrp []Hub

func (grp *HubGrp) AllNames() (names []string) {
	for _, elem := range *grp {
		names = append(names, elem.Name)
	}
	return
}

func (grp *HubGrp) Get(name string) IEle {
	for _, elem := range *grp {
		if elem.Name == name {
			return &elem
		}
	}
	return nil
}

func (grp *HubGrp) Add(newElem IEle) {
	*grp = append(*grp, *(newElem.(*Hub)))
}

func (grp *HubGrp) Update(name string, newElem IEle) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = *(newElem.(*Hub))
			return
		}
	}
}

func (grp *HubGrp) Delete(name string) {
	for i, elem := range *grp {
		if elem.Name == name {
			(*grp)[i] = (*grp)[len(*grp)-1]
			break
		}
	}
	*grp = (*grp)[:len(*grp)-1]
}

func (grp *HubGrp) Validate() error {
	for _, g := range *grp {
		if err := g.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (grp *HubGrp) Dispense() error {
	for _, g := range *grp {
		if err := g.Dispense(); err != nil {
			return err
		}
	}
	return nil
}
