package cli

import (
    "fmt"

    "github.com/andrejkoleshko/VSRPP-LAB/lab9/internal/domain/models"
    "github.com/andrejkoleshko/VSRPP-LAB/lab9/pkg/config"
)

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type WeatherInfo interface {
    GetTemperature(float64, float64) models.TempInfo
}

type cliApp struct {
    l    Logger
    wi   WeatherInfo
    conf config.Config
}

func New(l Logger, wi WeatherInfo, c config.Config) *cliApp {
    return &cliApp{l: l, wi: wi, conf: c}
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск информера погоды")

    lat := c.conf.L.Lat
    long := c.conf.L.Long
    c.l.Debug(fmt.Sprintf("Координаты: %.4f, %.4f", lat, long))

    tempInfo := c.wi.GetTemperature(lat, long)
    fmt.Printf("🌤 Температура воздуха: %.2f°C\n", tempInfo.Temp)

    return nil
}
