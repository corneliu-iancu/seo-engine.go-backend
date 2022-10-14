// @todo:
// add documentation.
// import logger instance.
package adaptor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"

	"fmt"
)

var TableName string = "SeoRules"

type RuleDynamoRepository struct {
	db *dynamodb.DynamoDB
}

type rule struct {
	Id        string
	ParentId  int
	Path      string
	CreatedAt string
	UpdatedAt string
}

func NewRuleDynamoRepository(db *dynamodb.DynamoDB) RuleDynamoRepository {
	return RuleDynamoRepository{
		db: db,
	}
}

// Finds one rule based on a string.
func (rdr RuleDynamoRepository) FindOne() error {
	return nil
}

// Persists one rule to the database.
func (rdr RuleDynamoRepository) Create() error {
	svc := rdr.db

	item := rule{
		Id:        uuid.New().String(),
		Path:      "https://www.superbet.ro",
		CreatedAt: "14-10-2022",
		UpdatedAt: "14-10-2022",
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		// log.Fatalf("Got error marshalling new movie item: %s", err)
		return err
	}

	tableName := TableName // @todo: refactor

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return err
		// log.Fatalf("Got error calling PutItem: %s", err)
	}

	return nil
}

// Finds one rule based on a string.
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

func (rdr RuleDynamoRepository) CreateRulesTable() error {
	svc := rdr.db
	tableName := TableName // @todo: refactor

	// Only Primary Key and Sort Key are required to create the dynamo table.
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Path"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ParentId"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Path"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("UId"),
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

	_, err := svc.CreateTable(input)
	if err != nil {
		// log.Fatalf("Got error calling CreateTable: %s", err)
		// fmt.Println("[ERROR]Got error calling CreateTable: %s", err)
		return err
	}

	fmt.Println("Created the table", tableName)
	return nil
}
