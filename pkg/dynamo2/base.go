package dynamo2

import (
	"github.com/vchain-dev/go-workers/config"
)

type MemberFactory = func(structure AWSStructure) interface{}

type DBMembers struct {
	instances  []interface{}
	Controller DynamoController
	factory    MemberFactory
}

func (members *DBMembers) FromDB(input []AWSStructure) {
	members.instances = make([]interface{}, 0)

	for idx := range input {
		value := members.factory(input[idx])
		if value != nil {
			members.instances = append(members.instances, value)
		}
	}
}

func (members *DBMembers) Iterator() <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		for _, val := range members.instances {
			ch <- val
		}
	}()

	return ch
}

func NewDBMembers(controller DynamoController, factory MemberFactory) *DBMembers {
	return &DBMembers{
		instances:  make([]interface{}, 0),
		Controller: controller,
		factory:    factory,
	}
}

func BuildMembers(tableName string, factory MemberFactory) *DBMembers {
	if config.TestRun {
		return NewDBMembers(
			NewDynamoControllerMock(tableName),
			factory,
		)
	}
	return NewDBMembers(
		NewDynamoController(tableName),
		factory,
	)
}
