package cache

import "time"

type Element struct {
	v        string
	deadline time.Time
}
type Cache struct {
	e map[string]Element
}

func NewCache() Cache {

	return Cache{make(map[string]Element)}
}

func (c Cache) Get(key string) (string, bool) {
	elem := c.e[key].v

	if c.e[key].deadline.After(time.Now()) {
		return elem, true
	} else {
		delete(c.e, key)
		return elem, false
	}

}

func (c *Cache) Put(key, value string) {

	newb := Element{value, time.Now().AddDate(200, 0, 0)}
	c.e[key] = newb
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0)
	for i := range c.e {
		if c.e[i].deadline.After(time.Now()) {
			keys = append(keys, i)
		} else {
			delete(c.e, i)
		}

	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	newb := Element{value, deadline}
	c.e[key] = newb

}
