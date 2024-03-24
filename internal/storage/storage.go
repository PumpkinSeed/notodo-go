package storage

import "context"

var storage map[string]interface{}

func get(_ context.Context, key string) (interface{}, error) {
	if storage == nil {
		return nil, ErrNotFound
	}
	if v, ok := storage[key]; ok {
		return v, nil
	}
	return nil, ErrNotFound
}

func set(_ context.Context, key string, value interface{}) {
	if storage == nil {
		storage = make(map[string]interface{})
	}
	storage[key] = value
}

func clean(_ context.Context) {
	storage = make(map[string]interface{})
}
