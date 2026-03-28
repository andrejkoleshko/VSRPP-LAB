package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/adapters/weather"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/pkg/flags"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/pkg/config"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/pkg/logger"
)

func main() {
    args := flags.Parse()

    r, err := os.Open(args.Path)
    if err != nil {
        panic(err)
    }

    cfg, err := config.Parse(r)
    if err != nil {
        panic(err)
    }

    l := logger.New()
    wi := getProvider(cfg, l)

    app := cli.New(l, wi, cfg)

    if err := app.Run(); err != nil {
        l.Error("Критическая ошибка выполнения", err)
        os.Exit(1)
    }

    l.Info("Программа завершена без ошибок")
}

func getProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
    switch c.P.Type {
    case "open-meteo":
        return weather.New(l)
    default:
        return weather.New(l)
    }
}
