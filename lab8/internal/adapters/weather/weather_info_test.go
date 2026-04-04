package weather

import (
    "testing"

    "github.com/andrejkoleshko/VSRPP-LAB/lab8/internal/domain/models"
)

type stubLogger struct{}

func (s *stubLogger) Info(string)              {}
func (s *stubLogger) Debug(string)             {}
func (s *stubLogger) Error(string, error)      {}

func TestWeatherInfo_GetTemperature_UsesLoadedValue(t *testing.T) {
    l := &stubLogger{}

    wi := &weatherInfo{
        c:        current{Temp: 4.2},
        l:        l,
        isLoaded: true,
    }

    res := wi.GetTemperature(0, 0)
    if res.Temp != 4.2 {
        t.Fatalf("expected 4.2, got %v", res.Temp)
    }
}

func TestWeatherInfo_GetTemperature_InitialLoad(t *testing.T) {
    l := &stubLogger{}

    wi := &weatherInfo{
        l:        l,
        isLoaded: false,
    }

    // Мы не тестируем HTTP, только то, что метод не падает
    _ = wi.getWeatherInfo(53.6694, 23.8131)

    res := wi.GetTemperature(53.6694, 23.8131)
    if (models.TempInfo{} == res) {
        t.Fatalf("expected non-zero TempInfo after load")
    }
}
