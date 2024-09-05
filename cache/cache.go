package cache

import "errors"

//const showExpectedResult = false;
//const showHints = false;

type Key string
type Value string

type ValueWriter interface {
	Set(k Key, v Value)
}
type ValueReader interface {
	Get(k Key) (Value, error)
}

// Make any required changes to the Cache struct.
type Cache struct {
	cache map[Key]Value
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[Key]Value),
	}
}

func (c *Cache) Set(k Key, v Value) {
	c.cache[k] = v
}

func (c *Cache) Get(k Key) (Value, error) {
	value, exists := c.cache[k]
	if exists {
		return value, nil
	} else {
		return Value(""), errors.New("key does not exist")
	}
}

func WriteValues(w ValueWriter, keys []Key, values []Value) {
	for i, k := range keys {
		w.Set(k, values[i])
	}
}

func ReadValues(r ValueReader, keys []Key) ([]Value, error) {
	values := make([]Value, len(keys))
	for i, k := range keys {
		v, err := r.Get(k)
		if err != nil {
			return nil, err
		}
		values[i] = v
	}

	return values, nil
}
