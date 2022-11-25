package factory

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"log"

	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbiface"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/adaptor"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/app"
	"go.uber.org/zap"
)

// Returns an application instance.
// > RulesService
// > MetadataService
func NewApplication(logger *zap.Logger) app.BusinessLogicImpl {
	fmt.Println("[DEBUG] create new application")

	// Create local dynamodb instance
	svc := createLocalClient()

	rulesRepository := adaptor.NewRuleDynamoRepository(*svc)

	return app.BusinessLogicImpl{
		Logger:       logger,
		RulesService: app.NewRulesService(rulesRepository),
	}
}

// Returns an instance of *dynamoDB type
func createLocalClient() *dynamodb.Client {
	// ====================================================
	// REMOTE CONNECTION.
	// ====================================================
	// Reads aws configuration from following files:
	// ~/.aws/credentials
	// ~/.aws/config

	/* sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
	 	})) */

	// ====================================================
	// LOCAL CONNECTION.
	// ====================================================

	// cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-central-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:8000"}, nil
			})),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	// Using the Config value, create the DynamoDB client
	return dynamodb.NewFromConfig(cfg)

	//sess, err := session.NewSession(&aws.Config{
	//	Region:   aws.String("eu-central-1"),
	//	Endpoint: aws.String("http://localhost:8000"),
	//})
	//
	//if err != nil {
	//	log.Println(err)
	//	// return
	//}
	//svc := dynamodb.New(sess)
	//
	//return dynamodbiface.DynamoDBAPI(svc)
}
