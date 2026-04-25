package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/flags"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/providers"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/pkg/config"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/pkg/logger"
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
    wi := providers.GetProvider(cfg, l)

    app := cli.New(l, wi, cfg)

    if err := app.Run(); err != nil {
        l.Error("Критическая ошибка выполнения", err)
        os.Exit(1)
    }

    l.Info("Программа завершена без ошибок")
}
