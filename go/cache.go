package isuports
import (
	"time"
	"github.com/patrickmn/go-cache"
)

// Cache は Type Parameterに対応したKey Value Store
type Cache[V any] struct {
	cache *cache.Cache
}

func NewCache[V any]() *Cache[V] {
	return &Cache[V]{cache: cache.New(cache.NoExpiration, cache.NoExpiration)}
}
func NewCacheWithExpire[V any](defaultExpiration, cleanupInterval time.Duration) *Cache[V] {
	return &Cache[V]{cache: cache.New(defaultExpiration, cleanupInterval)}
}

func (c *Cache[V]) Get(key string) (V, bool) {
	v, ok := c.cache.Get(key)
	if ok {
		return v.(V), true
	}
	var defaultValue V
	return defaultValue, false
}

func (c *Cache[V]) Set(k string, v V) {
	c.cache.Set(k, v, cache.NoExpiration)
}

func (c *Cache[V]) SetWithExpire(k string, v V, d time.Duration) {
	c.cache.Set(k, v, d)
}

// Flush はキャッシュをクリアします
func (c *Cache[V]) Flush() {
	c.cache.Flush()
}