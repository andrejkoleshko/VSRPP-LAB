package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab5-informer-pogody/internal/pkg/app/cli"
)

func main() {
    logger := cli.NewLogger()
    app := cli.New(logger)

    if err := app.Run(); err != nil {
        logger.Error("Приложение завершилось с ошибкой: " + err.Error())
        os.Exit(1)
    }

    logger.Info("Работа завершена успешно")
}