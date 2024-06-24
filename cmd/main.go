package main

import (
	"fmt"
	"time"

	. "github.com/ashpect/cache/pkg"
)

// Here's how you can use the cache package

func main() {

	// ----- BUILDING CACHES -----

	// Build standard empty cache with default size using director ConstructEmpty.
	builder := NewCacheBuilder()
	c := NewDirector(builder).ConstructEmpty()
	fmt.Println(c.GetAll())

	// Build standard cache using ConstructManual with initial map and size
	m := map[string]string{
		"watches": "2021-01-01 00:00:00",
		"shoes":   "2021-01-02 00:00:00",
		"bags":    "2021-01-03 00:00:00",
	}
	c = NewDirector(builder).ConstructManual(m, 3)
	fmt.Println(c.GetAll())

	// Building a cache with your needs instead of director
	c = NewCacheBuilder().SetSize(3).Build() // complex cache building if required, giving full freedom using builder
	fmt.Println(c.GetAll())

	// ----- HandlingIncomingData -----

	incoming := "laptops"
	response := c.LIFO(incoming).GetAll() // Standard Policies are inbuild methods
	fmt.Println(response)

	// To use your policy kindly follow the function NewPolicy Blueprint given below.
	// Below is an example of using caches with a custom policy as well as LIFO later.

	incoming = "ipads"
	c = NewPolicy(c, incoming)
	fmt.Println(c.GetAll()) // empty cache, so adds ipads as per Newpolicy

	time.Sleep(1.00 * time.Second)
	testingData := "watches"
	c = c.LIFO(testingData)
	fmt.Println(c.GetAll()) // cache not full yet, so adds watches

	time.Sleep(1 * time.Second)
	// now the cache is full
	testingData = "shoes"
	c = c.LIFO(testingData)
	fmt.Println(c.GetAll())

	time.Sleep(1 * time.Second)
	// should delete everything and add new value as per the new policy
	testingData = "bags"
	c = NewPolicy(c, testingData)
	fmt.Println(c.GetAll())
}

// BluePrint of how to define your policies
func NewPolicy(c *Cache, key string) *Cache {
	// A weird out little method to empty cache every time it gets filled and add incoming stuff irrespctively lol
	if c.IsCacheFull() {
		m := c.GetAll()
		for k := range m {
			c.Delete(k)
		}
	}

	// Value is optional if not given sets to time.Now()
	c.Set(key)
	return c
}
