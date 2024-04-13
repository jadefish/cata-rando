package cataRando

type Role int

//go:generate stringer -type=Role
const (
	Tank Role = iota
	Healer
	DPS
)
