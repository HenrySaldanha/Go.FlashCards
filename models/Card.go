package models

type Card struct {
	Id         string
	Phrase     string
	Meaning    string
	Difficulty Difficulty
	Frequency  Frequency
	Type       Type
}
