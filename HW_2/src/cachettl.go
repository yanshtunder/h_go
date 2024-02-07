package main

import (
	"fmt"
	"time"
)

type entry struct {
	value  interface{}
	expiry *time.Time
}

type Cache struct {
	timeTTL time.Duration
	cache   map[string]*entry
}

func New(defaultTTL time.Duration) *Cache {
	if defaultTTL < time.Second {
		defaultTTL = time.Second
	}

	cache := &Cache{
		timeTTL: defaultTTL,
		cache:   make(map[string]*entry),
	}
	return cache
}

func (cache *Cache) Add(key string, value interface{}, ttl time.Duration) {
	expiry := time.Now().Add(ttl)

	cache.cache[key] = &entry{
		value:  value,
		expiry: &expiry,
	}
}

func (cache *Cache) Get(key string) (interface{}, bool) {

	v, ok := cache.cache[key]

	if ok && v.expiry != nil && v.expiry.After(time.Now()) {
		return v.value, true
	}
	return nil, false
}

func (cache *Cache) GetKeys() []interface{} {

	keys := make([]interface{}, len(cache.cache))
	var i int
	for k := range cache.cache {
		keys[i] = k
		i++
	}
	return keys
}

func (cache *Cache) Count() int {
	return len(cache.cache)
}

func main() {
	cache := New(time.Second)
	key := "foo"
	val := 1
	ttl := time.Second
	cache.Add(key, val, ttl)

	keys := cache.GetKeys()
	v, ok := cache.Get(keys[0].(string))
	fmt.Printf("len: %d\tkey: %v\tval: %v\telIs: %v\n", len(keys), keys[0], v, ok)

	time.Sleep(2 * time.Second)

	keys = cache.GetKeys()
	v, ok = cache.Get(keys[0].(string))
	fmt.Printf("len: %d\tkey: %v\tval: %v\telIs: %v\n", len(keys), keys[0], v, ok)

}
