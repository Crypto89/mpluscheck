package main

import (
	"fmt"

	"github.com/Crypto89/mpluscheck/models"

	"encoding/json"
	"sync"
)



func (r *models.Realm) addDungeon(dungeon *models.Dungeon) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Dungeons = append(r.Dungeons, dungeon)
}

func main() {
	realm := &models.Realm{
		Name: "tarren-mill",
	}

	realm.Scrape()

	result, _ := json.MarshalIndent(realm, "", "  ")

	fmt.Printf("%s\n", string(result))
}
