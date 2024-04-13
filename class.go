package cataRando

type Class int

//go:generate stringer -type=Class
const (
	Druid Class = iota
	Hunter
	Mage
	Paladin
	Priest
	Rogue
	Shaman
	Warlock
	Warrior
)
