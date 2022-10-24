package app

import (
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// Application Services.
type Application struct {
	// Logger logging.Logger
	RulesService RulesService
}

// Used by our http handlers in order to process/retrieve data
type App interface {
	GetAllSegments() ([]rule.Segment, error)
	GetAllRules() ([]rule.Rule, error)
	AddRule(*url.URL) ([]rule.Segment, error) // @todo: rename to CreateRule, and return type.
	GetMatch(*url.URL) ([]rule.Rule, error)
}
