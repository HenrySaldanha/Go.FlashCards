package models

type Card struct {
	Id         string     `bson:"id" json:"id"`
	Phrase     string     `bson:"phrase" json:"phrase"`
	Meaning    string     `bson:"meaning" json:"meaning"`
	Difficulty Difficulty `bson:"difficulty" json:"difficulty"`
	Frequency  Frequency  `bson:"frequency" json:"frequency"`
	Type       Type       `bson:"type" json:"type"`
}
