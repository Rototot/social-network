package utils

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"testing"
)

func CleanAll(
	t *testing.T,
	conn *sql.DB,
) {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		CleanDB(t, conn)
		t.Log("mysql: clean completed")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		CleanRedis(t)

		t.Log("redis: clean completed")
	}()

	wg.Wait()
}

func CleanDB(t *testing.T, conn *sql.DB) {
	cnf := NewTestConfiguration(t)

	ctx := context.Background()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := tx.Query("SET FOREIGN_KEY_CHECKS=0;"); err != nil {
		t.Fatal(err)
	}
	rows, err := tx.Query(fmt.Sprintf(`
        SELECT TABLE_NAME
        FROM INFORMATION_SCHEMA.TABLES 
        WHERE TABLE_TYPE = 'BASE TABLE' 
          AND TABLE_SCHEMA = '%s'
          AND TABLE_NAME <> 'migrations'`, cnf.MySqlDatabase))
	if err != nil {
		t.Fatal(err)
	}

	var tables []string
	for rows.Next() {
		var table string
		err := rows.Scan(&table)
		if err != nil {
			t.Fatal(err)
		}

		tables = append(tables, table)
	}

	var wg sync.WaitGroup
	for _, table := range tables {
		wg.Add(1)
		go func(tb string) {
			defer wg.Done()
			_, err := tx.Query(fmt.Sprintf("TRUNCATE TABLE %s", tb))
			if err != nil {
				t.Fatal(err)
			}
		}(table)
	}

	wg.Wait()

	if _, err := tx.Query("SET FOREIGN_KEY_CHECKS=1;"); err != nil {
		t.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}

}

func CleanRedis(_ *testing.T) {

}
