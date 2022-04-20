package gredis

import (
	"encoding/json"
	"prow/library/log"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

// Setup Initialize the Redis instance
func Setup() error {
	RedisConn = &redis.Pool{
		MaxIdle:     g.Config().GetInt("redis.default.maxIdle"),
		MaxActive:   g.Config().GetInt("redis.default.maxActive"),
		IdleTimeout: g.Config().GetDuration("redis.default.idleTimeout"),
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", g.Config().GetString("redis.default.host"))
			if err != nil {
				return nil, err
			}
			if g.Config().GetString("redis.default.pwd") != "" {
				if _, err := c.Do("AUTH", g.Config().GetString("redis.default.pwd")); err != nil {
					c.Close()
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

	return nil
}

// Set a key/value
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

// Exists check a key
func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
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

func RedisDo(commandName string, args ...interface{}) (*gvar.Var, error) {
	conn := RedisConn.Get()
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Logger.Errorf("RedisDo close error: %v", err)
		}
	}()
	return resultToVar(conn.Do(commandName, args...))
}

// resultToVar converts redis operation result to gvar.Var.
func resultToVar(result interface{}, err error) (*gvar.Var, error) {
	if err == nil {
		if result, ok := result.([]byte); ok {
			return gvar.New(gconv.UnsafeBytesToStr(result)), err
		}
		// It treats all returned slice as string slice.
		if result, ok := result.([]interface{}); ok {
			return gvar.New(gconv.Strings(result)), err
		}
	}
	return gvar.New(result), err
}
