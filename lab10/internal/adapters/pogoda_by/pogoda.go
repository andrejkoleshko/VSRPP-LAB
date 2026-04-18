package pogodaby

import (
    "encoding/json"
    "net/http"

    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/domain/models"
)

const url = "https://pogoda.by/api/v2/weather-fact?station=26820"

type resp struct {
    Temp float32 `json:"t"`
}

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type pogoda struct {
    l Logger
}

func New(l Logger) *pogoda {
    return &pogoda{l: l}
}

// Реализация по твоей логике: возвращаем только TempInfo, без error
func (p *pogoda) GetTemperature(lat, long float64) models.TempInfo {
    response, err := http.Get(url)
    if err != nil {
        p.l.Error("Ошибка получения данных от pogoda.by", err)
        return models.TempInfo{}
    }
    defer response.Body.Close()

    var r resp
    if err := json.NewDecoder(response.Body).Decode(&r); err != nil {
        p.l.Error("Ошибка декодирования JSON", err)
        return models.TempInfo{}
    }

    return models.TempInfo{Temp: r.Temp}
}
