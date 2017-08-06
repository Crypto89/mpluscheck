package models

import "sync"

// Realm is a collection of dungeons
type Realm struct {
	Name     string     `json:"name"`
	Dungeons []*Dungeon `json:"dungeons"`

	lock   sync.RWMutex
	region *Region
}

// NewRealm returns a new realm instance
func NewRealm(region Region, name string) *Realm {
	return &Realm{
		region: region,
		Name:   name,
	}
}

func (r *Realm) AddDungeon(dungeon *Dungeon) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Dungeons = append(r.Dungeons, dungeon)
}
