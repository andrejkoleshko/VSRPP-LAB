package cli

import (
    "fmt"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/domain/models"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/pkg/config"
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
    l  Logger
    wi WeatherInfo
    c  config.Config
}

func New(l Logger, wi WeatherInfo, c config.Config) *cliApp {
    return &cliApp{
        l:  l,
        wi: wi,
        c:  c,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск информера погоды")

    lat := c.c.L.Lat
    long := c.c.L.Long

    c.l.Debug(fmt.Sprintf("Координаты из конфига: %.4f, %.4f", lat, long))

    temp := c.wi.GetTemperature(lat, long).Temp

    fmt.Printf("🌤 Температура сейчас: %.2f°C\n", temp)

    return nil
}
