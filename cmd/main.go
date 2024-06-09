package main

import (
	"crypto/subtle"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apperrors "github.com/namhq1989/vocab-booster-server-admin/core/error"
	"github.com/namhq1989/vocab-booster-server-admin/core/logger"
	"github.com/namhq1989/vocab-booster-server-admin/internal/caching"
	"github.com/namhq1989/vocab-booster-server-admin/internal/config"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monitoring"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-admin/internal/queue"
	"github.com/namhq1989/vocab-booster-server-admin/internal/searching"
	appfile "github.com/namhq1989/vocab-booster-server-admin/internal/utils/file"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/waiter"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff"
)

func main() {
	var err error

	// config
	cfg := config.Init()

	// logger
	logger.Init(cfg.Environment)

	// app error
	apperrors.Init()

	// file
	appfile.Init()

	// server
	a := app{}
	a.cfg = cfg

	// rest
	a.rest = initRest(cfg)

	// grpc
	a.rpc = initRPC()

	// jwt
	a.jwt, err = appjwt.Init(cfg.AccessTokenSecret, cfg.RefreshTokenSecret, time.Second*time.Duration(cfg.AccessTokenTTL), time.Second*time.Duration(cfg.RefreshTokenTTL))
	if err != nil {
		panic(err)
	}

	// database
	a.database = database.NewDatabaseClient(cfg.MongoURL, cfg.MongoDBName)

	// searching
	a.searching = searching.NewSearchingClient(cfg.MeilisearchHost, cfg.MeilisearchAPIKey)

	// queue
	a.queue = queue.Init(cfg.QueueRedisURL, cfg.QueueConcurrency)

	// init queue's dashboard
	a.rest.Any(fmt.Sprintf("%s/*", queue.DashboardPath), echo.WrapHandler(queue.EnableDashboard(cfg.QueueRedisURL)), middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if !cfg.IsEnvRelease {
			return true, nil
		}
		return subtle.ConstantTimeCompare([]byte(username), []byte(cfg.QueueUsername)) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte(cfg.QueuePassword)) == 1, nil
	}))

	// caching
	a.caching = caching.NewCachingClient(cfg.CachingRedisURL)

	// monitoring
	a.monitoring = monitoring.Init(a.Rest(), cfg.SentryDSN, cfg.SentryMachine, cfg.Environment)

	// waiter
	a.waiter = waiter.New(waiter.CatchSignals())

	// modules
	a.modules = []monolith.Module{
		&audit.Module{},
		&auth.Module{},
		&staff.Module{},
	}

	// start
	if err = a.startupModules(); err != nil {
		panic(err)
	}

	fmt.Println("--- started vocab-booster-server-admin application")
	defer fmt.Println("--- stopped vocab-booster-server-admin application")

	// wait for other service starts
	a.waiter.Add(
		a.waitForRest,
		a.waitForRPC,
	)
	if err = a.waiter.Wait(); err != nil {
		panic(err)
	}
}
