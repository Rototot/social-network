package utils

import (
	"context"
	"os"
	"social-network/pkg/users"
	"social-network/tests/factories"
	"testing"
)

import (
	"database/sql"
	testfixtures "github.com/go-testfixtures/testfixtures/v3"
)

func LoadFixtures(t *testing.T, conn *sql.DB, extraPath string) {
	path := extraPath
	if path == "" {
		path = "tests/fixtures"
	}

	os.Getwd()

	_, err := testfixtures.New(
		testfixtures.Database(conn),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(path),
	)
	if err != nil {
		t.Fatal(err)
	}
}

func LoadFactories(t *testing.T, conn *sql.DB) []*users.User {
	ctx := context.Background()

	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		t.Fatal(err)
	}

	var listUsers = make([]*users.User, 10)
	for i, _ := range listUsers {
		user, err := factories.UserFactory.CreateWithContext(
			context.WithValue(context.Background(), "tx", tx),
		)
		if err != nil {
			t.Fatal(err)
		}

		listUsers[i] = user.(*users.User)
	}

	return listUsers
}
