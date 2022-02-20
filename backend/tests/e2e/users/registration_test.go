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

type TestRegistrationSuite struct {
	utils.ApiTestSuite
}

func (s *TestRegistrationSuite) TestPostLoginWhenCorrectDataThenOK() {
	rawUser := utils.LoadFactory(s.T(), s.Conn, func(t *testing.T, ctx context.Context) (interface{}, error) {
		return factories.UserFactory.CreateWithContext(ctx)
	})
	user := rawUser.(*users.User)

	body, _ := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": "test_password",
	})

	apitest.New(). // configuration
			EnableNetworking().
			Post(fmt.Sprintf("%s/api/auth/register", s.ApiUrl)). // request
			JSON(body).
			Expect(s.T()). // expectations
			Body(`{"token": "4dff4ea340f0a823f15d3f4f01ab62eae0e5da579ccb851f8db9dfe84c58b2b37b89903a740e1ee172da793a6e79d560e5f7f9bd058a12a280433ed6fa46510a"}`).
			Status(http.StatusOK).
			End()
}

func TestTestRegistrationSuite(t *testing.T) {
	suite.Run(t, new(TestRegistrationSuite))
}
