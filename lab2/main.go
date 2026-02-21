package main

import (
    "database/sql"
    "fmt"
    _ "modernc.org/sqlite"
)

type product struct{
    id int
    model string
    company string
    price int
}

func main() { 

    db, err := sql.Open("sqlite", "store.db")
    _, err = db.Exec(` CREATE TABLE IF NOT EXISTS products ( id INTEGER PRIMARY KEY AUTOINCREMENT, model TEXT, company TEXT, price INTEGER ) `)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)", 
        "Apple", 72000)
    if err != nil{
        panic(err)
    }
    fmt.Println(result.LastInsertId())  // id последнего добавленного объекта
    fmt.Println(result.RowsAffected())  // количество добавленных строк


    
    rows, err := db.Query("select * from Products")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    products := []product{}
     
    for rows.Next(){
        p := product{}
        err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }

    for _, p := range products{
        fmt.Println(p.id, p.model, p.company, p.price)
    }


   
    result, err = db.Exec("update Products set price = $1 where id = $2", 69000, 1)
    if err != nil{
        panic(err)
    }
    fmt.Println(result.RowsAffected())  // количество обновленных строк


    
    result, err = db.Exec("delete from Products where id = $1", 1)
    if err != nil{
        panic(err)
    }
    fmt.Println(result.RowsAffected())  // количество удаленных строк
}
