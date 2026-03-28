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

type Cache interface {
    Get(key string) (float32, bool)
    Set(key string, value float32) error
}

type cliApp struct {
    l     Logger
    wi    WeatherInfo
    cache Cache
    c     config.Config
}

func New(l Logger, wi WeatherInfo, cache Cache, c config.Config) *cliApp {
    return &cliApp{
        l:     l,
        wi:    wi,
        cache: cache,
        c:     c,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Запуск информера погоды")

    lat := c.c.L.Lat
    long := c.c.L.Long

    key := fmt.Sprintf("%.4f:%.4f", lat, long)

    if temp, ok := c.cache.Get(key); ok {
        c.l.Info("Температура взята из кэша")
        fmt.Printf("🌤 Температура сейчас: %.2f°C\n", temp)
        return nil
    }

    temp := c.wi.GetTemperature(lat, long).Temp
    c.cache.Set(key, temp)

    fmt.Printf("🌤 Температура сейчас: %.2f°C\n", temp)
    return nil
}
