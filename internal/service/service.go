package service

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
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

// Returns an instance of *dynamoDB
func createLocalClient() *dynamodb.DynamoDB {
	// Reads aws configuration from following files:
	// ~/.aws/credentials
	// ~/.aws/config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess)
}
