package models

import "time"

// Rank is an entry on the leaderboard
type Rank struct {
	Position   int64         `json:"position"`
	Level      int64         `json:"level"`
	Duration   time.Duration `json:"duration"`
	Characters []*Character  `json:"characters"`
	Date       time.Time     `json:"date"`
}
