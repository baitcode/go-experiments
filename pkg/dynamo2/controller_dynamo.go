package dynamo2

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DBController struct {
	tableName string
}

func (d DBController) Query(keyCondition Filter, sortKeyCondition *Filter, filters ...Filter) []AWSStructure {

	builder := expression.NewBuilder()

	keyCondition.Apply(&builder)
	if sortKeyCondition != nil {
		(*sortKeyCondition).Apply(&builder)
	}

	for idx := range filters {
		filters[idx].Apply(&builder)
	}

	builtExpression, err := builder.Build()

	if err != nil {
		panic("pizda")
	}

	input := dynamodb.QueryInput{}
	input.
		SetTableName(d.tableName).
		SetKeyConditionExpression(*builtExpression.KeyCondition()).
		SetFilterExpression(*builtExpression.Filter()).
		SetExpressionAttributeNames(builtExpression.Names()).
		SetExpressionAttributeValues(builtExpression.Values())

	svc := getDynamo()
	result, err := svc.Query(&input)
	if err != nil {
		panic("pizda")
	}

	return result.Items
}

func (d DBController) Scan(filters ...Filter) []AWSStructure {
	builder := expression.NewBuilder()

	for idx := range filters {
		filters[idx].Apply(&builder)
	}

	builtExpression, err := builder.Build()

	if err != nil {
		panic("pizda")
	}

	input := dynamodb.ScanInput{}
	input.
		SetTableName(d.tableName).
		SetFilterExpression(*builtExpression.Filter()).
		SetExpressionAttributeNames(builtExpression.Names()).
		SetExpressionAttributeValues(builtExpression.Values())

	svc := getDynamo()
	result, err := svc.Scan(&input)
	if err != nil {
		panic("pizda")
	}

	return result.Items
}

func (d DBController) DeleteItem(key DynamoKey) []AWSStructure {
	input := dynamodb.DeleteItemInput{}
	input.
		SetTableName(d.tableName).
		SetKey(key.AsDynamoKey())

	svc := getDynamo()
	result, err := svc.DeleteItem(&input)
	if err != nil {
		panic("pizda")
	}

	return []AWSStructure{result.Attributes}
}

func (d DBController) PutItem(item AWSStructure, conditions ...Filter) []AWSStructure {
	builder := expression.NewBuilder()

	for idx := range conditions {
		conditions[idx].Apply(&builder)
	}

	builtExpression, err := builder.Build()

	input := dynamodb.PutItemInput{}
	input.
		SetTableName(d.tableName).
		SetConditionExpression(*builtExpression.Condition()).
		SetExpressionAttributeNames(builtExpression.Names()).
		SetExpressionAttributeValues(builtExpression.Values()).
		SetItem(item)

	svc := getDynamo()
	result, err := svc.PutItem(&input)
	if err != nil {
		panic("pizda")
	}

	return []AWSStructure{result.Attributes}
}

func (d DBController) GetItem(key DynamoKey) []AWSStructure {
	input := dynamodb.GetItemInput{}
	input.
		SetTableName(d.tableName).
		SetKey(key.AsDynamoKey())

	svc := getDynamo()
	result, err := svc.GetItem(&input)
	if err != nil {
		panic("pizda")
	}

	return []AWSStructure{result.Item}
}

func NewDynamoController(tableName string) DBController {
	return DBController{tableName}
}
