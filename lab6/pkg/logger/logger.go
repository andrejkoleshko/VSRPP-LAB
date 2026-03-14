package logger

import (
    "fmt"
    "time"
)

var (
    INFO  = "INFO"
    DEBUG = "DEBUG"
    ERROR = "ERROR"
)

type Logger struct {
    tag string
}

func New() *Logger {
    return &Logger{
        tag: "WeatherCLI",
    }
}

func (l *Logger) Info(msg string) {
    fmt.Printf(
        "\033[34m[%s] (%s) %s →\033[0m %s\n",
        time.Now().Format(time.RFC3339),
        l.tag,
        INFO,
        msg,
    )
}

func (l *Logger) Debug(msg string) {
    fmt.Printf(
        "\033[36m[%s] (%s) %s →\033[0m %s\n",
        time.Now().Format(time.RFC3339),
        l.tag,
        DEBUG,
        msg,
    )
}

func (l *Logger) Error(msg string, err error) {
    fmt.Printf(
        "\033[31m[%s] (%s) %s →\033[0m %s | %s\n",
        time.Now().Format(time.RFC3339),
        l.tag,
        ERROR,
        msg,
        err.Error(),
    )
}
