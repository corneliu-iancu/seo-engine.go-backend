package app

import "github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"

type RulesService struct {
	rulesRepository adaptor.RuleDynamoRepository
}

func NewRulesService(rulesRepository adaptor.RuleDynamoRepository) RulesService {
	return RulesService{
		rulesRepository: rulesRepository,
	}
}

// add RulesService methods here.
