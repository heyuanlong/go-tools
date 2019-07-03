package redis

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	ErrSingleLockInvalidRedisConn = errors.New("SingleLock : Invalid redis conn")
	ErrSingleLockOperationFailed  = errors.New("SingleLock : Operation failed")
	ErrSingleLockNotLocked        = errors.New("SingleLock : Not locked")
	ErrSingleLockInvalidLockValue = errors.New("SingleLock : Invalid lock value")
	ErrSingleLockLockIsUnlocked   = errors.New("SingleLock : Lock is unlocked")
	ErrSingleLockInvalidRedisConf = errors.New("SingleLock : Invalid redis conf")
)

type RedisPool struct {
	rp *redis.Pool
}

func NewRedisPool(host, port, auth string) (*RedisPool, error) {

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Println(addr)

	RedisClient := &redis.Pool{
		MaxIdle: 30,
		//MaxActive:   30,
		IdleTimeout: 60 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Println("redis open fail:", err)
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					c.Close()
					return nil, err
				}
			}
			// 选择db
			//c.Do("SELECT", REDIS_DB)
			return c, nil
		},
	}
	//懒加载
	return &RedisPool{rp: RedisClient}, nil
}

func (ts *RedisPool) GetRedis() redis.Conn {
	rc := ts.rp.Get()
	return rc
}
func (ts *RedisPool) CloseRedis(rc redis.Conn) {
	rc.Close()
}

func (ts *RedisPool) Set(key string, v interface{}, expire int64) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if err := rc.Send("SET", key, v); err != nil {
		log.Println(err)
		return err
	}

	if expire > 0 {
		if err := rc.Send("EXPIRE", key, expire); err != nil {
			log.Println(err)
			return err
		}
	}
	if err := rc.Flush(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *RedisPool) Incr(key string) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if _, err := rc.Do("INCR", key); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *RedisPool) Decr(key string) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if _, err := rc.Do("DECR", key); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *RedisPool) Del(key string) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if _, err := rc.Do("DEL", key); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *RedisPool) GetFloat64(key string) (float64, error) {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	v, err := redis.Float64(rc.Do("get", key))
	return v, err
}

func (ts *RedisPool) GetString(key string) (string, error) {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	v, err := redis.String(rc.Do("get", key))
	return v, err
}
func (ts *RedisPool) GetInt64(key string) (int64, error) {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	v, err := redis.Int64(rc.Do("get", key))
	return v, err
}

func (ts *RedisPool) GetIncr(key string) (int, error) {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	return redis.Int(rc.Do("incr", key))
}

func (ts *RedisPool) HgetAll(key string) (map[string]int64, error) {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	return redis.Int64Map(rc.Do("HGETALL", key))
}
func (ts *RedisPool) Hset(key string, subKey string, v interface{}) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if err := rc.Send("HSET", key, subKey, v); err != nil {
		log.Println(err)
		return err
	}
	if err := rc.Flush(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (ts *RedisPool) Hdel(key string, subKey string) error {
	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)
	if err := rc.Send("HDEL", key, subKey); err != nil {
		log.Println(err)
		return err
	}
	if err := rc.Flush(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (ts *RedisPool) Lock(keyName string, value string, milliseconds int64) error {
	if len(keyName) == 0 ||
		len(value) == 0 {
		return ErrSingleLockNotLocked
	}

	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)

	//	try to lock
	rpl, err := redis.String(rc.Do("SET", keyName, value, "NX", "PX", milliseconds))
	if nil != err {
		return err
	}
	if rpl != "OK" {
		return ErrSingleLockOperationFailed
	}

	return nil
}

func (ts *RedisPool) Unlock(keyName string, value string) error {
	if len(keyName) == 0 ||
		len(value) == 0 {
		return ErrSingleLockNotLocked
	}

	rc := ts.GetRedis()
	defer ts.CloseRedis(rc)

	//	try to unlock
	//	avoid to unlock a lock not belongs to the locker
	lockValue, err := redis.String(rc.Do("GET", keyName))
	if nil != err {
		return err
	}
	if lockValue != value {
		return ErrSingleLockInvalidLockValue
	}

	rpl, err := redis.Int(rc.Do("DEL", keyName))
	if nil != err {
		return err
	}

	if rpl != 1 {
		return ErrSingleLockLockIsUnlocked
	}

	return nil
}

func (ts *RedisPool) Test() error {
	rc := ts.rp.Get()
	_, err := redis.String(rc.Do("get", "key1"))
	rc.Close()
	if err != nil {
		return err
	}
	return nil
}
