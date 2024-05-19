package lru

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	mutex      sync.Mutex
	list       *list.List
	entries    map[any]*list.Element
	cacheTimer time.Duration
	ttl        time.Duration
	capacity   int
	count      int
}

type cacheEntry struct {
	Key   any
	Value any
	Time  time.Time
}

func NewCache(ttl time.Duration) *Cache {
	return &Cache{
		mutex:      sync.Mutex{},
		list:       list.New(),
		entries:    make(map[any]*list.Element),
		cacheTimer: time.Duration(1) * time.Second,
		ttl:        ttl,
		capacity:   10,
		count:      0,
	}
}

func (c *Cache) EvictExpiredItems() {
	for {
		time.Sleep(c.cacheTimer)
		now := time.Now()
		c.mutex.Lock()
		for e := c.list.Front(); e != nil; e = e.Next() {
			entry := e.Value.(cacheEntry)
			if now.Sub(entry.Time) > c.ttl {
				c.list.Remove(e)
				delete(c.entries, entry.Key)
				c.count--
			}
		}

		c.mutex.Unlock()
	}
}

func (c *Cache) Put(key, value any) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.entries[key]; ok {
		c.list.MoveToFront(val)
		c.list.Front().Value = cacheEntry{
			Key:   key,
			Value: value,
			Time:  time.Now(),
		}
	} else {
		c.entries[key] = c.list.PushFront(cacheEntry{
			Key:   key,
			Value: value,
			Time:  time.Now(),
		})
		if c.count >= c.capacity {
			entry := c.list.Back()
			if entry != nil {
				delete(c.entries, entry.Value.(cacheEntry).Key)
				c.list.Remove(entry)
			}

		}
		c.count++
	}
}

func (c *Cache) Get(key any) (any, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if val, ok := c.entries[key]; ok {
		c.list.MoveToFront(val)
		return val.Value.(cacheEntry).Value, true
	}
	return nil, false
}
