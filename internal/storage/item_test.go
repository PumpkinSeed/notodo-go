package storage

import (
	"context"
	"testing"

	"github.com/PumpkinSeed/notodo-go/internal/types"
)

func TestGetItems(t *testing.T) {
	// Arrange
	ctx := context.Background()
	clean(ctx)

	SetItem(ctx, "test2", types.Item{
		Name:        "test2",
		Description: "Test 2",
	})
	SetItem(ctx, "test3", types.Item{
		Name:        "test3",
		Description: "Test 3",
	})
	set(ctx, "test", "not_an_item")

	// Act
	items, err := GetItems(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if len(items) != 2 {
		t.Errorf("Items len should be 2, instead of %d", len(items))
	}
}

func TestGetItem(t *testing.T) {
	// Arrange
	ctx := context.Background()
	clean(ctx)

	SetItem(ctx, "test2", types.Item{
		Name:        "test2",
		Description: "Test 2",
	})

	// Act
	item, err := GetItem(ctx, "test2")
	if err != nil {
		t.Fatal(err)
	}

	// Assert
	if item.Name != "test2" {
		t.Errorf("Name should be 'test2', instead of %s", item.Name)
	}
}
