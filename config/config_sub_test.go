package config

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	ns := new(NatsStreaming)
	fmt.Println(ns.Validate())
}
