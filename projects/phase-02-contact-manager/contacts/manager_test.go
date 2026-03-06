package contacts

import (
	"errors"
	"testing"
)

func TestAddAndList(t *testing.T) {
	m := NewManager()

	c1, err := m.Add("Alice", "alice@example.com", "+94000000001")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	c2, err := m.Add("Bob", "bob@example.com", "+94000000002")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if c1.ID != 1 || c2.ID != 2 {
		t.Fatalf("expected sequential IDs 1 and 2, got %d and %d", c1.ID, c2.ID)
	}

	list := m.List()
	if len(list) != 2 {
		t.Fatalf("expected 2 contacts, got %d", len(list))
	}
	if list[0].Name != "Alice" || list[1].Name != "Bob" {
		t.Fatalf("unexpected list order/content: %+v", list)
	}
}

func TestAddValidation(t *testing.T) {
	m := NewManager()
	_, err := m.Add("", "alice@example.com", "123")
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("expected ErrInvalidInput, got %v", err)
	}
}

func TestUpdate(t *testing.T) {
	m := NewManager()
	created, err := m.Add("Alice", "alice@example.com", "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	updated, err := m.Update(created.ID, "Alicia", "alicia@example.com", "999")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.Name != "Alicia" || updated.Email != "alicia@example.com" || updated.Phone != "999" {
		t.Fatalf("unexpected updated value: %+v", updated)
	}
}

func TestUpdateNotFound(t *testing.T) {
	m := NewManager()
	_, err := m.Update(99, "A", "a@example.com", "123")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDelete(t *testing.T) {
	m := NewManager()
	c, err := m.Add("Alice", "alice@example.com", "123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if err := m.Delete(c.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if _, ok := m.Get(c.ID); ok {
		t.Fatalf("contact should have been removed")
	}
}

func TestDeleteNotFound(t *testing.T) {
	m := NewManager()
	err := m.Delete(100)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
