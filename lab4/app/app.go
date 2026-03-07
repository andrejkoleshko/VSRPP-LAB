package app

import (
    "net/http"
)

type Rows interface {
    Next() bool
    Scan(dest ...any) error
    Close() error
}

type DB interface {
    Exec(query string, args ...any) (interface{}, error)
    Query(query string, args ...any) (Rows, error)
}

type Logger interface {
    Info(string)
    Error(string, error)
}

type App struct {
    db DB
    l  Logger
}

func New(db DB, l Logger) *App {
    return &App{db: db, l: l}
}

func (a *App) Run() error {
    _, err := a.db.Exec(`CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        model TEXT,
        company TEXT,
        price INTEGER
    )`)
    if err != nil {
        a.l.Error("create table failed", err)
        return err
    }

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates/static"))))

    http.HandleFunc("/", a.handleIndex)
    http.HandleFunc("/create", a.handleCreate)
    http.HandleFunc("/delete", a.handleDelete)

    a.l.Info("Server started at http://localhost:8080")
    return http.ListenAndServe(":8080", nil)
}

