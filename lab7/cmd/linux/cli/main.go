package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab7/internal/adapters/weather"
    "github.com/andrejkoleshko/VSRPP-LAB/lab7/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab7/pkg/logger"
)

func main() {
    l := logger.New()
    wi := weather.New(l)
    app := cli.New(l, wi)

    if err := app.Run(); err != nil {
        l.Error("Критическая ошибка выполнения", err)
        os.Exit(1)
    }

    l.Info("Программа завершена без ошибок")
}
