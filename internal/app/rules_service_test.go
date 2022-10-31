package app

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/gusaul/go-dynamock"
	"log"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

type MyDynamo struct {
	Db dynamodbiface.DynamoDBAPI
}

var mock *dynamock.DynaMock
var Dyna *MyDynamo

// local dynamodb connection.
func init() {
	Dyna = new(MyDynamo)
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
	}

	svc := dynamodb.New(sess)

	Dyna.Db = dynamodbiface.DynamoDBAPI(svc)
	_, mock = dynamock.New()
}

func TestRulesService_AddRule(t *testing.T) {
	log.Println("implement add one rule.")
}

/**
 * Test to add @1500 rules records to the database.
 */
func TestRulesService_AddRulesSet(t *testing.T) {
	f, err := os.Open("./../../data/URIs.csv")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	ruleService := NewRulesService(adaptor.NewRuleDynamoRepository(Dyna.Db))

	for i, uri := range data {
		if i < 1500 {
			u, err := url.Parse(strings.TrimSpace(uri[0]))
			if err != nil {
				// fmt.Println("[ERROR] Got parse error: ", err)
			} else {
				// Extracted from `journey.go` file.
				pathParams := []string{}
				pathParams = append(pathParams, u.Host)
				pathParams = append(pathParams, strings.Split(u.Path, string('/'))[1:]...)

				segmentsList := ruleService.CreateFromListOfSegments(pathParams)

				fmt.Println(segmentsList)
			}
		}
	}
}

/**
 * Test to retrieve rules for a given URI.
 * test uri: https://superbet.ro/pariuri-sportive/fotbal
 */
func TestRulesService_GetMatch(t *testing.T) {
	u, err := url.Parse("https://superbet.ro/pariuri-sportive/fotbal/")
	if err != nil {
		panic(err)
	}

	// init rules service.
	ruleService := NewRulesService(adaptor.NewRuleDynamoRepository(Dyna.Db))

	firstDate := time.Now() // log time:
	r := ruleService.GetMatch(u)
	secondDate := time.Now() // log time:
	difference := secondDate.Sub(firstDate)
	fmt.Println("[TEST] Get Match took: ", difference)

	// expect to match certain node.
	fmt.Println("Found results: ", len(r))
}
