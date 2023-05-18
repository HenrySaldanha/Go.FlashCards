package models

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

type Type int

const (
	Word Type = iota
	Phrase
	Explanation
)
