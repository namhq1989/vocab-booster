package staff

import (
	"github.com/namhq1989/vocab-booster-server-admin/core/appcontext"
	"github.com/namhq1989/vocab-booster-server-admin/internal/monolith"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff/domain"
	"github.com/namhq1989/vocab-booster-server-admin/pkg/staff/infrastructure"
)

type Module struct{}

func (Module) Name() string {
	return "STAFF"
}

func (m Module) Startup(ctx *appcontext.AppContext, mono monolith.Monolith) error {
	var (
		// cfg = mono.Config()

		staffRepository = infrastructure.NewStaffRepository(mono.Database())

		// app
		// app = application.New()
	)

	// rest server
	// if err = rest.RegisterServer(ctx, app, mono.Rest(), mono.JWT(), cfg.IsEnvRelease); err != nil {
	// 	return err
	// }

	// grpc server
	// if err = grpc.RegisterServer(ctx, mono.RPC(), staffHub); err != nil {
	// 	return err
	// }

	// initialize
	m.initialize(ctx, staffRepository)

	return nil
}

func (m Module) initialize(ctx *appcontext.AppContext, staffRepository domain.StaffRepository) {
	m.addSelfStaff(ctx, staffRepository)
}

func (Module) addSelfStaff(ctx *appcontext.AppContext, staffRepository domain.StaffRepository) {
	const (
		name  = "Nam HQ"
		email = "namhq.1989@gmail.com"
	)

	total, err := staffRepository.CountByEmail(ctx, email)
	if err != nil {
		ctx.Logger().Error("[staff initializing] failed to count staff", err, appcontext.Fields{})
		panic(err)
	}

	if total > 0 {
		// ctx.Logger().Text(fmt.Sprintf("staff %s already exists", email))
		return
	}

	staff, err := domain.NewStaff(name, email, domain.StaffRoleAdmin.String(), domain.StaffStatusActive.String())
	if err != nil {
		ctx.Logger().Error("[staff initializing]  failed to create domain staff", err, appcontext.Fields{})
		panic(err)
	}

	if err = staffRepository.CreateStaff(ctx, *staff); err != nil {
		ctx.Logger().Error("[staff initializing]  failed to create db staff", err, appcontext.Fields{})
		panic(err)
	}
}
