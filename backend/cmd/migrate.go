package cmd

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"social-network/pkg/common/infrastructure"
	"social-network/pkg/common/infrastructure/configurator"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		var logger, _ = InitLogger()

		appConfig := configurator.NewAppConfig()

		//"user:password@tcp(host:port)/dbname?multiStatements=true"
		db, err := configurator.OpenMysqlConnection(appConfig)
		if err != nil {
			_ = logger.Log(infrastructure.LogError, err)
			os.Exit(1)
		}

		db.SetConnMaxLifetime(time.Minute * 1)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		driver, err := mysql.WithInstance(db, &mysql.Config{
			MigrationsTable: "migrations",
			DatabaseName:    appConfig.MySqlDatabase,
		})
		if err != nil {
			_ = logger.Log(infrastructure.LogError, err.Error())
			os.Exit(1)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://./db/migrations",
			"mysql",
			driver,
		)
		if err != nil {
			_ = logger.Log(infrastructure.LogError, err)
			os.Exit(1)
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			_ = logger.Log(infrastructure.LogError, err)

			os.Exit(1)
		}

		_ = logger.Log(infrastructure.LogMigrationsStatus, "Migrations completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
