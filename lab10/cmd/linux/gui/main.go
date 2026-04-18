package main

import (
    "os"

    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/pkg/app/gui"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/pkg/config"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/pkg/flags"
    fyneP "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/pkg/gui/fyne"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/pkg/providers"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/pkg/logger"
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
    provider := providers.GetProvider(cfg, l)

    p := fyneP.NewP()
    g := gui.New(l, p, provider, cfg)

    if err := g.Run(); err != nil {
        panic(err)
    }
}