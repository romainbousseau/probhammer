// package calculator provides the calculator object and its methods
package calculator

import (
	"math"
)

// Calculator describes the calculator object
type Calculator struct {
	Attacks   int32 `form:"attacks" binding:"required"`
	Damage    int32 `form:"damage" binding:"required"`
	HitSkill  int32 `form:"hitSkill" binding:"required"`
	Strength  int32 `form:"strength" binding:"required"`
	Toughness int32 `form:"toughness" binding:"required"`
	Save      int32 `form:"save" binding:"required"`
	Results   Results
}

// Results describes the results of the calculator
type Results struct {
	Hit   Result
	Wound Result
	Save  Result
	ExpectedDamage int32
}

// Result describes the result of a roll
type Result struct {
	Success int32
	Failure int32
	Prob    float64
}

// Calculate calculates the results for an attack
func (c *Calculator) Calculate() {
	// Hit roll
	c.Results.Hit = Roll(c.Attacks, c.HitSkill)

	// Wound roll
	woundOn := woundOn(c.Strength, c.Toughness)
	c.Results.Wound = Roll(c.Results.Hit.Success, woundOn)

	// Save roll
	c.Results.Save = Roll(c.Results.Wound.Success, c.Save)

	// Expected damage
	c.Results.ExpectedDamage = c.Results.Save.Failure * c.Damage
}

// Roll rolls dices and returns the result for a given treshold
func Roll(count int32, treshold int32) Result {
	var result Result

	result.Prob = diceProb(treshold)
	result.Success = int32(math.RoundToEven((float64(count) * result.Prob)))
	result.Failure = count - result.Success

	return result
}

func diceProb(skill int32) float64 {
	return (6 - float64(skill) + 1) / 6
}

func woundOn(strength int32, toughness int32) int32 {
	delta := strength - toughness

	switch {
	case delta >= toughness:
		return 2
	case delta > 0:
		return 3
	case delta == 0:
		return 4
	case delta <= -toughness:
		return 6
	case delta < 0:
		return 5
	}

	return 0
}
