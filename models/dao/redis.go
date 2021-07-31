package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Mmx233/VpsBrokerS/global"
	"github.com/go-redis/redis/v8"
	"time"
)

func init() {
	Redis.Init(&redis.Options{
		Addr:       global.Config.Redis.Addr,
		Password:   global.Config.Redis.Password,
		DB:         global.Config.Redis.DB,
		MaxConnAge: time.Hour * 5,
	})
}

type r struct {
	DB  *redis.Client
	ctx context.Context
}

var Redis r

func (s *r) Init(options *redis.Options) {
	s.DB = redis.NewClient(options)
	s.ctx = context.Background()
}

func (s *r) Cache(k string, v string, t time.Duration) error {
	return s.DB.Set(s.ctx, k, v, t).Err()
}

func (s *r) CacheStrut(k string, v interface{}, t time.Duration) error {
	j, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.Cache(k, string(j), t)
}

func (s *r) Read(k string) (string, error) {
	v, e := s.DB.Get(s.ctx, k).Result()
	if e != nil {
		return "", e
	}
	return v, nil
}

func (s *r) ReadStruct(k string, v interface{}) error {
	t, e := s.Read(k)
	if e != nil {
		return e
	}
	return json.Unmarshal([]byte(t), v)
}

func (s *r) Del(k string) {
	e := s.DB.Del(s.ctx, k).Err()
	if e != nil && e != redis.Nil {
		fmt.Println("redis删除失败\n", e)
	}
}

func (s *r) Flush() {
	s.DB.FlushDB(s.ctx)
}
