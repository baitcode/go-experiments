package dynamo2

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type AWSStructure = map[string]*dynamodb.AttributeValue

type DynamoKey interface {
	AsDynamoKey() AWSStructure
}

type DynamoItem interface {
	AsDynamoItem() AWSStructure
}

type DynamoController interface {
	Query(keyCondition Filter, sortKeyCondition *Filter, filters ...Filter) []AWSStructure
	Scan(filters ...Filter) []AWSStructure
	DeleteItem(key DynamoKey) []AWSStructure
	PutItem(item AWSStructure, conditions ...Filter) []AWSStructure
	GetItem(key DynamoKey) []AWSStructure
}

//type Table interface {
//	Query(filter QueryFilter)
//	Scan(filters ...Filter)
//	DeleteItem(key DynamoKey)
//	PutItem(item DynamoItem, conditionFilters ...Filter)
//}

type Filter interface {
	Apply(builder *expression.Builder)
}

type QueryFilter struct {
	KeyCondition     Filter
	SortKeyCondition *Filter
	Filters          []Filter
}

func NewQueryFilter(hash Filter, sort *Filter, filters ...Filter) *QueryFilter {
	return &QueryFilter{
		hash,
		sort,
		filters,
	}
}

type ScanFilter struct {
	KeyCondition     *Filter
	SortKeyCondition *Filter
	Filters          []Filter
}

func NewScanFilter(hash *Filter, sort *Filter, filters ...Filter) *ScanFilter {
	return &ScanFilter{
		hash,
		sort,
		filters,
	}
}
