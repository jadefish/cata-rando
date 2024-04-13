package cataRando

type Race int

//go:generate stringer -type=Race
const (
	NightElf Race = iota
	Worgen
	Human
	Dwarf
	Gnome
	Draenei
)
