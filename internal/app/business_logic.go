package app

import (
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"
	"go.uber.org/zap"
	"net/url"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// ================================================
// Application
// ================================================
// Our platform implementation for the business logic.
// Extended by:
// - business_logic_impl.go
// ================================================
type BusinessLogicImpl struct {
	Logger       *zap.Logger
	RulesService RulesService
}

// ================================================
// Business Logic.
// ================================================
// Used by our http handlers in order to process/retrieve data
// ================================================
type BusinessLogic interface {
	CreateRule(*url.URL, metadata.Metadata) ([]rule.Segment, error)
	GetAllRules() ([]rule.Rule, error)
	GetMatch(*url.URL) ([]rule.Rule, error)
	GetURLBySegmentId(segmentId string) ([]rule.Segment, error)
}
