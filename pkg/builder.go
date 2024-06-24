package ch

import (
	"fmt"
	"time"
)

// Go generics not really helful in this case
// type AcceptedKeys interface {
// 	int64
// }

// type AcceptedValues interface {
// 	int64 | string
// }

type Cache struct {
	cacheInitializationTime time.Time
	maxCacheSize            int
	cache                   map[string]string 
}

func (c *Cache) Set(key string, values ...string) {
	var value string
	if len(values) > 0 {
		value = values[0]
	} else {
		value = time.Now().Format(layout)
	}
	c.cache[key] = value
}

func (c *Cache) Delete(key string) {
	delete(c.cache, key)
}

func (c Cache) Get(key string) (string, bool) {
	value, found := c.cache[key]
	return value, found
}

func (c Cache) GetAll() map[string]string {
	return c.cache
}

func (c Cache) IsValueExists(key string) bool {
	_, ok := c.Get(key)
	if ok {
		return true
	} else {
		return false
	}
}

func (c Cache) IsCacheFull() bool {
	if len(c.cache) == c.maxCacheSize {
		return true
	} else {
		return false
	}
}

// CacheBuilder for choosy cache building, implemented builder and director pattern for scaling in case you need more fields in Cache types and more combo of cache building.
type CacheBuilder interface {
	Set(key string, value string) CacheBuilder
	SetSize(size int) CacheBuilder
	Build() *Cache
}

type cacheBuilder struct {
	cache *Cache
}

func (b *cacheBuilder) Set(key string, value string) CacheBuilder {
	b.cache.Set(key, value)
	return b
}

func (b *cacheBuilder) SetSize(size int) CacheBuilder {
	b.cache.maxCacheSize = size
	return b
}

func (b *cacheBuilder) Build() *Cache {
	return b.cache
}

func NewCacheBuilder() CacheBuilder {
	return &cacheBuilder{
		cache: &Cache{cache: make(map[string]string)},
	}
}

type Director struct {
	builder CacheBuilder
}

func (d *Director) ConstructEmpty() *Cache {
	cache := d.builder.Build()
	cache.maxCacheSize = defaultCacheSize
	return cache

}

func (d *Director) ConstructManual(m map[string]string, size int) *Cache {
	cache := d.builder.Build()
	if len(m) > size {
		fmt.Println("Cache size is less than the map size, not initialized.")
		return nil
	}

	cache.cache = m
	cache.maxCacheSize = size
	return cache
}

func NewDirector(builder CacheBuilder) *Director {
	return &Director{
		builder: builder,
	}
}
