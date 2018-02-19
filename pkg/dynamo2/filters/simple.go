package filters

import "github.com/aws/aws-sdk-go/service/dynamodb/expression"

type SimpleFilter struct {
	Key   string
	Value interface{}
}

func (f *SimpleFilter) Apply(builder *expression.Builder) {
	*builder = builder.WithFilter(
		expression.
			Name(f.Key).
			Equal(expression.Value(f.Value)),
	)
}

func FieldEquals(key string, value interface{}) *SimpleFilter {
	return &SimpleFilter{
		key,
		value,
	}
}
