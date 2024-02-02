package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type MemCacheClient struct {
	memCache *cache.Cache
}

func (c MemCacheClient) UpdateCache(key string, cacheData map[string]string, expireTime time.Duration) error {
	c.memCache.Set(key, cacheData, expireTime)
	return nil
}

func (c MemCacheClient) GetCache(key string) (map[string]string, error) {
	cacheData, found := c.memCache.Get(key)
	if !found {
		return nil, nil
	}
	return cacheData.(map[string]string), nil

}