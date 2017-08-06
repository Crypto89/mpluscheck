package models

// Region is a collection of realms
type Region struct {
	Name   string   `json:"name"`
	Realms []*Realm `json:"realms"`
}
