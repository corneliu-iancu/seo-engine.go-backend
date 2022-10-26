package factory

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"log"
)

// Returns an application instance.
// > RulesService
// > MetadataService
func NewApplication() app.Application {
	fmt.Println("[DEBUG] create new application")

	// Create local dynamodb instance
	svc := createLocalClient()

	rulesRepository := adaptor.NewRuleDynamoRepository(svc)

	return app.Application{
		RulesService: app.NewRulesService(rulesRepository),
	}
}

// Returns an instance of *dynamoDB type
func createLocalClient() dynamodbiface.DynamoDBAPI {
	// Reads aws configuration from following files:
	// ~/.aws/credentials
	// ~/.aws/config

	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	 SharedConfigState: session.SharedConfigEnable,
	// }))

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000")})

	if err != nil {
		log.Println(err)
		// return
	}
	svc := dynamodb.New(sess)

	return dynamodbiface.DynamoDBAPI(svc)
}
