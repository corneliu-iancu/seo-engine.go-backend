// @todo:
// add documentation.
// import logger instance.
package adaptor

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	model "github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"time"

	"log"
)

var tableName string = "SeoRules"

type RuleDynamoRepository struct {
	db dynamodb.Client
}

func NewRuleDynamoRepository(db dynamodb.Client) RuleDynamoRepository {
	return RuleDynamoRepository{
		db: db,
	}
}

// API METHODS
// Uses dynamodb.GetItem API Method.
func (rdr RuleDynamoRepository) GetSegmentByPathAndParent(path string, parent string) (model.Segment, error) {
	segment := model.Segment{Path: path, ParentId: parent}
	response, err := rdr.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       segment.GetKey(),
	})

	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", path, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &segment)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}

	return segment, err
}

// Retrieves all Segments.
// @todo: Check if it is useful.
// Uses dynamodb.Scan API Method.
func (rdr RuleDynamoRepository) GetSegments() ([]model.Segment, error) {
	segments := []model.Segment{}

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	response, err := rdr.db.Scan(context.TODO(), params)
	if err != nil {
		log.Printf("Couldn't scan for segments. Here's why: %v\n", err)
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &segments)

	return segments, err
}

// @todo: Check if it is useful.
// Uses dynamodb.Scan API Method
func (rdr RuleDynamoRepository) GetSegmentsByDomainName(domain string) ([]model.Segment, error) {
	segments := []model.Segment{}

	filter := expression.Name("Domain").Equal(expression.Value(domain))
	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		log.Printf("Couldn't build epxression for query. Here's why: %v\n", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	response, err := rdr.db.Scan(context.TODO(), params)
	if err != nil {
		log.Printf("Couldn't scan for segments by domain name %v. Here's why: %v\n",
			domain, err)
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &segments)
	if err != nil {
		log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
	}

	return segments, err
}

// ###################################################################################################################//
// Check if our table exists in the connection.
func (rdr RuleDynamoRepository) TableExists() (bool, error) {
	exists := true
	_, err := rdr.db.DescribeTable(
		context.TODO(), &dynamodb.DescribeTableInput{TableName: aws.String(tableName)},
	)
	if err != nil {
		var notFoundEx *types.ResourceNotFoundException
		if errors.As(err, &notFoundEx) {
			log.Printf("Table %v does not exist.\n", tableName)
			err = nil
		} else {
			log.Printf("Couldn't determine existence of table %v. Here's why: %v\n", tableName, err)
		}
		exists = false
	}
	return exists, err
}

// @todo: implement me
// Persists one segment to the database.
func (rdr RuleDynamoRepository) CreateSegment(segment *model.Segment) error {
	//svc := rdr.db
	//
	//if len(segment.Id) == 0 {
	//	segment.Id = uuid.New().String()[:8]
	//}
	//
	//fmt.Println("[DEBUG] New segmentId: ", segment.Id)
	//
	//av, err := dynamodbattribute.MarshalMap(segment)
	//if err != nil {
	//	// log.Fatalf("Got error marshalling new movie item: %s", err)
	//	return err
	//}
	//
	//input := &dynamodb.PutItemInput{
	//	Item:      av,
	//	TableName: aws.String(tableName),
	//}
	//
	//_, err = svc.PutItem(input)
	//if err != nil {
	//	// log.Fatalf("Got error calling PutItem: %s", err)
	//	return err
	//}
	//
	return nil
}

// @todo: implement me
func (rdr RuleDynamoRepository) UpdateSegment(segment *model.Segment) {

}

// Creates the rules table schema.
func (rdr RuleDynamoRepository) CreateRulesTable() (*types.TableDescription, error) {
	var tableDesc *types.TableDescription

	table, err := rdr.db.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{{
			AttributeName: aws.String("ParentId"),
			AttributeType: types.ScalarAttributeTypeN,
		}, {
			AttributeName: aws.String("Path"),
			AttributeType: types.ScalarAttributeTypeS,
		}},
		KeySchema: []types.KeySchemaElement{{
			AttributeName: aws.String("ParentId"),
			KeyType:       types.KeyTypeHash,
		}, {
			AttributeName: aws.String("Path"),
			KeyType:       types.KeyTypeRange,
		}},
		TableName: aws.String(tableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})

	if err != nil {
		log.Printf("Couldn't create table %v. Here's why: %v\n", tableName, err)
	} else {
		var waiter = dynamodb.NewTableExistsWaiter(&rdr.db)
		err = waiter.Wait(context.TODO(), &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName)}, 5*time.Minute)
		if err != nil {
			log.Printf("Wait for table exists failed. Here's why: %v\n", err)
		}
		tableDesc = table.TableDescription
	}
	return tableDesc, err

	// Only Primary Key and Sort Key are required to create the dynamo table.
	//input := &dynamodb.CreateTableInput{
	//	AttributeDefinitions: []*dynamodb.AttributeDefinition{
	//		{
	//			AttributeName: aws.String("Path"), // Rename me to UUID.
	//			AttributeType: aws.String("S"),
	//		},
	//		{
	//			AttributeName: aws.String("ParentId"),
	//			AttributeType: aws.String("S"),
	//		},
	//	},
	//	KeySchema: []*dynamodb.KeySchemaElement{
	//		{
	//			AttributeName: aws.String("Path"),
	//			KeyType:       aws.String("HASH"),
	//		},
	//		{
	//			AttributeName: aws.String("ParentId"),
	//			KeyType:       aws.String("RANGE"),
	//		},
	//	},
	//	ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
	//		ReadCapacityUnits:  aws.Int64(10),
	//		WriteCapacityUnits: aws.Int64(10),
	//	},
	//	TableName: aws.String(tableName),
	//}
	//
	//result, err := svc.CreateTable(input)
	//if err != nil {
	//	// log.Fatalf("Got error calling CreateTable: %s", err)
	//	// fmt.Println("[ERROR]Got error calling CreateTable: %s", err)
	//	return result, err
	//}
	//
	//fmt.Println("Created the table", tableName)
	//return result, nil
}
