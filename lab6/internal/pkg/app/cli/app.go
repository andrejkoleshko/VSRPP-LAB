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
    c.logger.Info("Инициализация процесса получения погодных данных")

    const lat = 53.6694
    const lon = 23.8131

    c.logger.Debug(fmt.Sprintf("Координаты запроса: lat=%.4f, lon=%.4f", lat, lon))

    type Current struct {
        Temp float32 `json:"temperature_2m"`
    }

    type Response struct {
        Current Current `json:"current"`
    }

    var respData Response

    query := fmt.Sprintf(
        "latitude=%f&longitude=%f&current=temperature_2m",
        lat,
        lon,
    )

    apiURL := "https://api.open-meteo.com/v1/forecast?" + query
    c.logger.Debug("Формируем HTTP‑запрос: " + apiURL)

    httpResp, err := http.Get(apiURL)
    if err != nil {
        base := errors.New("ошибка при выполнении HTTP‑запроса")
        c.logger.Error(base.Error())
        return errors.Join(base, err)
    }
    defer httpResp.Body.Close()

    if httpResp.StatusCode != http.StatusOK {
        errMsg := fmt.Errorf("получен неожиданный статус: %d", httpResp.StatusCode)
        c.logger.Error(errMsg.Error())
        return errMsg
    }

    raw, err := io.ReadAll(httpResp.Body)
    if err != nil {
        base := errors.New("не удалось прочитать тело ответа")
        c.logger.Error(base.Error())
        return errors.Join(base, err)
    }

    if err := json.Unmarshal(raw, &respData); err != nil {
        base := errors.New("ошибка разбора JSON‑данных")
        c.logger.Error(base.Error())
        return errors.Join(base, err)
    }

    c.logger.Info("Погодные данные успешно обработаны")

    fmt.Printf("🌤 Температура сейчас: %.2f°C\n", respData.Current.Temp)

    return nil
}
