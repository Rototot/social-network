package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"os"
	"os/signal"
	"social-network/pkg/common/infrastructure"
	"social-network/pkg/common/infrastructure/configurator"
	"social-network/pkg/ping"
	userEndpoints "social-network/pkg/users/endpoints"
	"social-network/pkg/users/infrastructure/persistance/mysql"
	redis2 "social-network/pkg/users/infrastructure/persistance/redis"
	userServices "social-network/pkg/users/services"
	userTransport "social-network/pkg/users/transport"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

const defaultHttpAddr = ":8000"
const defaultSessionTimeExpire = time.Hour * 24 * 30

func init() {
	rootCmd.AddCommand(serverCmd)
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		var httpAddr = flag.String("http.addr", defaultHttpAddr, "HTTP listen address")
		flag.Parse()

		// init loggers
		var logger, httpLogger = InitLogger()

		// init infra components
		var appConfig = configurator.NewAppConfig()
		var conn *sql.DB
		{
			db, err := configurator.OpenMysqlConnection(appConfig)
			if err != nil {
				_ = logger.Log(infrastructure.LogDatabase, err)
				os.Exit(1)
			}
			conn = db
		}
		defer conn.Close()

		var redisClient *redis.Client
		{
			client, err := configurator.OpenRedisConnection(appConfig)
			if err != nil {
				_ = logger.Log(infrastructure.LogRedis, err)
				os.Exit(1)
			}

			redisClient = client
		}

		// init http handlers
		var userHandler http.Handler
		var authHandler http.Handler
		var pingHandler http.Handler
		{
			// user and auth handlers
			{
				userRepository := mysql.NewUserRepository(conn)
				sessionStorage := redis2.NewSessionRepository(redisClient)
				endpoints, err := userEndpoints.MakeEndpoints(
					userServices.NewLoginService(
						userRepository,
						userServices.NewPasswordHasher(),
						sessionStorage,
						defaultSessionTimeExpire,
					),
					userServices.NewRegisterService(),
					// todo поправить
					userServices.NewLogoutService(sessionStorage),
					userRepository,
					sessionStorage,
				)
				if err != nil {
					_ = logger.Log(infrastructure.LogEndpoints, err)
					os.Exit(1)
				}

				userHandler = userTransport.MakeUserHttpHandler(
					endpoints,
					httpLogger,
				)

				authHandler = userTransport.MakeAuthHttpHandler(
					endpoints,
					httpLogger,
				)
			}

			pingHandler = ping.MakePingHttpHandler(
				ping.MakeEndpoints(),
				logger,
			)
		}

		// merge http handlers to server
		var mux = http.NewServeMux()
		mux.Handle("/api/auth/", authHandler)
		mux.Handle("/api/user/", userHandler)
		mux.Handle("/", pingHandler)

		h := http.DefaultServeMux
		h.Handle("/", mux)

		var server = &http.Server{
			Addr: *httpAddr,
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      h,
		}

		errs := make(chan error, 2)
		defer close(errs)

		go func() {
			_ = logger.Log(infrastructure.LogTransport, "http", "addr", *httpAddr)
			errs <- server.ListenAndServe()
		}()

		stopSignals := make(chan os.Signal, 1)
		defer close(stopSignals)

		go func() {
			signal.Notify(stopSignals, syscall.SIGINT, syscall.SIGTERM)
			errs <- fmt.Errorf("%s", <-stopSignals)
		}()

		_ = logger.Log(infrastructure.LogError, <-errs)

		wait := time.Second * 15
		ctx, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()

		_ = logger.Log(infrastructure.LogServer, "shutdown", infrastructure.LogError, server.Shutdown(ctx))
	},
}
