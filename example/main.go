package main

import (
	"fmt"
	cache "github.com/uzzalhcse/gcache"
)

func main() {
	c := cache.NewInMemoryCache()

	// Set a value with a TTL of 60 seconds
	c.Set("myKey", "myValue", 60)

	// Get the value
	value, err := c.Get("myKey")
	if err != nil {
		// Handle error (e.g., item not found)
		fmt.Println(err)
	} else {
		// Use the value
		fmt.Println(value)
	}
}
