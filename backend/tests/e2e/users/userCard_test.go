package users

import (
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"social-network/pkg/users"
	"social-network/tests/utils"
	"testing"
)

type TestUserCardSuite struct {
	utils.ApiTestSuite
}

func (s *TestUserCardSuite) TestPostLoginWhenCorrectDataThenOK() {
	expectedUserId := users.UserID(1)

	apitest.New(). // configuration
			EnableNetworking().
			Get(fmt.Sprintf("%s/api/user/card/%d", s.ApiUrl, expectedUserId)). // request
			Expect(s.T()).                                                     // expectations
			Body(`{"token":"testtoken"}`).
			Status(http.StatusOK).
			End()
}

func TestUserCardTestSuite(t *testing.T) {
	suite.Run(t, new(TestUserCardSuite))
}
