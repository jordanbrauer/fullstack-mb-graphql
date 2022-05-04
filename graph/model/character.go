package model

import "fullstackmb/app/characters"

var genders = map[string]Gender{
	"n/a":           GenderNone,
	"male":          GenderMale,
	"female":        GenderFemale,
	"hermaphrodite": GenderHermaphrodite,
}

type Character struct {
	Name string `json:"name"`
	// The character's height in centimetres.
	Height string `json:"height"`
	// The character's weight in kilograms.
	Weight string `json:"weight"`
	// The year which a character was born in.
	BornAt string `json:"bornAt"`
	// The character's gender (if known or possible).
	Gender Gender `json:"gender"`
}

func (character *Character) Fill(from *characters.Schema) *Character {
	character.Name = from.Name
	character.Height = from.Height
	character.Weight = from.Mass
	character.BornAt = from.BirthYear
	character.Gender = genders[from.Gender]

	return character
}
