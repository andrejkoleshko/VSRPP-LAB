package cli_test

import (
    "bytes"
    "testing"

    "github.com/golang/mock/gomock"

    mocks "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/mocks"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/pkg/app/cli"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/pkg/config"
    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/domain/models"
)

func TestCliApp_Run_FromCache(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLogger := mocks.NewMockLogger(ctrl)
    mockWeather := mocks.NewMockWeatherInfo(ctrl)
    mockCache := mocks.NewMockCache(ctrl)

    cfg := config.Config{
        L: config.Location{
            Lat:  53.6694,
            Long: 23.8131,
        },
    }

    key := "53.6694:23.8131"

    mockLogger.EXPECT().Info("Запуск информера погоды")
    mockCache.EXPECT().Get(key).Return(float32(6.9), true)
    mockLogger.EXPECT().Info("Температура взята из кэша")

    var buf bytes.Buffer

    app := cli.New(mockLogger, mockWeather, mockCache, cfg)
    app.Out = &buf

    if err := app.Run(); err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    expected := "🌤 Температура сейчас: 6.90°C\n"
    if buf.String() != expected {
        t.Fatalf("expected %q, got %q", expected, buf.String())
    }
}

func TestCliApp_Run_FromProvider(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLogger := mocks.NewMockLogger(ctrl)
    mockWeather := mocks.NewMockWeatherInfo(ctrl)
    mockCache := mocks.NewMockCache(ctrl)

    cfg := config.Config{
        L: config.Location{
            Lat:  53.6694,
            Long: 23.8131,
        },
    }

    key := "53.6694:23.8131"

    mockLogger.EXPECT().Info("Запуск информера погоды")
    mockCache.EXPECT().Get(key).Return(float32(0), false)
    mockWeather.EXPECT().GetTemperature(cfg.L.Lat, cfg.L.Long).Return(models.TempInfo{Temp: 3.5})
    mockCache.EXPECT().Set(key, float32(3.5))

    var buf bytes.Buffer

    app := cli.New(mockLogger, mockWeather, mockCache, cfg)
    app.Out = &buf

    if err := app.Run(); err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    expected := "🌤 Температура сейчас: 3.50°C\n"
    if buf.String() != expected {
        t.Fatalf("expected %q, got %q", expected, buf.String())
    }
}
