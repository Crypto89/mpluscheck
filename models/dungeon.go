package models

import "sync"

// Dungeon is a leaderboard
type Dungeon struct {
	realm *Realm
	Name  string  `json:"name"`
	Ranks []*Rank `json:"ranks"`

	lock sync.RWMutex
}

// NewDungeon returns new Dungeon struct
func NewDungeon(name string, realm *Realm) *Dungeon {
	return &Dungeon{
		Name:  name,
		realm: realm,
	}
}
