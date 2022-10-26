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

	// r := []rule.Rule{}

	host := u.Host
	p := strings.Split(u.Path, string('/')) // @todo: validate if the last path param is "/"
	p = p[1:]

	pathParams := []string{}
	pathParams = append(pathParams, host)

	// exclude first element, as the Path starts with an "/"
	pathParams = append(pathParams, p...)

	// STEP 1. Fetch Rules by domain name, as a tree struct.
	rules, _ := app.RulesService.GetRulesByDomain(pathParams[0])
	// STEP 2. Get Rules as tree.

	// findMatches(pathParams, rules, 0, &r)

	fmt.Println(rules)
	// matching algorithm.

	// Step 3. Iterate over rules, and check which nodes match

	// Step 4. Return as a list of Nodes.

	// Useful tips.
	// Only leaf nodes will register metadata for a needle uri.

	return rules, nil
}

// helper fn for reading tree.

func findMatches(urlPaths []string, tree []rule.Rule, treeLevel int, response *[]rule.Rule) {
	// fmt.Println("[DEBUG] Tree level: ", tree)

	for _, leaf := range tree {

		/* if leaf.Path == urlPaths[treeLevel] {
			*response = append(*response, rule.Rule{
				Id:       leaf.Id,
				Path:     leaf.Path,
				Type:     leaf.Type,
				Domain:   leaf.Domain,
				Children: []rule.Rule{},
			})
		} */

		if len(leaf.Children) > 0 {
			findMatches(urlPaths, leaf.Children, treeLevel+1, response)
		}

		fmt.Println("[DEBUG] Reading: ", leaf.Path, treeLevel, urlPaths[treeLevel])
	}
}

// @todo: add docs
func (app Application) CreateRulesTable() error {
	fmt.Println("[DEBUG] create rules table")

	//@todo: handle ResourceInUseException error.
	return app.RulesService.CreateRulesTable()
}
