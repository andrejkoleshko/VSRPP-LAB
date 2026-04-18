package providers

import (
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/adapters/pogoda_by"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/adapters/weather"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab10/pkg/config"
)

func GetProvider(c config.Config, l cli.Logger) cli.WeatherInfo {
    switch c.P.Type {
    case "open-meteo":
        return weather.New(l)
    case "pogoda":
        return pogodaby.New(l)
    default:
        return weather.New(l)
    }
}
