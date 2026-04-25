package gui

import (
    guisettings "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/domain/gui_settings"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/domain/models"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/pkg/config"
)

type Logger interface {
    Info(string)
    Debug(string)
    Error(string, error)
}

type WeatherInfo interface {
    GetTemperature(float64, float64) models.TempInfo
}

type Provider interface {
    CreateWindow(name string, size guisettings.WindowSize) (guisettings.Window, error)
    GetAppRunner() guisettings.AppRunner
    GetTextWidget(text string) guisettings.TextWidget
}

type guiApp struct {
    l  Logger
    p  Provider
    wi WeatherInfo
    c  config.Config
}

func New(l Logger, p Provider, wi WeatherInfo, c config.Config) *guiApp {
    return &guiApp{l: l, p: p, wi: wi, c: c}
}

func (g *guiApp) Run() error {
    g.l.Info("Запуск GUI информера погоды")

    win, err := g.p.CreateWindow("Weather GUI", guisettings.NewWS(400, 200))
    if err != nil {
        return err
    }

    tw := g.p.GetTextWidget("Загрузка...")
    win.SetTemperatureWidget(tw)

    temp := g.wi.GetTemperature(g.c.L.Lat, g.c.L.Long)
    win.UpdateTemperature(temp.Temp)

    win.Render()

    runner := g.p.GetAppRunner()
    runner.Run()

    return nil
}
