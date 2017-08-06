package scrape

import (
	"fmt"
	"sync"

	"github.com/Crypto89/mpluscheck/models"
)

type scraper struct {
	realm *models.Realm
}

// Scrape a realm
func (s *scraper) Scrape() {
	var wg sync.WaitGroup

	for _, name := range dungeons {
		wg.Add(1)
		dungeon := &models.Dungeon{
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
