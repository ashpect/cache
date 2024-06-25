package pkg

import (
	"fmt"
	"time"
)

func (c *Cache) LRU(key interface{}) *Cache {

	// Abstracting cache hit and updation is not required and would make things complicated, since cache value updation is dependent on policy.
	// Hence provided freedom to update logic in the policy.
	if !c.IsCacheFull() || c.IsValueExists(key) {
		c.Set(key)
		return c
	}

	var oldestTime time.Time
	var delkey interface{}
	m := c.GetAll()
	for k, v := range m {
		parsedTime, ok := v.(time.Time)
		if !ok {
			panic(fmt.Errorf("Value %v is not a time format: %v\n", k, v))
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

func (c *Cache) LIFO(key interface{}) *Cache {

	if c.IsValueExists(key) {
		// Time not updated to 'access time', kept as the added time
		return c
	}

	if !c.IsCacheFull() {
		c.Set(key)
		return c
	}

	var recentTime time.Time
	var delkey interface{}
	m := c.GetAll()
	for k, v := range m {
		parsedTime, ok := v.(time.Time)
		if !ok {
			panic(fmt.Errorf("Value %v is not a time format: %v\n", k, v))
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

func (c *Cache) FIFO(key interface{}) *Cache {

	if c.IsValueExists(key) {
		return c
	}

	var latestTime time.Time
	var oldestkey interface{}
	m := c.GetAll()
	for k, v := range m {
		parsedTime, ok := v.(time.Time)
		if !ok {
			panic(fmt.Errorf("Value %v is not a time format: %v\n", k, v))
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

func (c *Cache) LFU(key interface{}) *Cache {

	if c.IsValueExists(key) {
		v, _ := c.Get(key)
		intValue, ok := v.(int)
		if !ok {
			panic(fmt.Errorf("Value %v is not an int", v))
		}
		intValue++
		c.Set(key, intValue)
		return c
	}

	if !c.IsCacheFull() {
		c.Set(key, 1)
	} else {
		var lowestFreq int = 1e9
		var leastFreqKey interface{}
		m := c.GetAll()
		for k, v := range m {
			intValue, ok := v.(int)
			if !ok {
				panic(fmt.Errorf("Value for key %v is not a string: %v\n", k, v))
			}
			if intValue < lowestFreq {
				lowestFreq = intValue
				leastFreqKey = k
			}
		}
		c.Delete(leastFreqKey)
		c.Set(key, 1)
	}
	return c
}

func (c *Cache) MRU(key interface{}) *Cache {

	if !c.IsCacheFull() || c.IsValueExists(key) {
		c.Set(key)
		return c
	}

	var recentTime time.Time
	var delkey interface{}
	m := c.GetAll()
	for k, v := range m {
		parsedTime, ok := v.(time.Time)
		if !ok {
			panic(fmt.Errorf("Value %v is not a time format: %v\n", k, v))
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
