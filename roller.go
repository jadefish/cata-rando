package cataRando

import (
	"fmt"
	"math/rand"

	"golang.org/x/exp/slices"
)

var classes = []Class{
	Druid,
	Hunter,
	Mage,
	Paladin,
	Priest,
	Rogue,
	Shaman,
	Warlock,
	Warrior,
}
var raceClassMap = map[Class][]Race{
	Druid:   []Race{NightElf, Worgen},
	Hunter:  []Race{Draenei, Dwarf, Human, NightElf, Worgen},
	Mage:    []Race{Draenei, Dwarf, Gnome, Human, NightElf, Worgen},
	Paladin: []Race{Draenei, Dwarf, Human},
	Priest:  []Race{Draenei, Dwarf, Gnome, Human, NightElf, Worgen},
	Rogue:   []Race{Dwarf, Gnome, Human, NightElf, Worgen},
	Shaman:  []Race{Draenei, Dwarf},
	Warlock: []Race{Dwarf, Gnome, Human, Worgen},
	Warrior: []Race{Draenei, Dwarf, Gnome, Human, NightElf, Worgen},
}
var classSpecMap = map[Class][]Spec{
	Druid:   []Spec{Balance, Feral, Guardian, Restoration},
	Hunter:  []Spec{BeastMastery, Marksmanship, Survival},
	Mage:    []Spec{Arcane, Fire, Frost},
	Paladin: []Spec{Holy, Protection, Retribution},
	Priest:  []Spec{Discipline, Holy, Shadow},
	Rogue:   []Spec{Assassination, Combat, Subtlety},
	Shaman:  []Spec{Elemental, Enhancement, Restoration},
	Warlock: []Spec{Affliction, Demonology, Destruction},
	Warrior: []Spec{Arms, Fury, Protection},
}
var specRoleMap = map[Spec]Role{
	Guardian:    Tank,
	Protection:  Tank,
	Restoration: Healer,
	Holy:        Healer,
	Discipline:  Healer,
}

type Pick struct {
	Race  Race
	Class Class
	Spec  Spec
}

func (p Pick) Role() Role {
	if role, ok := specRoleMap[p.Spec]; ok {
		return role
	} else {
		return DPS
	}
}

func (p Pick) String() string {
	return fmt.Sprintf("%s %s %s", p.Race, p.Spec, p.Class)
}

type Roller struct {
	// races contains already-picked races.
	races []Race

	// classes contains already-picked classes.
	classes []Class
}

func NewRoller() Roller {
	return Roller{
		races:   []Race{},
		classes: []Class{},
	}
}

func (r Roller) Roll() Pick {
	n := rand.Intn(len(classes))
	class := classes[n]

	races := raceClassMap[class]
	n = rand.Intn(len(races))
	race := races[n]

	specs := classSpecMap[class]
	n = rand.Intn(len(specs))
	spec := specs[n]

	pick := Pick{race, class, spec}

	return pick
}

func (r Roller) pickRole(role Role) Pick {
	pick := r.Roll()

	for {
		if pick.Role() == role {
			break
		}

		pick = r.Roll()
	}

	return pick
}

// Roll picks a role with a race and class that haven't been already rolled.
func (r *Roller) RollRole(role Role) Pick {
	pick := r.pickRole(role)

	for {
		if slices.Contains(r.races, pick.Race) || slices.Contains(r.classes, pick.Class) {
			pick = r.pickRole(role)
		} else {
			break
		}
	}

	r.races = append(r.races, pick.Race)
	r.classes = append(r.classes, pick.Class)

	return pick
}
