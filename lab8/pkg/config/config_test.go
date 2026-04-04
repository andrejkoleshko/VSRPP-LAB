package config_test

import (
    "strings"
    "testing"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/pkg/config"
)

func TestParse_Config(t *testing.T) {
    raw := `
service:
  provider:
    type: open-meteo
  cache:
    type: redis
    addr: "localhost:6379"
  location:
    lat: 53.6694
    long: 23.8131
`
    r := strings.NewReader(raw)

    cfg, err := config.Parse(r)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if cfg.P.Type != "open-meteo" {
        t.Fatalf("expected provider open-meteo, got %s", cfg.P.Type)
    }

    if cfg.Cache.Type != "redis" {
        t.Fatalf("expected cache type redis, got %s", cfg.Cache.Type)
    }

    if cfg.Cache.Addr != "localhost:6379" {
        t.Fatalf("expected cache addr localhost:6379, got %s", cfg.Cache.Addr)
    }

    if cfg.L.Lat != 53.6694 || cfg.L.Long != 23.8131 {
        t.Fatalf("unexpected location: %+v", cfg.L)
    }
}
