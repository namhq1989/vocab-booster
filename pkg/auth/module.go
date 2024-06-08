package auth

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/grpcclient"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/application"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/grpc"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/infrastructure"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/auth/rest"
)

type Module struct{}

func (Module) Name() string {
	return "AUTH"
}

func (Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	staffGRPCClient, err := grpcclient.NewStaffClient(ctx, mono.Config().GRPCPort)
	if err != nil {
		return err
	}

	var (
		cfg = mono.Config()

		ssoRepository            = infrastructure.NewSSORepository(cfg.SSOGoogleClientID, cfg.SSOGoogleClientSecret)
		staffAuthTokenRepository = infrastructure.NewStaffAuthTokenRepository(mono.Database())
		jwtRepository            = infrastructure.NewJwtRepository(mono.JWT())
		staffHub                 = infrastructure.NewStaffHub(staffGRPCClient)

		// app
		app = application.New(staffAuthTokenRepository, ssoRepository, jwtRepository, staffHub)
	)

	// rest server
	if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), cfg.IsEnvRelease); err != nil {
		return err
	}

	// grpc server
	if err = grpc.RegisterServer(ctx, mono.RPC(), staffHub); err != nil {
		return err
	}

	return nil
}
