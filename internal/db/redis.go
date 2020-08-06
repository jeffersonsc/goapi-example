package db

import "github.com/mediocregopher/radix/v3"

// NewRedisConn create a new redis connection
func NewRedisConn(uri string) (*radix.Pool, error) {
	pool, err := radix.NewPool("tcp", uri, 10)
	if err != nil {
		return nil, err
	}

	err = pool.Do(radix.Cmd(nil, "PING"))
	if err != nil {
		return nil, err
	}

	return pool, err
}
