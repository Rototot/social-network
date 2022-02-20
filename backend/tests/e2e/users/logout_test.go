package users

import (
	"context"
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"social-network/pkg/users"
	"social-network/pkg/users/transport"
	"social-network/tests/factories"
	"social-network/tests/utils"
	"testing"
)

type TestLogoutSuite struct {
	utils.ApiTestSuite
}

func (s *TestLogoutSuite) TestPostLogoutOK() {
	rawUser := utils.LoadFactory(s.T(), s.Conn, func(t *testing.T, ctx context.Context) (interface{}, error) {
		return factories.UserFactory.CreateWithContext(ctx)
	})
	user := rawUser.(*users.User)
	sessionToken := utils.TestSessionGenerator(s.T(), user.ID)

	apitest.New(). // configuration
			EnableNetworking().
			Post(fmt.Sprintf("%s/api/auth/logout", s.ApiUrl)). // request
			Header(transport.HttpHeaderToken, sessionToken).
			Expect(s.T()). // expectations
			Status(http.StatusOK).
			End()

	apitest.New(). // configuration
			EnableNetworking().
			Get(fmt.Sprintf("%s/api/user/card/%d", s.ApiUrl, user.ID)). // request
			Header(transport.HttpHeaderToken, sessionToken).
			Expect(s.T()). // expectations
			Status(http.StatusUnauthorized).
			End()
}

func TestTestLogoutSuite(t *testing.T) {
	suite.Run(t, new(TestLogoutSuite))
}
