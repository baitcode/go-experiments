package filters

import "github.com/aws/aws-sdk-go/service/dynamodb/expression"

type RecordDoesNotExistOrFieldLessThenFilter struct {
	HashKey    string
	Value      interface{}
	OtherField string
	OtherValue interface{}
}

type filter = RecordDoesNotExistOrFieldLessThenFilter

func (f *filter) Apply(builder *expression.Builder) {
	*builder = builder.WithCondition(
		expression.Or(
			expression.AttributeNotExists(expression.Name(f.HashKey)),
			expression.And(
				expression.Name(f.HashKey).Equal(expression.Value(f.Value)),
				expression.Name(f.OtherField).LessThan(expression.Value(f.OtherValue)),
			),
		),
	)
}

func FieldNotExistsOrFieldLessThanValue(key string, value interface{}, otherField string, lessThan interface{}) *filter {
	return &filter{
		key,
		value,
		otherField,
		lessThan,
	}
}
