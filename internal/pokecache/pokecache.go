package pokecache

import (
    "sync"
    "time"
)

type Cache struct {
    entry map[string]cacheEntry
    mu    *sync.Mutex
}

type cacheEntry struct {
    createdAt time.Time
    value     []byte
}

func NewCache(interval time.Duration) Cache {
    c := Cache{
        entry: make(map[string]cacheEntry),
        mu:    &sync.Mutex{},
    }

    go c.reapLoop(interval)

    return c
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.entry[key] = cacheEntry{
        createdAt: time.Now(),
        value:     val,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    entry, exists := c.entry[key]
    return entry.value, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mu.Lock()
    defer c.mu.Unlock()
    for k, v := range c.entry {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.entry, k)
        }
    }
}
