package utils

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "testing"
)

func NewTestDbConnection(t *testing.T) *sql.DB {

    cnf := NewTestConfiguration(t)

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
        cnf.MySqlUser,
        cnf.MySqlPassword,
        cnf.MySqlHost,
        cnf.MySqlPort,
        cnf.MySqlDatabase,
    ))

    if err != nil {
        t.Fatal(err)
    }

    return db
}
