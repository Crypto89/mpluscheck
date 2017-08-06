package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Crypto89/mpluscheck/structs"
	"github.com/PuerkitoBio/goquery"
)

// Scrape the board for ranks
func (d *structs.Dungeon) Scrape() {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://worldofwarcraft.com/en-gb/game/pve/leaderboards/%s/%s", d.realm.Name, d.Name))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".SortTable .SortTable-body div.SortTable-row").Each(d.parseRank)
}

func (d *structs.Dungeon) addRank(rank *Rank) {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.Ranks = append(d.Ranks, rank)
}

func (d *structs.Dungeon) parseRank(i int, s *goquery.Selection) {
	col := s.Find("div.SortTable-col")
	if len(col.Nodes) != 5 {
		log.Fatalf("Expected 5 nodes but found %d\n", len(col.Nodes))
	}

	rank := &structs.Rank{}

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
