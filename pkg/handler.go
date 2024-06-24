package ch

import (
	"fmt"
	"strconv"
	"time"
)

func (c *Cache) LRU(key string) *Cache {

	// Abstracting cache hit and updation is not required and would make things complicated, since cache value updation is dependent on policy.
	// Hence provided freedom to update logic in the policy.
	if !c.IsCacheFull() || c.IsValueExists(key) {
		c.Set(key)
		return c
	}

	// The first value is the thing, example "watches", second is the time it was added.
	var oldestTime time.Time
	var delkey string
	for k, v := range c.cache {
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			fmt.Println(err)
		}
		if parsedTime.Before(oldestTime) {
			oldestTime = parsedTime
			delkey = k
		}
	}

	c.Delete(delkey)
	c.Set(key)

	return c
}

func (c *Cache) LIFO(key string) *Cache {

	if c.IsValueExists(key) {
		// Time not updated to 'access time', kept as the added time
		return c
	}

	if !c.IsCacheFull() {
		c.Set(key)
		return c
	}

	var recentTime time.Time
	var delkey string
	for k, v := range c.cache {
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			fmt.Println(err)
		}
		if parsedTime.After(recentTime) {
			recentTime = parsedTime
			delkey = k
		}
	}

	c.Delete(delkey)
	c.Set(key)

	return c
}

func (c *Cache) FIFO(key string) *Cache {

	if c.IsValueExists(key) {
		// Time not updated to 'access time', kept as the added time
		return c
	}

	var latestTime time.Time
	var oldestkey string
	for k, v := range c.cache {
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			fmt.Println(err)
		}
		if parsedTime.Before(latestTime) {
			latestTime = parsedTime
			oldestkey = k
		}
	}

	c.Delete(oldestkey)
	c.Set(key)

	return c
}

func (c *Cache) LFU(key string) *Cache {

	// The first value is the thing, example "shoes", second is the frquency.
	v, ok := c.Get(key)
	if ok {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		i++
		c.Set(key, strconv.Itoa(i))
		return c
	}

	var lowestFreq int
	var leastFreqKey string
	for k, v := range c.cache {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
		}
		if i < lowestFreq {
			lowestFreq = i
			leastFreqKey = k
		}
	}

	c.Delete(leastFreqKey)
	c.Set(key, "1")

	return c
}

func (c *Cache) MRU(key string) *Cache {

	if !c.IsCacheFull() || c.IsValueExists(key) {
		c.Set(key)
		return c
	}

	var recentTime time.Time
	var delkey string
	for k, v := range c.cache {
		parsedTime, err := time.Parse(layout, v)
		if err != nil {
			fmt.Println(err)
		}
		if parsedTime.After(recentTime) {
			recentTime = parsedTime
			delkey = k
		}
	}

	c.Delete(delkey)
	c.Set(key)

	return c
}
