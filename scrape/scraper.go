package scrape

import (
	"fmt"
	"sync"

	"github.com/Crypto89/mpluscheck/models"
)

type Scraper struct {
	region *models.Region
}

func NewScraper() *Scraper {
	return &Scraper{}
}

// Scrape a realm
func (s *Scraper) Scrape() {
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
