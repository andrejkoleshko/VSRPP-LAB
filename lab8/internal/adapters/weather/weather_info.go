package weather

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/domain/models"
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
    return &weatherInfo{
        l: l,
    }
}

func (wi *weatherInfo) getWeatherInfo(lat, long float64) error {
    var respData response

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat,
        long,
    )

    url := fmt.Sprintf("%s?%s", baseURL, params)
    wi.l.Debug("Формируем HTTP‑запрос: " + url)

    resp, err := http.Get(url)
    if err != nil {
        base := errors.New("ошибка при выполнении HTTP‑запроса")
        wi.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    defer func() {
        if err := resp.Body.Close(); err != nil {
            wi.l.Error("Ошибка закрытия тела ответа", err)
        }
    }()

    raw, err := io.ReadAll(resp.Body)
    if err != nil {
        base := errors.New("не удалось прочитать тело ответа")
        wi.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    wi.l.Debug(fmt.Sprintf("Ответ успешно прочитан, размер: %d байт", len(raw)))

    if err := json.Unmarshal(raw, &respData); err != nil {
        base := errors.New("ошибка разбора JSON‑данных")
        wi.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    wi.c = respData.Curr
    wi.isLoaded = true

    return nil
}

func (wi *weatherInfo) GetTemperature(lat, long float64) models.TempInfo {
    if !wi.isLoaded {
        wi.getWeatherInfo(lat, long)
    }

    return models.TempInfo{
        Temp: wi.c.Temp,
    }
}
