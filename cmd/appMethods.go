package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ze-redis-test/models"
)

func (a *app) Put(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var p models.Post

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Log.Log(err.Error())
		return
	}
	if err := a.DB.Put(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Log.Log(err.Error())
		return
	}

}
func (a *app) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var p models.Post

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Log.Log(err.Error())
		return
	}
	if _, err := a.DB.Get(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Log.Log(err.Error())
		return
	}

}
func (a *app) Count(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var p models.Post

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		a.Log.Log(err.Error())
		return
	}
	jsn, err := a.DB.Count(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		a.Log.Log(err.Error())
		return
	}
	fmt.Fprintf(w, jsn)
}
