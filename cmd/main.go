package main

import (
	"fmt"
	"net/http"
	"ze-redis-test/config"
	zlog "ze-redis-test/log"
	"ze-redis-test/models"
)

type app struct {
	Cnf config.Conf
	DB  models.DB
	Log zlog.Logger
}

func main() {
	app, err := newApp()
	if err != nil {
		fmt.Println("App initialization err: ", err)
		return
	}

	http.HandleFunc("/put", app.Put)
	http.HandleFunc("/get", app.Get)
	http.HandleFunc("/count", app.Count)

	app.Log.Log("Starting server at port " + app.Cnf.Params["app_port"])
	if err = http.ListenAndServe(":"+app.Cnf.Params["app_port"], nil); err != nil {
		app.Log.Log(err)

	}
	app.Log.Log("App stopped")
}
