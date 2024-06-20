package main

import (
	"fmt"
	"gocache/lib/cache"
)

func main() {
	cache := cache.NewLRUCache(2)
	cache.Set("a", "1")
	cache.Set("b", 2)

	val1, err := cache.Get("a")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val1)
	}

	val2, err := cache.Get("b")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val2)
	}

	// Adding a new key which should evict the least recently used key ("a")
	cache.Set("c", 3)

	val3, err := cache.Get("c")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val3)
	}

	val4, err := cache.Get("b")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val4)
	}

	// Deleting key "b"
	cache.Delete("b")

	val5, err := cache.Get("b")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val5)
	}

	// Trying to delete key "b" again to check error handling
	err = cache.Delete("b")
	fmt.Println(err)
}
