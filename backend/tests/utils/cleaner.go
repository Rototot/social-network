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
		t.Log("---- Clean Database ----")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		CleanRedis(t)

		t.Log("---- Clean Redis ----")
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
	tx.Query("SET FOREIGN_KEY_CHECKS=0;")
	tx.Query(fmt.Sprintf(`
        SELECT 'TRUNCATE TABLE ' + TABLE_NAME 
        FROM INFORMATION_SCHEMA.TABLES 
        WHERE TABLE_TYPE = 'BASE TABLE' 
          AND TABLE_SCHEMA = '%s'`, cnf.MySqlDatabase))
	tx.Query("SET FOREIGN_KEY_CHECKS=1;")

	tx.Commit()
}

func CleanRedis(_ *testing.T) {

}
