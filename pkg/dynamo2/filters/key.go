package filters

import "github.com/aws/aws-sdk-go/service/dynamodb/expression"

type KeyFilter struct {
	Key   string
	Value interface{}
}

func (f *KeyFilter) Apply(builder *expression.Builder) {
	*builder = builder.WithKeyCondition(
		expression.
			Key(f.Key).
			Equal(expression.Value(f.Value)),
	)
}

func KeyEquals(key string, value interface{}) *KeyFilter {
	return &KeyFilter{
		key,
		value,
	}
}