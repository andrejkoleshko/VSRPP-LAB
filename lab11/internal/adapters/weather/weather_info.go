package weather

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/domain/models"
)

const baseURL = "https://api.open-meteo.com/v1/forecast"

type current struct {
    Temp float32 `json:"temperature_2m"`
}

type response struct {
    Curr current `json:"current"`
}

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type weatherInfo struct {
    c        current
    l        Logger
    isLoaded bool
}

func New(l Logger) *weatherInfo {
    return &weatherInfo{l: l}
}

func (wi *weatherInfo) getWeatherInfo(lat, long float64) {
    if wi.isLoaded {
        return
    }

    params := fmt.Sprintf("latitude=%f&longitude=%f&current=temperature_2m", lat, long)
    url := fmt.Sprintf("%s?%s", baseURL, params)
    wi.l.Debug("Формируем HTTP‑запрос: " + url)

    resp, err := http.Get(url)
    if err != nil {
        wi.l.Error("Ошибка при выполнении HTTP‑запроса", err)
        return
    }
    defer resp.Body.Close()

    raw, err := io.ReadAll(resp.Body)
    if err != nil {
        wi.l.Error("Ошибка чтения тела ответа", err)
        return
    }

    var respData response
    if err := json.Unmarshal(raw, &respData); err != nil {
        wi.l.Error("Ошибка разбора JSON", err)
        return
    }

    wi.c = respData.Curr
    wi.isLoaded = true
}

func (wi *weatherInfo) GetTemperature(lat, long float64) models.TempInfo {
    wi.getWeatherInfo(lat, long)
    return models.TempInfo{Temp: wi.c.Temp}
}
