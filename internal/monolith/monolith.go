package monolith

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/caching"
	"github.com/namhq1989/vocab-booster-server-admin/internal/config"
	"github.com/namhq1989/vocab-booster-server-admin/internal/database"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monitoring"
	"github.com/namhq1989/vocab-booster-server-admin/internal/queue"
	"github.com/namhq1989/vocab-booster-server-admin/internal/searching"
	appjwt "github.com/namhq1989/vocab-booster-server-admin/internal/utils/jwt"
	"github.com/namhq1989/vocab-booster-server-admin/internal/utils/waiter"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.Server
	Database() *database.Database
	Searching() *searching.Searching
	Caching() *caching.Caching
	Rest() *echo.Echo
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	JWT() *appjwt.JWT
	Monitoring() *monitoring.Monitoring
	Queue() *queue.Queue
}

type Module interface {
	Name() string
	Startup(ctx *appcontext.AppContext, monolith Monolith) error
}
