package main

import (
    "database/sql"
    "lab3/app"
    _ "modernc.org/sqlite"
)

type RealLogger struct{}

func (RealLogger) Info(s string)  { println("INFO:", s) }
func (RealLogger) Error(s string, err error) { println("ERROR:", s, err.Error()) }

func main() {
    db, _ := sql.Open("sqlite", "store.db")
    logger := RealLogger{}
    a := app.New(app.NewSQLDB(db), logger)
    a.Run()
}
