package gredis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/mangmang/pkg/setting"
	"time"
)

var RedisConn *redis.Pool

func Setup() {
	RedisConn = &redis.Pool{
		MaxIdle:     setting.RedisSetting.MaxIdle,
		MaxActive:   setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

func Hset(key, field string, data interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("HSET", key, field, data)
	if err != nil {
		return err
	}
	return nil
}

func Hget(key, field string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("HGET", key, field))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Expire(key string, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()
	_, err := conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Hmset(key string, data map[string]interface{}) error {
	conn := RedisConn.Get()
	defer conn.Close()

	_, err := conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(data)...)
	if err != nil {
		return err
	}
	return nil

}

func Hgetall(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.StringMap(conn.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}
