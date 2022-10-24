package app

import (
	"fmt"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
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
func (rs RulesService) CreateFromListOfSegments(segments []string) []rule.Segment {
	result := []rule.Segment{}
	parent := ROOT //rootSegment.Id
	domain := segments[0]

	for _, segment := range segments {
		fmt.Println("[DEBUG] Looping segments >> ", segment)
		// pentru fiecare segment avem doua variante.
		s, error := rs.rulesRepository.GetSegmentByPathAndParent(segment, parent)
		if error != nil {
			fmt.Println("[ERROR] Failed to get by Path and ParentId", error)
		}

		if len(s.Id) > 0 { // 1. Exista deja segmentul
			fmt.Println("[DEBUG] Segment: ", segment, "exists:", s)
		} else { // 2. Nu exista segmentul.
			fmt.Println("[DEBUG] Segment: ", segment, "does not exists:", s)
			re := regexp.MustCompile(`{[a-zA-Z^0-9]*?\}`)

			// if re.Match(segment)

			s.Path = segment
			s.ParentId = parent
			s.Domain = domain
			s.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
			s.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

			if len(re.FindAllString(s.Path, -1)) == 1 {
				fmt.Println("[DEBUG] Segment: ", segment, "is variable type")
				s.Type = rule.VType
			} else {
				fmt.Println("[DEBUG] Segment: ", segment, "is fixed type")
				s.Type = rule.FType
			}

			err := rs.rulesRepository.CreateNode(&s)
			if err != nil {
				fmt.Println("[ERROR] Could not create node element with error: ", err)
			}

			fmt.Println("[DEBUG] Segment: ", s.Path, " has beed created.")
		}

		parent = s.Id
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

// PRIVATE METHODS BELLOW

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
		r := rule.Rule{Id: id, Path: segment.Path, Type: segment.Type, Domain: segment.Domain}
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
