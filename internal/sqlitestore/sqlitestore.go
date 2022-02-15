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

func NewItemStore(db *sql.DB) {
	q, err := db.Prepare("CREATE TABLE `data` (`key` VARCHAR(30) PRIMARY KEY, `timestamp` DATETIME NULL, `value` VAR CHAR(100))")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	q.Exec()
	fmt.Println("\\")
}

func AddItem(db *sql.DB, key, value string) {
	stmt, _ := db.Prepare("INSERT INTO data(key, timestamp, value) values(?,?,?)")
	_, err := stmt.Exec(key, time.Now().Format("2006-01-01T15:04:05Z"), value)
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
