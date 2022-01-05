package users

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/suite"
	"net/http"
	"social-network/tests/utils"
	"testing"
)

type TestLoginSuite struct {
	utils.ApiTestSuite
}

func (s *TestLoginSuite) TestPostLoginWhenCorrectDataThenOK() {
	body, _ := json.Marshal(map[string]string{
		"email":    "user-1-email@test.local",
		"password": "test_password",
	})

	apitest.New(). // configuration
			EnableNetworking().
			Post(fmt.Sprintf("%s/api/auth/login", s.ApiUrl)). // request
			JSON(body).
			Expect(s.T()). // expectations
			Body(`{"token": "4dff4ea340f0a823f15d3f4f01ab62eae0e5da579ccb851f8db9dfe84c58b2b37b89903a740e1ee172da793a6e79d560e5f7f9bd058a12a280433ed6fa46510a"}`).
			Status(http.StatusOK).
			End()
}

func TestLoginTestSuite(t *testing.T) {
	suite.Run(t, new(TestLoginSuite))
}
