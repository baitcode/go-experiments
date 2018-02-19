package dynamo2

type MockController struct {
	tableName string
}

func (d MockController) Query(keyCondition Filter, sortKeyCondition *Filter, filters ...Filter) []AWSStructure {
	return nil
}

func (d MockController) Scan(filters ...Filter) []AWSStructure {
	return nil
}

func (d MockController) DeleteItem(key DynamoKey) []AWSStructure {
	return nil
}

func (d MockController) PutItem(item AWSStructure, conditions ...Filter) []AWSStructure {
	return nil
}

func (d MockController) GetItem(key DynamoKey) []AWSStructure {
	return nil
}

func NewDynamoControllerMock(tableName string) MockController {
	return MockController{tableName}
}
