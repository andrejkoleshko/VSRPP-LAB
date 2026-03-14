package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab6/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab6/pkg/logger"
)

func main() {
    l := logger.New()
    app := cli.New(l)

    if err := app.Run(); err != nil {
        l.Error("Критическая ошибка выполнения", err)
        os.Exit(1)
    }

    l.Info("Программа завершена без ошибок")
}
