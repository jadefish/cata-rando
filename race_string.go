// Code generated by "stringer -type=Race"; DO NOT EDIT.

package cataRando

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NightElf-0]
	_ = x[Worgen-1]
	_ = x[Human-2]
	_ = x[Dwarf-3]
	_ = x[Gnome-4]
	_ = x[Draenei-5]
}

const _Race_name = "NightElfWorgenHumanDwarfGnomeDraenei"

var _Race_index = [...]uint8{0, 8, 14, 19, 24, 29, 36}

func (i Race) String() string {
	if i < 0 || i >= Race(len(_Race_index)-1) {
		return "Race(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Race_name[_Race_index[i]:_Race_index[i+1]]
}
