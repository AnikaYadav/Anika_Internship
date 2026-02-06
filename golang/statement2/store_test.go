package main

import "testing"

func TestAddAndCheckItem(t *testing.T) {
	items := make(map[string]bool)

	addItem(items, "hello")

	if !checkItem(items, "hello") {
		t.Errorf("expected 'hello' to exist")
	}
}

func TestRemoveItem(t *testing.T) {
	items := make(map[string]bool)
	items["hello"] = true

	removed := removeItem(items, "hello")

	if !removed {
		t.Errorf("expected item to be removed")
	}

	if checkItem(items, "hello") {
		t.Errorf("expected 'hello' to be removed")
	}
}

func TestRemoveNonExistingItem(t *testing.T) {
	items := make(map[string]bool)

	removed := removeItem(items, "missing")

	if removed {
		t.Errorf("did not expect removal of non-existing item")
	}
}

func TestListItems(t *testing.T) {
	items := make(map[string]bool)
	items["a"] = true
	items["b"] = true

	list := listItems(items)

	if len(list) != 2 {
		t.Errorf("expected 2 items, got %d", len(list))
	}
}
