package app

import (
	"fmt"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/helper"
	"net/url"
	"regexp"
	"time"
)

const ROOT = "ROOT_NODE"

// @todo: !! move to service package.
// @todo: add docs
type RulesService struct {
	rulesRepository adaptor.RuleDynamoRepository
}

// Creates a new Rules Service instance.
// Receives an repository as parameter.
func NewRulesService(rulesRepository adaptor.RuleDynamoRepository) RulesService {
	return RulesService{
		rulesRepository: rulesRepository,
	}
}

// Creates segments and persists them to db though repository layer.
func (rs RulesService) CreateFromListOfStrings(segments []string) []rule.Segment {
	result := []rule.Segment{}
	parentId := ROOT
	weight := 0
	domain := segments[0]

	for _, segment := range segments {
		fmt.Println("[DEBUG] Looping segments >> ", segment)
		// pentru fiecare segment avem doua variante.
		s, error := rs.rulesRepository.GetSegmentByPathAndParent(segment, parentId)

		if error != nil {
			fmt.Println("[ERROR] Failed to get by Path and ParentId", error)
		}

		if len(s.Id) == 0 { // 2. The segment does not exists.
			re := regexp.MustCompile(`{[a-zA-Z^0-9]*?\}`)

			if len(re.FindAllString(segment, -1)) == 1 {
				fmt.Println("[DEBUG] Variable type parameter: ", s.Path)
				s.Type = rule.VType
			} else {
				fmt.Println("[DEBUG] Fixed type parameter: ", s.Path)
				s.Type = rule.FType
			}

			s.Path = segment
			s.ParentId = parentId
			s.Domain = domain
			s.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			s.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			s.Weight = int8(weight)

			err := rs.rulesRepository.CreateNode(&s)
			if err != nil {
				fmt.Println("[ERROR] Could not create node element with error: ", err)
			}

			fmt.Println("[DEBUG] Segment: ", s, " has beed created")
		} else {
			fmt.Println("[DEBUG] Segment: ", s, " already exists.")
		}

		weight += int(s.Type)
		parentId = s.Id
		result = append(result, s)
	}
	return result
}

// Creates the rules tables. Returns error if the table already exists.
func (rs RulesService) CreateRulesTable() error {
	result, err := rs.rulesRepository.CreateRulesTable()
	fmt.Println(result)
	return err
}

// Returns all db segments.
func (rs RulesService) GetSegments() ([]rule.Segment, error) {
	results, err := rs.rulesRepository.GetSegments()

	return results, err
}

func (rs RulesService) GetRules() ([]rule.Rule, error) {
	segments, err := rs.rulesRepository.GetSegments()

	if err != nil {
		fmt.Println("[ERROR] failed to get all segments")
	}

	relations := relationMap(segments)
	rules := buildRules(relations[ROOT], relations, segments)

	return rules, nil
}

// Returns all rules by using the domain property as input.
func (rs RulesService) GetRulesByDomain(domain string) ([]rule.Rule, error) {

	segments, err := rs.rulesRepository.GetSegmentsByDomainName(domain)

	if err != nil {
		fmt.Println("[ERROR] failed to get all segments")
	}

	relations := relationMap(segments)

	rules := buildRules(relations[ROOT], relations, segments)

	return rules, nil
}

func (rs RulesService) GetMatch(u *url.URL) ([]rule.Rule, error) {
	pathParams := helper.GetURIAsSlice(u)
	fmt.Println("[DEBUG] Get match rules for ", pathParams)
	// costly operation. @todo: refactor to query for each param.
	// @todo: verify times for both aproches.
	rules, error := rs.GetRulesByDomain(pathParams[0])
	if error != nil {
		return nil, error
	}

	r := findMatches(rules, pathParams)
	return r, nil
}

func relationMap(segments []rule.Segment) map[string][]string {
	relations := make(map[string][]string)

	for _, segment := range segments {
		child, parent := segment.Id, segment.ParentId
		relations[parent] = append(relations[parent], child)
	}
	return relations
}

func buildRules(roots []string, relations map[string][]string, segments []rule.Segment) []rule.Rule {
	rules := make([]rule.Rule, len(roots))

	for i, id := range roots {

		segment := findSegment(id, segments)

		r := rule.Rule{Id: id, Path: segment.Path, Type: segment.Type, Domain: segment.Domain, ParentId: segment.ParentId, Weight: segment.Weight}

		if childIDs, ok := relations[id]; ok { // build children
			r.Children = buildRules(childIDs, relations, segments)
		}

		rules[i] = r
	}
	return rules
}

func findSegment(id string, segments []rule.Segment) rule.Segment {
	for _, segment := range segments {
		if segment.Id == id {
			return segment
		}
	}
	return rule.Segment{}
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
