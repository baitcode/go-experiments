package dynamo2

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/vchain-dev/go-workers/config"
)

func getDynamo() *dynamodb.DynamoDB {
	return dynamodb.New(config.AwsSession, config.DynamoDBConfig)
}
