package utils

import (
	"crypto/sha512"
	"fmt"
	"social-network/pkg/users"
	"testing"
)

func TestSessionGenerator(t *testing.T, id users.UserID) string {
	hasher := sha512.New()
	_, err := hasher.Write([]byte(id.String()))
	if err != nil {
		t.Fatal(err)
	}

	return fmt.Sprintf("%x", hasher.Sum(nil))
}
