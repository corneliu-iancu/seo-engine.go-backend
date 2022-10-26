package adaptor

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/domain/rule"
	"regexp"
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

	// repository := NewRuleDynamoRepository(Dyna.Db)

	//result, _ := repository.CreateRulesTable()
	//
	//mock.ExpectCreateTable().WillReturns(*result)

	toTestString := "{param1}"

	re := regexp.MustCompile(`{[a-zA-Z^0-9]*?\}`)

	if len(re.FindAllString(toTestString, -1)) == 1 {
		fmt.Println("matches the regex", rule.VType)
	} else {
		fmt.Println("does not match the regex", rule.FType)
	}
}
