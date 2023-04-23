package config

import (
	"testing"
)

func TestReadCfg(t *testing.T) {
	cfg, err := LoadConfig("app.toml")
	if err != nil {
		t.Logf(err.Error())
		return
	}
	t.Logf("%+v", *cfg)
}
