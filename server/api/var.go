package api

import (
	"reflect"
	"strings"

	lk "github.com/digisan/logkit"
	"github.com/nsip/otf-config/config"
)

var (
	log4get  = lk.Fac4GrpIdxLogF("GET", 0, lk.INFO, true)
	log4post = lk.Fac4GrpIdxLogF("POST", 0, lk.INFO, true)

	cfg          = config.GetConfig("../config.toml", "./config.toml")
	cfgNameQuery = "cfgName"
)

func WantedFieldsByType(a interface{}, prefix string) (fields []string) {
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := v.Type().Field(i).Name
		t := f.Type().String()
		if strings.HasPrefix(t, prefix) {
			fields = append(fields, n)
		}
	}
	return
}

func WantedFieldsByKind(a interface{}, kind string) (fields []string) {
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		n := v.Type().Field(i).Name
		k := f.Type().Kind().String()
		if k == kind {
			fields = append(fields, n)
		}
	}
	return
}
