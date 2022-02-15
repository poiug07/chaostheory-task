package store

import (
	"sort"
	"sync"
	"time"
)

type Item struct {
	Timestamp time.Time `json:"timestamp"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

type ItemStore struct {
	sync.Mutex
	items map[string]Item
}

func NewItemStore() *ItemStore {
	return &ItemStore{
		items: make(map[string]Item),
	}
}

func (is *ItemStore) AddItem(key, value string) {
	is.Lock()
	defer is.Unlock()

	item := Item{
		Timestamp: time.Now(),
		Key:       key,
		Value:     value,
	}

	is.items[key] = item
}

func (is *ItemStore) GetAllItems() []Item {
	is.Lock()
	defer is.Unlock()

	items := make([]Item, 0, len(is.items))
	for _, v := range is.items {
		items = append(items, v)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].Timestamp.After(items[j].Timestamp)
	})
	return items
}
