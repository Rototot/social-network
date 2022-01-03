package utils

import (
	"database/sql"
	"github.com/stretchr/testify/suite"
)

type ApiTestSuite struct {
	suite.Suite
	Conn   *sql.DB
	ApiUrl string
}

func (suite *ApiTestSuite) SetupSuite() {
	suite.Conn = NewTestDbConnection(suite.T())
	CleanAll(suite.T(), suite.Conn)
	LoadFactories(suite.T(), suite.Conn)

	suite.ApiUrl = "http://127.0.0.1:8000"

}

// TearDownAllSuite after all tests
func (suite *ApiTestSuite) TearDownSuite() {
	defer suite.Conn.Close()
}
