package cache

import (
	"CleanTodo/model"

	"github.com/gomodule/redigo/redis"
)

type redisTodoCache struct {
	c redis.Conn
}

func NewRedisTodoCache(r redis.Conn) model.TodoCache {
	return &redisTodoCache{
		c: r,
	}
}

func (r *redisTodoCache) CacheSet(listid, title, content string) error {
	_, err := r.c.Do("hset", listid, "title", title, "content", content)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisTodoCache) CacheDel(listid string) error {
	_, err := r.c.Do("del", listid)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisTodoCache) CacheGet(listid string) (model.List, error) {
	var list model.List
	title, err := redis.String(r.c.Do("hget", listid, "title"))
	if err != nil {
		return list, err
	}
	content, err := redis.String(r.c.Do("hget", listid, "content"))
	if err != nil {
		return list, err
	}
	list.Listid = listid
	list.Title = title
	list.Content = content
	return list, nil
}

func (r *redisTodoCache) CacheLoad(listid string) bool {
	len, _ := redis.Int(r.c.Do("hlen", listid))
	if len == 0 {
		return false
	}
	return true
}
