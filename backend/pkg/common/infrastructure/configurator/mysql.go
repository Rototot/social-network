package configurator

import (
	"database/sql"
	"fmt"
)

func OpenMysqlConnection(conf *AppConfig) (*sql.DB, error) {
	return sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		conf.MySqlUser,
		conf.MySqlPassword,
		conf.MySqlHost,
		conf.MySqlPort,
		conf.MySqlDatabase,
	))
}
