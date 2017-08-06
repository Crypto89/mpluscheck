package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Dungeon is a leaderboard
type Dungeon struct {
	realm *Realm
	Name  string  `json:"name"`
	Ranks []*Rank `json:"ranks"`

	lock sync.RWMutex
}

// Rank is an entry on the leaderboard
type Rank struct {
	Position   int64         `json:"position"`
	Level      int64         `json:"level"`
	Duration   time.Duration `json:"duration"`
	Characters []*Character  `json:"characters"`
	Date       time.Time     `json:"date"`
}

// Scrape the board for ranks
func (d *Dungeon) Scrape() {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://worldofwarcraft.com/en-gb/game/pve/leaderboards/%s/%s", d.realm.Name, d.Name))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".SortTable .SortTable-body div.SortTable-row").Each(d.parseRank)
}

func (d *Dungeon) addRank(rank *Rank) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.Ranks = append(d.Ranks, rank)
}

func (d *Dungeon) parseRank(i int, s *goquery.Selection) {
	col := s.Find("div.SortTable-col")
	if len(col.Nodes) != 5 {
		log.Fatalf("Expected 5 nodes but found %d\n", len(col.Nodes))
	}

	rank := &Rank{}

	if position, ok := col.Eq(0).Attr("data-value"); ok {
		p, err := strconv.ParseInt(position, 10, 64)
		if err != nil {
			panic(err)
		}

		rank.Position = p
	}

	if level, ok := col.Eq(1).Attr("data-value"); ok {
		l, err := strconv.ParseInt(level, 10, 64)
		if err != nil {
			panic(err)
		}

		rank.Level = l
	}

	if duration, ok := col.Eq(2).Attr("data-value"); ok {
		du, err := strconv.ParseInt(duration, 10, 64)
		if err != nil {
			panic(err)
		}

		rank.Duration = time.Duration(du)
	}

	// Characters is a bit more complex structure, do something with this later

	if date, ok := col.Eq(4).Attr("data-value"); ok {
		d, err := time.Parse(time.RFC3339, date)
		if err != nil {
			panic(err)
		}

		rank.Date = d
	}

	d.addRank(rank)
}
