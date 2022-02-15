package main

import (
	"chaostheory-task/internal/sqlitestore"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ItemServer struct {
	store *sql.DB
}

func NewItermServer(db *sql.DB) *ItemServer {
	sqlitestore.NewItemStore(db)
	return &ItemServer{
		store: db,
	}
}

func (is *ItemServer) ListHandler(w http.ResponseWriter, r *http.Request) {
	// allItems := is.store.GetAllItems()
	allItems := sqlitestore.GetAllItems(is.store)

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
	sqlitestore.AddItem(is.store, item.Key, item.Value)
	js, _ := json.Marshal(ItemType{item.Key, item.Value})
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}

func main() {
	os.Remove("test.db")
	log.Println("Creating ./data/test.db...")
	file, err := os.Create("test.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	db, _ := sql.Open("sqlite3", fmt.Sprintf("%s.db", "test"))
	defer db.Close()
	server := NewItermServer(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/list", server.ListHandler)
	r.Post("/add", server.AddHandler)
	http.ListenAndServe(":3000", r)
}
