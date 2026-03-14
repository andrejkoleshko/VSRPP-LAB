package cli

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
)

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type cliApp struct {
    l Logger
}

func New(l Logger) *cliApp {
    return &cliApp{
        l: l,
    }
}

func (c *cliApp) Run() error {
    c.l.Info("Инициализация процесса получения погодных данных")

    const lat = 53.6694
    const lon = 23.8131

    c.l.Debug(fmt.Sprintf("Координаты запроса: lat=%.4f, lon=%.4f", lat, lon))

    type Current struct {
        Temp float32 `json:"temperature_2m"`
    }

    type Response struct {
        Curr Current `json:"current"`
    }

    var response Response

    query := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat,
        lon,
    )

    url := "https://api.open-meteo.com/v1/forecast?" + query
    c.l.Debug("Формируем HTTP‑запрос: " + url)

    resp, err := http.Get(url)
    if err != nil {
        base := errors.New("ошибка при выполнении HTTP‑запроса")
        c.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    defer func() {
        if err := resp.Body.Close(); err != nil {
            c.l.Error("Ошибка закрытия тела ответа", err)
        }
    }()

    if resp.StatusCode != http.StatusOK {
        statusErr := fmt.Errorf("получен неожиданный статус: %d", resp.StatusCode)
        c.l.Error(statusErr.Error(), statusErr)
        return statusErr
    }

    raw, err := io.ReadAll(resp.Body)
    if err != nil {
        base := errors.New("не удалось прочитать тело ответа")
        c.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    c.l.Debug(fmt.Sprintf("Ответ успешно прочитан, размер: %d байт", len(raw)))

    if err := json.Unmarshal(raw, &response); err != nil {
        base := errors.New("ошибка разбора JSON‑данных")
        c.l.Error(base.Error(), err)
        return errors.Join(base, err)
    }

    c.l.Info("Погодные данные успешно обработаны")

    fmt.Printf("🌤 Температура сейчас: %.2f°C\n", response.Curr.Temp)

    return nil
}
