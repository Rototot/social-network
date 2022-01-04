package cmd

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"social-network/pkg/config"
	"social-network/pkg/users/endpoints"
	"social-network/pkg/users/persistance/mysql"
	"social-network/pkg/users/services"
	"social-network/pkg/users/transport"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
		flag.Parse()

		var logger log.Logger
		{
			logger = log.NewLogfmtLogger(os.Stderr)
			logger = log.With(logger, "ts", log.DefaultTimestampUTC)
			logger = log.With(logger, "caller", log.DefaultCaller)
		}

		fmt.Println("start server ...")

		// init conn
		var cnf = config.NewAppConfig()
		var conn *sql.DB
		{
			db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
				cnf.MySqlUser,
				cnf.MySqlPassword,
				cnf.MySqlHost,
				cnf.MySqlPort,
				cnf.MySqlDatabase,
			))
			if err != nil {
				logger.Log("database", err)
				os.Exit(1)
			}
			conn = db
		}

		defer conn.Close()

		var httpHandler http.Handler
		{
			var userRepository = mysql.NewUserRepository(conn)
			userEndpoints, err := endpoints.MakeEndpoints(
				services.NewLoginService(
					userRepository,
					services.NewPasswordHasher(),
					nil,
					// todo вынести в конфиг
					time.Hour*24*30,
				),
				services.NewRegisterService(),
				// todo поправить
				services.NewLogoutService(nil),
			)
			if err != nil {
				logger.Log("endpoints", err)
				os.Exit(1)
			}

			var r = mux.NewRouter()
			var rAPI = r.PathPrefix("/api").Subrouter()
			httpHandler = transport.NewHttpHandler(
				rAPI,
				userEndpoints,
				logger,
			)
		}

		errs := make(chan error)
		go listenStopSignal()(errs)

		go func() {
			logger.Log("transport", "HTTP", "addr", *httpAddr)
			errs <- http.ListenAndServe(*httpAddr, httpHandler)
		}()

		logger.Log("exit", <-errs)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func listenStopSignal() func(errs chan error) {
	return func(errs chan error) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}
}
