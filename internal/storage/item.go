package storage

import (
	"context"
	"strings"

	"github.com/PumpkinSeed/notodo-go/internal/types"
)

const itemStoragePrefix = "item::"

func GetItem(ctx context.Context, key string) (types.Item, error) {
	v, err := get(ctx, itemPrefix(key))
	if err != nil {
		return types.Item{}, err
	}
	if v != nil {
		if item, ok := v.(types.Item); ok {
			return item, nil
		}
	}
	return types.Item{}, ErrUnknownStorageError
}

func GetItems(ctx context.Context) ([]types.Item, error) {
	var items []types.Item
	for key := range storage {
		if keySplit := strings.Split(key, "::"); len(keySplit) > 0 && keySplit[0] == "item" {
			item, err := GetItem(ctx, cleanPrefix(key))
			if err == nil {
				items = append(items, item)
			}
		}
	}
	return items, nil
}

func SetItem(ctx context.Context, key string, value types.Item) {
	set(ctx, itemPrefix(key), value)
}

func itemPrefix(key string) string {
	return itemStoragePrefix + key
}

func cleanPrefix(key string) string {
	return strings.ReplaceAll(key, itemStoragePrefix, "")
}
