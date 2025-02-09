package mapper

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Mapper[T any] interface {
	MapToDomain(item *dynamodb.GetItemOutput) (T, error)
}
