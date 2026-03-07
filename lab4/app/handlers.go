package app

import (
    "html/template"
    "net/http"
    "strconv"
)

type Product struct {
    ID      int
    Model   string
    Company string
    Price   int
}

func (a *App) handleIndex(w http.ResponseWriter, r *http.Request) {
    rows, err := a.db.Query("SELECT id, model, company, price FROM products")
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var list []Product

    for rows.Next() {
        var p Product
        rows.Scan(&p.ID, &p.Model, &p.Company, &p.Price)
        list = append(list, p)
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, list)
}

func (a *App) handleCreate(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        tmpl := template.Must(template.ParseFiles("templates/create.html"))
        tmpl.Execute(w, nil)
        return
    }

    model := r.FormValue("model")
    company := r.FormValue("company")
    price, _ := strconv.Atoi(r.FormValue("price"))

    _, err := a.db.Exec(
        "INSERT INTO products (model, company, price) VALUES (?, ?, ?)",
        model, company, price,
    )
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (a *App) handleDelete(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))

    _, err := a.db.Exec("DELETE FROM products WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Для тестов
func (a *App) HandleIndexForTest(w http.ResponseWriter, r *http.Request) {
    a.handleIndex(w, r)
}

func (a *App) HandleCreateForTest(w http.ResponseWriter, r *http.Request) {
    a.handleCreate(w, r)
}

func (a *App) HandleDeleteForTest(w http.ResponseWriter, r *http.Request) {
    a.handleDelete(w, r)
}
