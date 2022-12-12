// @todo:
// add documentation.
// import logger instance.
package adaptor

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	model "github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/google/uuid"
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

// ################################################################ //
// ########################## API METHODS ######################### //
// ########################## GET METHODS ######################### //
// ################################################################ //
func (rdr RuleDynamoRepository) GetBySegmentId(segmentId string) (*model.Segment, error) {
	segments := []model.Segment{}

	filter := expression.Name("Id").Equal(expression.Value(segmentId))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		log.Printf("Couldn't build expression for query. Here's why: %v\n", err)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	response, err := rdr.db.Scan(context.TODO(), params)
	if err != nil {
		log.Printf("Couldn't scan for segments by segmentId %v. Here's why: %v\n",
			segmentId, err)
	}

	fmt.Println(len(response.Items))

	if len(response.Items) != 1 {
		log.Printf("Couldn't find a match for segment id %v. Here's why: %v\n", segmentId, err)
	}

	err = attributevalue.UnmarshalListOfMaps(response.Items, &segments)
	if err != nil {
		log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
	}

	// return segments, err

	return &segments[0], nil
}

func (rdr RuleDynamoRepository) GetSegmentByPathAndParent(path string, parent string) (model.Segment, error) {
	segment := model.Segment{Path: path, ParentId: parent}
	// fmt.Println("Key:: ", segment.GetKey()["Parent"])
	response, err := rdr.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       segment.GetKey(),
	})

	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v, %v\n", path, parent, err)

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

	filter := expression.Name("Domain").Equal(expression.Value(domain)).And(expression.Name("ParentId").Equal(expression.Value("ROOT")))

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

// Persists one segment to the database.
func (rdr RuleDynamoRepository) CreateSegment(segment *model.Segment) error {
	// generate Id.
	segment.Id = uuid.New().String()[:8]

	item, err := attributevalue.MarshalMap(segment)
	if err != nil {
		panic(err)
	}
	
	_, err = rdr.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName), Item: item,
	})
	if err != nil {
		log.Printf("Couldn't add item to table. Here's why: %v\n", err)
	}

	return err
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
			AttributeType: types.ScalarAttributeTypeS,
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
