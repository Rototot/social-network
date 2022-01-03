package mysql

import (
	"fmt"
	"social-network/pkg/users"
)

const Male = "male"
const Female = "female"

func GenderTransformerToDb(v users.Gender) (string, error) {
	if v == users.Male {
		return Male, nil
	}

	if v == users.Female {
		return Female, nil
	}

	return "", fmt.Errorf("unexpected value %d", v)
}

func GenderTransformer(v string) (users.Gender, error) {
	if v == Male {
		return users.Male, nil
	}

	if v == Female {
		return users.Female, nil
	}

	return 0, fmt.Errorf("unexpected value %s", v)
}
