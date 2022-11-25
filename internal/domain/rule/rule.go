package rule

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/metadata"
)

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
	Domain   string `json:"-"`
	Weight   int8
	Data     metadata.Metadata
	Children []Rule `json:",omitempty"`
}

// Used in persistance layer.
type Segment struct {
	Id        string            // Unique identifier
	ParentId  string            // Parent Id
	Domain    string            // Url domain
	Path      string            // Actual path param
	Data      metadata.Metadata // Metadata information
	CreatedAt string            // Created at
	UpdatedAt string            // Updated at
	Type      SegmentType       // Type of segment: fixed or variable
	Weight    int8              // Calculated by Segment position in URL.
}

// GetKey returns the composite primary key of the movie in a format that can be
// sent to DynamoDB.
func (segment Segment) GetKey() map[string]types.AttributeValue {
	path, err := attributevalue.Marshal(segment.Path)
	if err != nil {
		panic(err)
	}
	parentId, err := attributevalue.Marshal(segment.ParentId)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"Path": path, "ParentID": parentId}
}
