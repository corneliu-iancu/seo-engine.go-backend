package app

// Application Services.
type Application struct {
	// Logger logging.Logger
	RulesService RulesService
}

type App interface {
	GetAllRules() error
}
