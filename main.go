package main

import (
	"chaostheory-task/internal/store"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ItemServer struct {
	store store.ItemStore
}

func NewItermServer() *ItemServer {
	return &ItemServer{
		store: *store.NewItemStore(),
	}
}

func (is *ItemServer) ListHandler(w http.ResponseWriter, r *http.Request) {
	allItems := is.store.GetAllItems()

	js, err := json.Marshal(allItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}

func (is *ItemServer) AddHandler(w http.ResponseWriter, r *http.Request) {
	type ItemType struct {
		Key   string
		Value string
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var item ItemType
	if err := dec.Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	is.store.AddItem(item.Key, item.Value)
	js, _ := json.Marshal(ItemType{item.Key, item.Value})
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}

func main() {
	r := chi.NewRouter()
	server := NewItermServer()

	r.Use(middleware.Logger)
	r.Get("/list", server.ListHandler)
	r.Post("/add", server.AddHandler)
	http.ListenAndServe(":3000", r)
}
