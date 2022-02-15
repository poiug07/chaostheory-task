package sqlitestore

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Item struct {
	Timestamp time.Time `json:"timestamp"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

func NewDB(db *sql.DB) {
	q, err := db.Prepare("CREATE TABLE `data` (`key` VARCHAR(30) PRIMARY KEY, `timestamp` DATETIME NULL, `value` VAR CHAR(100))")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	q.Exec()
	fmt.Println("\\")
}

// Adds or overwrites item if already exists
func AddItem(db *sql.DB, key, value string) {
	stmt, _ := db.Prepare("INSERT or REPLACE INTO data(key, timestamp, value) values(?,?,?)")
	// fmt.Println(time.Now().Format("2006-01-02T15:04:05Z"))
	_, err := stmt.Exec(key, time.Now().Format("2006-01-02T15:04:05Z"), value)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetAllItems(db *sql.DB) []Item {
	items := make([]Item, 0)
	row, err := db.Query("SELECT * FROM data ORDER BY timestamp DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var item Item
		row.Scan(&item.Key, &item.Timestamp, &item.Value)
		items = append(items, item)
	}
	return items
}

// Try to get item by Key, return nil if not found.
func GetItemByKey(db *sql.DB, key string) *Item {
	stmt, err := db.Prepare("SELECT * FROM data WHERE key=?;")
	if err != nil {
		log.Fatal(err)
	}
	row, err := stmt.Query(key)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if row.Next() {
		var item Item
		row.Scan(&item.Key, &item.Timestamp, &item.Value)
		return &item
	} else {
		return nil
	}
}

// Gets all items with timestamps less than (<) provided params.
func GetItemsBeforeDate(db *sql.DB, year, month, day int) []Item {
	items := make([]Item, 0)
	stmt, err := db.Prepare("SELECT * FROM data WHERE timestamp<? ORDER BY timestamp DESC")
	if err != nil {
		log.Fatal(err)
	}
	row, err := stmt.Query(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var item Item
		row.Scan(&item.Key, &item.Timestamp, &item.Value)
		items = append(items, item)
	}
	return items
}

func GetItemsAfterDate(db *sql.DB, year, month, day int) []Item {
	items := make([]Item, 0)
	stmt, err := db.Prepare("SELECT * FROM data WHERE timestamp>=? ORDER BY timestamp DESC")
	if err != nil {
		log.Fatal(err)
	}
	row, err := stmt.Query(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var item Item
		row.Scan(&item.Key, &item.Timestamp, &item.Value)
		items = append(items, item)
	}
	return items
}
