package utils

import (
	"context"
	"social-network/pkg/common/infrastructure"
	"social-network/pkg/users"
	"social-network/tests/factories"
	"testing"
)

import (
	"database/sql"
	testfixtures "github.com/go-testfixtures/testfixtures/v3"
)

type FactoryCreator func(t *testing.T, ctx context.Context) (interface{}, error)

func LoadFixtures(t *testing.T, conn *sql.DB, extraPath string) {
	path := extraPath
	if path == "" {
		path = "tests/fixtures"
	}

	_, err := testfixtures.New(
		testfixtures.Database(conn),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(path),
	)
	if err != nil {
		t.Fatal(err)
	}
}

// LoadFactory use run LoadFactory(t, conn, func (t, ctx) { return factories.UserFactory.CreateWithContext(ctx)})
func LoadFactory(t *testing.T, conn *sql.DB, f FactoryCreator) interface{} {
	ctx := context.Background()
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	ctxTransaction := context.WithValue(ctx, infrastructure.CtxTransactionKey, tx)

	entity, err := f(t, ctxTransaction)
	if err != nil {
		t.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}

	t.Log("factory: loading completed")

	return entity
}

func LoadFactories(t *testing.T, conn *sql.DB) []*users.User {
	var listUsers = make([]*users.User, 10)
	for i := range listUsers {
		user := LoadFactory(t, conn, func(t *testing.T, ctx context.Context) (interface{}, error) {
			return factories.UserFactory.CreateWithContext(ctx)
		})

		listUsers[i] = user.(*users.User)
	}

	t.Log("factory: loading completed")
	return listUsers
}
