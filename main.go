package main

import (
	"chaostheory-task/internal/sqlitestore"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type ItemServer struct {
	store *sql.DB
}

func NewItemServer(db *sql.DB) *ItemServer {
	return &ItemServer{
		store: db,
	}
}

func (is *ItemServer) ListHandler(w http.ResponseWriter, r *http.Request) {
	allItems := sqlitestore.GetAllItems(is.store)

	renderJSON(w, allItems)
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
	renderJSON(w, ItemType{item.Key, item.Value})
}

func (is *ItemServer) KeyHandler(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	item := sqlitestore.GetItemByKey(is.store, key)

	if item == nil {
		http.Error(w, "Requested record does not exists", http.StatusNoContent)
		return
	}
	renderJSON(w, item)
}

func (is *ItemServer) DateBeforeHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day := parseYMD(w, r)
	items := sqlitestore.GetItemsBeforeDate(is.store, year, month, day)

	renderJSON(w, items)
}

func (is *ItemServer) DateAfterHandler(w http.ResponseWriter, r *http.Request) {
	year, month, day := parseYMD(w, r)
	items := sqlitestore.GetItemsAfterDate(is.store, year, month, day)

	renderJSON(w, items)
}

func (is *ItemServer) DeleteByKeyHandler(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	sqlitestore.DeleteByKey(is.store, key)
}

func parseYMD(w http.ResponseWriter, r *http.Request) (int, int, int) {
	y := chi.URLParam(r, "year")
	m := chi.URLParam(r, "month")
	d := chi.URLParam(r, "day")

	year, err := strconv.Atoi(y)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Expected integer year, month and day.", http.StatusBadRequest)
	}
	month, err := strconv.Atoi(m)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Expected integer year, month and day.", http.StatusBadRequest)
	}
	day, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
		http.Error(w, "Expected integer year, month and day.", http.StatusBadRequest)
	}
	return year, month, day
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(js)
}

const DBNAME = "test.db"

func main() {
	if _, err := os.Stat(DBNAME); errors.Is(err, os.ErrNotExist) {
		// If .db file does not exist
		log.Println(fmt.Sprintf("Creating %s...", DBNAME))
		file, err := os.Create(DBNAME)
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		db, _ := sql.Open("sqlite3", DBNAME)
		sqlitestore.NewDB(db)
		db.Close()
	}

	db, _ := sql.Open("sqlite3", DBNAME)
	defer db.Close()
	server := NewItemServer(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/list", server.ListHandler)
	r.Post("/add", server.AddHandler)
	r.Get("/key/{key}", server.KeyHandler)
	r.Delete("/key/{key}", server.DeleteByKeyHandler)
	r.Get("/date/before/{year}/{month}/{day}", server.DateBeforeHandler)
	r.Get("/date/after/{year}/{month}/{day}", server.DateAfterHandler)
	err := http.ListenAndServe(":3000", r)
	fmt.Println(err)
}
