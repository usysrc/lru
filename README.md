# LRU Cache

[![Build and Tests](https://github.com/usysrc/lru/actions/workflows/go.yml/badge.svg)](https://github.com/usysrc/lru/actions/workflows/go.yml)

This Go package implements a thread-safe Least Recently Used (LRU) cache with a time-to-live (TTL) feature. It allows storing, retrieving, and automatically evicting expired items from the cache.

## Features

- Thread-safe LRU caching
- Configurable TTL for cache entries
- Automatic eviction of expired items

## Installation

To install the package, run:

```sh
go get github.com/usysrc/lru
```

Then import it in your code:

```go
import "github.com/usysrc/lru"
```

## Usage

### Creating a Cache

Create a new cache with a specified TTL and capacity:

```go
ttl := 5 * time.Minute
capacity := 100
cache := lru.NewCache(ttl, capacity)
```

### Adding Items

Add an item to the cache with the `Put` method:

```go
cache.Put("key1", "value1")
```

### Retrieving Items

Retrieve an item from the cache with the `Get` method:

```go
value, ok := cache.Get("key1")
if ok {
    fmt.Println("Value:", value)
} else {
    fmt.Println("Key not found or expired")
}
```

### Evicting Expired Items

To automatically evict expired items, call the `EvictExpiredItems` method in a goroutine:

```go
go cache.EvictExpiredItems()
```

## API Reference

### `NewCache(ttl time.Duration, capacity int) *Cache`

Creates a new cache with the given TTL and capacity.

- `ttl`: The time-to-live duration for cache entries.
- `capacity`: The maximum number of entries the cache can hold.

### `func (c *Cache) Put(key, value any)`

Adds an item to the cache. If the key already exists, it updates the value and moves the item to the front.

- `key`: The key of the item.
- `value`: The value of the item.

### `func (c *Cache) Get(key any) (any, bool)`

Retrieves an item from the cache. If the key is found and the item is not expired, it moves the item to the front.

- `key`: The key of the item.
- Returns: The value and a boolean indicating if the key was found.

### `func (c *Cache) EvictExpiredItems()`

Evicts expired items from the cache. This method should be called in a goroutine to run periodically.

## Example

```go
package main

import (
	"fmt"
	"time"
	"github.com/usysrc/lru"
)

func main() {
	ttl := 5 * time.Minute
	capacity := 100
	cache := lru.NewCache(ttl, capacity)

	// Start the eviction goroutine
	go cache.EvictExpiredItems()

	// Add items to the cache
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Retrieve items from the cache
	value, ok := cache.Get("key1")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found or expired")
	}

	// Wait for items to expire
	time.Sleep(6 * time.Minute)

	// Try to retrieve expired item
	value, ok = cache.Get("key1")
	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found or expired")
	}
}
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

---

This README provides an overview of the package, installation instructions, usage examples, and API reference. For more details, refer to the package documentation and source code.
