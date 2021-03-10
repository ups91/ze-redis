package main

import (
	"fmt"
	"strconv"
	zlog "ze-redis-test/log"
	"ze-redis-test/models"

	"github.com/go-redis/redis"
)

func newApp() (app, error) {
	a := app{}
	cnf, err := getConfig() // HARDCODED STUB
	if err != nil {
		fmt.Println(err)
		return a, err
	}
	a.Cnf = cnf

	a.Log = &zlog.Log{} // stub - flush all to stdout

	switch a.Cnf.Params["db_type"] {
	case "redis":
		db, err := initRedis(a.Cnf.Params)
		if err != nil {
			return a, err
		}
		a.DB = db

	case "reindex": // will do
	case "postgres": // maybe will do

	default:
		return a, fmt.Errorf("Unknown database type")
	}
	return a, nil
}

func initRedis(params map[string]string) (models.DB, error) {
	dbnum, _ := strconv.Atoi(params["db_num"])

	client := redis.NewClient(&redis.Options{
		Addr:     params["db_addr"] + ":" + params["db_port"],
		Password: params["db_password"],
		DB:       dbnum,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &models.DBredis{Client: client}, nil
}
