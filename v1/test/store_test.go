package test

import (
	"testing"
	"todo-app/v1"
)

func TestAddItem(t *testing.T) {
	store := v1.NewStore()
	id := store.Add("test item", nil)

	if item, exists := store.Get(id); !exists || item.Title != "test item" {
		t.Fatalf("expected to find item with title 'test item'")
	}
}
