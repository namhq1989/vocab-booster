package application

type (
	Commands interface {
	}
	Queries interface {
	}
	Instance interface {
		Commands
		Queries
	}

	appCommandHandlers struct {
	}
	appQueryHandler struct {
	}
	Application struct {
		appCommandHandlers
		appQueryHandler
	}
)

var _ Instance = (*Application)(nil)

func New() *Application {
	return &Application{
		appCommandHandlers: appCommandHandlers{},
		appQueryHandler:    appQueryHandler{},
	}
}
