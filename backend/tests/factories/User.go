package factories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"social-network/pkg/common/infrastructure"
	"social-network/pkg/users"
	"social-network/pkg/users/infrastructure/persistance/mysql"
)
import "github.com/bluele/factory-go/factory"

var UserFactory = factory.NewFactory(
	&users.User{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return users.UserID(n), nil
}).
	Attr("Email", func(args factory.Args) (interface{}, error) {
		user := args.Instance().(*users.User)
		return fmt.Sprintf("user-%d-email@test.local", user.ID), nil

	}).
	Attr("Password", func(args factory.Args) (interface{}, error) {
		return users.HashedPassword("c8afeec4e9d29fa6307bc246965fe136a95bc47a9cfdedba0df256358eaa45ec0bf8d7a4333a4b13dc9a5508137d0f4d212272b27e64e41d4745a66b5f480759"), nil

	}).
	Attr("FirstName", func(args factory.Args) (interface{}, error) {
		user := args.Instance().(*users.User)
		return fmt.Sprintf("First Name - %d", user.ID), nil

	}).
	Attr("LastName", func(args factory.Args) (interface{}, error) {
		user := args.Instance().(*users.User)
		return fmt.Sprintf("Last Name - %d", user.ID), nil

	}).
	Attr("Age", func(args factory.Args) (interface{}, error) {
		return int8(rand.Intn(100) + 1), nil

	}).
	Attr("Gender", func(args factory.Args) (interface{}, error) {
		return users.Male, nil

	}).
	Attr("City", func(args factory.Args) (interface{}, error) {
		return "Moscow", nil

	}).
	Attr("Interests", func(args factory.Args) (interface{}, error) {
		user := args.Instance().(*users.User)
		return []string{
			fmt.Sprintf("Hobby-%d", user.ID),
			fmt.Sprintf("Hobby-%d", user.ID+1),
			fmt.Sprintf("Hobby-%d", user.ID+2),
		}, nil
	}).
	OnCreate(func(args factory.Args) error {
		user := args.Instance().(*users.User)

		tx := args.Context().Value(infrastructure.CtxTransactionKey).(*sql.Tx)

		interests, _ := json.Marshal(user.Interests)
		gender, err := mysql.GenderTransformerToDb(user.Gender)
		if err != nil {
			log.Fatalln(err)
		}

		return tx.QueryRowContext(args.Context(), `INSERT INTO users(
                  id, email, password, first_name, last_name, 
                  gender, age, interests, city
                  ) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			user.ID,
			user.Email,
			user.Password,
			user.FirstName,
			user.LastName,
			gender,
			user.Age,
			interests,
			user.City,
		).Err()
	})
