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

type TestLogoutSuite struct {
	utils.ApiTestSuite
}

func (s *TestLogoutSuite) TestPostLoginWhenCorrectDataThenOK() {
	body, _ := json.Marshal(map[string]string{
		"email":    "Karley_Dach@jasper.info",
		"password": "test_password",
	})

	apitest.New(). // configuration
			EnableNetworking().
			Post(fmt.Sprintf("%s/api/auth/login", s.ApiUrl)). // request
			JSON(body).
			Expect(s.T()). // expectations
			Body(`{"token": "testtoken"}`).
			Status(http.StatusOK).
			End()
}

func TestTestLogoutSuite(t *testing.T) {
	suite.Run(t, new(TestLogoutSuite))
}
