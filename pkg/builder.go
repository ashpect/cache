package ch

import "fmt"

type AcceptedKeys interface {
	int64
}

type AcceptedValues interface {
	int64 | string
}

// change this to cache handler very late which access cache directly somewhere
type Cache struct {
	cache map[string]string
}

func (c *Cache) Set(key string, value string) {
	c.cache[key] = value
}

func (c Cache) Get(key string) (string, bool) {
	value, found := c.cache[key]
	fmt.Println(len(c.cache))
	for k, v := range c.cache {
		fmt.Println(k, v)
	}
	return value, found
}

// CacheBuilder for choosy cache building
type CacheBuilder interface {
	Set(string, string) CacheBuilder
	Build() *Cache
}

type cacheBuilder struct {
	cache *Cache
}

func (b *cacheBuilder) Set(key string, value string) CacheBuilder {
	b.cache.Set(key, value)
	return b
}

func (b *cacheBuilder) Build() *Cache {
	return b.cache
}

// TODO : Handle building a fixed map size to emulate actual cache, write a wrapper
func NewCacheBuilder(size int) CacheBuilder {
	return &cacheBuilder{
		cache: &Cache{cache: make(map[string]string, size)},
	}
}

type Director struct {
	builder CacheBuilder
}

func (d *Director) ConstructStandard() *Cache {
	return d.builder.Build()
}

func (d *Director) ConstructManual(m map[string]string) *Cache {
	cache := d.builder.Build()
	cache.cache = m
	return cache
}

func NewDirector(builder CacheBuilder) *Director {
	return &Director{
		builder: builder,
	}
}
