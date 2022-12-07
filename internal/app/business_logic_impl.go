// Handles Business Logic of our application.

package app

import (
	"fmt"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/helper"
	"net/url"
)

// ################ Interface methods ################ //

func (app BusinessLogicImpl) GetAllRules() ([]rule.Rule, error) {
	app.Logger.Debug("[DEBUG] 💡 App(business layer) -> Get all rules.")
	return app.RulesService.GetRules()
}

// Adds a rule to the storage layer privided an URL.
// @returns []rule.Segment, Error
func (app BusinessLogicImpl) CreateRule(u *url.URL, meta metadata.Metadata) ([]rule.Segment, error) {
	app.Logger.Debug("[DEBUG] app.jurney / AddRule call")
	pathParams := helper.GetURIAsSlice(u)
	// @todo: Handle query parameters.
	// fmt.Println(meta.Title)
	// fmt.Println(meta.MetaTags)

	results := app.RulesService.CreateFromListOfStrings(pathParams, meta)
	return results, nil
}

// Returns a match Rule node based on a URL.
func (app BusinessLogicImpl) GetMatch(u *url.URL) ([]rule.Rule, error) {
	return app.RulesService.GetMatch(u)
}

func (app BusinessLogicImpl) GetURLBySegmentId(segmentId string) ([]rule.Segment, error) {
	return app.RulesService.GetURLBySegmentId(segmentId)
}

// ################ End Interface methods ################ //

// Creates the rules table.
func (app BusinessLogicImpl) CreateRulesTable() error {
	fmt.Println("[DEBUG] create rules table")

	//@todo: handle ResourceInUseException error.
	return app.RulesService.CreateRulesTable()
}
