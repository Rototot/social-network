package utils

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	"testing"
)

func CleanAll(
	t *testing.T,
	conn *sql.DB,
) {

	var wg sync.WaitGroup
	wg.Add(1)
	go func(t *testing.T) {
		defer wg.Done()
		CleanDB(t, conn)
		t.Log("mysql: clean completed")
	}(t)
	wg.Add(1)
	go func(t *testing.T) {
		defer wg.Done()
		CleanRedis(t)

		t.Log("redis: clean completed")
	}(t)

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
		go func(t *testing.T, tb string, tx *sql.Tx) {
			defer wg.Done()
			if _, err := tx.Query(fmt.Sprintf("TRUNCATE TABLE %s", tb)); err != nil {
				log.Fatalln(err)
			}

			if _, err := tx.Query(fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", tb)); err != nil {
				log.Fatalln(err)
			}

		}(t, table, tx)
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
