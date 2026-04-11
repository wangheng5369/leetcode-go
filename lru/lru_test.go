package lru

import (
	"testing"
)

func TestLRUCache_Basic(t *testing.T) {
	c := New(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	if c.Len() != 3 {
		t.Errorf("expected len 3, got %d", c.Len())
	}

	v, _ := c.Get("a")
	if v != 1 {
		t.Errorf("expected 1, got %v", v)
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	c := New(3)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("c", 3)

	// "a" is accessed, so "b" should be evicted when we add "d"
	c.Get("a")
	c.Put("d", 4)

	if c.Len() != 3 {
		t.Errorf("expected len 3, got %d", c.Len())
	}

	_, err := c.Get("b")
	if err == nil {
		t.Error("expected error for evicted key 'b'")
	}

	_, _ = c.Get("c") // should exist
	_, _ = c.Get("d") // should exist
}

func TestLRUCache_Update(t *testing.T) {
	c := New(2)

	c.Put("a", 1)
	c.Put("b", 2)
	c.Put("a", 10) // update "a"

	v, _ := c.Get("a")
	if v != 10 {
		t.Errorf("expected 10, got %v", v)
	}

	// add "c" - should evict "b" (a was made recently used)
	c.Put("c", 3)

	_, err := c.Get("b")
	if err == nil {
		t.Error("expected error for evicted key 'b'")
	}
}
