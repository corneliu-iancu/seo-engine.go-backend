// Handles Business Logic of our application.

package app

import (
	"fmt"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/helper"
	"net/url"
)

// @todo: Remove, as it won't be part of the public API.
//func (app BusinessLogicImpl) GetAllSegments() ([]rule.Segment, error) {
//	rules, err := app.RulesService.GetSegments()
//	return rules, err
//}

func (app BusinessLogicImpl) GetAllRules() ([]rule.Rule, error) {
	app.Logger.Debug("[DEBUG] 💡 App(business layer) -> Get all rules.")
	return app.RulesService.GetRules()
}

// Adds a rule to the storage layer privided an URL.
// @todo: rename, return type.
func (app BusinessLogicImpl) CreateRule(u *url.URL) ([]rule.Segment, error) {
	app.Logger.Debug("[DEBUG] app.jurney / AddRule call")
	pathParams := helper.GetURIAsSlice(u)
	// @todo: Handle query parameters.
	results := app.RulesService.CreateFromListOfStrings(pathParams)
	return results, nil
}

// Returns a match Rule node based on a URL.
func (app BusinessLogicImpl) GetMatch(u *url.URL) ([]rule.Rule, error) {
	// pathParams := helper.GetURIAsSlice(u)
	return app.RulesService.GetMatch(u)
}

// Creates the rules table.
func (app BusinessLogicImpl) CreateRulesTable() error {
	fmt.Println("[DEBUG] create rules table")

	//@todo: handle ResourceInUseException error.
	return app.RulesService.CreateRulesTable()
}
