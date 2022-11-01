package app

import (
	"go.uber.org/zap"
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// ================================================
// Application
// ================================================
// Our platform implementation for the business logic.
// Extended by:
// - journey.go
// ================================================
type Application struct {
	Logger       *zap.Logger
	RulesService RulesService
}

// ================================================
// Business Logic.
// ================================================
// Used by our http handlers in order to process/retrieve data
// ================================================
type App interface {
	GetAllSegments() ([]rule.Segment, error) // @todo: remove me.
	GetAllRules() ([]rule.Rule, error)
	AddRule(*url.URL) ([]rule.Segment, error) // @todo: rename to CreateRule, and return type.
	GetMatch(*url.URL) ([]rule.Rule, error)
}
