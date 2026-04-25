package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/app/gui"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/pkg/config"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/flags"
    fyneP "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/gui/fyne"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/internal/pkg/providers"
    "github.com/andrejkoleshko/VSRPP-LAB/lab11/pkg/logger"
)

func main() {
    args := flags.Parse()

    // Если путь не указан — используем системный
    configPath := args.Path
    if configPath == "" {
        configPath = "/usr/local/share/weather-gui/config.yaml"
    }

    r, err := os.Open(configPath)
    if err != nil {
        panic("cannot open config: " + err.Error())
    }

    cfg, err := config.Parse(r)
    if err != nil {
        panic("cannot parse config: " + err.Error())
    }

    l := logger.New()
    provider := providers.GetProvider(cfg, l)

    p := fyneP.NewP()
    g := gui.New(l, p, provider, cfg)

    if err := g.Run(); err != nil {
        panic(err)
    }
}