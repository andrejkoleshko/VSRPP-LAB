package app

import "database/sql"

type SQLDB struct {
    db *sql.DB
}

func NewSQLDB(db *sql.DB) *SQLDB {
    return &SQLDB{db: db}
}

func (s *SQLDB) Exec(query string, args ...any) (interface{}, error) {
    return s.db.Exec(query, args...)
}

func (s *SQLDB) Query(query string, args ...any) (Rows, error) {
    r, err := s.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    return &SQLRows{r}, nil
}

type SQLRows struct {
    r *sql.Rows
}

func (s *SQLRows) Next() bool { return s.r.Next() }
func (s *SQLRows) Scan(dest ...any) error { return s.r.Scan(dest...) }
func (s *SQLRows) Close() error { return s.r.Close() }
