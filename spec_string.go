// Code generated by "stringer -type=Spec"; DO NOT EDIT.

package cataRando

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Affliction-0]
	_ = x[Arcane-1]
	_ = x[Arms-2]
	_ = x[Assassination-3]
	_ = x[Balance-4]
	_ = x[BeastMastery-5]
	_ = x[Combat-6]
	_ = x[Demonology-7]
	_ = x[Destruction-8]
	_ = x[Discipline-9]
	_ = x[Elemental-10]
	_ = x[Enhancement-11]
	_ = x[Feral-12]
	_ = x[Fire-13]
	_ = x[Frost-14]
	_ = x[Fury-15]
	_ = x[Guardian-16]
	_ = x[Holy-17]
	_ = x[Marksmanship-18]
	_ = x[Protection-19]
	_ = x[Restoration-20]
	_ = x[Retribution-21]
	_ = x[Shadow-22]
	_ = x[Subtlety-23]
	_ = x[Survival-24]
}

const _Spec_name = "AfflictionArcaneArmsAssassinationBalanceBeastMasteryCombatDemonologyDestructionDisciplineElementalEnhancementFeralFireFrostFuryGuardianHolyMarksmanshipProtectionRestorationRetributionShadowSubtletySurvival"

var _Spec_index = [...]uint8{0, 10, 16, 20, 33, 40, 52, 58, 68, 79, 89, 98, 109, 114, 118, 123, 127, 135, 139, 151, 161, 172, 183, 189, 197, 205}

func (i Spec) String() string {
	if i < 0 || i >= Spec(len(_Spec_index)-1) {
		return "Spec(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Spec_name[_Spec_index[i]:_Spec_index[i+1]]
}
