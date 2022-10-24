// Handles Business Logic of our application.

package app

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// @todo: Route for retriving all segments.
func (app Application) GetAllSegments() ([]rule.Segment, error) {
	rules, err := app.RulesService.GetSegments()
	return rules, err
}

func (app Application) GetAllRules() ([]rule.Rule, error) {
	fmt.Println("[DEBUG] 💡 App(business layer) -> Get all rules.")

	return app.RulesService.GetRules()
}

// @todo: add docs
// @todo: rename, return type.
func (app Application) AddRule(u *url.URL) ([]rule.Segment, error) {
	fmt.Println("[DEBUG] app.jurney / AddRule call")

	fmt.Println(u)

	pathParams := []string{}
	pathParams = append(pathParams, u.Host)
	// exclude first element, as the Path starts with an "/"
	pathParams = append(pathParams, strings.Split(u.Path, string('/'))[1:]...)

	// @todo: Handle query parameters.

	results := app.RulesService.CreateFromListOfSegments(pathParams)
	return results, nil
}

func (app Application) GetMatch(u *url.URL) ([]rule.Rule, error) {
	fmt.Println("[DEBUG] Application - GetMatch")
	return nil, nil
}

// @todo: add docs
func (app Application) CreateRulesTable() error {
	fmt.Println("[DEBUG] create rules table")

	//@todo: handle ResourceInUseException error.
	return app.RulesService.CreateRulesTable()
}
