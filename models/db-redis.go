package models

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

const (
	HashPOSTS   = `postAuthor`
	AuthorPosts = `CountAuthorPosts`
)

// implement interface zdb.DB
type DBredis struct {
	Client *redis.Client
}

func (dbr *DBredis) Put(p *Post) error {
	key := p.PostName + ":" + p.Author + ":" + p.Date
	hkey := HashPOSTS + ":" + getHash(key)
	jsn, _ := json.Marshal(p)

	pipe := dbr.Client.Pipeline()
	pipe.Set(hkey, string(jsn), 0)
	//pipe.HSet(hkey, "name", p.Author)
	//pipe.HSet(hkey, "post", p.PostName)
	//pipe.HSet(hkey, "date", p.Date)
	pipe.HIncrBy(AuthorPosts, p.Author, 1)
	pipe.SAdd("author:"+p.Author, hkey)
	pipe.SAdd("post:"+p.PostName, hkey)
	_, err := pipe.Exec()
	return err
}

func (dbr *DBredis) Get(p *Post) (string, error) {

	var list *redis.StringSliceCmd
	switch {
	case p.Author != `` && p.PostName == ``:
		list = dbr.Client.SMembers("author:" + p.Author)
	case p.Author == `` && p.PostName != ``:
		list = dbr.Client.SMembers("post:" + p.PostName)
	default:
		list = dbr.Client.SInter("author:"+p.Author, "post:"+p.PostName)
	}
	if err := list.Err(); err != nil {
		return "", err
	}
	ids := list.Val()
	if len(ids) == 0 {
		return ``, nil
	}

	pp := dbr.Client.MGet(ids...)
	if err := pp.Err(); err != nil {
		return "", err
	}
	jsn, _ := json.Marshal(pp.Val())

	return string(jsn), nil
}

func (dbr *DBredis) Count(p *Post) (string, error) {
	var (
		authors map[string]string
		bs      []byte
		err     error
	)
	if p.Author == `` {
		a := dbr.Client.HGetAll(AuthorPosts)
		if a.Err() != nil {
			return "", a.Err()
		}
		authors = a.Val()
	} else {
		a := dbr.Client.HGet(AuthorPosts, p.Author)
		if a.Err() != nil {
			return "", a.Err()
		}
		authors = map[string]string{p.Author: a.Val()}
	}

	bs, err = json.Marshal(authors)
	return string(bs), err
}

func getHash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
