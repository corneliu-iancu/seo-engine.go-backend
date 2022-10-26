package rule

type SegmentType int8

const (
	FType SegmentType = 1 << iota
	VType
)

// From a Rule structure, may result one or many URIs.
type Rule struct {
	Id       string
	ParentId string
	Path     string
	Type     SegmentType
	Domain   string
	Children []Rule `json:",omitempty"`
}

// Used in persistance layer.
type Segment struct {
	Id        string      // Unique identifier
	ParentId  string      // Parent Id
	Domain    string      // Url domain
	Path      string      // Actual path param
	CreatedAt string      // Created at
	UpdatedAt string      // Updated at
	Type      SegmentType // Type of segment: fixed or variable
}
