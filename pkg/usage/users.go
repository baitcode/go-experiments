package usage

import (
	db "github.com/vchain-dev/go-workers/pkg/awstools/dynamo2"
	"github.com/aws/aws-sdk-go/aws"
)

func userFactory(structure db.AWSStructure) interface{} {
	return User{
		Id: *structure["Id"].S,
		Name: *structure["Name"].S,
	}
}

type User struct {
	Id   string
	Name string
}

func (u User) AsDynamoItem() db.AWSStructure {
	return db.AWSStructure{
		"Id": {S: aws.String(u.Id)},
		"Name": {S: aws.String(u.Name)},
	}
}

func (u User) AsKey() db.AWSStructure {
	return db.AWSStructure{
		"Id": {S: aws.String(u.Id)},
		"Name": {S: aws.String(u.Name)},
	}
}

type Users struct {
	*db.DBMembers
}

func (u Users) DeleteItem(key db.DynamoKey) {
	u.FromDB(u.Controller.DeleteItem(key))
}

func (u Users) Scan(filters ...db.Filter) {
	u.FromDB(u.Controller.Scan(filters...))
}

func (u Users) Query(f db.QueryFilter) {
	u.FromDB(u.Controller.Query(f.KeyCondition, f.SortKeyCondition, f.Filters...))
}

func (u Users) PutItem(item db.DynamoItem, conditionFilters ...db.Filter) {
	u.FromDB(u.Controller.PutItem(item.AsDynamoItem(), conditionFilters...))
}

func (u Users) GetItem(key db.DynamoKey) {
	u.FromDB(u.Controller.GetItem(key))
}

func (u Users) One() *User {
	for member := range u.DBMembers.Iterator() {
		user, ok := member.(User)
		if ok {
			return &user
		} else {
			// TODO: log
		}
	}
	return nil
}

func (u Users) Iterator() <-chan User {
	ch := make(chan User)

	go func() {
		for member := range u.DBMembers.Iterator() {
			user, ok := member.(User)
			if ok {
				ch <- user
			} else {
				// TODO: log
			}
		}
	}()

	return ch
}

func UsersDAO() Users {
	return Users{
		DBMembers: db.BuildMembers("Users", userFactory),
	}
}