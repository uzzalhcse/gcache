# gCache: In-Memory Cache Package
## Overview

This package provides a straightforward in-memory cache implementation. It allows you to store key-value pairs in memory with optional time-to-live (TTL) expiration.


### Features

-   **Simple Interface**: Easy-to-use functions for setting and getting cached items.
-   **TTL Expiration**: Set a time-to-live for cached items to automatically expire them.
-   **Thread-Safe**: Utilizes a sync.RWMutex for safe concurrent access.

## Installation
`go get -u github.com/uzzalhcse/gcache` 

## Usage
```
package main

import (
cache "github.com/uzzalhcse/gcache"
"time"
)

func main() {
// Create a new in-memory cache
c := cache.NewInMemoryCache()

	// Set a value with a TTL of 60 seconds
	c.Set("myKey", "myValue", 60)

	// Get the value
	value, err := c.Get("myKey")
	if err != nil {
		// Handle error (e.g., item not found)
	} else {
		// Use the value
	}

	// ...
}
``` 

----------

## Contributing

If you have any suggestions, improvements, or bug fixes, feel free to open an issue or create a pull request.

----------
