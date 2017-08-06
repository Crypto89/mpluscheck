package main

import (
	"fmt"

	"encoding/json"
	"sync"
)

var dungeons = []string{
	"black-rook-hold",
	"cathedral-of-eternal-night",
	"court-of-stars",
	"darkheart-thicket",
	"eye-of-azshara",
	"halls-of-valor",
	"maw-of-souls",
	"neltharions-lair",
	"return-to-karazhan-lower",
	"return-to-karazhan-upper",
	"the-arcway",
	"vault-of-the-wardens",
}

// Realm is a collection of dungeons
type Realm struct {
	Name     string     `json:"name"`
	Dungeons []*Dungeon `json:"dungeons"`

	lock sync.RWMutex
}

// Character describes a wow character
type Character struct {
	Name string `json:"name"`
}

// Scrape a realm
func (r *Realm) Scrape() {
	var wg sync.WaitGroup

	for _, name := range dungeons {
		wg.Add(1)
		dungeon := &Dungeon{
			Name:  name,
			realm: r,
		}

		fmt.Printf("Scraping dungeon: %s\n", name)
		go func() {
			dungeon.Scrape()
			r.addDungeon(dungeon)
			wg.Done()
		}()
	}

	wg.Wait()
}

func (r *Realm) addDungeon(dungeon *Dungeon) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Dungeons = append(r.Dungeons, dungeon)
}

func main() {
	realm := &Realm{
		Name: "tarren-mill",
	}

	realm.Scrape()

	result, _ := json.MarshalIndent(realm, "", "  ")

	fmt.Printf("%s\n", string(result))
}
