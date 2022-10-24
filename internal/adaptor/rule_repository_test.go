package adaptor

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"testing"
)

import (
	_ "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	dynamock "github.com/gusaul/go-dynamock"
)

type MyDynamo struct {
	Db dynamodbiface.DynamoDBAPI
}

var mock *dynamock.DynaMock
var Dyna *MyDynamo

func init() {
	Dyna = new(MyDynamo)
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)
	Dyna.Db = svc
	_, mock = dynamock.New()

}

func TestRuleRepository_CreateRulesTable(t *testing.T) {

	repository := NewRuleDynamoRepository(Dyna.Db)

	result, _ := repository.CreateRulesTable()

	mock.ExpectCreateTable().WillReturns(*result)
}
