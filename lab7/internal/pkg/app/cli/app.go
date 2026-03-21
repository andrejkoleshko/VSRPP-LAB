package cli

import (
    "fmt"

    "github.com/andrejkoleshko/VSRPP-LAB/lab7/internal/domain/models"
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
}

func New(l Logger, wi WeatherInfo) *cliApp {
    return &cliApp{
        l:  l,
        wi: wi,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск информера погоды")

    temp := c.wi.GetTemperature(53.6694, 23.8131).Temp

    fmt.Printf("🌤 Температура сейчас: %.2f°C\n", temp)

    return nil
}
