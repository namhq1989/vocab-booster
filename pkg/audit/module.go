package audit

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/application"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/grpc"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/infrastructure"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/rest"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/audit/worker"
)

type Module struct{}

func (Module) Name() string {
	return "AUDIT"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// infrastructure
		auditRepository = infrastructure.NewAuditRepository(mono.Database())

		// application
		app = application.New()
	)

	// rest server
	if err := rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT()); err != nil {
		return err
	}

	// grpc server
	if err := grpc.RegisterServer(ctx, mono.RPC()); err != nil {
		return err
	}

	// worker
	w := worker.New(mono.Queue(), auditRepository)
	w.Start()

	return nil
}
