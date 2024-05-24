package lru

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(time.Second*1, 100)
	go c.EvictExpiredItems()
	c.Put("foo", "bar")
	if _, ok := c.Get("foo"); !ok {
		t.Error("key not found")
	}
	if val, ok := c.Get("foo"); ok {
		if val != "bar" {
			t.Error("value not found")
		}
	}
	c.Put("foo", "bar-updated")
	if val, ok := c.Get("foo"); ok {
		if val != "bar-updated" {
			t.Error("value not updated")
		}
	}
	time.Sleep(time.Second * 10)
	if _, ok := c.Get("foo"); ok {
		t.Error("key found, expected to be not found")
	}

}

func TestCacheCapacity(t *testing.T) {
	c := NewCache(time.Second*100, 10)
	go c.EvictExpiredItems()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		c.Put(fmt.Sprintf("foo-%d", i), fmt.Sprintf("bar-%d", i))
	}
	// this should evict the oldes entry: foo-0
	c.Put("foo-new", "bar-new")
	if val, ok := c.Get("foo-0"); ok {
		t.Errorf("key was found, expected to not be found, value: %s", val)
	}
}
