package ch

import (
	"fmt"
	"sync"
	"time"
)

// Go generics but used interfaces cause it'll get dirty in methods
// type AcceptedKeys interface {
// 	int64
// }

// type AcceptedValues interface {
// 	int64 | string
// }

type Cache struct {
	cacheInitializationTime time.Time
	maxCacheSize            int
	cache                   map[interface{}]interface{}
	mutex                   sync.RWMutex
}

func (c *Cache) Set(key interface{}, values ...interface{}) {
	// if c.IsCacheFull() {
	// 	fmt.Println("Cache is full, cannot add more values. Evict some before adding.")
	// 	return
	// }
	var value interface{}
	if len(values) > 0 {
		value = values[0]
	} else {
		value = time.Now()
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
}

func (c *Cache) Delete(key interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.cache, key)
}

func (c *Cache) Get(key interface{}) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	value, found := c.cache[key]
	return value, found
}

func (c *Cache) GetAll() map[interface{}]interface{} {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	m := make(map[interface{}]interface{}, len(c.cache))
	for k, v := range c.cache {
		m[k] = v
	}
	return m
}

func (c *Cache) IsValueExists(key interface{}) bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	_, ok := c.Get(key)
	if ok {
		return true
	} else {
		return false
	}
}

func (c *Cache) IsCacheFull() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	// fmt.Println("Cache size: ", len(c.cache))
	// fmt.Println("Max Cache size: ", c.maxCacheSize)
	if len(c.cache) == c.maxCacheSize {
		// fmt.Println("true")
		return true
	} else {
		return false
	}
}

// Not implmenting thread safety here, as it is not required. Assuming cache is build once.
// CacheBuilder for choosy cache building, implemented builder and director pattern for scaling in case you need more fields in Cache types and more combo of cache building.
type CacheBuilder interface {
	Set(key interface{}, value interface{}) CacheBuilder
	SetSize(size int) CacheBuilder
	Build() *Cache
}

type cacheBuilder struct {
	cache *Cache
}

func (b *cacheBuilder) Set(key interface{}, value interface{}) CacheBuilder {
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
		cache: &Cache{cache: make(map[interface{}]interface{})},
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

func (d *Director) ConstructManual(m map[interface{}]interface{}, size int) *Cache {
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
