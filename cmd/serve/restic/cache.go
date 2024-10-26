package restic

import (
	"strings"
	"sync"

	"github.com/tamankuc/rclone_ui/fs"
)

// cache implements a simple object cache
type cache struct {
	mu           sync.RWMutex         // protects the cache
	items        map[string]fs.Object // cache of objects
	cacheObjects bool                 // whether we are actually caching
}

// create a new cache
func newCache(cacheObjects bool) *cache {
	return &cache{
		items:        map[string]fs.Object{},
		cacheObjects: cacheObjects,
	}
}

// find the object at remote or return nil
func (c *cache) find(remote string) fs.Object {
	if !c.cacheObjects {
		return nil
	}
	c.mu.RLock()
	o := c.items[remote]
	c.mu.RUnlock()
	return o
}

// add the object to the cache
func (c *cache) add(remote string, o fs.Object) {
	if !c.cacheObjects {
		return
	}
	c.mu.Lock()
	c.items[remote] = o
	c.mu.Unlock()
}

// remove the object from the cache
func (c *cache) remove(remote string) {
	if !c.cacheObjects {
		return
	}
	c.mu.Lock()
	delete(c.items, remote)
	c.mu.Unlock()
}

// remove all the items with prefix from the cache
func (c *cache) removePrefix(prefix string) {
	if !c.cacheObjects {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}
	if prefix == "/" {
		c.items = map[string]fs.Object{}
		return
	}
	for key := range c.items {
		if strings.HasPrefix(key, prefix) {
			delete(c.items, key)
		}
	}
}
