package config

import (
    "io"

    "gopkg.in/yaml.v3"
)

type ConfigFile struct {
    C Config `yaml:"service"`
}

type Provider struct {
    Type string `yaml:"type"`
}

type Cache struct {
    Type string `yaml:"type"`
    Addr string `yaml:"addr"`
}

type Location struct {
    Lat  float64 `yaml:"lat"`
    Long float64 `yaml:"long"`
}

type Config struct {
    P     Provider `yaml:"provider"`
    Cache Cache    `yaml:"cache"`
    L     Location `yaml:"location"`
}

func Parse(r io.Reader) (Config, error) {
    var c ConfigFile
    if err := yaml.NewDecoder(r).Decode(&c); err != nil {
        return Config{}, err
    }
    return c.C, nil
}
