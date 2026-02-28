package app

import "database/sql"

type Rows interface {
    Next() bool
    Scan(dest ...any) error
    Close() error
}

type DB interface {
    Exec(query string, args ...any) (sql.Result, error)
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

    _, err = a.db.Exec(
        "INSERT INTO products (model, company, price) VALUES ('iPhone X', $1, $2)",
        "Apple", 72000,
    )
    if err != nil {
        a.l.Error("insert failed", err)
        return err
    }

    rows, err := a.db.Query("SELECT * FROM products")
    if err != nil {
        a.l.Error("select failed", err)
        return err
    }
    rows.Close()

    _, err = a.db.Exec("UPDATE products SET price = $1 WHERE id = $2", 69000, 1)
    if err != nil {
        a.l.Error("update failed", err)
        return err
    }

    _, err = a.db.Exec("DELETE FROM products WHERE id = $1", 1)
    if err != nil {
        a.l.Error("delete failed", err)
        return err
    }

    a.l.Info("Success")
    return nil
}
