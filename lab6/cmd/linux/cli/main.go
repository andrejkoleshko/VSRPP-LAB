package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab6/internal/pkg/app/cli"
)

func main() {
    log := cli.NewLogger()
    app := cli.New(log)

    if err := app.Run(); err != nil {
        log.Error("Критическая ошибка выполнения: " + err.Error())
        os.Exit(1)
    }

    log.Info("Программа завершена без ошибок")
}
