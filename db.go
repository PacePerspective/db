package db

import (
	"context"
)

type NoSQLClient interface {
	Read(ctx context.Context, collection, id string) (map[string]interface{}, error)
	Insert(ctx context.Context, collection string, data interface{}) (string, error)
	InsertWithID(ctx context.Context, collection, id string, data interface{}) error
	Where(ctx context.Context, collection, key, value string) ([]map[string]interface{}, error)
}
