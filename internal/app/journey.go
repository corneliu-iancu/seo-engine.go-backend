// Handles Business Logic of our application.

package app

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
)

// @todo: Route for retriving all segments.
// @todo: Remove, as it won't be part of the public API.
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

	pathParams := []string{}
	pathParams = append(pathParams, u.Host)
	// exclude first element, as the Path starts with an "/"
	pathParams = append(pathParams, strings.Split(u.Path, string('/'))[1:]...)

	// @todo: Handle query parameters.

	results := app.RulesService.CreateFromListOfSegments(pathParams)
	return results, nil
}

func (app Application) GetMatch(u *url.URL) ([]rule.Rule, error) {
	host := u.Host
	p := strings.Split(u.Path, string('/')) // @todo: validate if the last path param is "/"
	p = p[1:]

	pathParams := []string{}
	pathParams = append(pathParams, host)

	// exclude first element, as the Path starts with an "/"
	pathParams = append(pathParams, p...)

	// STEP 1. Fetch Rules by domain name, as a tree struct.
	rules, _ := app.RulesService.GetRulesByDomain(pathParams[0])

	r := findMatches(rules, pathParams)

	return r, nil
}

// helper fn for reading tree.
func findMatches(tree []rule.Rule, urlPaths []string) []rule.Rule {
	fmt.Println("[DEBUG] Matching: ", urlPaths)
	result := []rule.Rule{}
	i := 0
	for i < len(tree) {
		// hard to read.
		// types defer (true) and type is fixed(1)(true) => true
		if tree[i].Path != urlPaths[0] && tree[i].Type == rule.FType {
			i++
			continue
		}

		if len(tree[i].Children) > 0 && len(urlPaths) > 1 {
			result = append(result, findMatches(tree[i].Children, urlPaths[1:])...)
		}

		if len(urlPaths) == 1 {
			fmt.Println("[DEBUG] Found: ", tree[i].Path)
			result = append(result, tree[i])
		}

		i++
	}
	return result
}

// @todo: add docs
func (app Application) CreateRulesTable() error {
	fmt.Println("[DEBUG] create rules table")

	//@todo: handle ResourceInUseException error.
	return app.RulesService.CreateRulesTable()
}
