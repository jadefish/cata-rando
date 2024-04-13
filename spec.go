package cataRando

type Spec int

//go:generate stringer -type=Spec
const (
	Affliction Spec = iota
	Arcane
	Arms
	Assassination
	Balance
	BeastMastery
	Combat
	Demonology
	Destruction
	Discipline
	Elemental
	Enhancement
	Feral
	Fire
	Frost
	Fury
	Guardian
	Holy
	Marksmanship
	Protection
	Restoration
	Retribution
	Shadow
	Subtlety
	Survival
)
