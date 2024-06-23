package main

import (
	"fmt"

	b "github.com/ashpect/cache/pkg"
)

func main() {
	cache := b.NewCacheBuilder(1).Set("key", "value").Build()
	fmt.Println(cache.Get("key"))
	cache.Set("keywsd", "valusde")
	cache.Set("sdfdsfsdf", "SDfdsf")
	fmt.Println(cache.Get("key"))

	// director := b.NewDirector(b.NewCacheBuilder(1))
	// lifo := director.ConstructLIFO()
	// fmt.Println(lifo.Get("key"))

}

func CustomSet(c *b.Cache, key string, value string) {
	// your logic
	c.Set(key, value)
}
