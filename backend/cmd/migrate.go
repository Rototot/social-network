package cmd

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"social-network/pkg/config"
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
		cnf := config.NewAppConfig()

		//"user:password@tcp(host:port)/dbname?multiStatements=true"
		db, _ := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
			cnf.MySqlUser,
			cnf.MySqlPassword,
			cnf.MySqlHost,
			cnf.MySqlPort,
			cnf.MySqlDatabase,
		))
		db.SetConnMaxLifetime(time.Minute * 1)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)

		driver, err := mysql.WithInstance(db, &mysql.Config{
			MigrationsTable: "migrations",
			DatabaseName:    cnf.MySqlDatabase,
		})
		if err != nil {
			log.Fatalln(err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://./db/migrations",
			"mysql",
			driver,
		)
		if err != nil {
			log.Fatalln(err)
		}

		err = m.Up()
		if err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalln(err)
		}

		log.Println("Migrations completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
