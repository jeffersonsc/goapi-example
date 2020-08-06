package product

import (
	"log"
	"os"

	"github.com/mediocregopher/radix/v3"
)

// Storage implements of cache strategy
type Storage interface {
	Get(key string) (*CacheData, error)
	Del(key string) error
	Set(key string, duration int64, content []byte) error
}

// Cache struct implement cache interface
type Cache struct {
	redis *radix.Pool
	log   *log.Logger
}

// CacheData represents data save on redis
type CacheData struct {
	ExpiresAt int64  `redis:"expires_at"`
	Content   []byte `redis:"content"`
}

// NewStorageCache instance cache
func NewStorageCache(redis *radix.Pool) Storage {
	logger := NewStorageCacheLogger()
	return &Cache{
		redis: redis,
		log:   logger,
	}
}

// NewStorageCacheLogger .
func NewStorageCacheLogger() *log.Logger {
	return log.New(os.Stdout, "[storage]", 0)
}

// Get data by key
func (c Cache) Get(key string) (*CacheData, error) {
	var data CacheData

	err := c.redis.Do(radix.Cmd(&data, "HGETALL", key))
	if err != nil {
		c.log.Println("Failed get data from cache - ", err.Error())
		return nil, err
	}

	return &data, nil
}

// Set data by key
func (c Cache) Set(key string, expiresAt int64, content []byte) error {
	var cacheData CacheData

	cacheData.Content = content
	cacheData.ExpiresAt = expiresAt

	err := c.redis.Do(radix.FlatCmd(nil, "HSET", key, cacheData))
	if err != nil {
		c.log.Println("Failed set data from cache - ", err.Error())
		return err
	}
	return nil
}

// Del delete data by key
func (c Cache) Del(key string) error {
	err := c.redis.Do(radix.Cmd(nil, "DEL", key))
	if err != nil {
		c.log.Println("Failed delete data from cache - ", err.Error())
		return err
	}

	return nil
}
