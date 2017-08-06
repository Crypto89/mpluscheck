package models

// Character describes a wow character
type Character struct {
	Name string `json:"name"`
}

// NewCharacter returns new Character struct
func NewCharacter(name string) *Character {
	return &Character{Name: name}
}
