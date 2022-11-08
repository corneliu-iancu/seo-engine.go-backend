// @todo:
// add documentation.
// import logger instance.
package adaptor

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	model "github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"github.com/google/uuid"
	"log"
)

var tableName string = "SeoRules"

type RuleDynamoRepository struct {
	db dynamodbiface.DynamoDBAPI
}

func NewRuleDynamoRepository(db dynamodbiface.DynamoDBAPI) RuleDynamoRepository {
	return RuleDynamoRepository{
		db: db,
	}
}

// API METHODS
func (rdr RuleDynamoRepository) GetSegmentByPathAndParent(path string, parent string) (model.Segment, error) {
	svc := rdr.db
	fmt.Println("[DEBUG][repository] Search in db for an entry based on Path and Parent", path, parent)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Path": {
				S: aws.String(path),
			},
			"ParentId": {
				S: aws.String(parent), //@todo: find docs to change parentId to be of number type.
			},
		},
	})

	if err != nil {
		fmt.Println("[FATAL] Got error calling GetItem:", err)
		return model.Segment{}, err
	}

	if len(result.Item) > 0 {
		s := model.Segment{}
		err = dynamodbattribute.UnmarshalMap(result.Item, &s)
		if err != nil {
			fmt.Println("[ERROR] Got error unmarshalling:", err)
			return model.Segment{}, nil
		}
		return s, nil
	}

	return model.Segment{}, nil
}

func (rdr RuleDynamoRepository) GetSegments() ([]model.Segment, error) {
	results := []model.Segment{}
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := rdr.db.Scan(params)
	if err != nil {
		return nil, err
	}

	for _, i := range result.Items {
		segment := model.Segment{}
		err = dynamodbattribute.UnmarshalMap(i, &segment)
		if err != nil {
			log.Fatalf("Got error unmashalling: %s", err)
		}
		results = append(results, segment)
	}

	return results, nil
}

func (rdr RuleDynamoRepository) GetSegmentsByDomainName(domain string) ([]model.Segment, error) {
	r := []model.Segment{}

	filter := expression.Name("Domain").Equal(expression.Value(domain))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		fmt.Println("[ERROR] Got error building expression:", err)
		return r, err
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	result, err := rdr.db.Scan(params)
	if err != nil {
		log.Fatalf("[FATAL] Got error unmashalling: %s", err)
		return nil, err
	}
	// handle error.
	for _, i := range result.Items {
		segment := model.Segment{}
		err = dynamodbattribute.UnmarshalMap(i, &segment)
		if err != nil {
			log.Fatalf("Got error unmashalling: %s", err)
			return nil, err
		}
		r = append(r, segment)
	}

	return r, nil
}

// Persists one segment to the database.
func (rdr RuleDynamoRepository) CreateNode(segment *model.Segment) error {
	svc := rdr.db

	if len(segment.Id) == 0 {
		segment.Id = uuid.New().String()[:8]
	}

	fmt.Println("[DEBUG] New segmentId: ", segment.Id)

	av, err := dynamodbattribute.MarshalMap(segment)
	if err != nil {
		// log.Fatalf("Got error marshalling new movie item: %s", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		// log.Fatalf("Got error calling PutItem: %s", err)
		return err
	}

	return nil
}

// Reads all the tables from the db connection.
func (rdr RuleDynamoRepository) GetTables() error {
	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	fmt.Printf("Tables:\n")

	svc := rdr.db

	for {
		// Get the list of tables
		result, err := svc.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return nil
		}

		for _, n := range result.TableNames {
			fmt.Println(*n)
		}

		// assign the last read tablename as the start for our next call to the ListTables function
		// the maximum number of table names returned in a call is 100 (default), which requires us to make
		// multiple calls to the ListTables function to retrieve all table names
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}

	return nil
}

// Creates the rules table schema.
func (rdr RuleDynamoRepository) CreateRulesTable() (*dynamodb.CreateTableOutput, error) {
	svc := rdr.db

	// Only Primary Key and Sort Key are required to create the dynamo table.
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Path"), // Rename me to UUID.
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ParentId"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Path"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("ParentId"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	result, err := svc.CreateTable(input)
	if err != nil {
		// log.Fatalf("Got error calling CreateTable: %s", err)
		// fmt.Println("[ERROR]Got error calling CreateTable: %s", err)
		return result, err
	}

	fmt.Println("Created the table", tableName)
	return result, nil
}

// @todo: remove me
func (rdr RuleDynamoRepository) GetRootSegment(segment model.Segment) model.Segment {
	// retrieve elements with ParentId == 0 and Path = segment.Path
	filter := expression.Name("ParentId").Equal(expression.Value("0")).And(expression.Name("Path").Equal(expression.Value(segment.Path)))

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		fmt.Println("[ERROR] Got error building expression:", err)
		return model.Segment{}
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	}

	result, err := rdr.db.Scan(params)

	if err != nil {
		fmt.Println("[ERROR] Query API call failed:", err)
		return model.Segment{}
	}

	if len(result.Items) > 0 {
		s := model.Segment{}
		err = dynamodbattribute.UnmarshalMap(result.Items[0], &s)
		if err != nil {
			fmt.Println("[ERROR] Got error unmarshalling:", err)
			return model.Segment{}
		}
		return s
	}

	return model.Segment{}
	//for _, item := range result.Items {
	//	segment := model.Segment{}
	//
	//	err = dynamodbattribute.UnmarshalMap(item, &segment)
	//	if err != nil {
	//		fmt.Println("[ERROR] Got error unmarshalling:", err)
	//	}
	//
	//	results = append(results, segment)
	//}

	//return results
}
