package cli

import (
    "fmt"
    "time"
)

type Logger interface {
    Info(msg string)
    Debug(msg string)
    Error(msg string)
}

type SimpleLogger struct {
    prefix string
}

func NewLogger() *SimpleLogger {
    return &SimpleLogger{
        prefix: "WeatherCLI",
    }
}

func (l *SimpleLogger) Info(msg string) {
    fmt.Printf("\033[34m[%s] [%s] INFO:\033[0m %s\n",
        time.Now().Format(time.RFC3339),
        l.prefix,
        msg,
    )
}

func (l *SimpleLogger) Debug(msg string) {
    fmt.Printf("\033[36m[%s] [%s] DEBUG:\033[0m %s\n",
        time.Now().Format(time.RFC3339),
        l.prefix,
        msg,
    )
}

func (l *SimpleLogger) Error(msg string) {
    fmt.Printf("\033[31m[%s] [%s] ERROR:\033[0m %s\n",
        time.Now().Format(time.RFC3339),
        l.prefix,
        msg,
    )
}