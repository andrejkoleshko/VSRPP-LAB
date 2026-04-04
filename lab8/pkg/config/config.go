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

type Location struct {
    Lat  float64 `yaml:"lat"`
    Long float64 `yaml:"long"`
}

type CacheConfig struct {
    Type string `yaml:"type"`
    Addr string `yaml:"addr"`
}

type Config struct {
    P     Provider    `yaml:"provider"`
    L     Location    `yaml:"location"`
    Cache CacheConfig `yaml:"cache"`
}

func Parse(r io.Reader) (Config, error) {
    var cf ConfigFile
    if err := yaml.NewDecoder(r).Decode(&cf); err != nil {
        return Config{}, err
    }
    return cf.C, nil
}
