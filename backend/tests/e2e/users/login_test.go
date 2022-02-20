package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"social-network/pkg/users"
	"social-network/tests/factories"
	"social-network/tests/utils"
	"testing"
)

type TestLoginSuite struct {
	utils.ApiTestSuite
}

func (s *TestLoginSuite) TestPostLoginWhenCorrectDataThenOK() {
	rawUser := utils.LoadFactory(s.T(), s.Conn, func(t *testing.T, ctx context.Context) (interface{}, error) {
		return factories.UserFactory.CreateWithContext(ctx)
	})
	user := rawUser.(*users.User)
	expectedToken := utils.TestSessionGenerator(s.T(), user.ID)

	body, _ := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": "test_password",
	})

	apitest.New(). // configuration
			EnableNetworking().
			Post(fmt.Sprintf("%s/api/auth/login", s.ApiUrl)). // request
			JSON(body).
			Expect(s.T()). // expectations
			Body(fmt.Sprintf(`{"token": "%s"}`, expectedToken)).
			Status(http.StatusOK).
			End()
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(TestLoginSuite))
}
