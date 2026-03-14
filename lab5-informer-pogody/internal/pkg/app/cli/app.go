package cli

import (
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
)

type cliApp struct {
    logger Logger
}

func New(logger Logger) *cliApp {
    return &cliApp{logger: logger}
}

func (c *cliApp) Run() error {
    c.logger.Info("Запуск информера погоды")

    const (
    latitude  = 53.6694
    longitude = 23.8131
    )


    c.logger.Debug(fmt.Sprintf("Используем координаты: %.4f, %.4f", latitude, longitude))

    type Current struct {
        Temp float32 `json:"temperature_2m"`
    }

    type Response struct {
        Current Current `json:"current"`
    }

    var response Response

    params := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        latitude,
        longitude,
    )

    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

    c.logger.Debug("Отправляем запрос: " + url)

    resp, err := http.Get(url)
    if err != nil {
        customErr := errors.New("не удалось выполнить запрос к API погоды")
        c.logger.Error(customErr.Error())
        return errors.Join(customErr, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        customErr := fmt.Errorf("сервер вернул статус %d", resp.StatusCode)
        c.logger.Error(customErr.Error())
        return customErr
    }

    data, err := io.ReadAll(resp.Body)
    if err != nil {
        customErr := errors.New("не удалось прочитать ответ сервера")
        c.logger.Error(customErr.Error())
        return errors.Join(customErr, err)
    }

    if err := json.Unmarshal(data, &response); err != nil {
        customErr := errors.New("не удалось разобрать JSON‑ответ")
        c.logger.Error(customErr.Error())
        return errors.Join(customErr, err)
    }

    c.logger.Info("Данные успешно получены")

    fmt.Printf(
        "🌤 Текущая температура в выбранной точке: %.2f°C\n",
        response.Current.Temp,
    )

    return nil
}